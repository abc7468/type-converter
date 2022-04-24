package data

import (
	"encoding/json"
	"fmt"
	"sixshop/apilog/configuration"
	"sixshop/apilog/utils"
	"strings"

	"github.com/clbanning/mxj/v2"
)

type Data struct {
	Body        string `json:"body"`
	Format      string `json:"format"`
	Original    map[string]interface{}
	ProductInfo ProductInfo
}

func (d *Data) MakeData() error {
	err := d.toMap()
	if err != nil {
		return err
	}
	err = d.makeProductInfo()
	if err != nil {
		return err
	}
	return nil
}

func (d *Data) makeProductInfo() error {
	var product string
	// var id string
	// var name string
	// var price string
	for _, p := range configuration.Conf.Field.Product {
		if _, ok := d.Original[p]; ok {
			product = p
			break
		}
	}
	tmp := d.Original[product].(map[string]interface{})
	if tmp == nil {
		return nil
	}
	for _, i := range configuration.Conf.Field.Ids {
		if _, ok := tmp[i]; ok {
			d.ProductInfo.Id = tmp[i].(string)
			break
		}
	}
	for _, n := range configuration.Conf.Field.Names {
		if _, ok := tmp[n]; ok {
			d.ProductInfo.Name = tmp[n].(string)
			break
		}
	}
	for _, p := range configuration.Conf.Field.Prices {
		if _, ok := tmp[p]; ok {
			d.ProductInfo.Price = tmp[p].(string)
			break
		}
	}
	fmt.Println(d.ProductInfo)
	return nil
}

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
