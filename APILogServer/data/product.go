package data

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ProductFieldList struct {
	Ids    []string `yaml:"ids"`
	Names  []string `yaml:"names"`
	Prices []string `yaml:"prices"`
}
