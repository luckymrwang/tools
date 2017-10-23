package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func UniToCn(textUnquoted string) (string, error) {
	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			return "", err
		}
		context += fmt.Sprintf("%c", temp)
	}

	return context, nil
}

func CnToUni(sText string) (string, error) {
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted, nil
}
