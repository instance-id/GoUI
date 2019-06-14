package appconfig

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type ConfigData struct {
	MainSettings MainSettings
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

func (m *MainSettings) GetConfig() *MainSettings {
	return m.loadConfig()
}

func (m *MainSettings) loadConfig() *MainSettings {

	config.AddDriver(yaml.Driver)
	filename := "./config/config.yml"

	err := config.LoadFiles(string(filename))
	if err != nil {
		panic(err)
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

func (m *MainSettings) saveConfig() (*MainSettings, error) {
	yml := New()

	yml.Set("settings", "system", "token", m.System.Token)
	yml.Set("settings", "system", "commandprefix", m.System.CommandPrefix)
	yml.Set("settings", "system", "requireemail", m.System.RequireEmail)
	yml.Set("settings", "system", "consoleloglevel", m.System.ConsoleLogLevel)
	yml.Set("settings", "system", "fileloglevel", m.System.FileLogLevel)

	yml.Set("settings", "integrations", "wordpress", m.Integrations.WordPress)
	yml.Set("settings", "integrations", "connection", m.Integrations.Connection)
	yml.Set("settings", "integrations", "webaddress", m.Integrations.WebAddress)

	yml.Set("settings", "discord", "guild", m.Discord.Guild)
	yml.Set("settings", "discord", "botusers", m.Discord.BotUsers)
	yml.Set("settings", "discord", "roles", m.Discord.Roles)

	yml.Set("settings", "assets", "datecompare", "No")
	yml.Set("settings", "assets", "comparedate", m.Assets.CompareDate)
	yml.Set("settings", "assets", "assetoriginal", m.Assets.AssetOriginal)
	yml.Set("settings", "assets", "assetreplacement", m.Assets.AssetReplacement)
	yml.Set("settings", "assets", "apikey", m.Assets.ApiKeys)
	yml.Set("settings", "assets", "package", m.Assets.Packages)

	err := yml.Write("./config/config.yml")
	if err != nil {
		panic(err)
	}

	return m, err
}
