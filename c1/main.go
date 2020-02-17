package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello")

	reader := strings.NewReader("Example of io.SectionReader\n")
	SectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, SectionReader)
}
