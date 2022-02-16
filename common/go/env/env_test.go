package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		env        string
		exportFunc func() bool
	}{
		{"prod", IsProd},
		{"staging", IsStaging},
		{"branch", IsBranchPreview},
		{"dev", IsDev},
	}

	for _, tc := range testCases {
		os.Setenv("ENV", tc.env)
		err := Parse()
		assert.NoError(t, err)
		assert.True(t, tc.exportFunc())
	}
}

func TestParseError(t *testing.T) {
	os.Unsetenv("ENV")
	err := Parse()
	assert.Error(t, err)
}
