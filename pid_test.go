package pid

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)
func TestPID(t *testing.T) {
	pf := NewFile("data/test.pid")
	require.NoError(t, pf.Create())
	require.Error(t, pf.Create())
	proc, err := pf.Process()
	require.NoError(t, err)
	require.Equal(t, os.Getpid(), proc.Pid)
	require.NoError(t, pf.Remove())
}
