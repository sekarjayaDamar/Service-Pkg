package utilities

import (
	"strings"
	"time"
)

func GenerateTimeRegister() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetTokenFromBearerHeader(authentication string) string {
	return strings.Replace(authentication, "Bearer ", "", -1)
}
