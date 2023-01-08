package repository

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

// Test ConnectDB
func TestConnectDb(t *testing.T) {
	ConnectDb(os.Getenv("POSTGRE_URI"))
}
