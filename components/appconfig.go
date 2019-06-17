package components

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

type ConnectionType int

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
		ConsoleLogLevel int    `json:"consoleloglevel"`
		FileLogLevel    int    `json:"fileloglevel"`
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
		AssetCodes       []string          `json:"assetcodes"`
		ReplaceDate      map[string]string `json:"replacedate"`
		AssetReplaced    map[string]string `json:"assetreplaced"`
		AssetReplacement map[string]string `json:"assetreplacement"`
		ApiKeys          map[string]string `json:"apikey"`
		Packages         map[string]string `json:"package"`
		Version          map[string]string `json:"version"`
	} `json:"assets"`
}

func (m *MainSettings) GetConfig() *MainSettings {
	return m.loadConfig()
}

func (m *MainSettings) loadConfig() *MainSettings {
	config.AddDriver(yaml.Driver)

	if _, err := os.Stat(configPath.Path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(configPath.FolderName, 0755)
			if err != nil {
				fmt.Printf("Error at creating new config %s", err)
			}
			_, err = Cntnrs.Wtr.NewConfig()
			if err != nil {
				fmt.Printf("Error at setting new config %s", err)
			}
		}
	}

	err := config.LoadFiles(string(configPath.Path))
	if err != nil {
		fmt.Printf("Error at loading config %s", err)
	}

	mainSettings := &MainSettings{
		System: struct {
			Token           string `json:"token"`
			CommandPrefix   string `json:"commandprefix"`
			RequireEmail    string `json:"requireemail"`
			ConsoleLogLevel int    `json:"consoleloglevel"`
			FileLogLevel    int    `json:"fileloglevel"`
		}{
			Token:           config.String("settings.system.token"),
			CommandPrefix:   config.String("settings.system.commandprefix"),
			RequireEmail:    config.String("settings.system.requireemail"),
			ConsoleLogLevel: config.Int("settings.system.consoleloglevel"),
			FileLogLevel:    config.Int("settings.system.fileloglevel"),
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
			AssetCodes       []string          `json:"assetcodes"`
			ReplaceDate      map[string]string `json:"replacedate"`
			AssetReplaced    map[string]string `json:"assetreplaced"`
			AssetReplacement map[string]string `json:"assetreplacement"`
			ApiKeys          map[string]string `json:"apikey"`
			Packages         map[string]string `json:"package"`
			Version          map[string]string `json:"version"`
		}{
			AssetCodes:       config.Strings("settings.assets.assetcodes"),
			ReplaceDate:      config.StringMap("settings.assets.replacedate"),
			AssetReplaced:    config.StringMap("settings.assets.assetreplaced"),
			AssetReplacement: config.StringMap("settings.assets.assetreplacement"),
			ApiKeys:          config.StringMap("settings.assets.apikey"),
			Packages:         config.StringMap("settings.assets.package"),
			Version:          config.StringMap("settings.assets.version"),
		}}

	return mainSettings
}
