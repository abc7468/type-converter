package data

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJsonToMap(t *testing.T) {
	d := Data{}
	d.Body = `
	{"prod":{"id": "abcd","name": "test name","price": "60000","category": "test"}}
	`
	d.Format = "json"
	d.toMap()

	want := map[string]interface{}{
		"prod": map[string]interface{}{
			"id":       "abcd",
			"name":     "test name",
			"price":    "60000",
			"category": "test",
		},
	}
	if !reflect.DeepEqual(d.Original, want) {
		fmt.Printf("original  : %#v\n", d.Original)
		fmt.Printf("want: %#v\n", want)
		t.Fatal("not DeepEqual")
	}
}
func TestXmlToMap(t *testing.T) {
	d := Data{}
	d.Body = `<product><pId>abcd</pId><pName>Jani</pName><pPrice>60000</pPrice><dDone>Done</dDone></product>`
	d.Format = "xml"
	d.toMap()

	want := map[string]interface{}{"product": map[string]interface{}{
		"pId":    "abcd",
		"pName":  "Jani",
		"pPrice": "60000",
		"dDone":  "Done",
	}}
	if !reflect.DeepEqual(d.Original, want) {
		fmt.Printf("NewMapXml, mv  : %#v\n", d.Original)
		fmt.Printf("NewMapXml, want: %#v\n", want)
		t.Fatal("not DeepEqual")
	}
}

func TestByteToMap(t *testing.T) {
	d := Data{}
	d.Body = `10 9 123 34 112 114 111 100 34 58 123 34 105 100 34 58 32 34 97 98 99 100 34 44 34 110 97 109 101 34 58 32 34 116 101 115 116 32 110 97 109 101 34 44 34 112 114 105 99 101 34 58 32 34 54 48 48 48 48 34 44 34 99 97 116 101 103 111 114 121 34 58 32 34 116 101 115 116 34 125 125 10 9`
	d.Format = "byte"
	d.toMap()

	want := map[string]interface{}{
		"prod": map[string]interface{}{
			"id":       "abcd",
			"name":     "test name",
			"price":    "60000",
			"category": "test",
		},
	}
	if !reflect.DeepEqual(d.Original, want) {
		fmt.Printf("NewMapXml, mv  : %#v\n", d.Original)
		fmt.Printf("NewMapXml, want: %#v\n", want)
		t.Fatal("not DeepEqual")
	}
}
