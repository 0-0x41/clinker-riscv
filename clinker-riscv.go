package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/0-0x41/clinker-riscv/pkg/utils/utils"
	"os"
)
  
func main() {
	if len(os.Args) < 2 {
		utils.Fatal("Error in parameters!")
	}  

	fmt.Printf("%v\n", os.Args)

	filename := os.Args[1]
	contents, err := os.ReadFile(filename)

	// if error, the program terminates
	utils.Error(err)

	// println(len(contents))

	if !bytes.HasPrefix(contents, []byte("177ELF")) {
		utils.Fatal("it's not an ELF file!")
	}
}
=======
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Error in parameters!")
		os.Exit(1)
	} else {
		fmt.Printf("%v\n", os.Args)
	}
}
>>>>>>> 0e45cb62222928f18dc41c6bdb9a94dd3436f051
