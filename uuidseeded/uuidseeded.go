package uuidseeded

import (
	"fmt"
	"hash/fnv"
	"io"
	"math/rand"

	"github.com/google/uuid"
)

// New generates UUID string using given seed reader.
func New(seedReader io.Reader) (string, error) {
	// calculate seed (int64)
	h := fnv.New64()
	_, err := io.Copy(h, seedReader)
	if err != nil {
		return "", fmt.Errorf("unexpected error occured when writing bytes from seed reader to hash: %w", err)
	}
	seed := int64(h.Sum64())

	// initialize rand reader from seed
	src := rand.NewSource(seed)
	var randReader io.Reader = rand.New(src)

	// generate UUID
	id, err := uuid.NewRandomFromReader(randReader)
	if err != nil {
		return "", fmt.Errorf("unexpected error occured when generating uuid: %w", err)
	}
	return id.String(), nil
}

