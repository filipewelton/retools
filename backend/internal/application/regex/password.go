package regex

import "regexp"

type passwordPatterns struct {
	HasLowercase         *regexp.Regexp
	HasUppercase         *regexp.Regexp
	HasNumbers           *regexp.Regexp
	HasSpecialCharacters *regexp.Regexp
}

var Password = passwordPatterns{
	HasLowercase:         regexp.MustCompile(`.*[a-z].*`),
	HasUppercase:         regexp.MustCompile(`.*[A-Z].*`),
	HasNumbers:           regexp.MustCompile(`.*\d.*`),
	HasSpecialCharacters: regexp.MustCompile(`.*[^a-zA-Z0-9].*`),
}
