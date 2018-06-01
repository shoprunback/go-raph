package raph

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsOne(s []string, e []string) bool {
	for _, ev := range e {
		for _, sv := range s {
			if sv == ev {
				return true
			}
		}
	}
	return false
}

func ContainsAll(s []string, e []string) bool {
	for _, ev := range e {
		found := false
		for _, sv := range s {
			if sv == ev {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Reverse(s []string) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

var alphabet = []rune("0123456789abcdef")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
