package repl

import (
	"arthur/lexer"
	"arthur/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "\n>> "

func Init(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		items := lexer.Lex(line)
		item := <-items
		if item.Type == token.EOF {
			break
		}
		fmt.Fprintf(out, "%+v\n", item)
	}
}
