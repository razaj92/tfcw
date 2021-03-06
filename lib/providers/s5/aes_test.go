package s5

import (
	"os"
	"testing"

	"github.com/mvisonneau/s5/cipher"
	"github.com/mvisonneau/tfcw/lib/schemas"
	"github.com/stretchr/testify/assert"
)

const (
	testAESKey string = "cc6af4c2bf251c1cce0aebdbd39dc91d"
)

func TestGetCipherEngineAES(t *testing.T) {
	cipherEngineType := schemas.S5CipherEngineTypeAES
	key := testAESKey

	// expected engine
	expectedEngine, err := cipher.NewAESClient(testAESKey)
	assert.Equal(t, err, nil)

	// all defined in client, empty variable config (default settings)
	v := &schemas.S5{}
	c := &Client{
		CipherEngineType: &cipherEngineType,
		CipherEngineAES: &schemas.S5CipherEngineAES{
			Key: &key,
		},
	}

	cipherEngine, err := c.getCipherEngine(v)
	assert.Equal(t, err, nil)
	assert.Equal(t, cipherEngine, expectedEngine)

	// all defined in variable, empty client config
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineAES: &schemas.S5CipherEngineAES{
			Key: &key,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	assert.Equal(t, err, nil)
	assert.Equal(t, cipherEngine, expectedEngine)

	// key defined in environment variable
	os.Setenv("S5_AES_KEY", testAESKey)
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
	}

	cipherEngine, err = c.getCipherEngine(v)
	assert.Equal(t, err, nil)
	assert.Equal(t, cipherEngine, expectedEngine)

	// other engine & key defined in client, overridden in variable
	otherCipherEngineType := schemas.S5CipherEngineTypeVault
	otherKey := "4177252ea44dea6b9d66815ab5dda08b"

	c = &Client{
		CipherEngineType: &otherCipherEngineType,
		CipherEngineAES: &schemas.S5CipherEngineAES{
			Key: &otherKey,
		},
	}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineAES: &schemas.S5CipherEngineAES{
			Key: &key,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	assert.Equal(t, err, nil)
	assert.Equal(t, cipherEngine, expectedEngine)
}
