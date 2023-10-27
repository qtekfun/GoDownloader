package main

import (
	"io"
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

	// Guarda el valor original de os.Args para restaurarlo más tarde
	originalArgs := os.Args

	// Modifica temporalmente os.Args para simular argumentos de línea de comandos
	os.Args = []string{"main", "-id=QlZNGcVfeF0&pp"}

	// Restaura os.Args al final de la prueba
	defer func() {
		os.Args = originalArgs
	}()

	// Ejecuta la función main
	main()

	// Lee la salida del programa
	w.Close()
	out, _ := io.ReadAll(r)

	// Restaura stdout y stderr
	os.Stdout = originalStdout
	os.Stderr = originalStderr

	// Verifica si la salida contiene el mensaje de éxito
	expectedOutput := ""
	if string(out) != expectedOutput {
		t.Errorf("Se esperaba '%s' pero se obtuvo '%s'", expectedOutput, string(out))
	}

	// Verifica si el archivo se ha creado correctamente
	if _, err := os.Stat("audio.mp4"); os.IsNotExist(err) {
		t.Errorf("El archivo audio.mp4 no se ha creado correctamente.")
	} else {
		// Elimina el archivo después de la prueba
		os.Remove("audio.mp4")
	}
}
