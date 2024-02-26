package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"soap-server/internal/entities"
	soaphandlers "soap-server/pkg/transport/handlers"
)

func main() {
	var buffer []byte
	client := http.Client{}
	resp, err := client.Post("http://127.0.0.1:8080/about", "application/soap+xml", bytes.NewBuffer(buffer))
	if err != nil {
		fmt.Println("post error:", err)
		return
	}
	defer resp.Body.Close()
	var envelope soaphandlers.Envelope[entities.About]
	err = xml.NewDecoder(resp.Body).Decode(&envelope)
	if err != nil {
		fmt.Println("decoder error:", err)
		return
	}
	fmt.Printf("%+v\n", envelope)

}
