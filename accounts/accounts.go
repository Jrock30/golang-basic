package accounts

import "errors"

// Account struct
type Account struct {
	// 앞글자를 대문자로 함으로써 밖에서 사용가능 (public, 소문자면 private)
	owner   string
	balance int
}

var errNoMany = errors.New("Can't withdraw")

// NewAccount creates Account (Constructor)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 복사본 말고 원본 리턴(주소값)
}

// Deposit receiver a (Method)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your balance
func (a Account) Balance() int {
	return a.balance
}

// Withdrew x amount from your account
func (a *Account) Withdrew(amount int) error {
	if a.balance < amount {
		//return error.Error()
		// GO 는 except 가 없다 에러가 터지는 것이 아니다.
		// 받을 곳에서 에러 메세지 처리
		//return errors.New("Can`t withdraw you are poor")
		return errNoMany
	}
	a.balance -= amount
	return nil
}
