package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Encode writes the given value as JSON and sets the provided http status code
func Encode[T any](w http.ResponseWriter, code int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("error: encoding json: %w", err)
	}
	return nil
}

// Decode reads and decoded JSON to a given type
func Decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error: decoding json: %w", err)
	}
	return data, nil
}

// IsEmpty reports wether all fields in a struct are nil pointers.
// Useful for rejecting patch requests where no fields were provided.
func IsEmpty(v any) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	for i := range rv.NumField() {
		f := rv.Field(i)
		if f.Kind() != reflect.Ptr || !f.IsNil() {
			return false
		}
	}
	return true
}
