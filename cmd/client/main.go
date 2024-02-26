package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	soaphandlers "soap-server/pkg/transport/handlers"
)

func main() {
	client := http.Client{}
	resp, err := client.Post("http://127.0.0.1:8080/about", "application/soap+xml", nil)
	if err != nil {
		fmt.Println("post error:", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading:", err)
	}
	// fmt.Printf("the data: %+v\n", string(respBody))
	var envelope soaphandlers.Envelope
	err = xml.Unmarshal(respBody, &envelope)
	if err != nil {
		fmt.Println("unmarshal error", err)
		return
	}
	// err = xml.NewDecoder(resp.Body).Decode(&envelope)
	// if err != nil {
	// 	fmt.Println("decoder error:", err)
	// 	return
	// }

	// d := xml.NewDecoder(resp.Body)

	// for t, _ := d.Token(); t != nil; t, _ = d.Token() {
	// 	switch se := t.(type) {
	// 	case xml.StartElement:
	// 		if se.Name.Local == foodElementName {
	// 			// d.DecodeElement(&food, &se)
	// 			// menu.Food = append(menu.Food, food)
	// 		}
	// 	}
	// }
	fmt.Printf("%+v\n", string(respBody))

	fmt.Printf("%+v\n", envelope)
}
