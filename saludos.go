package saludos

import (
	"errors"
	"fmt"
	"math/rand"
)

func Saludar(nombre string) (string, error) {
	if nombre == "" {
		return "", errors.New("El nombre no puede estar vacío")
	}
	return fmt.Sprintf(formatoAleatorio(), nombre), nil
}

func SaludosATodos(nombres []string) (map[string]string, error) {
	mensajes := make(map[string]string)
	for _, nombre := range nombres {
		mensaje, err := Saludar(nombre)
		if err != nil {
			return nil, err
		}
		mensajes[nombre] = mensaje
	}
	return mensajes, nil
}

func formatoAleatorio() string {
	formatos := []string{
		"Hola %s",
		"¿Qué tal %s?",
		"¡Buenos días %s!",
	}
	return formatos[rand.Intn(len(formatos))]
}
