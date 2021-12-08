package g

type GlobalConfig struct {
	Debug         bool          	    `json:"debug"`
	LogLevel   		string				`json:"loglevel"`
	LogFile 		string				`json:"logfile"`
	EtcdConfig		*EtcdConfig			`json:"etcdconfig"`
	EtcdSSL			*EtcdSSL			`json:"etcdssl"`
	Http			*HttpConfig			`json:"http"`
	TestLine		int					`json:"testline"`
	ROLE			string				`json:"role"`
}

type EtcdConfig struct {
	Host           []string  `json:"host"`
}

type EtcdSSL struct {
	CaFile           string          `json:"cafile"`
	CertFile         string          `json:"certfile"`
	CertKeyFile      string          `json:"keyfile"`
}

type HttpConfig struct {
	Enabled			bool		`json:"enabled"`
	Listen			string		`json:"listen"`
}

