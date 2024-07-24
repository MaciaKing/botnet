package database

import (
	"botnet/models" // Asegúrate de importar el paquete de modelos
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbport := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + dbport + " sslmode=" + sslmode
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection successful.")
}

func Migrate() {
	if err := DB.AutoMigrate(&models.Bot{}); err != nil {
		log.Fatal("Failed to migrate Library model:", err)
	}
	log.Println("Bot model migration successful.")

	createDefaultBoots()
}

func createDefaultBoots() {
	boots := []models.Bot{
		{Id: 1, Ip: "1.0.0.0"},
		{Id: 2, Ip: "2.0.0.0"},
		{Id: 3, Ip: "3.0.0.0"},
		{Id: 4, Ip: "4.0.0.0"},
	}

	for _, bot := range boots {
		if err := DB.Create(&bot).Error; err != nil {
			log.Printf("Error creating book %v: %v\n", bot, err)
		}
	}
	log.Println("Default boots created.")
}
