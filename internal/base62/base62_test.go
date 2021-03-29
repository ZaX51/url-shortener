package base62_test

import (
	"testing"

	"github.com/ZaX51/url-shortener/internal/base62"
)

func TestTobase64EmptyArray(t *testing.T) {
	result := base62.ToBase62([]byte{})

	if len(result) != 0 {
		t.Errorf("ToBase62 = %s; want empty string", result)
	}
}

func TestTobase64OneItemArray(t *testing.T) {
	result := base62.ToBase62([]byte{255})

	if result != "aaaBBBccc" {
		t.Errorf("ToBase62 = %s; want 'aaaBBBccc' string", result)
	}
}
