package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"monkey-language/evaluator"
	"monkey-language/lexer"
	"monkey-language/object"
	"monkey-language/parser"
	"os"
	"path/filepath"
)

const PROMPT = ">> "

// Wow, this is so cool!
const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())

			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const (
	fileExtension = ".mnk"
)

func StartFile(filename string) {
	out := os.Stdout
	env := object.NewEnvironment()

	if filepath.Ext(filename) != fileExtension {
		log.Fatalf("Invalid file extension. Expected %s", fileExtension)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.New(string(b))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
