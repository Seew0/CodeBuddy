package util

import (
	"log"
	"strings"
)

func GetCode(rawdata string) string {
	var extractedCode string
	startIndex := strings.Index(rawdata, "```")
	endIndex := strings.LastIndex(rawdata, "```")
	if startIndex != -1 && endIndex != -1 && startIndex < endIndex {
		extractedCode = rawdata[startIndex+3 : endIndex]
		return "Here is your solution in " + extractedCode
	}
	log.Printf("extractedCode is:    %v", extractedCode)
	return "Code extraction failed."
}
