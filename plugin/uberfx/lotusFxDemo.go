package uberfx

import (
	"go.uber.org/fx"
	"reflect"
)

type Settings struct {
	modules map[interface{}]fx.Option
	invokes []fx.Option
}

type Option func(*Settings) error
type invoke int
type special struct{ id int }


func Options(opts ...Option) Option {
	return func(s *Settings) error {
		for _ , opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}
		return nil
	}
}

func Override(typ, constructor interface{}) Option {
	return func(s *Settings) error {
		if i, ok := typ.(invoke); ok {
			s.invokes[i] = fx.Invoke(constructor)
			return nil
		}
		if c, ok := typ.(special); ok {
			s.modules[c] = fx.Provide(constructor)
			return nil
		}
		ctor := as(constructor, typ)
		rt := reflect.TypeOf(typ).Elem()
		s.modules[rt] = fx.Provide(ctor)
		return nil
	}
}

func as(in interface{}, as interface{}) interface{} {
	outType := reflect.TypeOf(as)

	if outType.Kind() != reflect.Ptr {
		panic("outType is not a pointer")
	}

	if reflect.TypeOf(in).Kind() != reflect.Func {
		ctype := reflect.FuncOf(nil, []reflect.Type{outType.Elem()}, false)

		return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
			out := reflect.New(outType.Elem())
			out.Elem().Set(reflect.ValueOf(in))

			return []reflect.Value{out.Elem()}
		}).Interface()
	}

	inType := reflect.TypeOf(in)

	ins := make([]reflect.Type, inType.NumIn())
	outs := make([]reflect.Type, inType.NumOut())

	for i := range ins {
		ins[i] = inType.In(i)
	}
	outs[0] = outType.Elem()
	for i := range outs[1:] {
		outs[i+1] = inType.Out(i + 1)
	}

	ctype := reflect.FuncOf(ins, outs, false)

	return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
		outs := reflect.ValueOf(in).Call(args)

		out := reflect.New(outType.Elem())
		if outs[0].Type().AssignableTo(outType.Elem()) {
			// Out: Iface = In: *Struct; Out: Iface = In: OtherIface
			out.Elem().Set(outs[0])
		} else {
			// Out: Iface = &(In: Struct)
			t := reflect.New(outs[0].Type())
			t.Elem().Set(outs[0])
			out.Elem().Set(t)
		}
		outs[0] = out.Elem()

		return outs
	}).Interface()
}
