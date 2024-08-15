package time_format

import (
    "strings"
    "testing"
)

func TestConvertJSONToXML(t *testing.T) {
    TimeFormat := `{
        "TimeType": "1",
        "TimeZoneID": "123"
    }`
    expectedTimeFormatXML := `<com:TimeFormat>
  <com:TimeType>1</com:TimeType>
  <com:TimeZoneID>123</com:TimeZoneID>
</com:TimeFormat>`

    TimeFormatXML, err := ConvertJSONToXML(TimeFormat)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Trim spaces and newlines for comparison
    expectedTimeFormatXML = strings.TrimSpace(expectedTimeFormatXML)
    TimeFormatXML = strings.TrimSpace(TimeFormatXML)

    if TimeFormatXML != expectedTimeFormatXML {
        t.Errorf("\n expected: \n %v, \ngot: \n %v", expectedTimeFormatXML, TimeFormatXML)
    }
}