package main

import (
	"crypto/rand"
	"os"

	"github.com/tjfoc/gmsm/sm2"

	mysm2 "github.com/jianlu8023/go-example/pkg/crypto/sm2"
	"github.com/jianlu8023/go-example/pkg/logger"
)

var log = logger.GetAppLogger()

func main() {

	privateKey, err := sm2.GenerateKey(rand.Reader)

	publicKeyPath := "testdata/sm2/public.pem"
	privateKeyPath := "testdata/sm2/private.pem"

	if err != nil {
		log.Errorf("GenerateKey error: %v", err)
		return
	}

	err = mysm2.PrivateKeySave(privateKey, "123", privateKeyPath)
	if err != nil {
		log.Errorf("PrivateKeySave error: %v", err)
		return
	}

	publicKey := privateKey.PublicKey

	err = mysm2.PublicKeySave(&publicKey, publicKeyPath)
	if err != nil {
		log.Errorf("PublicKeySave error: %v", err)
		return
	}
	original := "hello world this is a sm2 crypto test"

	log.Infof("original: %s", original)

	pubKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Errorf("ReadFile error: %v", err)
		return
	}

	encrypt, err := mysm2.Encrypt(string(pubKeyBytes), original, true)
	if err != nil {
		log.Errorf("Encrypt error: %v", err)
		return
	}
	log.Infof("encrypt: %s", encrypt)

	privKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Errorf("ReadFile error: %v", err)
		return
	}
	decrypt, err := mysm2.Decrypt(string(privKeyBytes), "123", string(encrypt), true)
	if err != nil {
		log.Errorf("Decrypt error: %v", err)
		return
	}
	log.Infof("decrypt: %s", decrypt)
}
