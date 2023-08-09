// Package database contains all the database related code
package database

// InitializeDatabase initializes the database
func InitializeDatabase() {
	setCore()
	setItems()
	setLocales()
	setLanguages()
	setHandbook()
	setTraders()
	setLocations()
	setQuests()
	setHideout()
	setWeather()
	setCustomization()
	setBots()
	setEditions()
	setFlea()
	setProfiles()
}

const (
	databaseLibPath       string = "assets/database"
	coreFilePath          string = databaseLibPath + "/core"
	botTemplateFilePath   string = coreFilePath + "/botTemplate.json"
	playerScavPath        string = coreFilePath + "/playerScav.json"
	clientSettingsPath    string = coreFilePath + "/client.settings.json"
	globalBotSettingsPath string = coreFilePath + "/__BotGlobalSettings.json"
	globalsFilePath       string = coreFilePath + "/globals.json"
	locationsFilePath     string = coreFilePath + "/locations.json"
	matchMetricsPath      string = coreFilePath + "/matchMetrics.json"
	serverConfigPath      string = coreFilePath + "/server.json"
	editionsDirPath       string = databaseLibPath + "/editions/"
	itemsPath             string = databaseLibPath + "/items.json"
	localesPath           string = databaseLibPath + "/locales"
	handbookPath          string = databaseLibPath + "/handbook.json"
	traderPath            string = databaseLibPath + "/traders/"
	questsPath            string = databaseLibPath + "/quests.json"
	hideoutPath           string = databaseLibPath + "/hideout/"
	weatherPath           string = databaseLibPath + "/weather.json"
	customizationPath     string = databaseLibPath + "/customization.json"
	botsPath              string = databaseLibPath + "/bot/"
	botsDirectory         string = botsPath + "bots/"
)
