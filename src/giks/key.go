package giks

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func SelectHashFunc(alg string) hash.Hash {
	if alg == "sha1" {
		return sha1.New()
	} else if alg == "sha256" {
		return sha256.New()
	} else if alg == "sha512" {
		return sha512.New()
	} else {
		return SelectHashFunc("sha256")
	}
}

func FileChecksum(filename string, alg string) string {
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	var hsh = SelectHashFunc(alg)
	if _, err := io.Copy(hsh, fd); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hsh.Sum(nil))
}

func GeneratePrivateKeyRSA(bits int) (*rsa.PrivateKey, error) {

	rsax, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not generate private key: %s", err))
	}
	return rsax, nil
}


func KeyPrivatePEM(key *rsa.PrivateKey, pwd string) (string, error) {
	// pem.Block
	pem_block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pemx := pem.EncodeToMemory(pem_block)
	if pwd != "" {
		pem_block, err := x509.EncryptPEMBlock(rand.Reader, pem_block.Type, pem_block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return "", err
		}
		pemx = pem.EncodeToMemory(pem_block)

	}
	return string(pemx), nil
}

func KeyPublicPEM(key *rsa.PrivateKey, pwd string) (string, error) {
	// pem.Block
	pem_block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	}

	pemx := pem.EncodeToMemory(pem_block)
	if pwd != "" {
		pem_block, err := x509.EncryptPEMBlock(rand.Reader, pem_block.Type, pem_block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return "", err
		}
		pemx = pem.EncodeToMemory(pem_block)

	}
	return string(pemx), nil
}
