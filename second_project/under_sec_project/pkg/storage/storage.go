package storage

import "github.com/Dimon228Pokemon/golang_pet.git/internal/storage"

//тут чтобы вытягивать приватные приколюхи наружу как я понял
//тобишь все коды клиентов, чтобы в дальнейшем подтягивать на сервер все проходит через этот файлик pkg
//как бы выдача подносов с компотиком чтобы сервер забрал и полакомился

func NewStrorage() *storage.Storage {
	return storage.NewStrorage()
}
