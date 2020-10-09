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

	e.tag = fmt.Sprintf("h%d", level)
	return
}

type elementStack struct {
	arr []element
}

func (s *elementStack) push(e element) {
	s.arr = append(s.arr, e)
}

func (s *elementStack) pop() (e element) {
	if len(s.arr) > 0 {
		l := len(s.arr)
		e = s.arr[l-1]
		s.arr = s.arr[0 : l-2]
	}

	return
}

func (s *elementStack) peek() (e element) {
	if len(s.arr) > 0 {
		l := len(s.arr)
		e = s.arr[l-1]
	}

	return
}

func encloseText(tag string, markdown string) (tagged string) {
	tagged = fmt.Sprintf("<%s>%s</%s>", tag, markdown, tag)
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

			case "\n":
				elems.push(element{isClosing: true})

				elem = element{}

			default:
				elems.push(elem)
				elem = element{}
				elem.kind = "text"
				elem.parent = elems.peek().tag
				elem.text += s.TokenText()
			}
		} else {
			elem.text += s.TokenText()

		}
	}

	html += encloseText(elem.tag, elem.text)
	return
}
