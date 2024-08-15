package common

import (
    "encoding/xml"
    "encoding/json"
	"fmt"
)

type Address struct {
    XMLName     xml.Name `xml:"com:Address" json:"-"`
    AddressId   string   `xml:"com:AddressId" json:"AddressId"`
    AddressType string   `xml:"com:AddressType,omitempty" json:"AddressType,omitempty"`
    Address1    string   `xml:"com:Address1,omitempty" json:"Address1,omitempty"`
    Address2    string   `xml:"com:Address2,omitempty" json:"Address2,omitempty"`
    Address3    string   `xml:"com:Address3,omitempty" json:"Address3,omitempty"`
    Address4    string   `xml:"com:Address4,omitempty" json:"Address4,omitempty"`
    Address5    string   `xml:"com:Address5,omitempty" json:"Address5,omitempty"`
    Address6    string   `xml:"com:Address6,omitempty" json:"Address6,omitempty"`
    Address7    string   `xml:"com:Address7,omitempty" json:"Address7,omitempty"`
    Address8    string   `xml:"com:Address8,omitempty" json:"Address8,omitempty"`
    Address9    string   `xml:"com:Address9,omitempty" json:"Address9,omitempty"`
    Address10   string   `xml:"com:Address10,omitempty" json:"Address10,omitempty"`
    Address11   string   `xml:"com:Address11,omitempty" json:"Address11,omitempty"`
    Address12   string   `xml:"com:Address12,omitempty" json:"Address12,omitempty"`
    PostCode    string   `xml:"com:PostCode,omitempty" json:"PostCode,omitempty"`
}

func ConvertJSONToXML(jsonInput string) (string, error) {
    var address Address
    err := json.Unmarshal([]byte(jsonInput), &address)
    if err != nil {
        return "", fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    xmlOutput, err := xml.MarshalIndent(address, "", "  ")
    if err != nil {
        return "", fmt.Errorf("error marshalling XML: %v", err)
    }

    return string(xmlOutput), nil
}