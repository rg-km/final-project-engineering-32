package api

import (
	"encoding/json"
	"net/http"
)

type BukuListErrorResponse struct {
	Error string `json:"error"`
}

type Bukujson struct {
	Judul_buku string `json:"judul_buku"`
}

type BukuListSuccessResponse struct {
	Books []Bukujson `json:"books"`
}

func (api *API) bukuList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := BukuListSuccessResponse{}
	response.Books = make([]Bukujson, 0)

	books, err := api.bukuRepo.GetListBuku()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(BukuListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, book := range books {
		response.Books = append(response.Books, Bukujson{
			Judul_buku: book.JudulBuku,
		})
	}

	encoder.Encode(response)
}