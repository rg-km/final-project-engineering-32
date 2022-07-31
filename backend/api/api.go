package api

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-32/backend/repository"
)

type API struct {
	usersRepo      repository.UserRepository
	bukuRepo       repository.BukuRepository
	peminjamanRepo repository.PeminjamanRepository
	mux            *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, bukuRepo repository.BukuRepository, peminjamanRepo repository.PeminjamanRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, bukuRepo, peminjamanRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))
	mux.Handle("/api/user/register", api.POST(http.HandlerFunc(api.register)))
	mux.Handle("/api/peminjaman/insert", api.POST(http.HandlerFunc(api.InsertPeminjaman)))

	mux.Handle("/api/buku", http.HandlerFunc(api.bukuList))

	// API with AuthMiddleware:
	mux.Handle("/api/buku/add", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addNewBuku))))
	// mux.Handle("/api/buku/delete", api.DELETE(api.AuthMiddleWare(http.HandlerFunc(api.deleteBuku))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
