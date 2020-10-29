package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/completion"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/configuration"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/creation"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/http/rest"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/listing"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/storage/aws"
)

func main() {
	var creator creation.Service
	var completor completion.Service
	var lister listing.Service

	config := configuration.BuildConfigFromFlags()

	s := aws.NewStorage(config.AwsConfig)

	creator = creation.NewService(s)
	completor = completion.NewService(s)
	lister = listing.NewService(s)

	// set up the HTTP server
	router := rest.Handler(creator, lister, completor)

	fmt.Println("The tasks server is running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
