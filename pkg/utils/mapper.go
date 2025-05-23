package utils

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

func MapStruct[T any](to *T, from any) (*T, error) {
	if from == nil {
		return to, errors.New("source struct is nil")
	}
	err := copier.Copy(to, from)
	if err != nil {
		return to, fmt.Errorf("failed to map struct: %w", err)
	}
	return to, nil
}
