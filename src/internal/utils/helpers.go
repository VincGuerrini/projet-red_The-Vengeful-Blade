package utils

import "strings"

func NormalizeName(s string) string {
    s = strings.TrimSpace(s)
    if s == "" {
        return "Inconnu"
    }
    lower := strings.ToLower(s)
    s = strings.ToUpper(lower[:1]) + lower[1:]
    return s
}

func ReadLine(prompt string) string

func ReadInt(prompt string, min, max int) (int, error)
