package coin

import (
	"math/rand"
	"crypto/x509"
	"crypto/rsa"
	"crypto/sha512"
	"encoding"
)

// KeyPair -
type KeyPair struct {
	PublicKey pem.Block
	PrivateKeyEncrypted pem.Block
}

const keyBitSize = 4096

func GenerateKeyPair(passphrase string) (KeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, keyBitSize)
	if err != nil {
		return KeyPair{}, err
	}
	publicKeyPEM := pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key.PublicKey)
	}

	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	privateKeyPEMEncrypted, err := x509.EncryptPEMBlock(
		rand.Reader,
		privateKeyPEM.Type,
		privateKeyPEM.Bytes,
		[]byte(passphrase),
		x509.PEMCipherAES256,
	)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{publicKeyPEM, privateKeyPEMEncrypted}, nil
}

func DecryptPrivateKey(privateKeyEncrypted pem.Block, passphrase string) (rsa.PrivateKey, error) {
	decryptedBytes, err := x509.DecryptPEMBlock(privateKeyEncrypted, []byte(passphrase))
	if err != nil {
		return pem.Block{}, errors.New("Could not decrypt private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedBytes)
	if err != nil {
		return pem.Block{}, errors.New("Invalid private key decrypted")
	}


}

func Sign(privateKey rsa.PrivateKey, data interface{}) ([]byte, error) {
	dataBytes, _ := data.(encoding.BinaryMarshaler).MarshalBinary();
	hash := sha512.Sum512(dataBytes)
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA512,
		hash[:],
	)
	if err != nil {
		return []byte{}, err
	}
	return signature, nil
}

func VerifySign(publicKey rsa.PublicKey, data interface{}, signature []byte) (bool, err) {
	dataBytes, _ := data.(encoding.BinaryMarshaler).MarshalBinary();
	hash := sha512.Sum512(dataBytes)
	err := rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA512,
		hash,
		signature,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}