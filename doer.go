package doer // import "github.com/nthlongtma/go-gomock"

import "io"

type Doer interface {
	io.Reader
	Close() error
	Do(int, string) error
}
