package snowlookupd

import (
	"log"
	"os"
	"snow/internal/app"
	"time"
)

type Options struct {
	Verbose          bool   `flag:"verbose"`
	LogPrefix        string `flag:"log-prefix"`
	TCPAddress       string `flag:tcp-address`
	HTTPAddress      string `flag:http-address`
	BroadcastAddress string `flag:"broadcast-address"`

	InactiveProducerTimeout time.Duration `flag:"inactive-producer-timeout"`
	TombstoneLifetime       time.Duration `flag:"tombstone-lifetime"`
	Logger                  app.Logger
}

func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return &Options{
		LogPrefix:               "[snowlookupd] ",
		TCPAddress:              "0.0.0.0:1201",
		HTTPAddress:             "0.0.0.0:1202",
		BroadcastAddress:        hostname,
		InactiveProducerTimeout: 300 * time.Second,
		TombstoneLifetime:       45 * time.Second,
	}
}
