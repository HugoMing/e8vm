package g8

import (
	"testing"

	"strings"

	"lonnie.io/e8vm/arch8"
)

func TestBareFunc_good(t *testing.T) {
	const N = 100000

	o := func(input, output string) {
		bs, es, _ := CompileBareFunc("main.g", input)
		if es != nil {
			t.Log(input)
			for _, e := range es {
				t.Log(e)
			}
			t.Error("compile failed")
			return
		}

		ncycle, out, e := arch8.RunImageOutput(bs, N)
		if ncycle == N {
			t.Log(input)
			t.Error("running out of time")
			return
		}
		if !arch8.IsHalt(e) {
			t.Log(input)
			t.Log(e)
			t.Error("did not halt gracefully")
			return
		}

		out = strings.TrimSpace(out)
		output = strings.TrimSpace(output)
		if out != output {
			t.Log(input)
			t.Log("expect: %s", output)
			t.Errorf("got: %s", out)
		}
	}

	o("printInt(0)", "0")
	o("printInt(3)", "3")
	o("printInt(-444)", "-444")
	o("printInt(2147483647)", "2147483647")
	o("printInt(-2147483647-1)", "-2147483648")
	o("printInt(300000000)", "300000000")
	o("printInt(4*5+3)", "23")
	o("printInt(3+4*5)", "23")
	o("printInt((3+4)*5)", "35")
	o("printInt((5*(3+4)))", "35")
	o("a:=3; if a==3 { printInt(5) }", "5")
	o("a:=5; if a==3 { printInt(5) }", "")
	o("a:=5; if a==3 { printInt(5) } else { printInt(10) }", "10")
	o("a:=3; for a>0 { printInt(a); a=a-1 }", "3\n2\n1")
	o("a:=0; for a<4 { printInt(a); a=a+1 }", "0\n1\n2\n3")
	o("a:=1; { a:=3; printInt(a) }", "3")
	o("true:=3; printInt(true)", "3")
	o("a,b:=3,4; printInt(a); printInt(b)", "3\n4")
	o("a,b:=3,4; { a,b:=b,a; printInt(a); printInt(b) }", "4\n3")
	o("var a int; printInt(a)", "0")
	o("var a,b = 3,4; printInt(a); printInt(b)", "3\n4")
	o("var a,b = 3,4; printInt(a); printInt(b)", "3\n4")
	o("var a,b int = 3,4; printInt(a); printInt(b)", "3\n4")
	o(` a,b:=3,4; { var a,b=b,a; printInt(a); printInt(b) }
	   	printInt(a); printInt(b)
	`, "4\n3\n3\n4")
	o("var i int; for i < 3 { printInt(i); i=i+1 }", "0\n1\n2")
	o("for true { break }; printInt(3)", "3")
	o("for true { if true break }; printInt(3)", "3")
	o("for { break }; printInt(33)", "33")
	o("i:=0; for i<3 { printInt(i); i=i+1; continue; break }", "0\n1\n2")
	o("printChar('x')", "x")
}

func TestBareFunc_bad(t *testing.T) {
	o := func(input string) {
		t.Log(input)
		_, es, _ := CompileBareFunc("main.g", input)
		if es == nil {
			t.Log(input)
			t.Error("should error")
			return
		}
	}

	o("a")                   // expression statement
	o("printInt")            // expression statement
	o("3+4")                 // expression statement
	o("a=3")                 // a not defined
	o("3=4")                 // non-addressable
	o("3=a")                 // non-addressable
	o("a:=3;a:=4")           // redefine
	o("printInt(true)")      // type mismatch
	o("printInt(3, 4)")      // arg count mismatch
	o("printInt()")          // arg count mismatch
	o("a := printInt(3, 4)") // mismatch
	o("a := 3, 4")           // count mismatch
	o("a, b := 3")           // count mismatch
	o("a, b := ()")          // invalid
	o("a()")                 // undefined function
	o("var a, b c")          // undefined type
	o("var a int; var b a")  // not a type
	o("break")               // not in for loop
	o("continue")            // not in for loop
	o("if true { break }")   // nothing to break
	o("if true break")       // nothing to break
}
