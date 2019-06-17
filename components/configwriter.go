package components

import (
	"fmt"
)

type ConfigWriter struct {
	ConfigPath *ConfigPath
}

// --- Create new config file ------------------------------------------------------------------------------------
func (c *ConfigWriter) NewConfig() (*ConfigWriter, error) {
	configPath.Path = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.FileName)
	yml := New()

	_ = yml.Set("settings", "system", "token", "")
	_ = yml.Set("settings", "system", "commandprefix", "")
	_ = yml.Set("settings", "system", "requireemail", "")
	_ = yml.Set("settings", "system", "consoleloglevel", 0)
	_ = yml.Set("settings", "system", "fileloglevel", 0)

	_ = yml.Set("settings", "integrations", "wordpress", "")
	_ = yml.Set("settings", "integrations", "connection", "")
	_ = yml.Set("settings", "integrations", "webaddress", "")

	_ = yml.Set("settings", "discord", "guild", "")
	_ = yml.Set("settings", "discord", "botusers", []string{""})
	_ = yml.Set("settings", "discord", "roles", map[string]string{"": ""})

	_ = yml.Set("settings", "assets", "assetcodes", []string{""})
	_ = yml.Set("settings", "assets", "assetreplaced", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "assetreplacement", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "replacedate", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "apikey", map[string]string{"": ""})
	_ = yml.Set("settings", "assets", "package", map[string]string{"": ""})

	err := yml.Write(configPath.Path)
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
	configPath.Path = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.FileName)
	yml := New()

	_ = yml.Set("settings", "system", "token", Cntnrs.Dac.System.Token)
	_ = yml.Set("settings", "system", "commandprefix", Cntnrs.Dac.System.CommandPrefix)
	_ = yml.Set("settings", "system", "requireemail", Cntnrs.Dac.System.RequireEmail)
	_ = yml.Set("settings", "system", "consoleloglevel", Cntnrs.Dac.System.ConsoleLogLevel)
	_ = yml.Set("settings", "system", "fileloglevel", Cntnrs.Dac.System.FileLogLevel)

	_ = yml.Set("settings", "integrations", "wordpress", Cntnrs.Dac.Integrations.WordPress)
	_ = yml.Set("settings", "integrations", "connection", Cntnrs.Dac.Integrations.Connection)
	_ = yml.Set("settings", "integrations", "webaddress", Cntnrs.Dac.Integrations.WebAddress)

	_ = yml.Set("settings", "discord", "guild", Cntnrs.Dac.Discord.Guild)
	_ = yml.Set("settings", "discord", "botusers", Cntnrs.Dac.Discord.BotUsers)
	_ = yml.Set("settings", "discord", "roles", Cntnrs.Dac.Discord.Roles)

	_ = yml.Set("settings", "assets", "assetcodes", Cntnrs.Dac.Assets.AssetCodes)
	_ = yml.Set("settings", "assets", "assetreplaced", Cntnrs.Dac.Assets.AssetReplaced)
	_ = yml.Set("settings", "assets", "assetreplacement", Cntnrs.Dac.Assets.AssetReplacement)
	_ = yml.Set("settings", "assets", "replacedate", Cntnrs.Dac.Assets.ReplaceDate)
	_ = yml.Set("settings", "assets", "apikey", Cntnrs.Dac.Assets.ApiKeys)
	_ = yml.Set("settings", "assets", "package", Cntnrs.Dac.Assets.Packages)

	err := yml.Write(configPath.Path)
	if err != nil {
		fmt.Printf("Error reloading new config %s", err)
	}

	return c, err
}

// --- Create new dbconfig file ----------------------------------------------------------------------------------

func (c *ConfigWriter) NewDbConfig() (*ConfigWriter, error) {
	configPath.DbPath = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.DbFileName)
	yml := New()

	_ = yml.Set("database", 0)
	_ = yml.Set("data", "address", "")
	_ = yml.Set("data", "username", "")
	_ = yml.Set("data", "password", "")
	_ = yml.Set("data", "dbname", "")
	_ = yml.Set("data", "tableprefix", "")

	err := yml.Write(configPath.DbPath)
	return c, err
}

// --- Save current data to dbconfig -----------------------------------------------------------------------------
func (c *ConfigWriter) SetDbConfig() (*ConfigWriter, error) {
	return c.saveDbConfig()
}

func (c *ConfigWriter) saveDbConfig() (*ConfigWriter, error) {
	configPath.DbPath = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.DbFileName)
	yml := New()

	_ = yml.Set("database", Cntnrs.Dbd.Database)
	_ = yml.Set("data", "address", Cntnrs.Dbd.Data.Address)
	_ = yml.Set("data", "username", Cntnrs.Dbd.Data.Username)
	_ = yml.Set("data", "password", Cntnrs.Dbd.Data.Password)
	_ = yml.Set("data", "dbname", Cntnrs.Dbd.Data.DbName)
	_ = yml.Set("data", "tableprefix", Cntnrs.Dbd.Data.TablePrefix)

	err := yml.Write(configPath.DbPath)
	if err != nil {
		fmt.Printf("Error writing new dbconfig %s", err)
	}
	return c, err
}
