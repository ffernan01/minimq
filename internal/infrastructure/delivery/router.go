package delivery

import (
	"errors"

	"github.com/ffernan01/minimq/internal/infrastructure/delivery/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	r *chi.Mux
)

// Config configurates routes, middlewares & delivery logic
func Config() error {
	if r != nil {
		return errors.New("Routes already configured")
	}

	r = chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", handler.GetSubscriptions)
		r.Get("/{subscriptionID}", handler.GetSubscription)
		r.Post("/", handler.PostSubscription)
		r.Put("/{subscriptionID}", handler.PutSubscription)
		r.Delete("/{subscriptionID}", handler.DeleteSubscription)
	})

	r.Route("/topics", func(r chi.Router) {
		r.Get("/", handler.GetTopics)
		r.Get("/{topicID}", handler.GetTopic)
		r.Post("/{topicID}/reset", handler.PostTopicReset)

		r.Route("/messages", func(m chi.Router) {
			r.Get("/{messageID}", handler.GetMessage)
			r.Post("/", handler.PostMessage)
		})
	})

	return nil
}
