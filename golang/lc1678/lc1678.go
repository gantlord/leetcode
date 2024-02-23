package lc1678

import "strings"

func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	return strings.ReplaceAll(command, "(al)", "al")
}
