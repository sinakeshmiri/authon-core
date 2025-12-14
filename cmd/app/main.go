package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sinakeshmiri/imcore/api/controller"
	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/infrastructure/database"
	"github.com/sinakeshmiri/imcore/repository"
	"github.com/sinakeshmiri/imcore/usecase"
)

func main() {
	postgres, err := database.OpenPostgres("")
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(postgres)
	useCase := usecase.NewCreateUserUsecase(userRepository, 3*time.Second)
	handler := controller.NewHandler(useCase)
	// 5. Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// 6. Register oapi handlers
	api.HandlerFromMux(&*handler, r)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
