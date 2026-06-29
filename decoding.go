package errors

import (
	"github.com/valyala/fastjson"
)

func UnmarshalJSON(data []byte) ([]error, error) {
	if len(data) == 0 {
		return nil, nil
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	v := val.Get("errors")
	if v == nil {
		return nil, wrap(nil, "invalid errors array")
	}
	items := make([]error, 0)
	switch v.Type() {
	case fastjson.TypeArray:
		for _, v := range v.GetArray() {
			status := v.GetInt("status")
			localErr := errorFromStatus(status)
			if err := localErr.UnmarshalJSON([]byte(v.String())); err == nil {
				items = append(items, localErr)
			}
		}
		return items, err
	case fastjson.TypeObject:
		status := v.GetInt("status")
		localErr := errorFromStatus(status)
		if err := localErr.UnmarshalJSON([]byte(v.String())); err == nil {
			items = append(items, localErr)
		}
	case fastjson.TypeString:
		it := new(Err)
		it.m = string(data)
		items = append(items, it)
	}
	return items, nil
}

func (e *Err) UnmarshalJSON(data []byte) error {
	if m := fastjson.GetString(data, "message"); len(m) > 0 {
		e.m = m
	}
	return nil
}
