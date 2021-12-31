package filegreeter

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"github.com/Tommy647/go_example/internal/greeter"

	"github.com/Tommy647/go_example/internal/logger"
	"go.uber.org/zap"
)

type fileName struct {
	hello  string
	coffee string
}

func New() *fileName {
	return &fileName{
		hello:  "greet.txt",
		coffee: "coffeegreet.txt",
	}
}

type FileGreeter struct {
	file    io.Reader
	content []string
	fileName
}

func (fG *FileGreeter) Greet(ctx context.Context) {
	logger.Info(ctx, "file greet called", zap.String("file", "file name"))
	basicGreeter := greeter.BasicGreeter{}
	file, err := readFile(fG.hello)
	if err != nil {
		return
	}

}

func readFile(file string) ([]string, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening the file to greet")
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err(); err != nil {
		return nil, err
	}
	return lines, err
}
