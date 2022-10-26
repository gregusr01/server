// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"time"

	"github.com/go-vela/server/tokenmanager"
	"github.com/golang-jwt/jwt"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// helper function to setup the queue from the CLI arguments.
func setupTokenManger(c *cli.Context) (tokenmanager.Service, error) {

	logrus.Debug("Creating tokenManger for server worker authentication")

	//THESE KEYS MEAN NOTHING - JUST TEST KEYS
	newKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAh98PryJvOO5SuO17VgNI1z+ikCiFUSPNWL6rTtnCWclDIX8dnc8+yppKR4vw56FZTO0ZdkVw8wFnkPkj/7HIkSy5/UimlGYTvjJ2JOBwH4BfCp4RRhLVoiTLSdAu1K0Jfl3OQKtEui3o4rxGe6MXlR+uT+annENh/2RUWgCanrZQt04PgZJ04XUiI+t9jbL5yx/fawaQlSRCFW1RE3oXhuT7muh01ZaZpuU1upu/hzbpMEPlwDeBMx9VzFaI1eY4aRNd2lyu7p7ILCynCceHde/Q0RadmobxQLKI5K+1r56II+wQRuurr4nsIhcG8T6XY/t/4Du0rShYN744+ljl2wIDAQABAoIBABLhcVmB7Hi5zW77OT9bl3yp0Bo+N0BuPDPP0xhS9EnryNNXybLLQMuAoz/L938IkdM5w2cHAUoTcOPZayI1/0wFLhc/SD6o7KdqdbZsJciK2yorivT02xD8Ee/A6TEOlpojyOx9oEBK7ujLBvRZVoaXb26U+8egKXcG5x3WpaXWNHvd8qgxfQ7k1alu4IDu49nqeRXPbzgTh2PzrVw1vIr2LHqVlIEYp2evbMWx/wYKErBE3QCeMYM631kC60fildJpHlAV1gf7r37lMlL6we6dOvvJxU+wkFYft34/TViWoXXNDIDVoZet+W+IRMtHi7otpuZADT0sjgon3pDgL/ECgYEA7mgVln7a1PxMlcKc3uEd0+drtzaElcpUAf08l/YIQBy0keUeZB6TLvs6m+y6IQVLfnVnCB28q00v7Q+f9V6oeuh8VWBnF/5ipWnvK0515eMnQFQXeSVsHQRz1dy0p6zxaALvyBsmyJvHfkdif4nq9Cre7Fk7oewp4PalWCpkxakCgYEAkeXoUM8aUAXwHY/e4iTiJeUZoo+gZOTg9n0FFlTufT7hEszPWPeiZOX//KLvdRJugKymzVhnYSwr/6LfU+LLc8TtetnsbCGwVH9f4lRDsOIMvjkeCGP2CDEcnMTUwHxIJvKPp8b2Q8SHZPUNhEvMFv0kTe3bZ1v/6+uBCNIDOeMCgYEAqewnzt9Vq86z614Nn1IGh31/qdNCxtyx6cUMBieHE+Mh1IbglW2xbCAGUxJ0S4rRly2opQFE8zeNvEKRuTqhjRDdZaDTeieHOez6WfyPTa4M0O3e2SsSFtCQm1K2tHgwi/jj3NV1XDCkDN5mVw7/Fs7jmsOzHCMOaliB2whEURkCgYA2ZaUXtBAYdA5Rx1mPsbbL8D59mNVxqNCjtntWFzaQZghfDRPmFPZsbkDifhGP8d5XhYfvmh15YpHJD3369d1rfaoZmvUGYA0xhAhJB6fxZGbh6cPj1vldlooXfV/hoLT6KIgdQxaAr97XanKut+ARVjLaB2w9flTOLpF+q2behwKBgQDldzmCXKGUZIk1tQAfqDaMAcrNtLKSEFt9CDhuwIe+MMuW5odsV/DPZmPOyN6cTN+2cTqGcWI2YGc25ourboI6ya/vkWwuRdHMv5DAFco0hn3uA41zaFM4qwNNyLsxcdFN2AyqXLJfKBLausQf5ZBfhseaRSDOHpAfiXYI9bnZcQ==\n-----END RSA PRIVATE KEY-----"

	signKey := newKey

	k, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(signKey))
	if err != nil {
		logrus.Trace("error parsing private key from PEM")
	}

	signKeyPublic := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAh98PryJvOO5SuO17VgNI1z+ikCiFUSPNWL6rTtnCWclDIX8dnc8+yppKR4vw56FZTO0ZdkVw8wFnkPkj/7HIkSy5/UimlGYTvjJ2JOBwH4BfCp4RRhLVoiTLSdAu1K0Jfl3OQKtEui3o4rxGe6MXlR+uT+annENh/2RUWgCanrZQt04PgZJ04XUiI+t9jbL5yx/fawaQlSRCFW1RE3oXhuT7muh01ZaZpuU1upu/hzbpMEPlwDeBMx9VzFaI1eY4aRNd2lyu7p7ILCynCceHde/Q0RadmobxQLKI5K+1r56II+wQRuurr4nsIhcG8T6XY/t/4Du0rShYN744+ljl2wIDAQAB\n-----END PUBLIC KEY-----"

	pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(signKeyPublic))
	if err != nil {
		logrus.Trace("error parsing public key from PEM")
	}

	_manager := &tokenmanager.Setup{
		Driver: "minter",
		//Database: d,
		PrivKey:           k,
		PubKey:            pk,
		SignMethod:        jwt.SigningMethodRS256,
		RegTokenDuration:  time.Minute * 10,
		AuthTokenDuration: time.Minute * 10,
	}

	return tokenmanager.New(_manager)
}
