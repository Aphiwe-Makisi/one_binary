package config

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`

	Logging struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logging"`

	Scripts struct {
		BackupPath string `mapstructure:"backup_path"`
	} `mapstructure:"scripts"`
}

var C Config
