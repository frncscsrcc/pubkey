package operator

type Config struct {
	KubeConfig            string
	DebugLevel            string
	WebServerPort         string
	Namespace             string
}

func GetDefaultConfig() Config {
	return Config{
		DebugLevel:            "INFO",
		KubeConfig:            "~/.kube/config",
		WebServerPort:         "8080",
		Namespace:             "default",
	}
}
