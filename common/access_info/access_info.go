package access_info

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type AccessInfo struct {
	ObjectIdType string `xml:"com:ObjectIdType" json:"ObjectIdType"`
	ObjectId     string `xml:"com:ObjectId" json:"ObjectId"`
}

type AccessInfoWrapper struct {
	XMLName xml.Name `xml:"com:AccessInfo"`
	AccessInfo
}

func ConvertJSONToXML(jsonInput string) (string, error) {
	var info AccessInfo
	if err := json.Unmarshal([]byte(jsonInput), &info); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	wrapper := AccessInfoWrapper{AccessInfo: info}
	xmlData, err := xml.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling XML: %v", err)
	}

	return string(xmlData), nil
}
