package soaphandlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"soap-server/internal/services"
)

type handlers struct {
	srvs services.Services
}

func InitSoapHandler(srvs services.Services) *http.ServeMux {
	handlers := &handlers{
		srvs: srvs,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/about", handlers.aboutHandle)
	mux.HandleFunc("/kzhk_projects", handlers.kzhkHandle)
	mux.HandleFunc("/astana_buildings", handlers.astanaBuildHandle)
	return mux
}

func (h *handlers) aboutHandle(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "not post request", http.StatusBadRequest)
		return
	}

	data := h.srvs.GiveAbout()
	resp := Envelope{
		Body: Body{
			Data: data,
		},
	}

	fmt.Println(data)

	if err := xml.NewEncoder(rw).Encode(resp); err != nil {
		http.Error(rw, "cannot encode xml", http.StatusInternalServerError)
		return
	}
}

func (h *handlers) kzhkHandle(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "not post request", http.StatusBadRequest)
		return
	}

	data, err := h.srvs.GiveKZHKProjects()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "could not get data", http.StatusInternalServerError)
		return

	}
	resp := Envelope{
		Body: Body{
			Data: data,
		},
	}
	fmt.Printf("%+v", resp)
	if err := xml.NewEncoder(rw).Encode(resp); err != nil {
		fmt.Printf("%+v", err)
		http.Error(rw, "cannot encode xml", http.StatusInternalServerError)
		return
	}

}

func (h *handlers) astanaBuildHandle(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "not post request", http.StatusBadRequest)
		return
	}

	data, err := h.srvs.GiveAstanaBuildings()
	if err != nil {
		fmt.Printf("%+v", err)
		http.Error(rw, "could not get data", http.StatusInternalServerError)
		return

	}

	resp := Envelope{
		Body: Body{
			Data: data,
		},
	}
	res, err := xml.Marshal(resp)
	if err != nil {
		http.Error(rw, "cannot encode xml", http.StatusInternalServerError)
		return
	}
	_, err = rw.Write(res)
	if err != nil {
		http.Error(rw, "cannot encode xml", http.StatusInternalServerError)
		return
	}
	// if err := xml.NewEncoder(rw).Encode(resp); err != nil {
	// 	http.Error(rw, "cannot encode xml", http.StatusInternalServerError)
	// 	return
	// }
}
