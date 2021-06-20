/**
main.go 는 프로젝트를 컴파일 하고 싶다면 만든다 (진입점, entry point)
## main 의 이름은 선택이 아니고 필수 규칙이다. (GO 언어는 선택을 할 수 없도록 표준 규칙이 많아서 심플하다.) ##
목적에 따라 필요 없을 수도 있다 공유를 위한 라이브러리를 만들거나 오픈소스 등등.
예를 들어 package learning 을 만들면 이 것은 컴파일이 되지 않는다.
  - 다른사람들이 사용하는 펑션같은 경우 main 외의 이름으로 패키지를 정하면 될 것 같다.
*/
package main

import (
	"fmt"
	"github.com/jrock30/learngo/accounts"
	"github.com/jrock30/learngo/basic"
	"github.com/jrock30/learngo/person"
	"log"
)

/**
main function
  - 시작점 entry point
*/
func main() {
	basic.Init()

	jrock := person.Person{}
	jrock.SetDetails("jrock", 20)
	// 여기서는 작동하지 않는다 set 한 데이터가. 이 것은 main 함수에 있는 jrock 객체
	// 즉 위의 객체는 복사가 되어 설정이 된다. 그래서 SetDetails 에서 객체를 *Person 포인터를 주어야한다.
	fmt.Println("Main Jrock", jrock)
	fmt.Println(">>> ", jrock.Name())

	account := accounts.NewAccount("jrock")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdrew(30)
	// 이렇게 하면 에러메세지가 발생한다.
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
