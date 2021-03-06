package main

import (
	"fmt"
	"github.com/andersondborges/locaweb/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Startint Project")
	fmt.Println("Server Running Port:9090")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/most_relevants/", controllers.ControllerMostRelevants)
	http.HandleFunc("/most_mentions/", controllers.ControllerMostMentions)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
