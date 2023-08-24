package registrywrapper

import (
	"errors"
	"fmt"
	"github.com/mindwingx/abstraction"
	registry "github.com/spf13/viper"
	"os"
)

type viper struct {
	*registry.Viper
}

func New() abstraction.Registry {
	return &viper{registry.New()}
}

func (v *viper) InitRegistry(configType string, configFilePath string) (err error) {
	v.SetConfigType(configType)

	config, err := os.Open(configFilePath)

	if err != nil {
		err = errors.New(fmt.Sprintf("failed to load config file: %s", err.Error()))
		return err
	}

	err = v.ReadConfig(config)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to load config file: %s", err.Error()))
		return err
	}
	return nil
}

func (v *viper) ValueOf(key string) abstraction.Registry {
	return &viper{v.Sub(key)}
}

func (v *viper) Parse(item interface{}) (err error) {
	err = v.Unmarshal(&item)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to load config file: %s", err.Error()))
		return err
	}

	return nil
}
