package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ericsolomon/pulsar/lexer"
	"github.com/ericsolomon/pulsar/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for t := lexer.NextToken(); t.Type != token.EOF; t = lexer.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}
