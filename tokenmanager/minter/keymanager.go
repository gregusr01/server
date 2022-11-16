package minter

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/sirupsen/logrus"
)

type signingKey struct {
	Kid        sql.NullString `sql:"kid"`
	PublicKey  sql.NullString `sql:"public_key"`
	ServerName sql.NullString `sql:"server_name"`
	Timestamp  sql.NullInt64  `sql:"timestamp"`
}

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) RefreshKeyCache() {
	logrus.Info("Minter:RefreshKeyCache Function Called")
	for {
		ticker := time.NewTicker(c.config.KeyCleanupTicker) //double check me
		for range ticker.C {
			logrus.Info("Retrieving list of valid public signing keys")

			//retrieve list of valid public signing keys

			if err := c.Database.UpdateKeyTTL(c.config.Kid); err != nil {
				logrus.Infof("unable to update key ttl for key kid: %s", c.config.Kid)
			}

			DBkeys, err := c.Database.ListSigningKeys()
			if err != nil {
				logrus.Info("Error retrieving list of public signing keys", err)
			}

			pubKeys := make(map[string]*rsa.PublicKey)

			//update cache
			for _, v := range DBkeys {

				if !v.PublicKey.Valid {
					logrus.Info("public key not valid")
					continue
				}

				pk, err := convertKeyString(v.PublicKey.String)
				if err != nil {
					logrus.Info("unable to decode public key", err)
					continue
				}
				if !v.Kid.Valid {
					logrus.Info("kid not valid")
					continue
				}

				//parse into pub key format
				pubKeys[v.Kid.String] = pk
			}

			//should really use mutex here
			c.config.PublicKeyCache = pubKeys
			for k, v := range pubKeys {
				logrus.Infof("public key cache kid %v: value %v", k, v)
			}
		}
	}
}

//takes a b64 encoded public key string and converts it to an *rsa.PublicKey for outside consumption
func convertKeyString(k string) (*rsa.PublicKey, error) {

	//decode public key
	unB64, err := base64.StdEncoding.DecodeString(k)

	if err != nil {
		return nil, err
	}

	//parse binary to *rsa.PublicKey
	return x509.ParsePKCS1PublicKey(unB64)
}
