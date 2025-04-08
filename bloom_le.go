//go:build !s390x

package parquet

import (
	"github.com/hbernardo/parquet-go/deprecated"
	"github.com/hbernardo/parquet-go/internal/unsafecast"
)

func unsafecastInt96ToBytes(src []deprecated.Int96) []byte {
	return unsafecast.Slice[byte](src)
}
