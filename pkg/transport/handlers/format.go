package soaphandlers

import (
	"encoding/xml"
	"soap-server/internal/entities"
)

type Envelope[V entities.Data] struct {
	XMLName xml.Name `xml:"env:Envelope"`
	Header  Header
	Body    Body[V]
}

type Header struct {
	XMLName xml.Name `xml:"env:Header"`
}

type Body[V entities.Data] struct {
	XMLName xml.Name `xml:"env:Body"`
	Data    V
}
