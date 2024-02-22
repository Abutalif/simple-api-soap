package soaphandlers

import "encoding/xml"

type AboutRequest struct {
	XMLName xml.Name `xml:"aboutRequest"`
}

type KzhkProjectsRequest struct {
	XMLName xml.Name `xml:"kzhkProjectsRequest"`
}

type AstanaBuildingsRequest struct {
	XMLName xml.Name `xml:"astanaBuildingsRequest"`
}
