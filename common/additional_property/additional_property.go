package additional_property

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
)

type AdditionalProperty struct {
    Code  string `xml:"com:Code" json:"Code"`
    Value string `xml:"com:Value" json:"Value"`
}

type AdditionalPropertyWrapper struct {
    XMLName xml.Name `xml:"com:AdditionalProperty"`
    AdditionalProperty
}

func ConvertJSONToXML(jsonInput string) (string, error) {
    var properties []AdditionalProperty
    if err := json.Unmarshal([]byte(jsonInput), &properties); err != nil {
        return "", fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    var xmlOutput string
    for _, property := range properties {
        wrapper := AdditionalPropertyWrapper{AdditionalProperty: property}
        xmlData, err := xml.MarshalIndent(wrapper, "", "  ")
        if err != nil {
            return "", fmt.Errorf("error marshalling XML: %v", err)
        }
        xmlOutput += string(xmlData) + "\n"
    }

    return xmlOutput, nil
}