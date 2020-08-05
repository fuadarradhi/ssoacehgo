package ssoacehgo

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"os"
	"time"
)

type SSOAcehResult struct {
	ID              string    `json:"id" bson:"id"`
	Nama            string    `json:"nm" bson:"nm"`
	Email           string    `json:"em" bson:"em"`
	EmailAlternatif string    `json:"ea" bson:"ea"`
	TelegramID      string    `json:"tl" bson:"tl"`
	HP              string    `json:"hp" bson:"hp"`
	NIK             string    `json:"nk" bson:"nk"`
	NIP             string    `json:"np" bson:"np"`
	Avatar          string    `json:"av" bson:"av"`
	DateTime        time.Time `json:"dt" bson:"dt"`
}

type SSOAcehJson struct {
	ApplicationID          string `json:"application_id" bson:"application_id"`
	ApplicationName        string `json:"application_name" bson:"application_name"`
	ApplicationDisplayName string `json:"application_display_name" bson:"application_display_name"`
	ApplicationDomain      string `json:"application_domain" bson:"application_domain"`
	SSOLoginUri            string `json:"sso_login_uri" bson:"sso_login_uri"`
	SSOLogoutUri           string `json:"sso_logout_uri" bson:"sso_logout_uri"`
	Base64RSAPrivateKey    string `json:"base64_rsa_private_key" bson:"base64_rsa_private_key"`
}

type SSOAcehClient interface {
	ParseToken(token string) (SSOAcehResult, error)
}

func NewSSOAcehClient(jsonPath string) (SSOAcehClient, error) {
	jsonFile, err := os.Open(jsonPath)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var jsonValue SSOAcehJson
	err = json.Unmarshal(byteValue, &jsonValue)
	if err != nil {
		return nil, err
	}
	return jsonValue, nil
}

func (j SSOAcehJson) ParseToken(token string) (SSOAcehResult, error) {
	var res = SSOAcehResult{}

	decodeToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return res, err
	}

	private, err := base64.URLEncoding.DecodeString(j.Base64RSAPrivateKey)
	if err != nil {
		return res, err
	}

	block, _ := pem.Decode([]byte(private))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	hash := sha1.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, key, decodeToken, nil)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(plaintext, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
