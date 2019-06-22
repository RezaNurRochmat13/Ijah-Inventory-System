package resources

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func SetupDatabaseConfiguration() *sql.DB {
	viper.SetConfigName("sqlite")
	viper.AddConfigPath("environment/development")

	handleReadingConfigFile := viper.ReadInConfig()

	if handleReadingConfigFile != nil {
		fmt.Printf("Error reading file config %s", handleReadingConfigFile.Error())
	}

	DBDriver := viper.GetString("sqlite.DriverDB")
	DBName := viper.GetString("sqlite.DBName")

	databaseConnectionConfiguration, errorDatabaseConfiguration := sql.Open(DBDriver, DBName)

	if errorDatabaseConfiguration != nil {
		fmt.Printf("Error connecting database %s", errorDatabaseConfiguration.Error())
	}

	databaseConnectionConfiguration.SetMaxOpenConns(10)
	databaseConnectionConfiguration.SetMaxIdleConns(10)

	return databaseConnectionConfiguration

}
