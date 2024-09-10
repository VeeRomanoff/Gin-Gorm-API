package gin_gorm_api

import (
	"github.com/VeeRomanoff/gin_gorm_api/models"
	"github.com/VeeRomanoff/gin_gorm_api/routers"
	"github.com/VeeRomanoff/gin_gorm_api/storage"
	"github.com/jinzhu/gorm"
	"log"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "dsa") // it automatically validates the connection AND PINGS DB
	if err != nil {
		log.Println("error connecting to the database", err)
	}

	defer storage.DB.Close()                  // let's not forget to close connection.
	storage.DB.AutoMigrate(&models.Article{}) // at this point gorm AUTOMATICALLY generates all queries, migrations and applies them

	// Configuring router
	r := routers.SetupRouter()

	// r - gin router
	r.Run()
}
