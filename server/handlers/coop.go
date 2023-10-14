package handlers

import (
	"MT-GO/services"
	"log"
	"net/http"
)

type coopInvites struct {
	Players []map[string]interface{} `json:"players"`
	Invite  []interface{}            `json:"invite"`
	Group   []interface{}            `json:"group"`
}

var coopStatusOutput = coopInvites{
	Players: make([]map[string]interface{}, 0),
	Invite:  make([]interface{}, 0),
	Group:   make([]interface{}, 0),
}

const routeNotImplemented = "Route is not implemented yet, using empty values instead"

func CoopServerStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Coop Server Match Status")
	services.ZlibReply(w, "")
}

func CoopGetInvites(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Coop Server Invites")
	services.ZlibJSONReply(w, coopStatusOutput)
}

func CoopServerDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting Coop Server")
	body := services.ApplyResponseBody(map[string]string{"response": "OK"})
	services.ZlibJSONReply(w, body)
}

func CoopConnect(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}

func CoopServerUpdate(w http.ResponseWriter, r *http.Request) {
	serverID := services.GetParsedBody(r).(map[string]interface{})["serverId"].(string)
	//GET COOP SERVER DATA HERE AND RETURN
	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerReadPlayers(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerJoin(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerExists(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerCreate(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerGetAllForLocation(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerFriendlyAI(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
func CoopServerSpawnPoint(w http.ResponseWriter, r *http.Request) {

	body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, body)
}
