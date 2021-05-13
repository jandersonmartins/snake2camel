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
		if char == SNAKE_CHAR && i > 0 {
			nextI := i + 1
			if nextI >= size {
				break
			}
			nextChar := string(snake[nextI])
			buff.WriteString(strings.ToUpper(nextChar))
			i = i + 2
		} else {
			buff.WriteString(char)
			i++
		}
	}
	return buff.String()
}
