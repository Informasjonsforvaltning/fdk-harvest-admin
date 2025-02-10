package repository

import "regexp"

func isValidID(id string) bool {
	// Define a regular expression for a valid ID (e.g., alphanumeric)
	var validID = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	return validID.MatchString(id)
}
