package parse

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/lex8"
)

func parseDataStr(p *parser, args []*lex8.Token) ([]byte, uint32) {
	buf := new(bytes.Buffer)

	for _, arg := range args {
		if arg.Type != String {
			p.Errorf(arg.Pos, "expect string, got %s", p.TypeStr(arg.Type))
			return nil, 0
		}

		s, e := strconv.Unquote(arg.Lit)
		if e != nil {
			p.Errorf(arg.Pos, "invalid string %s", arg.Lit)
			return nil, 0
		}
		buf.Write([]byte(s))
	}

	return buf.Bytes(), 0
}
