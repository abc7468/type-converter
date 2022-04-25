package data

import (
	"encoding/json"
	"sixshop/apilog/configuration"
	"sixshop/apilog/utils"
	"strings"
	"time"

	"github.com/clbanning/mxj/v2"
)

// 프로젝트에서 사용할 데이터 필드입니다.
type Data struct {
	Body        string `json:"body"`
	Format      string `json:"format"`
	Original    map[string]interface{}
	Json        string
	ProductInfo ProductInfo
	Time        time.Time
}

// Data 구조체를 초기화 하는 함수입니다.
func (d *Data) MakeData() error {
	err := d.toMap()
	if err != nil {
		return err
	}
	err = d.makeProductInfo()
	if err != nil {
		return err
	}
	d.Time = time.Now()
	return nil
}

//
func (d *Data) makeProductInfo() error {
	var product string
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

	case "xml":
		mv, err := mxj.NewMapXml([]byte(d.Body))
		if err != nil {
			return err
		}
		d.Original = mv

	}
	jsonData, _ := json.Marshal(d.Original)
	d.Json = string(jsonData)
	return nil
}
