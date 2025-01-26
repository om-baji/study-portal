package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/om-baji/initialisers"
	"github.com/om-baji/utils"
)

func init() {
	initialisers.ConnectDb()
	initialisers.SyncDatabase()
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		response := utils.Response{
			Message: "Health Ok!",
			Code:    200,
			Error:   nil,
		}

		utils.ToJSON(w, 200, response)
	})

	fmt.Println("Server Started on PORT 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Something went wrong!")
	}
}
