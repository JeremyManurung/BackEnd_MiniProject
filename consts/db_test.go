package consts_test

import (
	"minipro/consts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbNotMixed(t *testing.T) {
	assert.NotEqual(t, consts.DB_MAIN, consts.DB_TEST)
}
