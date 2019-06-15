package appconfig

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	. "github.com/instance-id/GoUI/dicontainer"
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
		AssetCodes       []string          `json:"assetcodes"`
		DateCompare      string            `json:"datecompare"`
		CompareDate      string            `json:"comparedate"`
		AssetOriginal    string            `json:"assetoriginal"`
		AssetReplacement string            `json:"assetreplacement"`
		ApiKeys          map[string]string `json:"apikey"`
		Packages         map[string]string `json:"package"`
	} `json:"assets"`
}

func (m *MainSettings) GetConfig() *MainSettings {
	return m.loadConfig()
}

func (m *MainSettings) loadConfig() *MainSettings {

	fmt.Printf("Config path from writer in di: %s \n", DiCon.Cnt.Wtr.ConfigData.Path)

	config.AddDriver(yaml.Driver)

	if _, err := os.Stat(DiCon.Cnt.Wtr.ConfigData.Path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(DiCon.Cnt.Wtr.ConfigData.FolderName, 0755)
			if err != nil {
				fmt.Printf("Error at creating new config %s", err)
			}
			_, err = DiCon.Cnt.Wtr.NewConfig()
			if err != nil {
				fmt.Printf("Error at setting new config %s", err)
			}
		}
	}

	err := config.LoadFiles(string(DiCon.Cnt.Wtr.ConfigData.Path))
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
			AssetCodes       []string          `json:"assetcodes"`
			DateCompare      string            `json:"datecompare"`
			CompareDate      string            `json:"comparedate"`
			AssetOriginal    string            `json:"assetoriginal"`
			AssetReplacement string            `json:"assetreplacement"`
			ApiKeys          map[string]string `json:"apikey"`
			Packages         map[string]string `json:"package"`
		}{
			AssetCodes:       config.Strings("settings.assets.assetcodes"),
			DateCompare:      config.String("settings.assets.datecompare"),
			CompareDate:      config.String("settings.assets.comparedate"),
			AssetOriginal:    config.String("settings.assets.assetoriginal"),
			AssetReplacement: config.String("settings.assets.assetreplacement"),
			ApiKeys:          config.StringMap("settings.assets.apikey"),
			Packages:         config.StringMap("settings.assets.package"),
		}}

	return mainSettings
}
