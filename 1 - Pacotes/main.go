package main

import (
	"fmt"
	"modulo/master"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da main")
	master.EscrevendoMaster()
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
