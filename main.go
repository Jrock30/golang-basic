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

// 여러개의 아규먼트를 받는 방법2
func repeatMe2(a ...int) int {
	var total int            // total := 0
	for _, item := range a { // index, item
		total += item
	}
	return total
}

func part1() {
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

	// 여러개의 아규먼트2
	result := repeatMe2(2, 3, 4, 5, 6, 7)
	fmt.Println(result)

	// print format
	name01 := "jrock ! ! ! ! ! ! Is my name"
	for _, letter := range name01 {
		//fmt.Println(string(letter)) // 바이트 코드가 나오기 때문에 string 으로 감싸준다
		//fmt.Printf("%o", letter) // octal noation -> 8진수
		//fmt.Printf("%b", letter) // binary -> 2진수
		//fmt.Printf("%d", letter) // digit -> 10진수
		//fmt.Printf("%x", letter) // hex -> 16진수
		//fmt.Printf("%U", letter) // unicode

		// Sprintf -> 단순히 콘솔에 print 해주는게 아니라 format 된 string 을 return 해준다
		xAsBinary := fmt.Sprintf("%b", letter)
		fmt.Println(xAsBinary)
	}
}

func part2() {
	// GO 의 배열은 길이가 한정되어 있다. element 개수를 미리 명시해야한다.
	firstName := [3]string{"kim", "park", "jung"}
	for _, name := range firstName {
		fmt.Println(name)
	}
	for i := 0; i < len(firstName); i++ {
		fmt.Println(firstName[i])
	}

	// 배열은 길이가 한정되어 있지만 slice 는 한정 되어 있지 않다. (array 를 만들고 공간이 필요하면 더 만들어줌)
	firstName2 := []string{"kim", "park", "jung"}
	fmt.Printf("%v\n", firstName2)
	// array 에 사용할 수 없고 slice 에서만 사용가능
	firstName2 = append(firstName2, "gyu")
	// 복사를 원한다면
	//firstName3 := append(firstName2, "gyu")
	fmt.Printf("%v\n", firstName2)
}

/**
Pointer
*/
func part3() {
	a := 2
	b := a
	/**
	어떤 변수에든 &를 더해주면 memory address 를 얻을 수 있다.
	복사한 데이터 이기 때문에 주소값이 다르다.
	*/
	fmt.Println(&b, &a)
	b = 4
	fmt.Println(a, b)

	// 이렇게 하면 메모리 주소가 저장 된다.
	c := 3
	d := &c
	fmt.Println(d, &c)
	c = 6
	// * 를 붙이면 주소 값의 value 를 보여준다
	fmt.Println(c, *d)
}

func part4() {
	/**
	struct 생성
	name, age 는 생략 가능하다.(순서는 지켜야한다)
	*/
	jrock := person{name: "jrock", age: 33}
	// struct 메소드 콜
	jrock.sayHello()
}

/**
struct
class 와 유사하다고 보자
struct 의 장점은 receiver 함수다, Go 에서의 메소드 같은 것
*/
type person struct {
	name string
	age  int
}

/**
receiver 함수
struct 안에 메소드를 생성하는 것이 아닌 함수 인자에 struct 를 주입? 한다.
이렇게 하면 GO 는 이 것은 person struct 메소드라는 것을 알게 된다.
*/
func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
}

/**
main function
  - 시작점 entry point
*/
func main() {
	// const, var, func
	part1()
	// array
	part2()
	// pointer
	part3()
	// struct
	part4()
}
