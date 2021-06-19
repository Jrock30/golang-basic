package person

import "fmt"

// Person
/**
대문자로 선언하면 다른 패키지에서 Import 할 수 있다.
필드도 마찬가지로 대문자로 하여야 하지만 receiver function 를 사용하면 대문자를 사용하지 않아도 된다.
*/
type Person struct {
	name string
	age  int
}

// SetDetails
/**
receiver function
  - 밖에서 사용할 것 이기 때문에 대문자로 시작.(export)
*Person 으로 포인터를 주면 복사가 되지 않고 main 에서의 jrock.SetDetails 를 그대로 수정이 가능하다.
*/
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	// p 는 이 method 를 부르면 person 을 부르는 variable 을 뜻함
	fmt.Println("SeeDetails' jrock:", p)
}

// Name
/**
복사본을 사용함. *Person 포인터 미사용
상황에 따라 복사본을 사용할 것인지 판단해서 사용할 것
*/
func (p Person) Name() string {
	return p.name
}
