package main

import (
	"log"
	"net/http"
	"support-api/db"

	// Para manejar cada table y organizarlo lo colocamos en handlers para cada una
	"support-api/handlers"
)

// Para las peticiones
func withCors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Encabezados CORS
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Manejar preflight OPTIONS
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Llama al handler original
		handler(w, r)
	}
}

func main() {
	db.Connect()

	// Rutas de usuarios
	http.HandleFunc("/user/login", withCors(handlers.Login))

	// Para mensajes
	http.HandleFunc("/messages-by-case", withCors(handlers.GetMessagesByCase))
	http.HandleFunc("/create-messages", withCors(handlers.CreateMessage))

	// Rutas para las categorias
	http.HandleFunc("/categories", withCors(handlers.GetCategories))
	http.HandleFunc("/category", withCors(handlers.GetCategoryById))

	// Ruta para las incidencias
	// Obtiene las incidencias con un limit 10
	http.HandleFunc("/cases", withCors(handlers.GetUserCase))
	http.HandleFunc("/casesByCategory", withCors(handlers.GetCasesByCategory))
	http.HandleFunc("/case", withCors(handlers.GetCaseById))
	http.HandleFunc("/case/create", withCors(handlers.CreateCase))
	http.HandleFunc("/case/set-status", withCors(handlers.SetCaseStatus))

	// El puerto puede que este ocupado por otra cosa que tengas
	log.Println("Servidor escuchando en http://localhost:8080/categories")
	log.Println("Servidor escuchando en http://localhost:8080/user/login")
	http.ListenAndServe(":8080", nil)
}
