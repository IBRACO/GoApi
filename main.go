package main

import "fmt"

var v string

func main() {

	v,_ = GenerateRsa()

	tok, err := GenerateJWT(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(tok)
	db = connexion()
	Roots()

}
