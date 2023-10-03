package server

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"MT-GO/database"
	"MT-GO/services"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func upgradeToWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	sessionID := r.URL.Path[28:] //mongoID is 24 chars
	database.SetConnection(sessionID, conn)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [" + r.Method + "] Request URL: [" + r.URL.Path + "] on [" + strings.TrimPrefix(r.Host, "127.0.0.1") + "]")

		if r.Header.Get("Connection") == "Upgrade" && r.Header.Get("Upgrade") == "websocket" {
			upgradeToWebsocket(w, r)
		} else {
			buffer := &bytes.Buffer{}

			if r.Header.Get("Content-Type") != "application/json" {
				next.ServeHTTP(w, r)
				return
			}

			buffer = services.ZlibInflate(r)
			if buffer == nil || buffer.Len() == 0 {
				next.ServeHTTP(w, r)
				return
			}

			var parsedData map[string]interface{}
			decoder := json.NewDecoder(buffer)
			if err := decoder.Decode(&parsedData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			/// Store the parsed data in the request's context
			ctx := context.WithValue(r.Context(), services.ParsedBodyKey, parsedData)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		}
	})
}

func startHTTPSServer(serverReady chan<- struct{}, certs *services.Certificate, mux *muxt) {
	mux.initRoutes(mux.mux)

	httpsServer := &http.Server{
		Addr: mux.address,
		TLSConfig: &tls.Config{
			RootCAs:      nil,
			Certificates: []tls.Certificate{certs.Certificate},
		},
		Handler: logAndDecompress(mux.mux),
	}

	fmt.Println("Started " + mux.serverName + " HTTPS server on " + mux.address)
	serverReady <- struct{}{}

	err := httpsServer.ListenAndServeTLS(certs.CertFile, certs.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
}

type muxt struct {
	mux        *http.ServeMux
	address    string
	serverName string
	initRoutes func(mux *http.ServeMux)
}

func SetHTTPSServer() {
	srv := database.GetServerConfig()

	cert := services.GetCertificate(srv.IP)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
	cert.Certificate = certs

	fmt.Println()

	// serve static content
	mainServeMux := http.NewServeMux()
	ServeStaticMux(mainServeMux)

	muxers := []*muxt{
		{
			mux: mainServeMux, address: database.GetMainIPandPort(),
			serverName: "Main", initRoutes: setMainRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetTradingIPandPort(),
			serverName: "Trading", initRoutes: setTradingRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetMessagingIPandPort(),
			serverName: "Messaging", initRoutes: setMessagingRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetRagFairIPandPort(),
			serverName: "RagFair", initRoutes: setRagfairRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetLobbyIPandPort(),
			serverName: "Lobby", initRoutes: setLobbyRoutes, // Embed the route initialization function
		},
	}

	serverReady := make(chan struct{})

	for _, muxData := range muxers {
		go startHTTPSServer(serverReady, cert, muxData)
	}

	for range muxers {
		<-serverReady
	}

	close(serverReady)
	fmt.Println()
}
