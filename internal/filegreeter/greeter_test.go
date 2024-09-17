package filegreeter

import (
	"bytes"
	"context"
	"testing"
)

func TestFileGreeter_Greet(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("jimmy")
	content, _ := readFile(&buffer)
	tests := []struct {
		name   string
		in     string
		expect string
	}{
		{
			name:   "Happy - Name matches file",
			in:     "jimmy",
			expect: "Hello, jimmy!",
		},
		{
			name:   "Sad - Name does not match file",
			in:     "jim",
			expect: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &FileGreeter{
				file:    &buffer,
				content: content,
			}
			if got := g.Greet(context.Background(), tt.in); got != tt.expect {
				t.Errorf("Greet() = %v, want %v", got, tt.expect)
			}
		})
	}
}
