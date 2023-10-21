package main

import (
	"bytes"
)

func normalize(phonenum string) string{
	var buf bytes.Buffer
	for _, ch:= range phonenum{
		if ch>='0' && ch <='9'{
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}