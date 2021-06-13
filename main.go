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
	"github.com/jrock30/learngo/something"
	"strings"
)

//func multiply(a int, b int) int {
func multiply(a, b int) int { // 두개의 타입이 같으면 뒤에 하나만 적어주면 된다.
	return a * b
}

// GO Lang 은 여러개의 Return 값을 가질 수 있다.
func leAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 여러개의 아규먼트를 받는 방법
func repeatMe(words ...string) {
	fmt.Println(words) // 배열로 받아짐.
}

/**
main function
  - 시작점 entry point
*/
func main() {
	/*
		fmt
		  - 포맷팅을 위한 패키지
	*/
	fmt.Println("Hello world")
	something.SayHello() // 앞에 대문자 메서드를 입력하면 이렇게 가져다 쓸 수 있다. (소문자는 안됨)

	// 변수 (변수를 사용하지 않으면 컴파일 에러가 발생하네?)
	var name1 string = "jrock2"
	fmt.Println(name1)
	/**
	  (축약형)변수를 이렇게 짧게도 선언할 수 있다.
	  이렇게 하면 타입도 자동으로 찾아준다.
	  func 밖에서는 사용 불가 var 를 사용해야함.
	*/
	name := "jrock3"
	fmt.Println(name)

	// 상수
	const name2 string = "jrock"
	fmt.Println(name2)

	fmt.Println(multiply(2, 2))

	// 여러개 다른 타입을 리턴하는 함수
	//totalLength, upperName := leAndUpper("jrock")
	// 두번째 값을 무시하려면 생략이 아닌 _ 로 작성해야한다.
	totalLength, _ := leAndUpper("jrock")
	//fmt.Println(totalLength, upperName)
	fmt.Println(totalLength)

	// 여러개의 아규먼트
	repeatMe("jung", "hong", "kim", "park")
}
