package main

import (
	"fmt"
	"net/http"
)

func getCards() {

}

func getCard() {

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// В зависимости от метода HTTP-запроса вызываем соответствующий обработчик
	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)
	case http.MethodPut:
		handlePUT(w, r)
	case http.MethodDelete:
		handleDELETE(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// Обработчик для GET-запросов
func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это GET-запрос! ", r.URL.Path)
	switch r.URL.Path {
	case "/cards":
		fmt.Print("list cards")
	case "/card":
		fmt.Print("one card")
	}
}

// Обработчик для POST-запросов
func handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это POST-запрос!")
}

// Обработчик для PUT-запросов
func handlePUT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это PUT-запрос!")
}

// Обработчик для DELETE-запросов
func handleDELETE(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это DELETE-запрос!")
}
