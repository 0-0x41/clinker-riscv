package utils

import (
	"fmt"
	"os"
)

func Fatal(v any) {
	fmt.Printf("rvld: \033[0;1;31mfatal:\033[0m %v\n", v)
	os.Exit(1)
}

func Error(err error) {
	if err != nil {
		Fatal(err)
	}
}