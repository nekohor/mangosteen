package gradeklass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstr(t *testing.T) {
	 sub := Substr("MRTRG00301", 7, 3)
	 assert.Equal(t, "301", sub)
}