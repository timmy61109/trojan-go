package mux

import "github.com/Potterli20/trojan-go-fork/config"

type MuxConfig struct {
	Enabled     bool `json:"enabled" yaml:"enabled"`
	IdleTimeout int  `json:"idle_timeout" yaml:"idle-timeout"`
	Concurrency int  `json:"concurrency" yaml:"concurrency"`
}

type Config struct {
	Mux MuxConfig `json:"mux" yaml:"mux"`
}

func init() {
	config.RegisterConfigCreator(Name, func() any {
		return &Config{
			Mux: MuxConfig{
				Enabled:     false,
				IdleTimeout: 30,
				Concurrency: 8,
			},
		}
	})
}
