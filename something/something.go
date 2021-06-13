package something

import "fmt"

func sayBye() {
	fmt.Println("Bye")
}

// SayHello /**
/**
메서드에 대문자를 붙이면 다른 곳에서 가져다 쓸 수 있다.
export 가 가능해진다.
*/
func SayHello() {
	fmt.Println("Hello")
}
