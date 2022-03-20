package test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// PickStdout 渡された関数のstdoutへの出力を取得する
func PickStdout(t *testing.T, callback func()) string {
	t.Helper()

	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()

	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdout = w

	callback()

	err = w.Close()
	assert.NoError(t, err)

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(r)
	assert.NoError(t, err)

	s := buffer.String()

	return s[:len(s)-1]
}
