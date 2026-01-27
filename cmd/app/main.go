package main

import (
	"log"
	"net/http"
	"strings"

	"hitalentTest/internal/db"
	"hitalentTest/internal/handler"
	"hitalentTest/internal/repository"
	"hitalentTest/internal/service"
)

func main() {
	database := db.Connect()

	chatRepo := repository.NewChatRepository(database)
	msgRepo := repository.NewMessageRepository(database)

	chatService := service.NewChatService(chatRepo)
	msgService := service.NewMessageService(msgRepo, chatRepo)

	chatHandler := handler.NewChatHandler(chatService, msgService)
	msgHandler := handler.NewMessageHandler(msgService)

	mux := http.NewServeMux()

	mux.HandleFunc("/chats/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/chats/")

		// POST /chats/
		if path == "" || path == "/" {
			if r.Method == http.MethodPost {
				chatHandler.CreateChat(w, r)
				return
			}
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// POST /chats/{id}/messages/
		if strings.HasSuffix(path, "/messages/") {
			if r.Method == http.MethodPost {
				msgHandler.SendMessage(w, r)
				return
			}
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// GET /chats/{id}
		// DELETE /chats/{id}
		switch r.Method {
		case http.MethodGet:
			chatHandler.GetChat(w, r)
		case http.MethodDelete:
			chatHandler.DeleteChat(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
