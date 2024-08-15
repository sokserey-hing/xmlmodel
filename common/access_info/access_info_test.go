package access_info

import (
    "strings"
    "testing"
)

func TestConvertJSONToXML(t *testing.T) {
    AccessInfo := `{
        "ObjectIdType": "4",
        "ObjectId": "123"
    }`

    expectedAccessInfoXML := `<com:AccessInfo>
  <com:ObjectIdType>4</com:ObjectIdType>
  <com:ObjectId>123</com:ObjectId>
</com:AccessInfo>`

    AccessInfoXML, err := ConvertJSONToXML(AccessInfo)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Trim spaces and newlines for comparison
    expectedAccessInfoXML = strings.TrimSpace(expectedAccessInfoXML)
    AccessInfoXML = strings.TrimSpace(AccessInfoXML)

    if AccessInfoXML != expectedAccessInfoXML {
        t.Errorf("expected %v, got %v", expectedAccessInfoXML, AccessInfoXML)
    }
}