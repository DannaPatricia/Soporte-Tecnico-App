package handlers

import (
	// Para usar fuciones de conversion
	"strconv"
	// Para el formato json
	"encoding/json"
	"net/http"

	// se importa la variable db de la base de datos
	"support-api/db"
	// se importan los modelos
	"support-api/models"
)

// Función para obtener un listado de las incidencias con un limit (de momento)
func GetUserCase(w http.ResponseWriter, r *http.Request) {
	// Definir la cabecera para obtener resultados en formato json
	w.Header().Set("Content-Type", "application/json")

	// se obtienen los datos por parametro, el rol y el id de usuario
	userRole := r.URL.Query().Get("userRole")
	userIDStr := r.URL.Query().Get("userId")

	// Se crea un slice para almecenar las incidencias, almacenara datos de tipo Case definido en el modelo
	var cases []models.SupportCase
	var err error

	// verificacion del rol, se ejecutaran diferentes querys segun el rol
	if userRole == "support" {
		// Se muestra un listado
		// Se hace referencia al slcie con el "&" ya que está modificando los datos
		err = db.DB.Find(&cases).Error
	} else if userRole == "user" {
		// verificacion del parametro id para realizar la query
		if userIDStr == "" {
			http.Error(w, "Falta el parámetro 'userId'", http.StatusBadRequest)
			return
		}

		userId, convErr := strconv.Atoi(userIDStr)
		if convErr != nil {
			http.Error(w, "userId inválido", http.StatusBadRequest)
			return
		}

		// Usuario solo ve sus incidencias
		err = db.DB.Where("user_id = ?", userId).Find(&cases).Error
	} else {
		http.Error(w, "Rol no reconocido", http.StatusBadRequest)
		return
	}

	// Se comprueba si se devolvió un error
	if err != nil {
		http.Error(w, "Error al consultar las incidencias ", http.StatusInternalServerError)
		return
	}
	// Se convierten los datos en formato json, no se hace referencia con el "&" porque solo está leyendo datos, no modifcando
	err = json.NewEncoder(w).Encode(cases)
	if err != nil {
		http.Error(w, "Error al convertir los datos a JSON ", http.StatusInternalServerError)
		return
	}
}

// Funcion para obtener incidencias filtradas por categoria y rol de usuario
func GetCasesByCategory(w http.ResponseWriter, r *http.Request) {
	// Definir la cabecera como JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtener los parámetros de la URL
	categoryIdStr := r.URL.Query().Get("categoryId")
	userRole := r.URL.Query().Get("userRole")
	userIDStr := r.URL.Query().Get("userId")

	// Validar categoryId
	if categoryIdStr == "" {
		http.Error(w, "Falta el parámetro 'categoryId'", http.StatusBadRequest)
		return
	}
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		http.Error(w, "El categoryId debe ser un número válido", http.StatusBadRequest)
		return
	}

	var cases []models.SupportCase

	// verificacion del rol
	if userRole == "support" {
		// Soporte ve todos los casos de esa categoia
		err = db.DB.Where("category_id = ?", categoryId).Find(&cases).Error
	} else if userRole == "user" {
		// Validar userId
		if userIDStr == "" {
			http.Error(w, "Falta el parámetro 'userId'", http.StatusBadRequest)
			return
		}
		userId, convErr := strconv.Atoi(userIDStr)
		if convErr != nil {
			http.Error(w, "El userId debe ser un número válido", http.StatusBadRequest)
			return
		}

		// Usuario solo ve sus propios casos sobre esa categoria
		err = db.DB.Where("category_id = ? AND user_id = ?", categoryId, userId).Find(&cases).Error
	} else {
		http.Error(w, "Rol no reconocido", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Error al consultar las incidencias", http.StatusInternalServerError)
		return
	}

	// Devolver resultado en formato JSON
	err = json.NewEncoder(w).Encode(cases)
	if err != nil {
		http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
	}
}


// Función para obtner una incidencia por ID
func GetCaseById(w http.ResponseWriter, r *http.Request) {
	// Definir la cabecera como para verla como un json
	w.Header().Set("Content-Type", "application/json")

	// se obtiene el id desde la url
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Falta el parámetro 'id'", http.StatusBadRequest)
		return
	}

	// Se pasa de string de int
	caseId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "El ID debe ser un número valido", http.StatusBadRequest)
		return
	}

	// Crear una variable para almacenar la incidencia
	var report models.SupportCase

	// Obtener la incidencia correspondiente mediante una condicion
	dbResponse := db.DB.Where("id = ?", caseId).First(&report).Error
	if dbResponse != nil {
		http.Error(w, "Error al consultar la incidencia a partir del id", http.StatusInternalServerError)
		return
	}

	// Se convierten los datos en formato json, no se hace referencia con el "&" porque solo está leyendo datos, no modifcando
	dbResponse = json.NewEncoder(w).Encode(report)
	if dbResponse != nil {
		http.Error(w, "Error al convertir los datos a JSON", http.StatusInternalServerError)
		return
	}
}

// Funcion para insertar una nueva incidencia
func CreateCase(w http.ResponseWriter, r *http.Request) {
	// Definir la cabecera como para verla como un json
	w.Header().Set("Content-Type", "application/json")

	// se declara una variable par aalmacenar la nueva incidencia
	var newCase models.SupportCase

	// se decodifica el body recibido del json y el contenido se inserta en la variable anterior
	err := json.NewDecoder(r.Body).Decode(&newCase)

	// Validacion basica
	if newCase.Title == "" || newCase.Description == "" || newCase.UserID == 0 || newCase.CategoryID == 0 {
		http.Error(w, "Faltan campos requeridos", http.StatusBadRequest)
		return
	}

	// Crear la incidencia en la base de datos
	// Se omiten los campos insertados por la base de datos
	result := db.DB.Omit("Status", "CreatedAt").Create(&newCase)
	err = result.Error
	if err != nil {
		http.Error(w, "Error al crear la incidencia", http.StatusInternalServerError)
		return
	}

	// Devuelve la incidencia creada como JSON
	json.NewEncoder(w).Encode(newCase)
}

func SetCaseStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Struct temporal para asegurara el tipo de dato
	var input struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}

	// Pasar el JSON
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Usamos GORM para hacer el update dejo en enlace por que no lo entiendo al completo todavia
	// https://gorm.io/docs/update.html#Update-Single-Column
	if err := db.DB.Model(&models.SupportCase{}).
		Where("id = ?", input.ID).
		Update("status", input.Status).Error; err != nil {
		http.Error(w, "Error al actualizar el estado: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta correcta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Estado actualizado correctamente",
	})
}
