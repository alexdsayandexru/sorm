package main

import (
	"fmt"
	"github.com/alexdsayandexru/sorm/internal/validator"
)

func main() {
	fmt.Println(validator.Validate("Привет мир - hello World").Regex("^[A-Za-zА-Яа-я -]+$").GetResult())
}
