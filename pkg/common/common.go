package common

import "os"

func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return defaultVal
	}
}
