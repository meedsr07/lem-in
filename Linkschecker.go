package main

import (
	"fmt"
	"strings"
)

func LinkChecker(lines []string) bool {
	foundlinks := false
	for _, v := range lines {
		line := strings.TrimSpace(v)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") {
			foundlinks = true
		}
		
		if foundlinks {
			part := strings.Split(line, "-")
			if len(part) != 2 {
				fmt.Println("Error invalid forma")
				return false
			}
			if part[0] == "" || part[1] == "" {
				fmt.Println("error invalid form")
				return false
			}
			
		}
	}
	return  true
}
