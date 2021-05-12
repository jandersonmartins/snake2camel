package s2c

import (
	"bytes"
	"strings"
)

const SNAKE_CHAR = "_"

func ConvertToCamel(snake string) string {
	buff := bytes.Buffer{}
	size := len(snake)
	i := 0
	for {
		if i >= size {
			break
		}
		char := string(snake[i])
		if char == SNAKE_CHAR {
			char = string(snake[i+1])
			buff.WriteString(strings.ToUpper(char))
			i = i + 2
		} else {
			buff.WriteString(char)
			i++
		}
	}
	return buff.String()
}
