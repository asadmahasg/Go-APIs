// Generates random strings.
//
// Example usage:
//
//	s := random.NewUri() // s is now "apHCJBl7L1OmC57n"
//
// A standard string created by NewUri() is 16 bytes in length and consists of
// Latin upper and lowercase letters, and numbers (from the set of 62 allowed
// characters), which means that it has ~95 bits of entropy. To get more
// entropy, you can use NewLenUri(UUIDLenUri), which returns 20-byte string, giving
// ~119 bits of entropy, or any other desired length.
//
// Functions read from crypto/rand random source, and panic if they fail to read from it.
package random

import "crypto/rand"

const (
	// StdLenUri is the standard length of a URI string to achieve ~95 bits of entropy.
	StdLenUri = 16

	// UUIDLenUri is the length of a URI string to achieve ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLenUri = 20
)

// StdCharsUri is a set of standard characters allowed in a URI string.
var StdCharsUri = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// NewUri returns a new random string of the standard length, consisting of
// standard characters.
func NewUri() string {
	return NewLenCharsUri(StdLenUri, StdCharsUri)
}

// NewLenUri returns a new random string of the provided length, consisting of
// standard characters.
func NewLenUri(length int) string {
	return NewLenCharsUri(length, StdCharsUri)
}

// NewLenCharsUri returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenCharsUri(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("random: wrong charset length for NewLenCharsUri")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("random: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
