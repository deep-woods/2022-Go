package Users

import "github.com/go-chi/chi"

func UsersRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/getUser", GetHandler)
	r.Post("/updateUser", PostHandler)
	return r
}
