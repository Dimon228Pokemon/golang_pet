package main

import (
	"fmt"
	"log"

	"github.com/Dimon228Pokemon/golang_pet/second_project/under_sec_project/internal/storage"
)

func main() {
	st := storage.NewStorage()

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
