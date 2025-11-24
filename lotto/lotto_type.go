package lotto

type LottoType int

const (
	UserLotto LottoType = iota
	WinningLotto
)

func (lt LottoType) String() string {
	return [...]string{"사용자 로또", "당첨 로또", "보너스 로또"}[lt]
}
