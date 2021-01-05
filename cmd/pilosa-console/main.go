package main

import (
	"flag"
	"log"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

var bind = flag.String("bind", ":800", "bind address")

func main() {
	var port string
	flag.Parse()
	re := regexp.MustCompile(":(\\d{1,5})$")
	match := re.FindStringSubmatch(*bind)
	if len(match) == 2 {
		port = match[1]
	} else {
		log.Fatalf("Invalid bind address: %s", *bind)
	}
	webAddress := "http://localhost:" + port

	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Serving Pilosa Console at " + webAddress)
	err := http.ListenAndServe(*bind, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "running ListenAndServe"))
	}
}
