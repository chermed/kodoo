package data

import "github.com/chermed/kodoo/internal/config"

func NewAlias(key string) config.Alias {
	return config.Alias{
		Key:   key,
		Value: key,
	}
}
func GetAliasValue(cfg config.Config, key, model string) config.Alias {
	for _, alias := range cfg.Aliases {
		if alias.Key == key {
			if alias.Model == "" || alias.Model == model {
				return alias
			}
		}
	}
	return NewAlias(key)
}
