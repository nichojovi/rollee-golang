package config

import (
	"log"
	"os"

	"gopkg.in/gcfg.v1"
)

type MainConfig struct {
	Server struct {
		Name string
		Port string
	}

	DBConfig struct {
		SlaveDSN      string
		MasterDSN     string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
	}
}

func ReadModuleConfig(cfg interface{}, path string, module string) bool {
	environ := os.Getenv("ENV")
	if environ == "" {
		environ = "development"
	}

	fname := path + "/" + module + "." + environ + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	return err == nil
}

func ReadConfig(cfg interface{}, module string) interface{} {
	ok := ReadModuleConfig(cfg, "files/etc/rollee", module)
	if !ok {
		log.Fatalln("failed to read config for ", module)
	}

	return cfg
}
