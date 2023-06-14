package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init() {
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

// Handler for POST /api/data
func createData(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ColumnName string `json:"column_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newData := Data{ColumnName: requestData.ColumnName}
	result := db.Create(&newData)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
}

// Handler for GET /api/data
func getData(w http.ResponseWriter, r *http.Request) {
	var data []Data
	result := db.Find(&data)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Handler for GET /api/data/{id}
func getDataByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data Data
	result := db.First(&data, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Handler for PUT /api/data/{id}
func updateData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData struct {
		ColumnName string `json:"column_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data Data
	result := db.First(&data, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	data.ColumnName = requestData.ColumnName
	result = db.Save(&data)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

// Handler for DELETE /api/data/{id}
func deleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data Data
	result := db.First(&data, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	result = db.Delete(&data)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/data", createData).Methods("POST")
	router.HandleFunc("/api/data", getData).Methods("GET")
	router.HandleFunc("/api/data/{id}", getDataByID).Methods("GET")
	router.HandleFunc("/api/data/{id}", updateData).Methods("PUT")
	router.HandleFunc("/api/data/{id}", deleteData).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
