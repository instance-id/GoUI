package components

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/instance-id/GoUI/appconfig"
	"github.com/instance-id/GoUI/appconfig/writer"
	"github.com/sarulabs/di/v2"
)

//var DatabaseData *DatabaseDetails

var DatabaseData = &DatabaseDetails{Providers: []string{"mysql", "postgres", "mssql", "sqlite"}}

type Containers struct {
	Dac *appconfig.MainSettings
	Dbd *appconfig.DbSettings
	Dba *xorm.Engine
	Wtr *writer.ConfigWriter
}

var Cntnrs *Containers

func CmdInitialize() {
	app := DISetup()
	defer app.Delete()

	Cntnrs.Dac = DataAccessContainer(app)
	Cntnrs.Dbd = DatabaseContainer(app)
	Cntnrs.Dba = DatabaseAccessContainer(app)
	Cntnrs.Wtr = WriterAccessContainer(app)

	fmt.Printf("Test1 %s \n", Cntnrs.Dbd.Data.Address)

	fmt.Printf("Test3 %s \n", Cntnrs.Dac.Assets.AssetCodes)
	Cntnrs.Dac.Assets.AssetCodes = []string{"SCT", "UFPS1"}
	fmt.Printf("Test3 %s \n", Cntnrs.Dac.Assets.AssetCodes)

	InitAssetData()
}

func DISetup() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(Services...)
	app := builder.Build()

	return app
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func DataAccessContainer(di di.Container) *appconfig.MainSettings {
	d, _ := di.Get("configData").(*appconfig.MainSettings)
	return d
}

// --- Provides global access to configuration data via dependency injection container -------------------------------------
func WriterAccessContainer(di di.Container) *writer.ConfigWriter {
	d, _ := di.Get("configWriter").(*writer.ConfigWriter)
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
	database := db.Get("db").(*XormDB).Engine
	return database
}
