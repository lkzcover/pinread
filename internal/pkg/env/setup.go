package env

import (
	"fmt"
	"os"
)

// Setup a struct with environment parameters
var Setup ServiceEnvironment

// ServiceEnvironment list of ENV-variables required for the service to work.
type ServiceEnvironment struct {
	DBConn string

	TelegramToken string
}

func Load() error {
	Setup.DBConn = os.Getenv("DB_CONN")
	if len(Setup.DBConn) == 0 {
		return fmt.Errorf("empty DB_CONN variable")
	}

	Setup.TelegramToken = os.Getenv("TG_TOKEN")
	if len(Setup.DBConn) == 0 {
		return fmt.Errorf("empty TG_TOKEN variable")
	}

	return nil
}
