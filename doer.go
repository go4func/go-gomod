package doer // import "github/com/mytest/doer"

import "io"

type Doer interface {
	io.Reader
	Close() error
	Do(int, string) error
}
