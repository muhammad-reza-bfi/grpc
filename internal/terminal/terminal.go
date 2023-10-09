package terminal

import (
	"bufio"
	"os"
)

// Terminal is ...
type Terminal struct {
	scanner *bufio.Scanner
}

// New ...
func New() *Terminal {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	return &Terminal{
		scanner: scanner,
	}
}

// ValText ...
func (t *Terminal) ValText() (string, bool) {
	if t.scanner.Scan() {
		return t.scanner.Text(), true
	}

	return "", false
}
