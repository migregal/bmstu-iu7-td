package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Debug bool `yaml:"debug"`
	HTTP  struct {
		Address         string        `yaml:"address"`
		GracefulTimeout time.Duration `yaml:"gracefull_timeout"`
		RedirectHost    string        `yaml:"redirect_host"`
		JWTSecret       string        `yaml:"jwt_secret"`
	} `yaml:"http"`
	UserDB struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Passsword string `yaml:"password"`
		Name      string `yaml:"name"`
	} `yaml:"user_db"`
	FilesDB struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Passsword string `yaml:"password"`
		Name      string `yaml:"name"`
	} `yaml:"files_db"`
	Render struct {
		Styles   map[string]string `yaml:"styles"`
		Wrappers map[string]struct {
			Begin string `yaml:"begin"`
			End   string `yaml:"end"`
		} `yaml:"wrappers"`
	}
}

func New(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file content: %w", err)
	}

	c := Config{}
	if err = yaml.UnmarshalStrict(data, &c); err != nil {
		return Config{}, fmt.Errorf("failed to parse file content: %w", err)
	}

	return c, nil
}
