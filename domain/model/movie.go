package model

import (
	"github.com/jalozanot/demoCeiba/domain/validators"
)

const (
	StatusActive = "active"
)

type Movie struct {
	Id           int64  `json:"Id"`
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	CodigoBarras string `json:"CodigoBarras"`
}

type Creature struct {
	Name     string
	Greeting string
}


func (movie *Movie) CreateMovil(nombre string, categoria string, codigoBarras string) (Movie, error) {

	if err := validators.ValidateRequired(nombre, "Nombre should have some value"); err != nil {
		return Movie{}, err
	}

	if err := validators.ValidateRequired(categoria, "Categoria should have some value"); err != nil {
		return Movie{}, err
	}

	if err := validators.ValidateRequired(categoria, "codigoBarras should have some value"); err != nil {
		return Movie{}, err
	}

	return Movie{
		Nombre:       nombre,
		Categoria:    categoria,
		CodigoBarras: codigoBarras,
	}, nil
}
