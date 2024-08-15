package time_format

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
)

type TimeFormat struct {
    TimeType   string `xml:"com:TimeType" json:"TimeType"`
    TimeZoneID string `xml:"com:TimeZoneID" json:"TimeZoneID"`
}

type TimeFormatWrapper struct {
    XMLName xml.Name `xml:"com:TimeFormat"`
    TimeFormat
}

func ConvertJSONToXML(jsonInput string) (string, error) {
    var format TimeFormat
    if err := json.Unmarshal([]byte(jsonInput), &format); err != nil {
        return "", fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    wrapper := TimeFormatWrapper{TimeFormat: format}
    xmlData, err := xml.MarshalIndent(wrapper, "", "  ")
    if err != nil {
        return "", fmt.Errorf("error marshalling XML: %v", err)
    }
    fmt.Println(string(xmlData))
    return string(xmlData), nil
}