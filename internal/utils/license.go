package utils

import "regexp"

const uuidPattern = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

// IsUUID checks if the given string is a valid UUID. All license keys should be UUIDs.
func IsUUID(license string) bool {
	return regexp.MustCompile(uuidPattern).MatchString(license)
}
