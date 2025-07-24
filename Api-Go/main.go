package main

import (
	"log"
	"net/http"
	"os"
	"support-api/db"

	//Para variables de entorno para el Railway
	// Para manejar cada table y organizarlo lo colocamos en handlers para cada una
	"support-api/handlers"
)

// Para las peticiones
func withCors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedOrigins := map[string]bool{
			"http://localhost:5173":           true,
			"https://dannapatricia.github.io": true,
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Encabezados CORS
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

	// Railway se encargara de asginar el puerto adecuado para el despliegue, en local pillara el 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Por defecto en local
	}

	log.Printf("Servidor escuchando en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
