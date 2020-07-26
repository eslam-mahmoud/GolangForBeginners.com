package shortener

import (
	"math"
	"shortener/repo"
	filesystemStorage "shortener/repo/filesystem"
	"strings"
)

// Service struct have functions Shorten & Expand
// use NewService as constructor
type Service struct {
	storage repo.Repo
}

// NewService constructor for Service
func NewService() (*Service, error) {
	s, err := filesystemStorage.NewStorage("/tmp/urls")
	if err != nil {
		return nil, err
	}
	return &Service{
		storage: s,
	}, nil
}

// Shorten takes long url
// return short url id, nil or "", error
func (s *Service) Shorten(longURL string) (string, error) {
	return s.storage.Shorten(longURL)
}

// Expand takes short url id
// return long full url and error
func (s *Service) Expand(ID string) (string, error) {
	return s.storage.Expand(ID)
}

var base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func decToBase62(x int, out *string) {
	d := x / 62
	r := x % 62
	*out = string(base62[r]) + *out
	if d > 62 {
		decToBase62(d, out)
	} else {
		*out = string(base62[d]) + *out
	}
}

func base62ToDec(s string) (result int) {
	power := 0
	s = reverse(s)
	for pos := range s {
		result += strings.Index(base62, string(s[pos])) * int(math.Pow(62, float64(power)))
		power++
	}
	return
}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
