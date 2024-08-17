package datastructures

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (st *Stack) Push(value string) {
	st.items = append(st.items, value)
}

func (st *Stack) Pop() (string, error) {
	length := len(st.items)
	if length == 0 {
		return "", fmt.Errorf("no more items in the stack")
	}
	value := st.items[length-1]
	st.items = st.items[:length-1]
	return value, nil
}

func (st Stack) ToString(separator string) string {
	return strings.Join(st.items,separator)

}