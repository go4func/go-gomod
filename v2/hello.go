package hello // import "github.com/nthlongtma/go-gomod/v2"

import "rsc.io/quote/v3"

func Hello() string {
	return "hello from version 2"
}

func Proverb() string {
	return quote.Concurrency()
}

func Quote() string {
	return "versioning is func."
}
