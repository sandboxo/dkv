package main

import (
	"github.com/go-akka/configuration"
	"log"
	"os"
	"strings"
)

var DKVConfig *configuration.Config

func LoadConfig(filepath string) {
	DKVConfig = configuration.LoadConfig(filepath)
}

func ValidateConfig() {

	dbListenAddr := DKVConfig.GetString("listen-addr")
	replMasterAddr := DKVConfig.GetString("replication.master-addr")
	statsdAddr := DKVConfig.GetString("statsd-addr")
	dbEngineIni := DKVConfig.GetString("db.engine-ini")

	if dbListenAddr != "" && strings.IndexRune(dbListenAddr, ':') < 0 {
		log.Panicf("given listen address: %s is invalid, must be in host:port format", dbListenAddr)
	}
	if replMasterAddr != "" && strings.IndexRune(replMasterAddr, ':') < 0 {
		log.Panicf("given master address: %s for replication is invalid, must be in host:port format", replMasterAddr)
	}
	if statsdAddr != "" && strings.IndexRune(statsdAddr, ':') < 0 {
		log.Panicf("given StatsD address: %s is invalid, must be in host:port format", statsdAddr)
	}

	if DKVConfig.GetBoolean("db.diskless") && strings.ToLower(DKVConfig.GetString("db.engine")) == "rocksdb" {
		log.Panicf("diskless is available only on Badger storage")
	}

	if strings.ToLower(DKVConfig.GetString(confKeyDBRole)) == slaveRole && replMasterAddr == "" {
		log.Panicf("repl-master-addr must be given in slave mode")
	}

	if dbEngineIni != "" {
		if _, err := os.Stat(dbEngineIni); err != nil && os.IsNotExist(err) {
			log.Panicf("given storage configuration file: %s does not exist", dbEngineIni)
		}
	}
}
//func main2() {
//	conf := configuration.ParseString(configText)
//	fmt.Println("config.one-second:", conf.GetTimeDuration("config.one-second"))
//	fmt.Println("config.one-day:", conf.GetTimeDuration("config.one-day"))
//	fmt.Println("config.array:", conf.GetStringList("config.array"))
//	fmt.Println("config.bar:", conf.GetString("config.bar"))
//	fmt.Println("config.foo:", conf.GetString("config.foo"))
//	fmt.Println("config.number:", conf.GetInt64("config.number"))
//	fmt.Println("config.object.a:", conf.GetString("config.object.a"))
//	fmt.Println("config.object.c.d:", conf.GetString("config.object.c.d"))
//	fmt.Println("config.object.c.f:", conf.GetString("config.object.c.f"))
//	fmt.Println("self-ref:", conf.GetInt64List("self-ref"))
//	fmt.Println("byte-size:", conf.GetByteSize("byte-size"))
//	fmt.Println("home:", conf.GetString("home"))
//	fmt.Println("default:", conf.GetString("none", "default-value"))
//	fmt.Println("plus-equal:", conf.GetString("plus-equal"))
//	fmt.Println("plus-equal-array:", conf.GetStringList("plus-equal-array"))
//}