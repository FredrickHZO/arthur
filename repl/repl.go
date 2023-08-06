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
		l := lexer.NewLexer(line)

		for item := l.NextToken(); item.Type != token.EOF; item = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", item)
		}
		fmt.Println()
	}
}
