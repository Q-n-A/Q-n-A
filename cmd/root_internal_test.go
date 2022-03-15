package cmd

import (
	"testing"

	"github.com/Q-n-A/Q-n-A/util/test"
	"github.com/stretchr/testify/assert"
)

func Test_printBanner(t *testing.T) {
	t.Parallel()

	want := `Q'n'A - traP Anonymous Question Box Service
   ____  _       _  ___
  / __ \( )____ ( )/   |
 / / / /|// __ \|// /| |
/ /_/ /  / / / / / ___ |
\___\_\ /_/ /_/ /_/  |_|
`

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		printed := test.PickStdout(t, printBanner)
		assert.Equal(t, want, printed)
	})
}
