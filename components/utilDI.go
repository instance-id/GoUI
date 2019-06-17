package components

import (
	"github.com/sarulabs/di/v2"
)

type Containers struct {
	Dac *MainSettings
	Dbd *DbSettings
	Wtr *ConfigWriter
}

var Cntnrs = new(Containers)

func CmdInitialize() {
	app := DISetup()
	defer app.Delete()

	Cntnrs.Dac = DataAccessContainer(app)
	Cntnrs.Dbd = DatabaseContainer(app)
	Cntnrs.Wtr = WriterAccessContainer(app)
}

func DISetup() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(Services...)
	app := builder.Build()

	return app
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func DataAccessContainer(di di.Container) *MainSettings {
	d, _ := di.Get("configData").(*MainSettings)
	return d
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func WriterAccessContainer(di di.Container) *ConfigWriter {
	d, _ := di.Get("configWriter").(*ConfigWriter)
	return d
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func DatabaseContainer(di di.Container) *DbSettings {
	d, _ := di.Get("dbData").(*DbSettings)
	return d
}
