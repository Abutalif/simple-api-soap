package soaphandlers

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  Header
	Body    Body
}

type Header struct {
	XMLName xml.Name `xml:"Header"`
}

type Body struct {
	XMLName xml.Name    `xml:"Body"`
	Data    interface{} `xml:"data"`
}
