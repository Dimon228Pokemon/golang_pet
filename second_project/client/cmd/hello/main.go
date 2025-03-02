package main

import (
	"fmt"
	"log"

	"github.com/Dimon228Pokemon/golang_pet/second_project/under_sec_project/pkg/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("text.txt", nil)
	if err != nil {
		log.Fatal(err)
	}

	//file,err := st.Upload("text.txt",nil)
	fmt.Println("Client", file)
}
