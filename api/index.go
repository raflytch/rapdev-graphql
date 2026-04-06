package handler

import (
	"net/http"

	"rapdev-graphql/internal/app"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var httpHandler http.HandlerFunc

func init() {
	application := app.NewApp()
	httpHandler = adaptor.FiberApp(application)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	httpHandler(w, r)
}
