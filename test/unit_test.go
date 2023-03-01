package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fajarcandraaa/simple-golang-unit-testing/config/app"
	"github.com/fajarcandraaa/simple-golang-unit-testing/config/routers"
	"github.com/fajarcandraaa/simple-golang-unit-testing/model"
	"github.com/fajarcandraaa/simple-golang-unit-testing/src/user"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = routers.Serve{}

//========================== Set connection for testing ==========================

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	log.Printf("Before calling m.Run() !!!")
	ret := m.Run()
	log.Printf("After calling m.Run() !!!")
	os.Exit(ret)
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}

}

//========================== =========================== ==========================

func refreshTable() error {
	registry := app.SetMigrationTable()
	server.DB.Exec("SET foreign_key_checks=0")
	err := server.DB.Debug().DropTableIfExists(registry...).Error
	if err != nil {
		return err
	}
	server.DB.Exec("SET foreign_key_checks=1")
	err = server.DB.Debug().AutoMigrate(registry...).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	log.Printf("refreshUserTable routine OK !!!")
	return nil
}

//============================ initiate services ============================

func userRepository() user.Service {
	r := model.NewRepository(server.DB)
	s := user.NewService(r)
	return s
}

//============================ ================== ============================
