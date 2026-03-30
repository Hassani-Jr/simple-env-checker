package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
)

func main(){
	envName := os.Getenv("ENVIRONMENT_NAME")
	if envName == ""{
		envName = "Unknown"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! %s", envName)
	})

	fmt.Println("listening on port 8080")
	log.Fatalln(http.ListenAndServe(":8080",nil))
}