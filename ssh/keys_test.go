package ssh

import (
	"encoding/pem"
	"testing"
)

func TestNewKeyPair(t *testing.T) {
	pair, err := NewKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	if privPem := pem.EncodeToMemory(&pem.Block{"RSA PRIVATE KEY", nil, pair.PrivateKey}); len(privPem) == 0 {
		t.Fatal("No PEM returned")
	}

	if fingerprint := pair.Fingerprint(); len(fingerprint) == 0 {
		t.Fatal("Unable to generate fingerprint")
	}
}
