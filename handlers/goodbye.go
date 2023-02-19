package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func SayGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (bye *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	bye.logger.Println("goodbye motherfucker")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "fix the error asshole", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "pawan is saying %s", data)
}
