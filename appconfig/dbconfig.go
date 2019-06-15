package appconfig

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	. "github.com/instance-id/GoUI/dicontainer"
	"os"
)

var dbData DbData

type DbData struct {
	DbSettings DbSettings
	folderName string
	fileName   string
	path       string
}

// --- Maps dbconfig.yml fields to DbSettings fields -------------------------------------------------------------------
type DbSettings struct {
	Database int `json:"database"`
	Data     struct {
		Address     string `json:"address"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		DbName      string `json:"dbname"`
		TablePrefix string `json:"tableprefix"`
	} `json:"data"`
}

// --- Gets called from Services and returns DbSettings to Dependency Injection container ------------------------------
func (d *DbSettings) GetDbConfig() *DbSettings {
	return d.loadDbConfig()
}

// --- Populates the DbSettings struct from dbconfig.yml file and returns the data for use -----------------------------
func (d *DbSettings) loadDbConfig() *DbSettings {

	config.AddDriver(yaml.Driver)

	fmt.Printf("Db path: %s \n", dbData.path)

	if _, err := os.Stat(dbData.path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dbData.folderName, 0755)
			if err != nil {
				fmt.Printf("Error at creating new dbconfig %s", err)
			}
			_, err = DiCon.Cnt.Wtr.NewDbConfig()
			if err != nil {
				fmt.Printf("Error at setting new dbconfig %s", err)
			}
		}
	}

	err := config.LoadFiles(string(dbData.path))
	if err != nil {
		fmt.Printf("Error loading dbconfig %s", err)
	}

	dbSettings := &DbSettings{
		Database: config.Int("database"),
		Data: struct {
			Address     string `json:"address"`
			Username    string `json:"username"`
			Password    string `json:"password"`
			DbName      string `json:"dbname"`
			TablePrefix string `json:"tableprefix"`
		}{
			Address:     config.String("data.address"),
			Username:    config.String("data.username"),
			Password:    config.String("data.password"),
			DbName:      config.String("data.dbname"),
			TablePrefix: config.String("data.tableprefix"),
		},
	}
	return dbSettings
}
