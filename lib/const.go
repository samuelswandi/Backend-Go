package lib

import (
	"fmt"
)

var (
	SUCCESS = "SUCCESS"
)

func GenericError(error string) string {
	return fmt.Sprintf("[ERROR]: %s", error)
}
