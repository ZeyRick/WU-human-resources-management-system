package validate

import "regexp"

func IsValidGoogleCoordinate(s string) bool {
	// Regular expression for Google coordinates (latitude, longitude)
	regex := `^[-+]?\d+(?:\.\d+)?,\s*[-+]?\d+(?:\.\d+)?$`

	// Match the string against the regular expression
	return regexp.MustCompile(regex).MatchString(s)
}
