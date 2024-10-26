package helper

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
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

func ParseDate(dateStr string) time.Time {
	date, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return time.Time{}
	}
	return date
}

func ParseFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		v = strings.ToLower(v)
		v = strings.ReplaceAll(v, "rp", "")
		v = strings.ReplaceAll(v, ".", "")
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return f
	default:
		return 0
	}
}

func ParseDecimal(value interface{}) decimal.Decimal {
	switch v := value.(type) {
	case float64:
		return decimal.NewFromFloat(v)
	case string:
		v = strings.ToLower(v)
		v = strings.ReplaceAll(v, "rp", "")
		v = strings.ReplaceAll(v, ".", "")
		f, err := decimal.NewFromString(v)
		if err != nil {
			return decimal.Zero
		}
		return f
	default:
		return decimal.Zero
	}
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
