package stages

import (
	"fmt"
	"reflect"

	"github.com/vito/booklit"
	"github.com/vito/booklit/ast"
)

type Evaluate struct {
	Plugins []booklit.Plugin

	Result booklit.Content
}

func (eval *Evaluate) VisitString(str ast.String) {
	eval.Result = booklit.Append(eval.Result, booklit.String(str))
}

func (eval *Evaluate) VisitSequence(seq ast.Sequence) {
	for _, node := range seq {
		node.Visit(eval)
	}
}

func (eval *Evaluate) VisitParagraph(node ast.Paragraph) {
	newContent := eval.Result

	para := booklit.Paragraph{}
	for _, sentence := range node {
		eval.Result = nil

		sentence.Visit(eval)

		if eval.Result != nil {
			para = append(para, eval.Result)
		}
	}

	eval.Result = nil

	if len(para) == 0 {
		// paragraph resulted in no content (e.g. an invoke with no return value)
		return
	}

	if len(para) == 1 && !para[0].IsSentence() {
		// paragraph resulted in block content (e.g. a section)
		eval.Result = booklit.Append(newContent, para[0])
		return
	}

	eval.Result = booklit.Append(newContent, para)
}

func (eval *Evaluate) VisitInvoke(invoke ast.Invoke) {
	var method reflect.Value
	for _, p := range eval.Plugins {
		value := reflect.ValueOf(p)
		method = value.MethodByName(invoke.Method)
		if method.IsValid() {
			break
		}
	}

	if !method.IsValid() {
		panic(fmt.Errorf("undefined method: %s", invoke.Method))
	}

	argContent := make([]booklit.Content, len(invoke.Arguments))
	for i, arg := range invoke.Arguments {
		eval := &Evaluate{
			Plugins: eval.Plugins,
		}

		arg.Visit(eval)

		argContent[i] = eval.Result
	}

	argc := method.Type().NumIn()
	if method.Type().IsVariadic() {
		argc--

		if len(argContent) < argc {
			panic(fmt.Errorf("argument count mismatch for %s: given %d, need at least %d", invoke.Method, len(argContent), argc))
		}
	} else {
		if len(argContent) != argc {
			panic(fmt.Errorf("argument count mismatch for %s: given %d, need %d", invoke.Method, argc, len(argContent)))
		}
	}

	argv := make([]reflect.Value, argc)
	for i := 0; i < argc; i++ {
		t := method.Type().In(i)
		argv[i] = eval.convert(t, argContent[i])
	}

	if method.Type().IsVariadic() {
		variadic := argContent[argc:]
		variadicType := method.Type().In(argc)

		subType := variadicType.Elem()
		for _, arg := range variadic {
			argv = append(argv, eval.convert(subType, arg))
		}
	}

	result := method.Call(argv)
	switch len(result) {
	case 0:
		return
	case 1:
		last := result[0]
		switch v := last.Interface().(type) {
		case error:
			if v != nil {
				panic(v)
			}
		case booklit.Content:
			eval.Result = booklit.Append(eval.Result, v)
		default:
			panic(fmt.Errorf("unknown return type: %T", v))
		}
	case 2:
		first := result[0]
		switch v := first.Interface().(type) {
		case booklit.Content:
			eval.Result = booklit.Append(eval.Result, v)
		default:
			panic(fmt.Errorf("unknown first return type: %T", v))
		}

		last := result[1]
		switch v := last.Interface().(type) {
		case error:
			if v != nil {
				panic(v)
			}
		default:
			panic(fmt.Errorf("unknown second return type: %T", v))
		}
	default:
		panic(fmt.Errorf("expected 0-2 return values from %s, got %d", invoke.Method, len(result)))
	}
}

func (eval Evaluate) convert(to reflect.Type, content booklit.Content) reflect.Value {
	switch reflect.New(to).Interface().(type) {
	case *string:
		return reflect.ValueOf(content.String())
	case *booklit.Content:
		return reflect.ValueOf(content)
	default:
		panic(fmt.Errorf("unsupported argument type: %s", to))
	}
}
