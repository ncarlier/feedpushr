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

// setFlagEnvString set string value from flag or env with fallback
func setFlagEnvString(p *string, key, desc, fallback string) {
	if val := envValue(key); val != nil {
		fallback = *val
	}
	flag.StringVar(p, key, fallback, envDesc(key, desc))
}

// setFlagEnvBool set bool value from flag or env with fallback
func setFlagEnvBool(p *bool, key, desc string, fallback bool) {
	if val := envValue(key); val != nil {
		fallback, _ = strconv.ParseBool(*val)
	}
	flag.BoolVar(p, key, fallback, envDesc(key, desc))
}

// setFlagEnvInt set int value from flag or env with fallback
func setFlagEnvInt(p *int, key, desc string, fallback int) {
	if val := envValue(key); val != nil {
		fallback, _ = strconv.Atoi(*val)
	}
	flag.IntVar(p, key, fallback, envDesc(key, desc))
}

// setFlagEnvDuration set duration value form flag or env with fallback
func setFlagEnvDuration(p *time.Duration, key, desc string, fallback time.Duration) {
	if val := envValue(key); val != nil {
		fallback, _ = time.ParseDuration(*val)
	}
	flag.DurationVar(p, key, fallback, envDesc(key, desc))
}

// setFlagEnvArray set array value from flag or env with fallback
func setFlagEnvArray(p *ArrayFlags, key, desc string, fallback []string) {
	if val := envValue(key + "s"); val != nil {
		fallback = strings.Split(*val, ",")
	}
	result := &ArrayFlags{
		fallback: fallback,
	}
	flag.Var(result, key, envDesc(key+"s", desc+" (comma separated list when using env variable)"))
}

// setFlagString set string value from flag with fallback
func setFlagString(p *string, key, desc, fallback string) {
	flag.StringVar(p, key, fallback, desc)
}

// setFlagBool set bool value from flag with fallback
func setFlagBool(p *bool, key, desc string, fallback bool) {
	flag.BoolVar(p, key, fallback, desc)
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
