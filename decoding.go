package errors

import (
	"github.com/buger/jsonparser"
)

func UnmarshalJSON(data []byte) ([]Err, error) {
	_, typ, _, err := jsonparser.Get(data)
	if err != nil {
		return nil, err
	}
	items := make([]Err, 0)
	switch typ {
	case jsonparser.Array:
		var err error
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			it := Err{}
			err = it.UnmarshalJSON(value)
			if err != nil {
				return
			}
			items = append(items, it)
		})
		if err != nil {
			return items, err
		}
	case jsonparser.Object:
		it := Err{}
		err = it.UnmarshalJSON(data)
		items = append(items, it)
	case jsonparser.String:
		it := Err{}
		it.m = string(data)
		items = append(items, it)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}

/*
	c error
	m string
	t []byte
	l int
	f string
}
 */
func (e *Err) UnmarshalJSON(data []byte) error {
	if m, err := jsonparser.GetString(data, "m"); err == nil {
		e.m = m
	}
	if f, err := jsonparser.GetString(data, "f"); err == nil {
		e.f = f
	}
	if l, err := jsonparser.GetInt(data, "l"); err == nil {
		e.l = int(l)
	}

	return nil
}
