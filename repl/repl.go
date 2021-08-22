package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = "Monkey> "

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprint(writer, PROMPT)
		line := scanner.Text()
		lexical := lexer.New(line)
		for tok := lexical.NextToken(); tok.Type != token.EOF; tok = lexical.NextToken() {
			fmt.Fprintf(writer, "%+v\n", tok)
		}
	}
}
