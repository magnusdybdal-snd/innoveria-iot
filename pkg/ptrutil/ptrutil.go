// Package ptrutil contains utility functions for managing pointers.
package ptrutil

// Deref returns the value pointed to by s, or empty string if s is nil.
func Deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
