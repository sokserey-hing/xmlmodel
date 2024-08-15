package address

import (
    "testing"
	"fmt"
)

func TestConvertJSONToXML(t *testing.T) {
    jsonInput := `{
        "AddressId": "123",
        "AddressType": "Home",
        "Address1": "123 Main St",
        "Address2": "Apt 4B",
        "Address3": "Building 5",
        "Address4": "",
        "Address5": "",
        "Address6": "",
        "Address7": "",
        "Address8": "",
        "Address9": "",
        "Address10": "",
        "Address11": "",
        "Address12": "",
        "PostCode": "12345"
    }`

    expectedXMLOutput := `<com:Address>
  <com:AddressId>123</com:AddressId>
  <com:AddressType>Home</com:AddressType>
  <com:Address1>123 Main St</com:Address1>
  <com:Address2>Apt 4B</com:Address2>
  <com:Address3>Building 5</com:Address3>
  <com:PostCode>12345</com:PostCode>
</com:Address>`

    xmlOutput, err := ConvertJSONToXML(jsonInput)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
	fmt.Println(xmlOutput)

    if xmlOutput != expectedXMLOutput {
        t.Errorf("expected %v, got %v", expectedXMLOutput, xmlOutput)
    }
}