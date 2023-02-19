package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("hello i'm cool...")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "oops! you've got an error ", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "data: %s", data)
}
