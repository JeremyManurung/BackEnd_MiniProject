package consts_test

import (
	"minipro/consts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbNotMixed(t *testing.T) {
	// Pastikan db testing dan db main tidak sama
	assert.NotEqual(t, consts.DB_MAIN, consts.DB_TEST)
}
