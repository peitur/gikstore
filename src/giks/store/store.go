package store

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func GeratePrivateKeyRSA(bits int, pwd string) (string, error) {

	rsax, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not generate private key: %s", err))
	}

	// pem.Block
	pem_block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsax),
	}

	if pwd != "" {
		pem_block, err = x509.EncryptPEMBlock(rand.Reader, pem_block.Type, pem_block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return "", err
		}
	}
	// Private key in PEM format
	pem := pem.EncodeToMemory(pem_block)

	return string(pem), nil
}
