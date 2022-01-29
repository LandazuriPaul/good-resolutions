package utils

import (
	log "github.com/LandazuriPaul/good-resolutions/api/pkg/logger"
	"os"
	"strconv"
)

func mustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.MissingArg(k)
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// MustGetString will return the env as string or panic if it is not present
func MustGetString(k string) string {
	return mustGet(k)
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := mustGet(k)
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.MissingArg(k)
		log.Panic("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}
