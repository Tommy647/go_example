package filegreeter

import (
	"bufio"
	"context"
	"github.com/Tommy647/go_example/internal/greeter"
	"io"
	"log"
	"os"
)

// New returns a new instance of our file greeter
func New(path string) *FileGreeter {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic("Error opening file" + err.Error())
	}
	content := readFile(file)
	return &FileGreeter{
		file:    file,
		content: content,
	}

}

// FileGreeter our file greeter
type FileGreeter struct {
	file    io.Reader
	content []string
}

// readFile returns lines in a slice from and io.Reader
func readFile(reader io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		readLine := scanner.Text()
		lines = append(lines, readLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// Greet provides our hello request, checks the file to see
// if `in` exists, if so, uses it in the greeting, otherwise sets empty string
func (g *FileGreeter) Greet(ctx context.Context, in string) string {
	basicGreeter := greeter.New()
	fileContent := g.content

	for _, v := range fileContent {
		if v == in {
			return basicGreeter.Greet(ctx, v)
		}
	}

	return basicGreeter.Greet(ctx, "")

}
