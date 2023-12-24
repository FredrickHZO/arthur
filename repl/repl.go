package repl

import (
	"arthur/lexer"
	"arthur/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Init(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			return
		}

		l := lexer.NewLexer(line)

		for item := l.Tokenize(); item.Type != token.EOF; item = l.Tokenize() {
			fmt.Fprintf(out, "%+v\n", item)
		}
		fmt.Println()
	}
}
