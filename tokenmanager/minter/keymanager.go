package minter

import (
	"time"

	"github.com/sirupsen/logrus"
)

type signingKey struct {
	Kid 				sql.NullString `sql:"kid"`
	PublicKey 	sql.NullString `sql:"public_key"`
	ServerName	sql.NullString `sql:"server_name"`
	Timestamp 	sql.NullInt64  `sql:"timestamp"`
}

// CleanInvalidTokens cleans old entries to the invalid token db
func (c *client) RefreshKeyCache() {
	logrus.Info("Minter:RefreshKeyCache Function Called")
	for {
		ticker := time.NewTicker(c.config.TickerInterval) //double check me
		for range ticker.C {
			logrus.Info("Retrieving list of valid public signing keys")

      //retrieve list of valid public signing keys
      var DBkeys []signingKey
      DBkeys, err := c.Database.ListSigningKeys()
      if err != nil {
        logrus.Info("Error retrieving list of public signing keys", err)
      }

      pubKeys := make(map[string]*rsa.Pubkey)

      //update cache
      for k,v range := DBkeys {

        pk, err := convertKeyString(v.PublicKey)
        if err != nil {
          logrus.Info("unable to decode public key", err)
        }

        //parse into pub key format
        pubKeys[v.Kid] := pk
      }

      //should really use mutex here
      c.config.PublicKeyCache = &pubkeys
		}
	}
}

//takes a b64 encoded public key string and converts it to an *rsa.PublicKey for outside consumption
func convertKeyString(k string) (*rsa.PublicKey, error) {

  //decode public key
  unB64, err := base64.StdEncoding.DecodeString(v.PublicKey)

  if err != nil {
    return nil, err
  }

  //parse binary to *rsa.PublicKey
  return x509.ParsePKCS1PublicKey(unB64)
}
