package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	result := jsonError("test")
	require.Equal(t, []byte(`{"message":"test"}`), result)
}
