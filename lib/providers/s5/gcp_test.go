package s5

import (
	"os"
	"testing"

	"github.com/mvisonneau/go-helpers/test"
	"github.com/mvisonneau/s5/cipher"
	cipherGCP "github.com/mvisonneau/s5/cipher/gcp"
	"github.com/mvisonneau/tfcw/lib/schemas"
)

const (
	testGCPKMSKeyName string = "foo"
)

func TestGetCipherEngineGCP(t *testing.T) {
	cipherEngineType := schemas.S5CipherEngineTypeGCP
	kmsKeyName := testGCPKMSKeyName

	// expected engine
	expectedEngine, err := cipher.NewGCPClient(kmsKeyName)
	test.Expect(t, err, nil)

	// all defined in client, empty variable config (default settings)
	v := &schemas.S5{}
	c := &Client{
		CipherEngineType: &cipherEngineType,
		CipherEngineGCP: &schemas.S5CipherEngineGCP{
			KmsKeyName: &kmsKeyName,
		},
	}

	cipherEngine, err := c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherGCP.Client).Config, expectedEngine.Config)

	// all defined in variable, empty client config
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineGCP: &schemas.S5CipherEngineGCP{
			KmsKeyName: &kmsKeyName,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherGCP.Client).Config, expectedEngine.Config)

	// key defined in environment variable
	os.Setenv("S5_GCP_KMS_KEY_NAME", testGCPKMSKeyName)
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherGCP.Client).Config, expectedEngine.Config)

	// other engine & key defined in client, overridden in variable
	otherCipherEngineType := schemas.S5CipherEngineTypeVault
	otherKmsKeyName := "bar"

	c = &Client{
		CipherEngineType: &otherCipherEngineType,
		CipherEngineGCP: &schemas.S5CipherEngineGCP{
			KmsKeyName: &otherKmsKeyName,
		},
	}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineGCP: &schemas.S5CipherEngineGCP{
			KmsKeyName: &kmsKeyName,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherGCP.Client).Config, expectedEngine.Config)
}
