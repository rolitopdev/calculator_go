# 🧮 Mini API en Go para Operaciones Matemáticas

¡Hola! 👋 Este es un mini proyecto hecho en Go (Golang) que levanta un servidor web básico para hacer operaciones matemáticas simples (suma, resta, multiplicación y división).

## 🚀 ¿Qué hace?

Levanta un servidor local en el puerto `6001` con un endpoint `/operacion`, que recibe un JSON con dos números y un operador... y ¡te devuelve el resultado! 😎

### 🔢 Operaciones que soporta

- Suma (`+`)
- Resta (`-`)
- Multiplicación (`*`)
- División (`/`) *(con validación para que no se divida entre cero)*

## 📦 ¿Cómo se usa?

### 1. Clona este repo

```bash
git clone https://github.com/rolitopdev/calculator_go.git
cd calculator_go
go run main.go
