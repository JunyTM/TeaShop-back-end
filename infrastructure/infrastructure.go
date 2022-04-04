package infrastructure

import (
	"flag"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	APPPORT    = ""
	DBHOST     = ""
	DBPORT     = ""
	DBUSER     = ""
	DBPASSWORD = ""
	DBNAME     = ""
	DBURL = ""
	HTTPSWAGGER = ""
	ROOTPATH    = ""
)

var (
	appPort    string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbURL 	string

	httpSwagger string
	rootPath    string
	storagePath string

	db *gorm.DB
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	}
	return defaultValue
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func loadEnvParameters() {
	root, _ := os.Getwd()
	appPort = getStringEnvParameter(APPPORT, goDotEnvVariable(("APPPORT")))
	dbHost = getStringEnvParameter(DBHOST, goDotEnvVariable(("DBHOST")))
	dbPort = getStringEnvParameter(DBPORT, goDotEnvVariable("DBPORT"))
	dbUser = getStringEnvParameter(DBUSER, goDotEnvVariable("DBUSER"))
	dbPassword = getStringEnvParameter(DBPASSWORD, goDotEnvVariable("DBPASSWORD"))
	dbName = getStringEnvParameter(DBNAME, goDotEnvVariable("DBNAME"))
	dbURL = getStringEnvParameter(DBURL, goDotEnvVariable("DATABASE_URL"))
	httpSwagger = getStringEnvParameter(HTTPSWAGGER, goDotEnvVariable("HTTPSWAGGER"))
	rootPath = getStringEnvParameter(ROOTPATH, root)
	storagePath = rootPath + "/storage/"

}

func init() {
	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")
	flag.Parse()

	loadEnvParameters()

	if err := InitDatabase(initDB); err != nil {
		log.Println(err)
	}
}


// GetDB get database instance
func GetDB() *gorm.DB {
	return db
}

// GetDBName get database name
func GetDBName() string {
	return dbName
}

// GetDBURL get database URL
func GetDBURL() string {
	return dbURL
}

// GetAppPort export app port
func GetAppPort() string {
	return appPort
}

// GetHTTPSwagger export link swagger
func GetHTTPSwagger() string {
	return httpSwagger
}

// GetRootPath export root path system
func GetRootPath() string {
	return rootPath
}

func GetStoragePath() string {
	return storagePath
}