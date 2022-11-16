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
	c.Logger.Tracef("Starting RefreshKeyCache function on loop.  Will run every %v to update the signing key cache with the latest entries", c.config.KeyCleanupTicker)
	for {
		ticker := time.NewTicker(c.config.KeyCleanupTicker) //double check me
		for range ticker.C {
			c.Logger.Tracef("Refreshing list of valid public signing keys")

			//retrieve list of valid public signing keys
			if err := c.Database.UpdateKeyTTL(c.config.Kid); err != nil {
				c.Logger.Warningf("unable to update key ttl for key kid: %s", c.config.Kid)
			}

			DBkeys, err := c.Database.ListSigningKeys()
			if err != nil {
				c.Logger.Warning("Error retrieving list of public signing keys", err)
			}

			pubKeys := make(map[string]*rsa.PublicKey)

			//update cache
			for _, v := range DBkeys {

				if !v.PublicKey.Valid {
					c.Logger.Warningf("Public key not valid.  Value: %v", v.PublicKey)
					continue
				}

				pk, err := convertKeyString(v.PublicKey.String)
				if err != nil {
					c.Logger.Warning("Unable to decode public key: ", err)
					continue
				}
				if !v.Kid.Valid {
					c.Logger.Warningf("Kid not valid. Value: %v", v.Kid)
					continue
				}

				//parse into pub key format
				pubKeys[v.Kid.String] = pk
			}

			c.config.PublicKeyCache = pubKeys
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
