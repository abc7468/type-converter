package data

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Data struct {
	Body        string `json:"body"`
	Format      string `json:"format"`
	Original    map[string]interface{}
	productInfo Product
}

func (d *Data) MakeProductInfo() {
	d.makeResult()

}

func (d *Data) makeResult() error {
	var result interface{}
	switch d.Format {
	case "json":
		err := json.Unmarshal([]byte(d.Body), &result)
		if err != nil {
			return err
		}
		d.Original = result.(map[string]interface{})
		fmt.Println(d.Original["id"])
	case "byte":

	case "xml":
		err := xml.Unmarshal([]byte(d.Body), &result)
		if err != nil {
			return err
		}
	}

	return nil
}
