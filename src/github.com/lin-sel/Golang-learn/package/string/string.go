package string

import (
	"strings"
)

// IsContains Return boolean type
func IsContains(str, search string) bool {

	// return true if str conatain search string.
	return strings.Contains(str, search)
}
