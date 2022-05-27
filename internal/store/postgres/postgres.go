package postgres

import (
	"fmt"
	"os"

	"github.com/MalukiMuthusi/orders-b/internal/logger"
	"github.com/MalukiMuthusi/orders-b/internal/models"
	"github.com/MalukiMuthusi/orders-b/internal/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres database
type Postgres struct {
	Db *gorm.DB
}

// Initialize a connection to the database, and automigrate tables
func New() (Postgres, error) {

	// read database configurations

	dbUser := viper.GetString(utils.DbUser)
	dbPwd := viper.GetString(utils.DbPwd)
	dbName := viper.GetString(utils.DbName)
	dbPort := viper.GetString(utils.DbPort)
	dbHost := viper.GetString(utils.DbHost)

	// construct a URI connection string

	var dbURI string

	if viper.GetBool(utils.DbHostedOnCloud) {
		instanceConnectionName := viper.GetString(utils.DbConnectionName)
		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}

		dbURI = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	} else {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPwd, dbPort, dbName)
	}

	// open a connection to the database

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("Unable to connect to database", err)
	}

	p := Postgres{Db: db}

	// Confirm a successful connection.
	if err := p.Db.AutoMigrate(&models.Order{}); err != nil {
		logger.Log.Fatalf("failed to establish database connection", err)
	}

	logger.Log.Info("successfully established a database connection")

	return p, nil
}
