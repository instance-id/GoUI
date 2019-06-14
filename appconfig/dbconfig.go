package appconfig

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

type DbData struct {
	DbSettings *DbSettings
	folderName string
	fileName   string
	path       string
}

// --- Maps dbconfig.yml fields to DbSettings fields -------------------------------------------------------------------
type DbSettings struct {
	Database string `json:"database"`
	Data     struct {
		Address     string `json:"address"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		DbName      string `json:"dbname"`
		TablePrefix string `json:"tableprefix"`
	} `json:"data"`
}

var dbData DbData

// --- Gets called from Services and returns DbSettings to Dependency Injection container ------------------------------
func (d *DbSettings) GetDbConfig() *DbSettings {
	return d.loadDbConfig()
}

// --- Populates the DbSettings struct from dbconfig.yml file and returns the data for use -----------------------------
func (d *DbSettings) loadDbConfig() *DbSettings {
	config.AddDriver(yaml.Driver)
	dbData.folderName = "config"
	dbData.fileName = "dbconfig.yml"
	dbData.path = fmt.Sprintf("./%s/%s", dbData.folderName, dbData.fileName)

	fmt.Printf("Db path: %s \n", dbData.path)

	if _, err := os.Stat(dbData.path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dbData.folderName, 0755)
			if err != nil {
				fmt.Printf("Error at creating new dbconfig %s", err)
			}
			_, err = d.newDbConfig()
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
		Database: config.String("database"),
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

func (d *DbSettings) SetDbConfig() (*DbSettings, error) {
	return d.saveDbConfig()
}

func (d *DbSettings) newDbConfig() (*DbSettings, error) {
	yml := New()

	_ = yml.Set("database", "mysql")
	_ = yml.Set("data", "address", "")
	_ = yml.Set("data", "username", "")
	_ = yml.Set("data", "password", "")
	_ = yml.Set("data", "dbname", "")
	_ = yml.Set("data", "tableprefix", "")

	err := yml.Write(dbData.path)
	if err != nil {
		fmt.Printf("Error writing new dbconfig %s", err)
	}
	return d, err
}

func (d *DbSettings) saveDbConfig() (*DbSettings, error) {
	yml := New()

	_ = yml.Set("database", d.Database)
	_ = yml.Set("data", "address", d.Data.Address)
	_ = yml.Set("data", "username", d.Data.Username)
	_ = yml.Set("data", "password", d.Data.Password)
	_ = yml.Set("data", "dbname", d.Data.DbName)
	_ = yml.Set("data", "tableprefix", d.Data.TablePrefix)

	err := yml.Write(dbData.path)
	if err != nil {
		fmt.Printf("Error writing new dbconfig %s", err)
	}
	return d, err
}
