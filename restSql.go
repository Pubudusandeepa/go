package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Define a global variable for the database connection
var db *gorm.DB

func init() {
	// Replace "user:password@tcp(localhost:3306)/dbname" with your MySQL connection string
	dsn := "root:root@tcp(localhost:3306)/todo_list"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	gorm.Model
	ColumnName string `gorm:"column:column_name"`
}

func migrateSchema() {
	err := db.AutoMigrate(&Data{})
	if err != nil {
		log.Fatal(err)
	}
}

// Handler for GET /api/data
func getData(w http.ResponseWriter, r *http.Request) {
	var data []Data
	result := db.Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Handler for POST /api/data
func createData(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newData := Data{ColumnName: requestData.Data}
	result := db.Create(&newData)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/users", getData)
	http.HandleFunc("/users/add", createData)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8082", nil))
}
