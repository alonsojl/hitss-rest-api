package cert

import (
	"crypto/rsa"
	"hitss/pkg/helper/logger"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func Load() {
	err := readCertificate()
	if err != nil {
		panic(err)
	}
}

func readCertificate() error {
	privateCer, err := ioutil.ReadFile(os.Getenv("ROOT") + "/" + os.Getenv("PRIVATE_KEY"))
	if err != nil {
		logger.Write(err)
		return err
	}

	publicCer, err := ioutil.ReadFile(os.Getenv("ROOT") + "/" + os.Getenv("PUBLIC_KEY"))
	if err != nil {
		logger.Write(err)
		return err
	}

	return parseKeys(privateCer, publicCer)
}

func parseKeys(privateCer, publicCer []byte) error {
	var err error
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateCer)
	if err != nil {
		logger.Write(err)
		return err
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicCer)
	if err != nil {
		logger.Write(err)
		return err
	}

	return nil
}
