package handlers

import (
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.logger.Println("POST method called")
		return
	}

	// handle an update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	listofProducts := data.GetProducts()
	err := listofProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to mashal json", http.StatusInternalServerError)
	}
	// rw.Write(listofProducts)
}
