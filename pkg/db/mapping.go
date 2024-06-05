package db

import (
	"github.com/faith0831/easygen/pkg/config"
)

// GetMappingType GetMappingType
func GetMappingType(providerName string, lang string, typ string, isNull bool) string {
	conf, err := config.LoadMappingConfig(lang, providerName)
	if err != nil {
		return "unknown"
	}

	if item, ok := conf[typ]; ok {
		if !isNull {
			return item.Underlying
		} else {
			return item.Nullable
		}
	}

	return "unknown"
}
