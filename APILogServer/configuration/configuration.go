package configuration

var Conf = Configuration{}

type Configuration struct {
	Field FieldList `yaml:"field"`
	Kafka KafkaConf `yaml:"kafka"`
	Mysql DbConf    `yaml:"mysql"`
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
	MaxBufSize       int    `yaml:"maxBufSize"`
}
type DbConf struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
