package cfn

import (
	"strings"
)

// StripShebang strips the first line of input and returns the rest, an empty string will be returned if there is only
// one line.
func StripShebang(input string) string {
	all := strings.SplitN(input, "\n", 2)

	// Only splits the first newline
	if len(all) != 2 {
		return ""
	}

	return all[1]
}
