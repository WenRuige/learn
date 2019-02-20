package main

import "fmt"

type fileError struct {
	str string
}

func (fe *fileError) Error() string {
	return fe.str
}

func openFile() ([]byte, error) {
	return nil, &fileError{str: "error Fucking"}
}

func main() {

	str, err := openFile()
	fmt.Println(str, err.Error())

}
