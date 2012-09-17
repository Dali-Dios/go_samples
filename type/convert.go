package main

import ("fmt")

func main() {
	// TODO code application logic here
    var a int
    var b uint32
    var c uint32

    a = 100
    b = 200
    //c = a + b  // Error: (mismatched types int and uint32)
    c = uint32(a) + b

    fmt.Printf("a + b = %d\n", c)

    var s1 = "Hello World!"
    s2 := []byte(s1)
    s2[4] = 'X'

    fmt.Printf("New string: %s\n", s2)

    d := []byte{65,66,67,68};
    s3 := string(d)
    fmt.Printf("s3: %s, d=%s\n", s3, d)
}