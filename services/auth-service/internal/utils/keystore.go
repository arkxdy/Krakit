package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Key struct {
	Kid        string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	IsActive   bool
	CreatedAt  time.Time
}

type KeyStore struct {
	keys map[string]*Key
	mu   sync.RWMutex
}

func NewKeyStore() *KeyStore {
	return &KeyStore{
		keys: make(map[string]*Key),
	}
}

func (ks *KeyStore) Rotate() {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	for _, k := range ks.keys {
		k.IsActive = false
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	k := &Key{
		Kid:        uuid.NewString(),
		PrivateKey: priv,
		PublicKey:  &priv.PublicKey,
		IsActive:   true,
		CreatedAt:  time.Now(),
	}

	ks.keys[k.Kid] = k
}

func (ks *KeyStore) GetActiveKey() *Key {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	for _, k := range ks.keys {
		if k.IsActive {
			return k
		}
	}
	return nil
}

func (ks *KeyStore) GetKey(kid string) *Key {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	return ks.keys[kid]
}

func (ks *KeyStore) GetAll() []*Key {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	var out []*Key
	for _, k := range ks.keys {
		out = append(out, k)
	}
	return out
}
