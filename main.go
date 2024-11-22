package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Функция для обработки главной страницы
func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Обработка всех статических файлов в директории /frontend
	fs := http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend")))
	r.Handle("/frontend/*", fs)

	// Маршрут для главной страницы
	r.Get("/", mainPage)

	// Лог об успешном запуске сервера
	log.Println("Сервер запущен на порту 8080")

	// Лог о директории из которой запущен сервер
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Не удалось получить текущую директорию: %v", err)
	}
	log.Printf("Сервер запущен из директории: %s", dir)

	http.ListenAndServe(":8080", r)
}
