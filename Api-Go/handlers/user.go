package handlers

import (
	"encoding/json"
	"net/http"
	"support-api/db"
	"support-api/models"
)

// Para validar login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Validacion de metodo post
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Validacion los datos que quiero y su tipo
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Buscar el usuario por su nombre de usuario
	var user models.User
	// por si ayuda de algo como usar estas cosas
	// El where https://gorm.io/docs/query.html#Conditions
	// El first https://gorm.io/docs/query.html#Retrieving-a-single-object
	result := db.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Comparar contraseñas (sin encriptar aqui ya veremos como)
	if user.Password != req.Password {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	// Si esta bien enviar respuesta
	// no se como va todavia asi que dejo esto pa otro dia https://pkg.go.dev/encoding/json#Encoder.Encode
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login correcto",
		// regresar los datos para pinia
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
