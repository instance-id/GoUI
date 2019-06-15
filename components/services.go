package components

import (
	"fmt"
	"github.com/instance-id/GoUI/appconfig"
	"github.com/instance-id/GoUI/appconfig/writer"
	"github.com/sarulabs/di/v2"
)

var Services = []di.Def{
	{
		Name:  "configData",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var cfg appconfig.MainSettings
			config := cfg.GetConfig()
			return config, nil
		}},
	{
		Name:  "configWriter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var writer = new(writer.ConfigWriter)
			writer.ConfigData.FolderName = "config"
			writer.ConfigData.FileName = "config.yml"
			writer.ConfigData.DbFileName = "dbconfig.yml"
			writer.ConfigData.Path = fmt.Sprintf("./%s/%s", writer.ConfigData.FolderName, writer.ConfigData.FileName)
			writer.ConfigData.DbPath = fmt.Sprintf("./%s/%s", writer.ConfigData.FolderName, writer.ConfigData.DbFileName)
			return writer, nil
		}},
	{
		// --- Creates database connection object ----------------------------------------------------------------------
		Name:  "dbData",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var db appconfig.DbSettings
			dbConfig := db.GetDbConfig()
			return dbConfig, nil
		}},
	{
		// --- Creates database connection object ----------------------------------------------------------------------
		Name:  "dbConn",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var conn DbConfig
			dbConfig := ctn.Get("dbData").(*appconfig.DbSettings)
			dbConn := conn.ConnectDB(dbConfig)
			return dbConn, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*DbConfig).Xorm.Close()
		}},
	{
		// --- Uses database connection object and returns a connection session ----------------------------------------
		Name:  "db",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			conn := ctn.Get("dbConn").(*DbConfig).Xorm
			return conn, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*DbConfig).Xorm.Close()
		}},
}
