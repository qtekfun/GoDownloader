package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	// Almacena el estado actual de los valores de las variables de salida
	originalStdout := os.Stdout
	originalStderr := os.Stderr
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	// Redirige stdout y stderr para capturar la salida del programa
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	// Ejecuta la función main
	go main()

	// Lee la salida del programa
	w.Close()
	out, _ := ioutil.ReadAll(r)

	// Restaura stdout y stderr
	os.Stdout = originalStdout
	os.Stderr = originalStderr

	// Verifica si la salida contiene el mensaje de éxito
	expectedOutput := ""
	if string(out) != expectedOutput {
		t.Errorf("Se esperaba '%s' pero se obtuvo '%s'", expectedOutput, string(out))
	}

	// Verifica si el archivo se ha creado correctamente
	if _, err := os.Stat("video.mp4"); os.IsNotExist(err) {
		t.Errorf("El archivo video.mp4 no se ha creado correctamente.")
	} else {
		// Elimina el archivo después de la prueba
		os.Remove("video.mp4")
	}
}
