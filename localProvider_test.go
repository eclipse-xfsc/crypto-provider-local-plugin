package main

import (
	"context"
	"testing"

	core "github.com/eclipse-xfsc/crypto-provider-core"
	"github.com/eclipse-xfsc/crypto-provider-core/types"
)

func Test_signing_rsa4096(t *testing.T) {
	localProvider := new(LocalCryptoProvider)
	if !core.Sign_Testing_Rsa4096(localProvider) {
		t.Error()
	}
}

func Test_encryption_aes256(t *testing.T) {
	localProvider := new(LocalCryptoProvider)
	if !core.Encryption_Testing_Aes256(localProvider) {
		t.Error()
	}
}

func Test_encryption_ed(t *testing.T) {
	localProvider := new(LocalCryptoProvider)
	if !core.Sign_Testing_Ed(localProvider) {
		t.Error()
	}
}

func Test_GetKeys(t *testing.T) {
	localProvider := new(LocalCryptoProvider)
	b, err := core.GetKeys_Test(localProvider)

	if !b {
		t.Error(err)
	}
}

func Test_Extension(t *testing.T) {
	localProvider := new(LocalCryptoProvider)
	ctx := types.CryptoContext{
		Namespace: "1",
		Group:     "2",
		Context:   context.Background(),
	}
	err := localProvider.CreateCryptoContext(ctx)
	if err != nil {
		t.Fail()
	}

	idf := types.CryptoIdentifier{
		KeyId:         "123",
		CryptoContext: ctx,
	}

	err = localProvider.GenerateKey(types.CryptoKeyParameter{
		Identifier: idf,
		KeyType:    types.Ed25519,
	})

	if err != nil {
		t.Fail()
	}

	key, err := localProvider.GetKey(idf)

	if err != nil {
		t.Fail()
	}

	_, err = key.GetJwk()

	if err != nil {
		t.Fail()
	}
}
