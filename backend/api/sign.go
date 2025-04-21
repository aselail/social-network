package api

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const (
	pubKeyFile  = "id_ed25519.pub"
	privKeyFile = "id_ed25519.pem"
)

var publicKey ed25519.PublicKey
var privateKey ed25519.PrivateKey

func GenOrLoadKey() {
	err := loadKeys()
	if err != nil {
		log.Println("Generating new key pair...")
		generateKeys()
		if err := saveKeys(); err != nil {
			log.Fatalf("Error saving keys: %v", err)
		}
		log.Println("Key pair generated and saved.")
	} else {
		log.Println("Key pair loaded from disk.")
	}
}

func generateKeys() {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatalf("Error generating key: %v", err)
	}
	privateKey = priv
	publicKey = pub
}

func saveKeys() error {
	if err := os.WriteFile(pubKeyFile, publicKey, 0644); err != nil {
		return fmt.Errorf("error writing public key: %w", err)
	}
	if err := os.WriteFile(privKeyFile, privateKey, 0600); err != nil {
		return fmt.Errorf("error writing private key: %w", err)
	}
	return nil
}

func loadKeys() error {
	privBytes, err := os.ReadFile(privKeyFile)
	if err != nil {
		return err
	}
	pubBytes, err := os.ReadFile(pubKeyFile)
	if err != nil {
		return err
	}

	if len(privBytes) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid private key size")
	}
	if len(pubBytes) != ed25519.PublicKeySize {
		return fmt.Errorf("invalid public key size")
	}

	privateKey = privBytes
	publicKey = pubBytes
	return nil
}

func Sign(data string) string {
	dataBytes := []byte(data)
	sig := ed25519.Sign(privateKey, dataBytes)
	return hex.EncodeToString(sig)
}

func Verify(data string, signature string) bool {

	dataBytes := []byte(data)
	signatureBytes, err := hex.DecodeString(signature)

	if err != nil {
		return false
	}

	return ed25519.Verify(publicKey, dataBytes, signatureBytes)
}
