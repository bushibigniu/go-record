package main

import "fmt"

func main()  {

	str := "cwn fbu8w vnwinkwji"
	var by [4]byte
	by[0] = 'a'
	by[1] = 'b'

	//string to byte
	b := []byte(str)
	fmt.Println(b)

	//byte to string
	s := string(by[:])
	fmt.Printf("%v",s)
	fmt.Printf("%s", s)

}