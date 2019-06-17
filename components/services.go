package components

import (
	"github.com/sarulabs/di/v2"
)

var Services = []di.Def{
	{
		Name:  "configData",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var cfg MainSettings
			config := cfg.GetConfig()
			return config, nil
		}},
	{
		Name:  "configWriter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var writer ConfigWriter
			return writer, nil
		}},
	{
		// --- Get DB Data ----------------------------------------------------------------------
		Name:  "dbData",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var db DbSettings
			dbConfig := db.GetDbConfig()
			return dbConfig, nil
		}},
	{
		// --- Creates database connection object ----------------------------------------------------------------------
		Name:  "dbConn",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var conn DbConfig
			dbConfig := ctn.Get("dbData").(*DbSettings)
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
			conn := ctn.Get("dbConn").(*DbConfig)
			return conn, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*DbConfig).Xorm.Close()
		}},
}
