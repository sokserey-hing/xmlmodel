package req_header

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/sokserey-hing/xmlmodel/common/time_format"
	"github.com/sokserey-hing/xmlmodel/common/additional_property"

)

type ReqHeader struct {
    Version            string              `xml:"com:Version,omitempty" json:"Version,omitempty"`
    BusinessCode       string              `xml:"com:BusinessCode,omitempty" json:"BusinessCode,omitempty"`
    TransactionId      string              `xml:"com:TransactionId" json:"TransactionId"`
    Channel            string              `xml:"com:Channel,omitempty" json:"Channel,omitempty"`
    PartnerId          string              `xml:"com:PartnerId,omitempty" json:"PartnerId,omitempty"`
    BrandId            string              `xml:"com:BrandId,omitempty" json:"BrandId,omitempty"`
    ReqTime            string              `xml:"com:ReqTime" json:"ReqTime"`
    TimeFormat         time_format.TimeFormat `xml:"com:TimeFormat,omitempty" json:"TimeFormat,omitempty"`
    AccessUser         string              `xml:"com:AccessUser" json:"AccessUser"`
    AccessPassword     string              `xml:"com:AccessPassword" json:"AccessPassword"`
    OperatorId         string              `xml:"com:OperatorId,omitempty" json:"OperatorId,omitempty"`
    AdditionalProperty []additional_property.AdditionalProperty `xml:"com:AdditionalProperty,omitempty" json:"AdditionalProperty,omitempty"`
}

type ReqHeaderWrapper struct {
	XMLName xml.Name `xml:"com:ReqHeader"`
	ReqHeader
}

func ConvertJSONToXML(jsonInput string) (string, error) {
	var header ReqHeader
	if err := json.Unmarshal([]byte(jsonInput), &header); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	wrapper := ReqHeaderWrapper{ReqHeader: header}
	xmlData, err := xml.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling XML: %v", err)
	}

	return string(xmlData), nil
}	