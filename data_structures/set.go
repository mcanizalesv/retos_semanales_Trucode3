package datastructures

import (
	"errors"
)

type Set struct {
	items map[string]bool
}

//Methods

func (s *Set) Append(item string) error {
	//Crea el map si es que no existe y se está intentando ingresar a él
	if s.items == nil {
		s.items = make(map[string]bool)
	}
	//Evalúa si el elemento ya existe en el mapa
	if _, exists := s.items[item]; exists {
		return errors.New("item already exists")
	}

	s.items[item] = true
	return nil
}

func (s Set) Exists(item string) bool {
	_,exists := s.items[item]
	return exists
}

func (s *Set) Delete(item string)  {
	delete(s.items, item)
}


