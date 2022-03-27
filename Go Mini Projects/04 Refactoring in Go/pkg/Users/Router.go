package Users

import "github.com/go-chi/chi"

func UsersRoutes(s Service) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/getUser", GetHandler)
	router.Post("/updateUser", PostHandler(s))
	return router
}
