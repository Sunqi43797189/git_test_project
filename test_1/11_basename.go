package main

import (
	"bytes"
	"strings"
)

func baseName(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func baseNameV2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1 :]
	if dot :=strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaV2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteString(string(v))
		if (i+1) == n {
			break
		}
		if (i+1)%3 == 0 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
