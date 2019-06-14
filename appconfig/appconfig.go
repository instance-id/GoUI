package appconfig

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

type ConfigData struct {
	MainSettings *MainSettings
	folderName   string
	fileName     string
	path         string
}

type ConnectionType int

const (
	Oauth  ConnectionType = 0
	Oauth2 ConnectionType = 1
	JWT    ConnectionType = 2
)

type WordPressSettings struct {
	ConnectionType    int
	ConnectionDetails struct {
		SiteAddress  string `json:"site_address"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Redirect     string `json:"redirect"`
		AppName      string `json:"app_name"`
		AccessUrl    string `json:"access_url"`
		AuthorizeUrl string `json:"authorize_url"`
		GrantType    string `json:"grant_type"`
	}
	Assets map[string]string `json:"assets"`
}

type MainSettings struct {
	System struct {
		Token           string `json:"token"`
		CommandPrefix   string `json:"commandprefix"`
		RequireEmail    string `json:"requireemail"`
		ConsoleLogLevel string `json:"consoleloglevel"`
		FileLogLevel    string `json:"fileloglevel"`
	} `json:"system"`
	Integrations struct {
		WordPress  string `json:"wordpress"`
		Connection string `json:"connection"`
		WebAddress string `json:"webaddress"`
	} `json:"integrations"`
	Discord struct {
		Guild    string            `json:"guild"`
		BotUsers []string          `json:"botusers"`
		Roles    map[string]string `json:"roles"`
	} `json:"discord"`
	Assets struct {
		DateCompare      string            `json:"datecompare"`
		CompareDate      string            `json:"comparedate"`
		AssetOriginal    string            `json:"assetoriginal"`
		AssetReplacement string            `json:"assetreplacement"`
		ApiKeys          map[string]string `json:"apikey"`
		Packages         map[string]string `json:"package"`
	} `json:"assets"`
}

var configData ConfigData

func (m *MainSettings) GetConfig() *MainSettings {
	return m.loadConfig()
}

func (m *MainSettings) loadConfig() *MainSettings {

	config.AddDriver(yaml.Driver)
	configData.folderName = "config"
	configData.fileName = "config.yml"
	configData.path = fmt.Sprintf("./%s/%s", configData.folderName, configData.fileName)

	fmt.Printf("Config path: %s \n", configData.path)

	if _, err := os.Stat(configData.path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(configData.folderName, 0755)
			if err != nil {
				fmt.Printf("Error at creating new config %s", err)
			}
			_, err = m.newConfig()
			if err != nil {
				fmt.Printf("Error at setting new config %s", err)
			}
		}
	}

	err := config.LoadFiles(string(configData.path))
	if err != nil {
		fmt.Printf("Error at loading config %s", err)
	}

	mainSettings := &MainSettings{
		System: struct {
			Token           string `json:"token"`
			CommandPrefix   string `json:"commandprefix"`
			RequireEmail    string `json:"requireemail"`
			ConsoleLogLevel string `json:"consoleloglevel"`
			FileLogLevel    string `json:"fileloglevel"`
		}{
			Token:           config.String("settings.system.token"),
			CommandPrefix:   config.String("settings.system.commandprefix"),
			RequireEmail:    config.String("settings.system.requireemail"),
			ConsoleLogLevel: config.String("settings.system.consoleloglevel"),
			FileLogLevel:    config.String("settings.system.fileloglevel"),
		},
		Integrations: struct {
			WordPress  string `json:"wordpress"`
			Connection string `json:"connection"`
			WebAddress string `json:"webaddress"`
		}{
			WordPress:  config.String("settings.integrations.wordpress"),
			Connection: config.String("settings.integrations.connection"),
			WebAddress: config.String("settings.integrations.webaddress"),
		},
		Discord: struct {
			Guild    string            `json:"guild"`
			BotUsers []string          `json:"botusers"`
			Roles    map[string]string `json:"roles"`
		}{
			Guild:    config.String("settings.discord.guild"),
			BotUsers: config.Strings("settings.discord.botusers"),
			Roles:    config.StringMap("settings.discord.roles"),
		},
		Assets: struct {
			DateCompare      string            `json:"datecompare"`
			CompareDate      string            `json:"comparedate"`
			AssetOriginal    string            `json:"assetoriginal"`
			AssetReplacement string            `json:"assetreplacement"`
			ApiKeys          map[string]string `json:"apikey"`
			Packages         map[string]string `json:"package"`
		}{
			DateCompare:      config.String("settings.assets.datecompare"),
			CompareDate:      config.String("settings.assets.comparedate"),
			AssetOriginal:    config.String("settings.assets.assetoriginal"),
			AssetReplacement: config.String("settings.assets.assetreplacement"),
			ApiKeys:          config.StringMap("settings.assets.apikey"),
			Packages:         config.StringMap("settings.assets.package"),
		}}

	return mainSettings
}

func (m *MainSettings) SetConfig() (*MainSettings, error) {
	return m.saveConfig()
}

func (m *MainSettings) newConfig() (*MainSettings, error) {
	yml := New()

	_ = yml.Set("settings", "system", "token", "")
	_ = yml.Set("settings", "system", "commandprefix", "")
	_ = yml.Set("settings", "system", "requireemail", "")
	_ = yml.Set("settings", "system", "consoleloglevel", "")
	_ = yml.Set("settings", "system", "fileloglevel", "")

	_ = yml.Set("settings", "integrations", "wordpress", "")
	_ = yml.Set("settings", "integrations", "connection", "")
	_ = yml.Set("settings", "integrations", "webaddress", "")

	_ = yml.Set("settings", "discord", "guild", "")
	_ = yml.Set("settings", "discord", "botusers", []string{""})
	_ = yml.Set("settings", "discord", "roles", map[string]string{"": ""})

	_ = yml.Set("settings", "assets", "datecompare", "")
	_ = yml.Set("settings", "assets", "comparedate", "")
	_ = yml.Set("settings", "assets", "assetoriginal", "")
	_ = yml.Set("settings", "assets", "assetreplacement", "")
	_ = yml.Set("settings", "assets", "apikey", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "package", map[string]string{"": ""})

	err := yml.Write(configData.path)
	if err != nil {
		fmt.Printf("Error creating new config %s", err)
	}

	return m, err
}

func (m *MainSettings) saveConfig() (*MainSettings, error) {
	yml := New()

	_ = yml.Set("settings", "system", "token", m.System.Token)
	_ = yml.Set("settings", "system", "commandprefix", m.System.CommandPrefix)
	_ = yml.Set("settings", "system", "requireemail", m.System.RequireEmail)
	_ = yml.Set("settings", "system", "consoleloglevel", m.System.ConsoleLogLevel)
	_ = yml.Set("settings", "system", "fileloglevel", m.System.FileLogLevel)

	_ = yml.Set("settings", "integrations", "wordpress", m.Integrations.WordPress)
	_ = yml.Set("settings", "integrations", "connection", m.Integrations.Connection)
	_ = yml.Set("settings", "integrations", "webaddress", m.Integrations.WebAddress)

	_ = yml.Set("settings", "discord", "guild", m.Discord.Guild)
	_ = yml.Set("settings", "discord", "botusers", m.Discord.BotUsers)
	_ = yml.Set("settings", "discord", "roles", m.Discord.Roles)

	_ = yml.Set("settings", "assets", "datecompare", m.Assets.DateCompare)
	_ = yml.Set("settings", "assets", "comparedate", m.Assets.CompareDate)
	_ = yml.Set("settings", "assets", "assetoriginal", m.Assets.AssetOriginal)
	_ = yml.Set("settings", "assets", "assetreplacement", m.Assets.AssetReplacement)
	_ = yml.Set("settings", "assets", "apikey", m.Assets.ApiKeys)
	_ = yml.Set("settings", "assets", "package", m.Assets.Packages)

	err := yml.Write(configData.path)
	if err != nil {
		fmt.Printf("Error reloading new config %s", err)
	}

	return m, err
}
