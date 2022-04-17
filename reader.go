package main

//The io package specifies the io.Reader interface, which represents the read end of a stream of data
import (
	"fmt"
	"io"
	"strings"
)

/*
The io.Reader interface has a Read method:
func (T) Read(b []byte) (n int, err error)
*/

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {

		//Read populates the given byte slice with data and returns the number of bytes populated and an error value
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		//It returns an io.EOF error when the stream ends
		if err == io.EOF {
			break
		}
	}
}
