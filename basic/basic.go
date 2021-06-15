package basic

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
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 리턴 값을 초기화 하고 그 값을 수정 후 자동 리턴 (naked return) return 을 굳이 명시하지 않아도 됨. 가독성이 좋지 않을 수 있다.
func lenAndUpper2(name string) (length int, uppercase string) {
	// function 의 실행이 끝나고 defer 가 실행된다.
	defer fmt.Println("I'm done")
	fmt.Println("===TEST 이게 위의 defer 보다 먼저 실행===")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
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

// loop
func supperAdd(numbers ...int) int {
	// 조건에 따라 반복 실행을 함 forEach
	total := 0
	for _, number := range numbers {
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		//fmt.Println(numbers[i])
	}
	return total
}

// if else
func canIDring(age int) bool {
	// if 조건 전에 variable 을 만들 수 있다. 그 후 바로 사용
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return true
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
	totalLength, _ := lenAndUpper("jrock")
	//fmt.Println(totalLength, upperName)
	fmt.Println(totalLength)

	// naked return
	length, uppercase := lenAndUpper2("jrock")
	fmt.Println(length, uppercase)

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
	favFood := []string{"kimchi", "ramen"}
	jrock := person{name: "jrock", age: 33, favFood: favFood}
	// struct 메소드 콜
	jrock.sayHello()
}

// loop
func part5() {
	total := supperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}

// if else switch
func part6() {
	fmt.Println(canIDring(16))
}

// map
func part7() {
	// map[key type]value type
	jrock := map[string]string{"name": "jrock", "age": "20"}
	for key, value := range jrock {
		fmt.Println(key, value)
	}
}

/**
struct
class 와 유사하다고 보자
struct 의 장점은 receiver 함수다, Go 에서의 메소드 같은 것
*/
type person struct {
	name    string
	age     int
	favFood []string
}

/**
receiver 함수
struct 안에 메소드를 생성하는 것이 아닌 함수 인자에 struct 를 주입? 한다.
이렇게 하면 GO 는 이 것은 person struct 메소드라는 것을 알게 된다.
*/
func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
	fmt.Println(p.favFood)
}

// Init
// 다른 곳에서 부르려면 대문자 입력
func Init() {
	// const, var, func, defer, naked return
	//part1()

	// array
	//part2()

	// pointer
	//part3()

	// struct
	//part4()

	// loop
	//part5()

	// if else, switch
	//part6()

	// map, object
	part7()
}
