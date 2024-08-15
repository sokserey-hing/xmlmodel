
package req_header

import (
    "strings"
    "testing"
    
    
)

func TestConvertJSONToXML(t *testing.T) {
    jsonInput := `{
        "Version": "1.0",
        "BusinessCode": "BC123",
        "TransactionId": "TX123456",
        "Channel": "Web",
        "PartnerId": "P123",
        "BrandId": "B123",
        "ReqTime": "2023-10-01T12:00:00Z",
        "TimeFormat": {
            "TimeType": "UTC",
            "TimeZoneID": "TZ123"
        },
        "AccessUser": "user123",
        "AccessPassword": "pass123",
        "OperatorId": "OP123",
        "AdditionalProperty": [
            {
                "Code": "AP1",
                "Value": "Value1"
            },
            {
                "Code": "AP2",
                "Value": "Value2"
            }
        ]
    }`

    expectedXML := `<com:ReqHeader>
  <com:Version>1.0</com:Version>
  <com:BusinessCode>BC123</com:BusinessCode>
  <com:TransactionId>TX123456</com:TransactionId>
  <com:Channel>Web</com:Channel>
  <com:PartnerId>P123</com:PartnerId>
  <com:BrandId>B123</com:BrandId>
  <com:ReqTime>2023-10-01T12:00:00Z</com:ReqTime>
  <com:TimeFormat>
    <com:TimeType>UTC</com:TimeType>
    <com:TimeZoneID>TZ123</com:TimeZoneID>
  </com:TimeFormat>
  <com:AccessUser>user123</com:AccessUser>
  <com:AccessPassword>pass123</com:AccessPassword>
  <com:OperatorId>OP123</com:OperatorId>
  <com:AdditionalProperty>
    <com:Code>AP1</com:Code>
    <com:Value>Value1</com:Value>
  </com:AdditionalProperty>
  <com:AdditionalProperty>
    <com:Code>AP2</com:Code>
    <com:Value>Value2</com:Value>
  </com:AdditionalProperty>
</com:ReqHeader>`

    xmlOutput, err := ConvertJSONToXML(jsonInput)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Trim spaces and newlines for comparison
    expectedXML = strings.TrimSpace(expectedXML)
    xmlOutput = strings.TrimSpace(xmlOutput)

    if xmlOutput != expectedXML {
        t.Errorf("\nexpected: \n%s\n\ngot: \n%s", expectedXML, xmlOutput)
    }
}