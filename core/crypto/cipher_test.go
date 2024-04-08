package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CipherError_Error(t *testing.T) {
	t.Parallel()
	t.Run("ErrInvalidIVLen value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidIVLen
		expMsg := "ErrInvalidIVLen: invalid IV length."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrInvalidIVFixedLen value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidIVFixedLen
		expMsg := "ErrInvalidIVFixedLen: invalid IV 'fixed' length."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrInvalidRawIVLen value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidRawIVLen
		expMsg := "ErrInvalidRawIVLen: invalid raw IV length."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrInvalidKeyLen value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidKeyLen
		expMsg := "ErrInvalidKeyLen: invalid key length."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrInvalidADLen value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidADLen
		expMsg := "ErrInvalidADLen: invalid additional data length."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrInvalidBuffer value", func(t *testing.T) {
		t.Parallel()
		err := ErrInvalidBuffer
		expMsg := "ErrInvalidBuffer: the buffer structure cannot be read."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("ErrCipherAuthFailed value", func(t *testing.T) {
		t.Parallel()
		err := ErrCipherAuthFailed
		expMsg := "ErrCipherAuthFailed: failed to decrypt the data."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
	t.Run("Unknown value", func(t *testing.T) {
		t.Parallel()
		err := CipherError(957361)
		expMsg := "Error: unknown."
		msg := err.Error()
		assert.Equal(t, expMsg, msg)
	})
}
