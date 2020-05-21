package epinapi

import (
	"regexp"
)

//var jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
var jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(.*)?json)`)
