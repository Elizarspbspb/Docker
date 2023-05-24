package main

import (
    "fmt"
    "net/http" // HTTP клиент-сервер
)

func main() {
    fmt.Println("Hello, World from Docker!")
    http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) { // регистрирует функцию обработчика
        fmt.Fprintf(w, "<h1>Hello from Docker container!</h1>") // форматированный вывод
    })

    http.ListenAndServe(":7080", nil) // Запуск сервера с адресом и нулевым обработчиком
}
