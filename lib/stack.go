package markdown

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
		s.arr = s.arr[0 : l-1]
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

func (s *elementStack) isEmpty() (e bool) {
	e = len(s.arr) == 0
	return
}

func (s *elementStack) getArray() (elems []element) {
	elems = s.arr
	return
}
