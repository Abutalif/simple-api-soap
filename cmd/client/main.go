package main

import (
	"context"
	"log"
	"net/http"
	"soap-server/internal/entities"
	soaphandlers "soap-server/pkg/transport/handlers"

	"github.com/globusdigital/soap"
)

func main() {
	var (
		client *soap.Client
		resp   *http.Response
		err    error
	)
	ctx := context.Background()

	client = soap.NewClient("http://127.0.0.1:8080/about", nil)
	about := entities.About{}
	resp, err = client.Call(ctx, "aboutRequest", &soaphandlers.AboutRequest{}, &about)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(about, resp.Status)

	client = soap.NewClient("http://127.0.0.1:8080/astana_buildings", nil)
	astanaBuildings := entities.AstanaBuilding{}
	resp, err = client.Call(ctx, "astanaBuildingsRequest", &soaphandlers.AstanaBuildingsRequest{}, &astanaBuildings)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(astanaBuildings, resp.Status)

	client = soap.NewClient("http://127.0.0.1:8080/kzhk_projects", nil)
	kzhkProjects := entities.KHCAllowedProject{}
	resp, err = client.Call(ctx, "kzhkProjectsRequest", &soaphandlers.KzhkProjectsRequest{}, &kzhkProjects)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(kzhkProjects, resp.Status)
}
