package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/dasm8"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

func exit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	os.Exit(-1)
}

func printErrs(es []*lex8.Error) {
	if len(es) == 0 {
		return
	}

	for _, e := range es {
		fmt.Println(e)
	}
	exit(nil)
}

func irTest() {
	p := ir.NewPkg("_")
	f := p.NewFunc("main")
	b := f.NewBlock()

	v1 := f.NewTemp(4)

	b.Assign(v1, ir.Snum(3))
	b.Arith(v1, v1, "+", ir.Snum(4))

	lib := ir.BuildPkg(p)

	buf := new(bytes.Buffer)
	e := link8.LinkMain(lib, buf)
	if e != nil {
		exit(e)
	}

	lines := dasm8.Dasm(buf.Bytes(), arch8.InitPC)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	task := flag.String("task", "ir", "the testing task to do")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		exit(errors.New("need exactly one input input file"))
	}

	fname := args[0]

	fin, e := os.Open(fname)
	if e != nil {
		exit(e)
	}
	defer func() {
		e := fin.Close()
		if e != nil {
			exit(e)
		}
	}()

	switch *task {
	case "toks":
		toks, es := parse.Tokens(fname, fin)
		printErrs(es)
		for _, t := range toks {
			fmt.Printf("%s: %s %q\n", t.Pos, parse.TypeStr(t.Type), t.Lit)
		}
	case "exprs":
		exprs, es := parse.Exprs(fname, fin)
		printErrs(es)
		for i, expr := range exprs {
			fmt.Printf("%d> ", i+1)
			fmt.Println(parse.PrintExpr(expr))
		}
	case "stmts":
		stmts, es := parse.Stmts(fname, fin)
		printErrs(es)
		fmt.Print(parse.PrintStmts(stmts))
	case "ir":
		irTest()
	}
}
