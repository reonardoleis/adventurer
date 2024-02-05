package utils

import (
	"fmt"
	"strings"
)

func SprintfSeparator(
	format string,
	separator string,
	sepCount int,
	values ...interface{},
) string {
	sepText := strings.Repeat(separator, sepCount)

	return fmt.Sprintf(
		sepText+"\n"+format+"\n"+sepText,
		values...,
	)
}
