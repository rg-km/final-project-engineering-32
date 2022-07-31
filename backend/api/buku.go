package api

import (
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
	_ "github.com/rg-km/final-project-engineering-32/backend/repository"
)

type BukuListErrorResponse struct {
	Error string `json:"error"`
}

type Bukujson struct {
	ID_buku      int64  `json:"id_buku"`
	Pengarang    string `json:"pengarang"`
	Judul_buku   string `json:"judul_buku"`
	Tahun_terbit int    `json:"tahun_terbit"`
	No_Buku      string `json:"no_buku"`
	List_Buku    string `json:"list_buku"`
}

type BukuForm struct {
	Pengarang    string `json:"pengarang"`
	Judul_buku   string `json:"judul_buku"`
	Tahun_terbit int    `json:"tahun_terbit"`
	No_Buku      string `json:"no_buku"`
	List_Buku    string `json:"list_buku"`
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
			ID_buku:      book.ID,
			Pengarang:    book.Pengarang,
			Judul_buku:   book.JudulBuku,
			Tahun_terbit: book.TahunTerbit,
			No_Buku:      book.NoBuku,
			List_Buku:    book.ListBuku,
		})
	}

	encoder.Encode(response)
}

func (api *API) addNewBuku(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var requestBuku BukuForm
	encoder := json.NewEncoder(w)
	err := json.NewDecoder(req.Body).Decode(&requestBuku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = api.bukuRepo.AddNewBuku(requestBuku.Pengarang, requestBuku.Judul_buku, requestBuku.Tahun_terbit, requestBuku.No_Buku, requestBuku.List_Buku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BukuListErrorResponse{Error: err.Error()})
		return
	}
	w.Write([]byte("add buku successful"))

}

// func (api *API) DeleteBuku(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(req)["id"]
// if api.bukuRepo.CheckIfProductExists(productId) == false {
// 	w.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(w).Encode("Product Not Found!")
// 	return
// }
// var product entities.Product
// database.Instance.Delete(&product, productId)
// json.NewEncoder(w).Encode("Product Deleted Successfully!")

// }
