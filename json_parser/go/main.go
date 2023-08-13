package json_parser_recursive

// This package uses recursive approach to implement json parser, by making use of states.

import "fmt"

type Phase int // new custom type Phase is dependant on underlying type int
// States at different points in input string for json parser
const (
	OPEN  Phase = iota // 0
	KEY                // 1
	VALUE              // 2
	CLOSE              // 3
)

type JsonParser struct {
	index int
}

func main() {
	parser := JsonParser{}
	input := `{"name":"Ishan", "age": 20, "city" : "new york"}`
	res, err := parser.ParseJson(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}

func (p *JsonParser) ParseJson(input string) (res map[string]interface{}, err error) {
	p.index = 0
	res, err = p.recursive(input)
	if err != nil {
		return nil, err
	}

	if p.index != len(input) {
		return nil, fmt.Errorf("cannot traverse all chars and stop at %d/%d", p.index, len(input))
	}

	return res, nil
}

func (p *JsonParser) recursive(input string) (res map[string]interface{}, err error) {

}

// there are two util methods which will be required by recursive method repeatedly - removeWhiteSpaces and extractString from key
