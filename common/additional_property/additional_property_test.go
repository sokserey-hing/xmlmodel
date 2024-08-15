package additional_property

import (
    "testing"
    "fmt"
)

func TestConvertJSONToXML(t *testing.T) {
    Address := `
    [
    {
        "Code":"123",
        "Value":"Home"
    },
    {
        "Code":"1236",
        "Value":"Office"
    }
    ]
    `

    expectedAddressXML := `<com:AdditionalProperty>
  <com:Code>123</com:Code>
  <com:Value>Home</com:Value>
</com:AdditionalProperty>
<com:AdditionalProperty>
  <com:Code>1236</com:Code>
  <com:Value>Office</com:Value>
</com:AdditionalProperty>
`

    AddressXML, err := ConvertJSONToXML(Address)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    fmt.Println(AddressXML)

    if AddressXML != expectedAddressXML {
        t.Errorf("expected %v, got %v", expectedAddressXML, AddressXML)
    }
}