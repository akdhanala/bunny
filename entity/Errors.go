package entity

import (
	"fmt"
)

type CommandNotFound struct {
	Command string
}

func (e CommandNotFound) Error() string {
	return fmt.Sprintf("Command (%s) not recognized", e.Command)
}