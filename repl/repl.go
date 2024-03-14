package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/VivekAlhat/Nova/lexer"
	"github.com/VivekAlhat/Nova/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Print(PROMPT)
		cmd := scanner.Scan()
		if !cmd {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			return
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
