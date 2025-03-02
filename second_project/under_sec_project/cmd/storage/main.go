package main

import (
	"fmt"
	"log"

	"github.com/Dimon228Pokemon/golang_pet/internal/storage"
)

func main() {
	st := storage.NewStrorage()

	file, err := st.Upload("test.txt", []byte("hello mohnatka"))
	if err != nil {
		log.Fatal(err)
	}

	restoredfile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it is restored", restoredfile)
}
