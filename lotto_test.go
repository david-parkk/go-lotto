package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLotto_Success(t *testing.T) {
	numbers := []int{5, 3, 12, 1, 45, 23}
	lotto, err := NewLotto(numbers, UserLotto)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3, 5, 12, 23, 45}, lotto.Numbers)
}

func TestNewLotto_WrongLength(t *testing.T) {
	numbers := []int{1, 2, 3}
	_, err := NewLotto(numbers, WinningLotto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "로또 번호는 6개")
}

func TestNewLotto_OutOfRangeNumber(t *testing.T) {
	numbers := []int{0, 2, 3, 4, 5, 6}
	_, err := NewLotto(numbers, BonusLotto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "1~45")
}
