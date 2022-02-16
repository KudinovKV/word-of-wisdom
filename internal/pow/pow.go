package pow

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"strconv"
)

const targetBits = 24

type Manager struct {
	barrier *big.Int
}

func NewManager() *Manager {
	barrier := big.NewInt(1)
	barrier.Lsh(barrier, uint(256-targetBits))
	return &Manager{barrier: barrier}
}

func (m *Manager) Calculate(challenge string) int64 {
	var (
		result big.Int
		nonce  int64
	)
	for {
		data := append([]byte(challenge), []byte(strconv.FormatInt(nonce, 16))...)
		hash := sha256.Sum256(data)
		result.SetBytes(hash[:])
		if m.barrier.Cmp(&result) == 1 {
			break
		}
		nonce++
	}
	return nonce
}

func (m *Manager) Validate(challenge string, nonce int64) error {
	var (
		data   = append([]byte(challenge), []byte(strconv.FormatInt(nonce, 16))...)
		result big.Int
	)
	hash := sha256.Sum256(data)
	result.SetBytes(hash[:])

	if m.barrier.Cmp(&result) != 1 {
		return errors.New("challenge failed, invalid nonce")
	}
	return nil
}
