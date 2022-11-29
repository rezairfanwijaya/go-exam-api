package connection_test

import (
	"testing"

	"github.com/rezairfanwijaya/go-exam-api.git/connection"
)

func TestDB(t *testing.T) {
	connection.DB("../.env")
}
