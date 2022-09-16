package random_test

import (
	"testing"

	"project/libgo/random"
)

func validateCharsUri(t *testing.T, u string, chars []byte) {
	for _, c := range u {
		var present bool
		for _, a := range chars {
			if rune(a) == c {
				present = true
			}
		}
		if !present {
			t.Fatalf("chars not allowed in %q", u)
		}
	}
}

func TestNewUri(t *testing.T) {
	u := random.NewUri()
	// Check length
	if len(u) != random.StdLenUri {
		t.Fatalf("wrong length: expected %d, got %d", random.StdLenUri, len(u))
	}
	// Check that only allowed characters are present
	validateCharsUri(t, u, random.StdCharsUri)

	// Generate 1000 uniuris and check that they are unique
	uris := make([]string, 1000)
	for i := range uris {
		uris[i] = random.NewUri()
	}
	for i, u := range uris {
		for j, u2 := range uris {
			if i != j && u == u2 {
				t.Fatalf("not unique: %d:%q and %d:%q", i, u, j, u2)
			}
		}
	}
}

func TestNewLenUri(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := random.NewLenUri(i)
		if len(u) != i {
			t.Fatalf("request length %d, got %d", i, len(u))
		}
	}
}

func TestNewLenCharsUri(t *testing.T) {
	length := 10
	chars := []byte("01234567")
	u := random.NewLenCharsUri(length, chars)

	// Check length
	if len(u) != length {
		t.Fatalf("wrong length: expected %d, got %d", random.StdLenUri, len(u))
	}
	// Check that only allowed characters are present
	validateCharsUri(t, u, chars)

	// Check that two generated strings are different
	u2 := random.NewLenCharsUri(length, chars)
	if u == u2 {
		t.Fatalf("not unique: %q and %q", u, u2)
	}
}

func TestNewLenCharsUriMaxLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("didn't panic")
		}
	}()
	chars := make([]byte, 257)
	random.NewLenCharsUri(32, chars)
}

func TestBias(t *testing.T) {
	chars := []byte("abcdefghijklmnopqrstuvwxyz")
	slen := 100000
	s := random.NewLenCharsUri(slen, chars)
	counts := make(map[rune]int)
	for _, b := range s {
		counts[b]++
	}
	avg := float64(slen) / float64(len(chars))
	for k, n := range counts {
		diff := float64(n) / avg
		if diff < 0.95 || diff > 1.05 {
			t.Errorf("Bias on '%c': expected average %f, got %d", k, avg, n)
		}
	}
}
