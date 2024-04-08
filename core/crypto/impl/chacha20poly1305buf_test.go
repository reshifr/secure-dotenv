package crypto_impl

import (
	"bytes"
	"testing"

	"github.com/reshifr/secure-env/core/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_MakeChaCha20Poly1305Buf(t *testing.T) {
	t.Parallel()
	t.Run("ErrInvalidRawIVLen error", func(t *testing.T) {
		t.Parallel()
		rawIV := bytes.Repeat([]byte{0xaa}, 8)
		ad := bytes.Repeat([]byte{0xbb}, ChaCha20Poly1305ADLen)
		ciphertext := bytes.Repeat([]byte{0xcc}, 7)
		var expBuf *ChaCha20Poly1305Buf = nil
		expErr := crypto.ErrInvalidRawIVLen
		buf, err := MakeChaCha20Poly1305Buf(rawIV, ad, ciphertext)
		assert.Equal(t, expBuf, buf)
		assert.Equal(t, err, expErr)
	})
	t.Run("ErrInvalidADLen error", func(t *testing.T) {
		t.Parallel()
		rawIV := bytes.Repeat([]byte{0xaa}, IV96Len)
		ad := bytes.Repeat([]byte{0xbb}, 8)
		ciphertext := bytes.Repeat([]byte{0xcc}, 7)
		var expBuf *ChaCha20Poly1305Buf = nil
		expErr := crypto.ErrInvalidADLen
		buf, err := MakeChaCha20Poly1305Buf(rawIV, ad, ciphertext)
		assert.Equal(t, expBuf, buf)
		assert.Equal(t, err, expErr)
	})
	t.Run("Succeed", func(t *testing.T) {
		t.Parallel()
		rawIV := bytes.Repeat([]byte{0xaa}, IV96Len)
		ad := bytes.Repeat([]byte{0xbb}, ChaCha20Poly1305ADLen)
		ciphertext := bytes.Repeat([]byte{0xcc}, 7)
		block := []byte{
			0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
			0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
			0xcc, 0xcc, 0xcc,
		}
		expBuf := &ChaCha20Poly1305Buf{block: block}
		buf, err := MakeChaCha20Poly1305Buf(rawIV, ad, ciphertext)
		assert.Equal(t, expBuf, buf)
		assert.Equal(t, err, nil)
	})
}

func Test_LoadChaCha20Poly1305Buf(t *testing.T) {
	t.Parallel()
	t.Run("ErrInvalidBuffer error", func(t *testing.T) {
		t.Parallel()
		block := []byte{
			0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
			0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb,
		}
		var expBuf *ChaCha20Poly1305Buf = nil
		expErr := crypto.ErrInvalidBuffer
		buf, err := LoadChaCha20Poly1305Buf(block)
		assert.Equal(t, expBuf, buf)
		assert.ErrorIs(t, err, expErr)
	})
	t.Run("Succeed", func(t *testing.T) {
		t.Parallel()
		block := []byte{
			0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
			0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
			0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
			0xcc, 0xcc, 0xcc,
		}
		expBuf := &ChaCha20Poly1305Buf{block: block}
		buf, err := LoadChaCha20Poly1305Buf(block)
		assert.Equal(t, expBuf, buf)
		assert.ErrorIs(t, err, nil)
	})
}

func Test_ChaCha20Poly1305Buf_RawIV(t *testing.T) {
	t.Parallel()
	block := []byte{
		0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
		0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
		0xcc, 0xcc, 0xcc,
	}
	buf, _ := LoadChaCha20Poly1305Buf(block)
	expIV := bytes.Repeat([]byte{0xaa}, IV96Len)
	iv := buf.RawIV()
	assert.Equal(t, expIV, iv)
}

func Test_ChaCha20Poly1305Buf_AD(t *testing.T) {
	t.Parallel()
	block := []byte{
		0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
		0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
		0xcc, 0xcc, 0xcc,
	}
	buf, _ := LoadChaCha20Poly1305Buf(block)
	expAD := bytes.Repeat([]byte{0xbb}, ChaCha20Poly1305ADLen)
	ad := buf.AD()
	assert.Equal(t, expAD, ad)
}

func Test_ChaCha20Poly1305Buf_Ciphertext(t *testing.T) {
	t.Parallel()
	block := []byte{
		0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
		0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
		0xcc, 0xcc, 0xcc,
	}
	buf, _ := LoadChaCha20Poly1305Buf(block)
	expCiphertext := bytes.Repeat([]byte{0xcc}, 7)
	ciphertext := buf.Ciphertext()
	assert.Equal(t, expCiphertext, ciphertext)
}

func Test_ChaCha20Poly1305Buf_Block(t *testing.T) {
	t.Parallel()
	expBlock := []byte{
		0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa,
		0xaa, 0xaa, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb,
		0xbb, 0xbb, 0xbb, 0xbb, 0xcc, 0xcc, 0xcc, 0xcc,
		0xcc, 0xcc, 0xcc,
	}
	buf, _ := LoadChaCha20Poly1305Buf(expBlock)
	block := buf.Block()
	assert.Equal(t, expBlock, block)
}
