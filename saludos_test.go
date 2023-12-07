package saludos

import (
	"regexp"
	"testing"
)

func TestSaludoAUnaPersona(t *testing.T) {
	nombre := "Gladys"
	esperado := regexp.MustCompile(`\b` + nombre + `\b`)
	mensaje, err := Saludar(nombre)
	if !esperado.MatchString(mensaje) || err != nil {
		t.Fatalf("Saludar(%s) = %q, se esperaba saludar a %s", nombre, mensaje, nombre)
	}
}

func TestSaludoVacio(t *testing.T) {
	mensaje, err := Saludar("")
	if mensaje != "" || err == nil {
		t.Fatalf(`Saludar("") = %q, %v, se esperaba un error`, mensaje, err)
	}
}
