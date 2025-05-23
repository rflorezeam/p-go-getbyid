package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/rflorezeam/libro-read-one/config"
    "github.com/rflorezeam/libro-read-one/repositories"
    "github.com/rflorezeam/libro-read-one/services"
)

type Handler struct {
    service services.LibroService
}

func NewHandler(service services.LibroService) *Handler {
    return &Handler{
        service: service,
    }
}

// Handler para obtener un libro por ID
func (h *Handler) ObtenerLibroPorID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    id := vars["id"]

    libro, err := h.service.ObtenerLibroPorID(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Libro no encontrado"})
        return
    }

    json.NewEncoder(w).Encode(libro)
}

func main() {
    // Inicializar la base de datos
    config.ConectarDB()

    // Inicializar las capas
    repo := repositories.NewLibroRepository()
    service := services.NewLibroService(repo)
    handler := NewHandler(service)

    // Configurar el router
    router := mux.NewRouter()
    router.HandleFunc("/libros/{id}", handler.ObtenerLibroPorID).Methods("GET")

    // Configurar el puerto
    port := os.Getenv("PORT")
    if port == "" {
        port = "8083"
    }

    fmt.Printf("Servicio de lectura individual de libros corriendo en puerto %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}