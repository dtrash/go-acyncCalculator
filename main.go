package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ops = map[string]func(x, y int) int{
		"+": func(x, y int) int { return x + y },
		"-": func(x, y int) int { return x - y },
		"*": func(x, y int) int { return x * y },
		"/": func(x, y int) int { return x / y },
	}
)

type task struct {
	id           int
	x, y, result int
	op           string
	err          error
}

func (t *task) calculate() {
	if f, ok := ops[t.op]; ok {
		t.result = f(t.x, t.y)
	} else {
		t.err = errors.New(fmt.Sprintf("Operation %q is not supported", t.op))
	}
}

func calcInRoutine(t *task, out chan *task) {
	t.calculate()
	out <- t
}

type byId []*task

func (s byId) Len() int {
	return len(s)
}

func (s byId) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s byId) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	slInput := strings.Split(input, ",")

	out := make(chan *task, len(slInput))
	for idx, v := range slInput {
		parts := strings.Fields(v)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[2])
		t := &task{
			id: idx,
			x:  x,
			y:  y,
			op: parts[1],
		}
		go calcInRoutine(t, out)
	}

	result := make([]*task, len(slInput))
	for idx := range slInput {
		result[idx] = <-out
	}

	sort.Sort(byId(result))

	for _, t := range result {
		if t.err == nil {
			fmt.Printf("%d %s %d = %d\n", t.x, t.op, t.y, t.result)
		} else {
			fmt.Println(t.err)
		}
	}
}
