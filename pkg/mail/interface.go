package mail

type Config struct {
	ServerHostName string `yaml:"server_host_name"`
	ServerPort     string `yaml:"server_port"`
	Account        string `yaml:"account"`
	Password       string `yaml:"password"`
}
