package BlockChain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

type Wallet struct {
	PubKey     string
	RandomWord []string
}

func CreateWallet(userName string) *pem.Block {

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	_ = os.MkdirAll("../../internal/wallets/" + userName + "/", 0755);
	privatePem, _ := os.Create("../../internal/wallets/" + userName + "/Private.pem")

	err := pem.Encode(privatePem, privateKeyBlock)

	if err != nil {
		log.Println("we have an error in private pem encoding");
	}


	publicKeyByets, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: publicKeyByets,
	}	
	publicPem, _ := os.Create("../../internal/wallets/" + userName +  "/Public.pem")
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		log.Println("we have an error in public pem encoding");
	}
	return publicKeyBlock
}
