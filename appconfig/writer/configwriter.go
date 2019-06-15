package writer

import (
	"fmt"
	. "github.com/instance-id/GoUI/dicontainer"
)

type ConfigWriter struct {
	ConfigData *Data
}

type Data struct {
	FolderName string
	FileName   string
	DbFileName string
	Path       string
	DbPath     string
}

var (
	ConfigData = Data{}
)

// --- Create new config file ------------------------------------------------------------------------------------
func (c *ConfigWriter) NewConfig() (*ConfigWriter, error) {
	c.ConfigData.Path = fmt.Sprintf("./%s/%s", ConfigData.FolderName, ConfigData.FileName)
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

	_ = yml.Set("settings", "assets", "assetcodes", "")
	_ = yml.Set("settings", "assets", "datecompare", "")
	_ = yml.Set("settings", "assets", "comparedate", "")
	_ = yml.Set("settings", "assets", "assetoriginal", "")
	_ = yml.Set("settings", "assets", "assetreplacement", "")
	_ = yml.Set("settings", "assets", "apikey", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "package", map[string]string{"": ""})

	err := yml.Write(c.ConfigData.Path)
	if err != nil {
		fmt.Printf("Error creating new config %s", err)
	}

	return c, err
}

// --- Save current data to config -------------------------------------------------------------------------------
func (c *ConfigWriter) SetConfig() (*ConfigWriter, error) {
	return c.saveConfig()
}

func (c *ConfigWriter) saveConfig() (*ConfigWriter, error) {
	c.ConfigData.Path = fmt.Sprintf("./%s/%s", ConfigData.FolderName, ConfigData.FileName)
	yml := New()

	_ = yml.Set("settings", "system", "token", DiCon.Cnt.Dac.System.Token)
	_ = yml.Set("settings", "system", "commandprefix", DiCon.Cnt.Dac.System.CommandPrefix)
	_ = yml.Set("settings", "system", "requireemail", DiCon.Cnt.Dac.System.RequireEmail)
	_ = yml.Set("settings", "system", "consoleloglevel", DiCon.Cnt.Dac.System.ConsoleLogLevel)
	_ = yml.Set("settings", "system", "fileloglevel", DiCon.Cnt.Dac.System.FileLogLevel)

	_ = yml.Set("settings", "integrations", "wordpress", DiCon.Cnt.Dac.Integrations.WordPress)
	_ = yml.Set("settings", "integrations", "connection", DiCon.Cnt.Dac.Integrations.Connection)
	_ = yml.Set("settings", "integrations", "webaddress", DiCon.Cnt.Dac.Integrations.WebAddress)

	_ = yml.Set("settings", "discord", "guild", DiCon.Cnt.Dac.Discord.Guild)
	_ = yml.Set("settings", "discord", "botusers", DiCon.Cnt.Dac.Discord.BotUsers)
	_ = yml.Set("settings", "discord", "roles", DiCon.Cnt.Dac.Discord.Roles)

	_ = yml.Set("settings", "assets", "assetcodes", DiCon.Cnt.Dac.Assets.AssetCodes)
	_ = yml.Set("settings", "assets", "datecompare", DiCon.Cnt.Dac.Assets.DateCompare)
	_ = yml.Set("settings", "assets", "comparedate", DiCon.Cnt.Dac.Assets.CompareDate)
	_ = yml.Set("settings", "assets", "assetoriginal", DiCon.Cnt.Dac.Assets.AssetOriginal)
	_ = yml.Set("settings", "assets", "assetreplacement", DiCon.Cnt.Dac.Assets.AssetReplacement)
	_ = yml.Set("settings", "assets", "apikey", DiCon.Cnt.Dac.Assets.ApiKeys)
	_ = yml.Set("settings", "assets", "package", DiCon.Cnt.Dac.Assets.Packages)

	err := yml.Write(c.ConfigData.Path)
	if err != nil {
		fmt.Printf("Error reloading new config %s", err)
	}

	return c, err
}

// --- Create new dbconfig file ----------------------------------------------------------------------------------

func (c *ConfigWriter) NewDbConfig() (*ConfigWriter, error) {
	c.ConfigData.Path = fmt.Sprintf("./%s/%s", ConfigData.FolderName, ConfigData.FileName)
	yml := New()

	_ = yml.Set("database", 0)
	_ = yml.Set("data", "address", "")
	_ = yml.Set("data", "username", "")
	_ = yml.Set("data", "password", "")
	_ = yml.Set("data", "dbname", "")
	_ = yml.Set("data", "tableprefix", "")

	err := yml.Write(c.ConfigData.Path)
	if err != nil {
		fmt.Printf("Error writing new dbconfig %s", err)
	}
	return c, err
}

// --- Save current data to dbconfig -----------------------------------------------------------------------------
func (c *ConfigWriter) SetDbConfig() (*ConfigWriter, error) {
	return c.saveDbConfig()
}

func (c *ConfigWriter) saveDbConfig() (*ConfigWriter, error) {
	c.ConfigData.Path = fmt.Sprintf("./%s/%s", ConfigData.FolderName, ConfigData.FileName)
	yml := New()

	_ = yml.Set("database", DiCon.Cnt.Dbd.Database)
	_ = yml.Set("data", "address", DiCon.Cnt.Dbd.Data.Address)
	_ = yml.Set("data", "username", DiCon.Cnt.Dbd.Data.Username)
	_ = yml.Set("data", "password", DiCon.Cnt.Dbd.Data.Password)
	_ = yml.Set("data", "dbname", DiCon.Cnt.Dbd.Data.DbName)
	_ = yml.Set("data", "tableprefix", DiCon.Cnt.Dbd.Data.TablePrefix)

	err := yml.Write(c.ConfigData.Path)
	if err != nil {
		fmt.Printf("Error writing new dbconfig %s", err)
	}
	return c, err
}
