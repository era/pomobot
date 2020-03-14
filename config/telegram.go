package config
import (
	"os"
)
func GetToken() string {
	return os.Getenv("TELEGRAM_TOKEN")
}
