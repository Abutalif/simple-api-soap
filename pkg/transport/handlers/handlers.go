package soaphandlers

import (
	"net/http"
	"soap-server/internal/services"

	"github.com/globusdigital/soap"
)

func InitSoapHandler(srvs services.Services) *soap.Server {
	srv := soap.NewServer()
	// srv.R
	srv.RegisterHandler(
		"/about",
		"aboutRequest",
		"aboutRequest",
		func() interface{} {
			return &AboutRequest{}
		},
		func(request interface{}, rw http.ResponseWriter, r *http.Request) (response interface{}, err error) {
			return srvs.GiveAbout(), nil
		},
	)
	srv.RegisterHandler(
		"/kzhk_projects",
		"kzhkProjectsRequest",
		"kzhkProjectsRequest",
		func() interface{} {
			return &KzhkProjectsRequest{}
		},
		func(request interface{}, rw http.ResponseWriter, r *http.Request) (response interface{}, err error) {
			return srvs.GiveKZHKProjects()
		},
	)
	srv.RegisterHandler(
		"/astana_buildings",
		"astanaBuildingsRequest",
		"astanaBuildingsRequest",
		func() interface{} {
			return &AstanaBuildingsRequest{}
		},
		func(request interface{}, rw http.ResponseWriter, r *http.Request) (response interface{}, err error) {
			return srvs.GiveAstanaBuildings()
		},
	)
	return srv
}
