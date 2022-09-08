package api

import "encoding/xml"

type BookmakerDetails struct {
	XMLName      xml.Name `xml:"bookmaker_details"`
	ResponseCode string   `xml:"response_code,attr"`
	ExpireAt     string   `xml:"expire_at,attr"`
	BookmakerID  string   `xml:"bookmaker_id,attr"`
	VirtualHost  string   `xml:"virtual_host,attr"`
}
