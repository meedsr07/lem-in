package main

import (
	"fmt"
	"strings"
)

func LinkChecker(lines []string) bool {
	foundLink := false

	for _, v := range lines {
		line := strings.TrimSpace(v)

		// تجاهل الأسطر الفارغة و التعليقات
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// نفحص فقط الأسطر التي تحتوي على "-"
		if strings.Contains(line, "-") {
			foundLink = true

			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				fmt.Println("Error: invalid link format")
				return false
			}

			if parts[0] == "" || parts[1] == "" {
				fmt.Println("Error: invalid link format")
				return false
			}
		}
	}

	if !foundLink {
		fmt.Println("Error: no links found")
		return false
	}

	return true
}