package utils

import (
	"github.com/instance-id/GoUI/appconfig"
	"github.com/instance-id/GoUI/services"
	"github.com/sarulabs/di/v2"
)

var (
	Dac *appconfig.MainSettings
)

func CmdInitialize(di di.Container) {
	Dac = DataAccessContainer(di)
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
