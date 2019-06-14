package utils

import (
	"github.com/go-xorm/xorm"
	"github.com/instance-id/GoUI/appconfig"
	"github.com/instance-id/GoUI/components"
	"github.com/instance-id/GoUI/services"
	"github.com/sarulabs/di/v2"
)

var (
	Dac *appconfig.MainSettings
	Dbd *appconfig.DbSettings
	Dba *xorm.Engine
)

func CmdInitialize(di di.Container) {
	Dac = DataAccessContainer(di)
	Dbd = DatabaseContainer(di)
	Dba = DatabaseAccessContainer(di)
}

func DISetup() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(services.Services...)
	app := builder.Build()

	return app
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func DataAccessContainer(di di.Container) *appconfig.MainSettings {
	d, _ := di.Get("configData").(*appconfig.MainSettings)
	return d
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func DatabaseContainer(di di.Container) *appconfig.DbSettings {
	d, _ := di.Get("dbData").(*appconfig.DbSettings)
	return d
}

// --- Provides global database access via dependency injection container --------------------------------------------------
func DatabaseAccessContainer(di di.Container) *xorm.Engine {
	db, _ := di.SubContainer()
	database := db.Get("db").(*components.XormDB).Engine
	return database
}
