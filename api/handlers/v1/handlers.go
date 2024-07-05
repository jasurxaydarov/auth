package v1

import "github.com/jasurxaydarov/auth/storage"

type handlers struct{
	storage storage.StorageI
}

type Handler struct{
	Storage storage.StorageI
}


func NewHandler(h Handler)*handlers{

	return &handlers{
		storage: h.Storage,
	}
}