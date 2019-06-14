package services

import (
	"github.com/instance-id/GoUI/appconfig"
	"github.com/instance-id/GoUI/components"

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
		// --- Creates database connection object ----------------------------------------------------------------------
		Name:  "dbData",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var db appconfig.DbSettings
			dbConfig := db.GetDbConfig()
			return dbConfig, nil
		},
	},
	{
		// --- Creates database connection object ----------------------------------------------------------------------
		Name:  "dbConn",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var conn components.DbConfig
			dbConfig := ctn.Get("dbData").(*appconfig.DbSettings)
			dbConn := conn.ConnectDB(dbConfig)
			return dbConn, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*components.DbConfig).Xorm.Close()
		},
	},
	{
		// --- Uses database connection object and returns a connection session ----------------------------------------
		Name:  "db",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			conn := ctn.Get("dbConn").(*components.DbConfig).Xorm
			return conn, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*components.DbConfig).Xorm.Close()
		},
	},
}
