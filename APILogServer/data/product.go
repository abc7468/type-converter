package data

// 시나리오에서 유의미하다고 생각되는 Product에 대한 정보를 담기위한 구조체 입니다.
type ProductInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

// 각 포맷별로 같은 역할을 하는 다른 필드명을 conf.yaml에 입력하면 프로젝트 실행 시 메모리에 저장됩니다.
// 그리고 저장된 후보 데이터 값 중 존재하는 필드명을 통해 ProductInfo를 구성합니다.
type ProductFieldList struct {
	Ids    []string `yaml:"ids"`
	Names  []string `yaml:"names"`
	Prices []string `yaml:"prices"`
}
