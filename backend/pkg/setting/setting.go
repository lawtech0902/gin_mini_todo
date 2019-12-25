package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string `json:"name"`
}

// Init init config
func Init(cfg string) error {
	c := Config{Name: cfg}
	if err := c.initConfig(); err != nil {
		return err
	}
	
	c.watchConfig()
	return nil
}

// initConfig init config helper
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("conf")
	}
	
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("api")
	
	replacer := strings.NewReplacer(",", "_")
	viper.SetEnvKeyReplacer(replacer)
	
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	
	return nil
}

// watchConfig monitor the config file change and hot reload
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed: %s", in.Name)
	})
}


