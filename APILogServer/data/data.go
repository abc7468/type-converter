package data

import (
	"encoding/json"
	"fmt"
	"sixshop/apilog/utils"
	"strings"

	"github.com/clbanning/mxj/v2"
)

type Data struct {
	Body        string `json:"body"`
	Format      string `json:"format"`
	Original    map[string]interface{}
	productInfo Product
}

func (d *Data) MakeData() {
	d.toMap()

}

// func (d *Data) makeProductInfo() error {

// }

func (d *Data) toMap() error {
	var result interface{}

	switch d.Format {
	case "json":
		err := json.Unmarshal([]byte(d.Body), &result)
		if err != nil {
			return err
		}
		d.Original = result.(map[string]interface{})
		fmt.Println(d.Original)

	case "byte":
		byteString := strings.Split(d.Body, " ")
		bytes := make([]byte, len(byteString))
		for i, s := range byteString {
			bytes[i] = uint8(utils.StringToInt(s))
		}
		err := json.Unmarshal(bytes, &result)
		if err != nil {
			return err
		}
		d.Original = result.(map[string]interface{})
		fmt.Println(d.Original)

	case "xml":
		mv, err := mxj.NewMapXml([]byte(d.Body))
		if err != nil {
			return err
		}
		d.Original = mv
		fmt.Println(d.Original)

	default:
		fmt.Println(d.Format)

	}

	return nil
}
