package crypto_impl

import (
	"bytes"
	"errors"
	"testing"

	"github.com/reshifr/secure-env/core/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_NewAutoRNG(t *testing.T) {
	t.Parallel()
	fn := crypto.FnCSPRNG{}
	expRNG := AutoRNG{fnCSPRNG: fn}

	rng := NewAutoRNG(fn)
	assert.Equal(t, expRNG, rng)
}

func Test_AutoRNG_Block(t *testing.T) {
	t.Parallel()
	t.Run("ErrReadEntropyFailed error", func(t *testing.T) {
		t.Parallel()
		fn := crypto.FnCSPRNG{
			Read: func([]byte) (int, error) {
				return 0, errors.New("")
			},
		}
		var expBlock []byte = nil
		expErr := crypto.ErrReadEntropyFailed

		rng := NewAutoRNG(fn)
		block, err := rng.Block(8)
		assert.Equal(t, expBlock, block)
		assert.ErrorIs(t, err, expErr)
	})
	t.Run("Succeed", func(t *testing.T) {
		t.Parallel()
		expBlock := bytes.Repeat([]byte{0xff}, 8)
		fn := crypto.FnCSPRNG{
			Read: func(block []byte) (n int, err error) {
				copy(block, expBlock)
				return len(block), nil
			},
		}

		rng := NewAutoRNG(fn)
		block, err := rng.Block(8)
		assert.Equal(t, expBlock, block)
		assert.ErrorIs(t, err, nil)
	})
}

func Test_AutoRNG_Read(t *testing.T) {
	t.Parallel()
	t.Run("ErrReadEntropyFailed error", func(t *testing.T) {
		t.Parallel()
		fn := crypto.FnCSPRNG{
			Read: func([]byte) (int, error) {
				return 0, errors.New("")
			},
		}
		block := [8]byte{}
		expBlock := [8]byte{}
		expErr := crypto.ErrReadEntropyFailed

		rng := NewAutoRNG(fn)
		err := rng.Read(block[:])
		assert.Equal(t, expBlock, block)
		assert.ErrorIs(t, err, expErr)
	})
	t.Run("Succeed", func(t *testing.T) {
		t.Parallel()
		block := make([]byte, 8)
		expBlock := bytes.Repeat([]byte{0xff}, 8)
		fn := crypto.FnCSPRNG{
			Read: func(block []byte) (int, error) {
				copy(block, expBlock)
				return len(block), nil
			},
		}

		rng := NewAutoRNG(fn)
		err := rng.Read(block[:])
		assert.Equal(t, expBlock, block)
		assert.ErrorIs(t, err, nil)
	})
}
