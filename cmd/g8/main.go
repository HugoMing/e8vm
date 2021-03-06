package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/dasm8"
	"lonnie.io/e8vm/g8"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
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

func main() {
	bare := flag.Bool("bare", false, "parse as bare function")
	parseAST := flag.Bool("parse", false, "parse only and print out the ast")
	ir := flag.Bool("ir", false, "prints out the IR")
	dasm := flag.Bool("d", false, "deassemble the image")
	ncycle := flag.Int("n", 100000, "maximum number of cycles")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		exit(errors.New("need exactly one input input file"))
	}
	fname := args[0]
	input, e := ioutil.ReadFile(fname)
	if e != nil {
		exit(e)
	}

	if *bare {
		if *parseAST {
			stmts, es := parse.Stmts(fname, bytes.NewBuffer(input))
			printErrs(es)
			ast.FprintStmts(os.Stdout, stmts)
		} else {
			bs, es, irLog := g8.CompileBareFunc(fname, string(input))
			printErrs(es)
			printIRLog(irLog, *ir)
			runImage(bs, *dasm, *ncycle)
		}
	} else {
		if *parseAST {
			f, es := parse.File(fname, bytes.NewBuffer(input))
			printErrs(es)
			ast.FprintFile(os.Stdout, f)
		} else {
			bs, es, irLog := g8.CompileSingleFile(fname, string(input))
			printErrs(es)
			printIRLog(irLog, *ir)
			runImage(bs, *dasm, *ncycle)
		}
	}
}

func runImage(bs []byte, dasm bool, n int) {
	if dasm {
		lines := dasm8.Dasm(bs, arch8.InitPC)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	if len(bs) == 0 {
		fmt.Println("(the image is empty)")
		return
	}

	ncycle, e := arch8.RunImage(bs, n)
	fmt.Printf("(%d cycles)\n", ncycle)
	if e != nil {
		fmt.Println(e)
	}
}

func printIRLog(irLog []byte, ir bool) {
	if !ir {
		return
	}
	if irLog == nil {
		fmt.Println("(no IR log produced)")
	} else {
		fmt.Println(string(irLog))
	}

}
