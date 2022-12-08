package helpers

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	largestuint64 = 18446744073709551615
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// seededRand is a rand.Rand that is seeded once.
var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

// StringWithCharset returns a random string of length n consisting of characters in charset.
func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

// String returns a random string of length n consisting of characters in charset.
func String(length int) string {
  return StringWithCharset(length, charset)
}

//EnforceHTTP adds http:// to the beginning of a url if it is not present
func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

//RemoveDomainError removes the domain from the url and checks if it is the same as the domain
func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}
	return true
}