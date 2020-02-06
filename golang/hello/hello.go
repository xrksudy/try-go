package hello

import "rsc.io/quote/v3"

// Hello ...
func Hello() string {
	return quote.HelloV3()
}

// Proverb ...
func Proverb() string {
	return quote.Concurrency()
}
