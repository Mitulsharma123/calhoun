package main

import (
	"regexp"
)

func normalize(phonenum string) string{
	// Compile a regular expression pattern that matches any non-digit character
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phonenum, "")
}


/*func normalize(phonenum string) string{
	var buf bytes.Buffer
	for _, ch:= range phonenum{
		if ch>='0' && ch <='9'{
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}*/