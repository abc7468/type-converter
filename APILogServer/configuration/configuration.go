package configuration

var Conf = Configuration{}

type Configuration struct {
	Field FieldList `yaml:"field"`
	Kafka KafkaConf `yaml:"kafka"`
}

type FieldList struct {
	Product []string `yaml:"product"`
	Ids     []string `yaml:"ids"`
	Names   []string `yaml:"names"`
	Prices  []string `yaml:"prices"`
}

type KafkaConf struct {
	BootstrapServers string `yaml:"bootStrapServers"`
	GroupId          string `yaml:"groupId"`
	AutoOffsetReset  string `yaml:"autoOffsetReset"`
}
