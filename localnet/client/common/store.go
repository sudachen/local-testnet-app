package common

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/spacemeshos/ed25519"
	"sudachen.xyz/pkg/localnet/fu"
)

type accountKeys struct {
	PubKey  string `json:"pubkey"`
	PrivKey string `json:"privkey"`
}

type Store map[string]accountKeys

func (s Store) CreateAccount(alias string) *LocalAccount {
	sPub, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fu.Error("cannot create account: %s", err)
		return nil
	}
	acc := &LocalAccount{Name: alias, PubKey: sPub, PrivKey: key}
	s[alias] = accountKeys{PubKey: hex.EncodeToString(sPub), PrivKey: hex.EncodeToString(key)}
	return acc
}
