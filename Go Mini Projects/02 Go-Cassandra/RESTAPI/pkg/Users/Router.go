package Users

import "github.com/go-chi/chi"

func UsersRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/getUser", getHandler)
	r.Post("/updateUser", postHandler)
	return r
}
