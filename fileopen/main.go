package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
	// os.Args = append(os.Args, )

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error opening file", err)
		os.Exit(1)
	}

	bs := make([]byte, 999)

	br, err := f.Read(bs)
	if err != nil {
		fmt.Println("error opening file", err)
	}
	// io.Copy(os.Stdout, f)
	fmt.Println(br, "bytes read", string(bs[:br]))
	f.Close()
}
