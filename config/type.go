package config

import "time"

type (
	AppConfig struct {
		Secret string `json:"secret"`
		Debug  bool   `json:"debug" default:"true"`
	}

	ServerConfig struct {
		RunMode         string        `json:"run_mode"`
		ListenAddr      string        `json:"listen_addr"`
		LimitConnection int           `json:"limit_connection"`
		ReadTimeout     time.Duration `json:"read_timeout"`
		WriteTimeout    time.Duration `json:"write_timeout"`
		IdleTimeout     time.Duration `json:"idle_timeout"`
		MaxHeaderBytes  int           `json:"max_header_bytes"`
	}

	LoggerConfig struct {
		Level          string        `json:"level"`
		Formatter      string        `json:"formatter"`
		DisableConsole bool          `json:"disable_console"`
		Write          bool          `json:"write"`
		Path           string        `json:"path"`
		FileName       string        `json:"file_name"`
		MaxAge         time.Duration `json:"max_age"`
		RotationTime   time.Duration `json:"rotation_time"`
		Debug          bool          `json:"debug"`
	}

	Configuration struct {
		App    AppConfig    `json:"app"`
		Server ServerConfig `json:"server"`
		Logger LoggerConfig `json:"logger"`
	}
)
