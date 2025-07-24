package handlers

import (
	"encoding/json"
	"net/http"

	"support-api/db"
	"support-api/models"
)

// Estructura para recibir el body con el ID
type SupportCaseRequest struct {
	SupportCaseId int `json:"supportCaseId"`
}

// Función para obtener los mensajes de un caso específico
func GetMessagesByCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Solo permitimos POST (como en el login)
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el json recibido
	var req SupportCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Validar el ID recibido conm lo del struc
	if req.SupportCaseId == 0 {
		http.Error(w, "ID de caso de soporte inválido", http.StatusBadRequest)
		return
	}

	// Buscar todos los mensajes con ese SupportCaseId
	var mensajes []models.Message
	if err := db.DB.Where("support_case_id = ?", req.SupportCaseId).Find(&mensajes).Error; err != nil {
		http.Error(w, "Error al obtener los mensajes", http.StatusInternalServerError)
		return
	}

	// Devolver el json
	if err := json.NewEncoder(w).Encode(mensajes); err != nil {
		http.Error(w, "Error al codificar los mensajes", http.StatusInternalServerError)
		return
	}
}

// Funcion para insertar una nueva incidencia
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Solo permitimos POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificamos el mensaje del body
	var newMessage models.Message
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Validaciones básicas
	if newMessage.Content == "" || newMessage.SupportCaseId == 0 || newMessage.UserId == 0 || newMessage.Sender == "" {
		http.Error(w, "Faltan campos requeridos", http.StatusBadRequest)
		return
	}

	// Insertar el mensaje en la base de datos
	if err := db.DB.Create(&newMessage).Error; err != nil {
		http.Error(w, "Error al guardar el mensaje", http.StatusInternalServerError)
		return
	}

	// Devolver el mensaje creado
	if err := json.NewEncoder(w).Encode(newMessage); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		return
	}
}
