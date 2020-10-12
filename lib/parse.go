package markdown

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

const (
	text       = "text"
	tagHeading = "heading"
)

type element struct {
	isClosing bool
	parent    string
	kind      string
	tag       string
	open      string
	prev      string
	text      string
}

func (e *element) setHeading() (err error) {
	if e.kind == "" {
		e.kind = tagHeading
		e.tag = "h1"
		return
	}

	if e.kind != tagHeading {
		err = errors.New("this element is not a heading tag")
		return
	}

	level, err := strconv.Atoi(string(e.tag[1]))

	if level > 5 {
		err = errors.New("the heading is already at it's maximum depth")
		return
	}

	e.tag = fmt.Sprintf("h%d", level+1)
	return
}

// Markdown takes a markdown file and returns HTML
func Markdown(markdown string) (html string) {
	elems := elementStack{}
	elem := element{}
	var s scanner.Scanner

	s.Init(strings.NewReader(markdown))
	s.Whitespace = 1<<'\r' | 1<<'\t'

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if elem.kind != "text" {
			switch s.TokenText() {
			case "#":
				elem.setHeading()

			case " ":
				elems.push(elem)
				elem = element{}

			case "\n":
				ok, lastElem := elems.peek()
				var openingTag string
				var openingKind string

				if ok {
					if lastElem.kind == "text" {
						openingTag = lastElem.parent
					} else {
						openingTag = lastElem.tag
						openingKind = lastElem.kind
					}

					elems.push(element{isClosing: true, kind: openingKind, tag: openingTag})
					elem = element{}
				}

			default:
				ok, peek := elems.peek()
				elem.kind = "text"

				if ok {
					elem.parent = peek.tag
					elem.text += s.TokenText()
				}
			}
		} else {
			if s.TokenText() == "\n" || tok == scanner.EOF {
				ok, lastElem := elems.peek()
				openingKind := lastElem.kind
				openingTag := lastElem.tag

				if ok {
					elems.push(elem)
					elems.push(element{isClosing: true, tag: openingTag, kind: openingKind})

					elem = element{}
				}

			}

			elem.text += s.TokenText()

		}
	}

	for _, elem := range elems.getArray() {
		if elem.kind == "text" {
			html += elem.text
		} else if elem.isClosing {
			html += fmt.Sprintf("</%s>", elem.tag)

		} else {
			html += fmt.Sprintf("<%s>", elem.tag)
		}
	}

	return
}
