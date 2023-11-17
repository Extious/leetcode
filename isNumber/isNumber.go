package isNumber

import (
	"errors"
	"strings"
)

// number ::= decimal ['e'|'E' integer]
// decimal ::= ['-'|'+'] (normalFloat | dotFloat)
// integer ::= ['-'|'+'] x+
// normalFloat ::= xxx.[xxx]
// dotFloat ::= .x+

type NumberParser struct {
	Source string
	Pos    int
	Len    int
}

func New(s string) *NumberParser {
	return &NumberParser{
		Source: s,
		Pos:    0,
		Len:    len(s),
	}
}

func (p NumberParser) AtSign() bool {
	c := p.Source[p.Pos]
	return c == '-' || c == '+'
}

func (p NumberParser) AtDigit() bool {
	c := p.Source[p.Pos]
	return c >= '0' && c <= '9'
}

func (p NumberParser) AtDot() bool {
	return p.Source[p.Pos] == '.'
}
func (p NumberParser) AtE() bool {
	c := p.Source[p.Pos]
	return c == 'e' || c == 'E'
}

func (p NumberParser) IsEnd() bool {
	return p.Pos >= p.Len
}

func (p *NumberParser) Advance() {
	p.Pos++
}

// number ::= decimal ['e'|'E' integer]
func (p *NumberParser) ParseNumber() error {
	return p.ParseDecimal()
}

// decimal ::= ['-'|'+'] (normalFloat | dotFloat)
func (p *NumberParser) ParseDecimal() error {
	// eat '-'/ '+'
	if !p.IsEnd() && p.AtSign() {
		p.Advance()
	}

	if !p.IsEnd() {
		if p.AtDigit() {
			return p.ParseNormalFloat()
		} else if p.AtDot() {
			return p.ParseDotFloat()
		} else {
			return errors.New("unkown symbol at decimal")
		}
	}
	return errors.New("null number")
}

// normalFloat ::= xxx.[xxx]
func (p *NumberParser) ParseNormalFloat() error {
	// eat xxx
	for !p.IsEnd() && p.AtDigit() {
		p.Advance()
	}

	// eat .xxxx
	if !p.IsEnd() && p.AtDot() {
		p.Advance() // eat .

		// eat xxx
		for !p.IsEnd() && p.AtDigit() {
			p.Advance()
		}
	}

	// parse power
	if !p.IsEnd() {
		if p.AtE() {
			p.Advance() // eat 'e'
			return p.ParseInteger()
		} else {
			return errors.New("unkown symbol after decimal")
		}
	}

	return nil
}

// dotFloat ::= .x+
func (p *NumberParser) ParseDotFloat() error {
	// eat .
	if !p.IsEnd() && p.AtDot() {
		p.Advance()
	}

	if !p.IsEnd() && p.AtDigit() {
		p.Advance()
	} else {
		return errors.New("null dot")
	}

	// eat xxx
	for !p.IsEnd() && p.AtDigit() {
		p.Advance()
	}

	// parse power
	if !p.IsEnd() {
		if p.AtE() {
			p.Advance() // eat 'e'
			return p.ParseInteger()
		} else {
			return errors.New("unkown symbol after decimal")
		}
	}

	return nil
}

// integer ::= ['-'|'+'] x+
func (p *NumberParser) ParseInteger() error {
	// eat '-'/ '+'
	if !p.IsEnd() && p.AtSign() {
		p.Advance()
	}

	if !p.IsEnd() && p.AtDigit() {
		p.Advance()
	} else {
		return errors.New("null power")
	}

	// eat xxx
	for !p.IsEnd() {
		if !p.AtDigit() {
			return errors.New("unkown symbol at power")
		}
		p.Advance()
	}

	return nil
}

func isNumber(s string) bool {
	parser := New(strings.TrimSpace(s))

	err := parser.ParseNumber()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return err == nil
}
