package main

import (
	"encoding/json" // Para trabajar con JSON (entrada y salida)
	"fmt"           // Para imprimir cosas en consola
	"log"           // Para manejar errores del servidor
	"net/http"      // Para levantar el servidor HTTP
)

// Estructura que representa una operación matemática
type Operacion struct {
	Numero1  float64 `json:"numero1"`  // Primer número
	Numero2  float64 `json:"numero2"`  // Segundo número
	Operador string  `json:"operador"` // Operación: "+", "-", "*", "/"
}

// Estructura para la respuesta
type Respuesta struct {
	Resultado float64 `json:"resultado"`
	Mensaje   string  `json:"mensaje"`
}

// Función que hace la magia: recibe la operación y la resuelve
func calcular(op Operacion) (float64, string) {
	switch op.Operador {
	case "+":
		return op.Numero1 + op.Numero2, "Suma realizada correctamente"
	case "-":
		return op.Numero1 - op.Numero2, "Resta realizada correctamente"
	case "*":
		return op.Numero1 * op.Numero2, "Multiplicación realizada correctamente"
	case "/":
		// Validamos que no se divida entre cero
		if op.Numero2 == 0 {
			return 0, "Error: no se puede dividir entre cero"
		}
		return op.Numero1 / op.Numero2, "División realizada correctamente"
	default:
		return 0, "Operador no válido. Usa: +, -, * o /"
	}
}

// Esta función maneja las peticiones que lleguen al endpoint /operacion
func manejarOperacion(w http.ResponseWriter, r *http.Request) {
	// Solo aceptamos POST, porque vamos a recibir datos
	if r.Method != "POST" {
		http.Error(w, "Método no permitido. Usa POST", http.StatusMethodNotAllowed)
		return
	}

	var op Operacion

	// Leemos el cuerpo del request y lo convertimos a nuestra estructura
	err := json.NewDecoder(r.Body).Decode(&op)
	if err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	// Llamamos a la función calcular
	resultado, mensaje := calcular(op)

	// Creamos una estructura de respuesta
	respuesta := Respuesta{
		Resultado: resultado,
		Mensaje:   mensaje,
	}

	// Enviamos la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

// Punto de entrada del programa
func main() {
	fmt.Println("Servidor iniciado en http://localhost:6001")

	// Registramos la ruta /operar con su función manejadora
	http.HandleFunc("/operacion", manejarOperacion)

	// Iniciamos el servidor en el puerto 6001
	err := http.ListenAndServe(":6001", nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
