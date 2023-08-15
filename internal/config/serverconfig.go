package config

type ServerConfig struct {
	Host        string `json:"host"`
	Port        int64  `json:"port"`
	SSlCertFile string `json:"ssl_cert_file"`
	SSLKeyFile  string `json:"ssl_key_file"`
}
