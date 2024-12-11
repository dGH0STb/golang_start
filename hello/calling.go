package hello

import (
	"fmt"

	"github.com/golang"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
