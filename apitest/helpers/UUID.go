package helpers

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GenerateAuthCode() string {
	rnd, _ := uuid.NewV4()
	return fmt.Sprintf("%s", rnd)
}
