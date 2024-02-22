package main

import (
	"context"
	"log"
	"soap-server/internal/entities"
	soaphandlers "soap-server/pkg/transport/handlers"

	"github.com/globusdigital/soap"
)

func main() {
	ctx := context.Background()
	client := soap.NewClient("http://127.0.0.1:8080/", nil)

	about := entities.About{}
	httpResp, err := client.Call(ctx, "about", &soaphandlers.AboutRequest{}, &about)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(about, httpResp.Status)

	astanaBuildings := entities.AstanaBuilding{}
	httpResp, err = client.Call(ctx, "astanaBuildingsRequest", &soaphandlers.AboutRequest{}, &astanaBuildings)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(astanaBuildings, httpResp.Status)

	kzhkProjects := entities.KHCAllowedProject{}
	httpResp, err = client.Call(ctx, "kzhkProjectsRequest", &soaphandlers.AboutRequest{}, &kzhkProjects)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(kzhkProjects, httpResp.Status)
}
