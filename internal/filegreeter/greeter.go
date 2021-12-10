package filegreeter

import (
	"bufio"
	"context"
	"io"
	"os"

	"github.com/Tommy647/go_example/internal/greeter"
)

// New returns a new instance of our file greeter
func New(path string) (*FileGreeter, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := readFile(file)
	if err != nil {
		return nil, err
	}

	return &FileGreeter{
		file:    file,
		content: content,
	}, nil

}

// FileGreeter our file greeter
type FileGreeter struct {
	file    io.Reader
	content []string
}

// readFile returns lines in a slice from and io.Reader
func readFile(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		readLine := scanner.Text()
		lines = append(lines, readLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Greet provides our hello request, checks the file to see
// if `in` exists, if so, uses it in the greeting, otherwise sets empty string
func (g *FileGreeter) Greet(ctx context.Context, in string) string {
	basicGreeter := greeter.New()

	for _, v := range g.content {
		if v == in {
			return basicGreeter.Greet(ctx, v)
		}
	}

	return basicGreeter.Greet(ctx, "")

}
