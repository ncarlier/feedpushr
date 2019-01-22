package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ncarlier/feedpushr/pkg/strcase"
)

const envPrefix = "APP_"

// FlagEnvString returns flag or env string value with fallback
func FlagEnvString(key, desc, fallback string) *string {
	if val := envValue(key); val != nil {
		fallback = *val
	}
	return flag.String(key, fallback, envDesc(key, desc))
}

// FlagEnvBool returns flag or env bool value with fallback
func FlagEnvBool(key, desc string, fallback bool) *bool {
	if val := envValue(key); val != nil {
		fallback, _ = strconv.ParseBool(*val)
	}
	return flag.Bool(key, fallback, envDesc(key, desc))
}

// FlagEnvInt returns flag or env int value with fallback
func FlagEnvInt(key, desc string, fallback int) *int {
	if val := envValue(key); val != nil {
		fallback, _ = strconv.Atoi(*val)
	}
	return flag.Int(key, fallback, envDesc(key, desc))
}

// FlagEnvDuration returns flag or env duration value with fallback
func FlagEnvDuration(key, desc string, fallback time.Duration) *time.Duration {
	if val := envValue(key); val != nil {
		fallback, _ = time.ParseDuration(*val)
	}
	return flag.Duration(key, fallback, envDesc(key, desc))
}

// FlagEnvArray returns flag or env array value with fallback
func FlagEnvArray(key, desc string, fallback []string) *ArrayFlags {
	result := new(ArrayFlags)
	if val := envValue(key + "s"); val != nil {
		fallback = strings.Split(*val, ",")
	}
	flag.Var(result, key, envDesc(key+"s", desc+" (comma separated list when using env variable)"))
	return result
}

// FlagString returns flag string value with fallback
func FlagString(key, desc, fallback string) *string {
	return flag.String(key, fallback, desc)
}

// FlagBool returns env bool value with fallback
func FlagBool(key, desc string, fallback bool) *bool {
	return flag.Bool(key, fallback, desc)
}

func envDesc(key, desc string) string {
	envKey := strings.ToUpper(strcase.ToSnake(key))
	return fmt.Sprintf("%s (env: %s%s)", desc, envPrefix, envKey)
}

func envValue(key string) *string {
	envKey := strings.ToUpper(strcase.ToSnake(key))
	if value, ok := os.LookupEnv(envPrefix + envKey); ok {
		return &value
	}
	return nil
}
