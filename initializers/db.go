package initializers

// this will be the db connection code
import (
	"fmt"
	"go-mvc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to DB")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Tire{})
	DB.AutoMigrate(&models.User{})
}
