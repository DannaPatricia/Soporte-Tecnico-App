package handlers

import (
	// Para usar fuciones de conversion
	"strconv"
	// Para el formato json
	"encoding/json"
	"net/http"

	// La variable de la db
	"support-api/db"

	// usamos los objeto/estructuras de models
	"support-api/models"
)

// Función para obtener todos los usuarios
func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Definir la cabecera como para verla como un json
	w.Header().Set("Content-Type", "application/json")

	// Crear un slice para almacenar las categoriass
	var categories []models.Category

	// Consultar todos los usuarios usando la conexion global db.DB y la tabla se lo sabe gorm, como ? ni idea, lo hace automáticamente usando reglas de naming convention.
	dbResponse := db.DB.Find(&categories).Error
	if dbResponse != nil {
		http.Error(w, "Error al consultar los categorias", http.StatusInternalServerError)
		return
	}

	// Convertir los datos de los usuarios a formato JSON
	dbResponse = json.NewEncoder(w).Encode(categories)
	if dbResponse != nil {
		http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
		return
	}
}

// Función para obtner una categoria por ID
func GetCategoryById(w http.ResponseWriter, r *http.Request){
	// Definir la cabecera como para verla como un json
	w.Header().Set("Content-Type", "application/json")

	// se obtiene el id desde la url
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el parámetro 'id'", http.StatusBadRequest)
		return
	}

	// Se pasa de string de int
	categoryId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "El ID debe ser un número valido", http.StatusBadRequest)
		return
	}

	// Crear una variable para almacenar la categoria
	var category models.Category

	// Obtener la cateoria correspondiente mediante una condicion
	dbResponse := db.DB.Where("id = ?", categoryId).First(&category).Error
	if dbResponse != nil {
		http.Error(w, "Error al consultar la categoria a partir del id", http.StatusInternalServerError)
		return
	}

	// Convertir los datos de los usuarios a formato JSON
	dbResponse = json.NewEncoder(w).Encode(category)
	if dbResponse != nil {
		http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
		return
	}
}