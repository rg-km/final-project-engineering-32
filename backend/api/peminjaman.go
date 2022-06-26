package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type PeminjamanRequest struct {
	IdBuku         int64  `json:"id_buku"`
	IdUser         int64  `json:"id_user"`
	TanggalPinjam  string `json:"tgl_pinjam"`
	TanggalKembali string `json:"tgl_kembali"`
}

type PeminjamanResponse struct {
	Message string `json:"message"`
}

func (api *API) InsertPeminjaman(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var peminjamanRequest PeminjamanRequest
	err := json.NewDecoder(req.Body).Decode(&peminjamanRequest)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = api.peminjamanRepo.InsertPeminjaman(peminjamanRequest.IdBuku, peminjamanRequest.IdUser, peminjamanRequest.TanggalPinjam, peminjamanRequest.TanggalKembali)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(PeminjamanResponse{
		Message: "berhasil ditambahkan",
	})
}
