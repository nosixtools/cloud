package config

import "time"

type RegisterConfig struct {
	Address []string
	Timeout time.Duration
}