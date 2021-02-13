package server

import "web-stash-api/config"

func Init() {
	r := NewRouter()
	r.Run(getBindAddress())
}

func getBindAddress() string {
	if config.Cfg.Server.BindAddress == "" {
		return "localhost:10000"
	}
	return config.Cfg.Server.BindAddress
}
