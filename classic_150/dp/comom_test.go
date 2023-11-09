package dp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInterleave(t *testing.T) {
	require := require.New(t)

	ok := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	require.True(ok)
}
