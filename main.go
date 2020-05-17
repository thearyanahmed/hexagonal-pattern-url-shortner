package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	_ "strconv"
	"syscall"

	h "github.com/thearyanahmed/url-shortener/api"
	_ "github.com/thearyanahmed/url-shortener/repository/mongo"
	rr "github.com/thearyanahmed/url-shortener/repository/redis"

	"github.com/thearyanahmed/url-shortener/shortener"
)

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		return nil
		// mongoURL := os.Getenv("MONGO_URL")
		// mongodb := os.Getenv("MONGO_DB")
		// mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		// repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// return repo
	}
	return nil
}
