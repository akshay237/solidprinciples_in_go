package main

import (
	"fmt"

	"github.com/akshay237/solidprinciples_in_go/isp"
	"github.com/akshay237/solidprinciples_in_go/lsp"
	"github.com/akshay237/solidprinciples_in_go/ocp"
	"github.com/akshay237/solidprinciples_in_go/srp"
)

func main() {
	fmt.Println("Run Program of Solid Principles")
	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
}
