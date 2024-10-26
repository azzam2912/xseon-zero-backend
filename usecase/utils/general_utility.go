package utils

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func CleanFilePath(cleanPath string) (string, error) {
	cleanPath = strings.ReplaceAll(cleanPath, "'", " ")
	cleanPath = strings.TrimPrefix(cleanPath, "file://")
	// Replace %20 with spaces
	cleanPath = strings.ReplaceAll(cleanPath, "%20", " ")
	cleanPath = strings.TrimPrefix(cleanPath, " ")

	// Decode other URL-encoded characters
	decodedPath, err := url.QueryUnescape(cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to decode path: %w", err)
	}

	// Ensure the path starts with a forward slash if it's not already absolute
	if !filepath.IsAbs(decodedPath) {
		decodedPath = "/" + decodedPath
	}
	// Clean the path (removes . and .. elements)
	cleanedPath := filepath.Clean(decodedPath)

	// On Windows, convert forward slashes to backslashes
	if filepath.Separator != '/' {
		cleanedPath = filepath.FromSlash(cleanedPath)
	}
	log.Println(cleanedPath)
	return cleanedPath, nil
}

// Helper function to parse date strings
func ParseDate(dateStr string) time.Time {
	date, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		log.Printf("Error parsing date %s: %v", dateStr, err)
		return time.Time{}
	}
	return date
}

// Helper function to parse float values
func ParseFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Printf("Error parsing float %v: %v", v, err)
			return 0
		}
		return f
	default:
		log.Printf("Unexpected type for float value: %T", v)
		return 0
	}
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
