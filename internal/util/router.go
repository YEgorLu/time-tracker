package util

import "fmt"

// Formats Mux route with method
func Rpm(basePath string) func(method, path string) string {
	return func(method, path string) string {
		if path == "" {
			return fmt.Sprintf("%s %s", method, basePath)
		}
		return fmt.Sprintf("%s %s/%s", method, basePath, path)
	}
}
