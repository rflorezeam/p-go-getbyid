package services

import (
	"github.com/rflorezeam/libro-read-one/models"
	"github.com/rflorezeam/libro-read-one/repositories"
)

type LibroService interface {
	ObtenerLibroPorID(id string) (*models.Libro, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) ObtenerLibroPorID(id string) (*models.Libro, error) {
	return s.repo.ObtenerLibroPorID(id)
} 