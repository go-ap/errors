package errors

import (
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

func UnmarshalJSON(data []byte) ([]Error, error) {
	if len(data) == 0 {
		return nil, nil
	}
	val, typ, _, err := jsonparser.Get(data, "errors")
	if err != nil {
		return nil, err
	}
	items := make([]Error, 0)
	switch typ {
	case jsonparser.Array:
		var err error
		jsonparser.ArrayEach(val, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			status, err := jsonparser.GetInt(value, "status")
			if err == nil {
				localErr := errorFromStatus(int(status))
				if err := localErr.UnmarshalJSON(value); err == nil {
					items = append(items, localErr)
				}
			}
		})
		return items, err
	case jsonparser.Object:
		if status, err := jsonparser.GetInt(data, "status"); err == nil {
			localErr := errorFromStatus(int(status))
			if err := localErr.UnmarshalJSON(data); err == nil {
				items = append(items, localErr)
			}
		}
	case jsonparser.String:
		it := new(Err)
		it.m = string(data)
		items = append(items, it)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (e *Err) UnmarshalJSON(data []byte) error {
	if m, err := jsonparser.GetString(data, "message"); err == nil {
		e.m = m
	}
	if val, err := jsonparser.GetString(data, "location"); err == nil {
		pieces := strings.Split(val, ":")
		if len(pieces) > 0 {
			e.f = pieces[0]
			if len(pieces) > 1 {
				l, _ := strconv.ParseInt(pieces[1], 10, 32)
				e.l = int(l)
			}
		}
	}

	return nil
}

func (n *notFound) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (m *methodNotAllowed) UnmarshalJSON(data []byte) error {
	return m.Err.UnmarshalJSON(data)
}
func (n *notValid) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (f *forbidden) UnmarshalJSON(data []byte) error {
	return f.Err.UnmarshalJSON(data)
}
func (n *notImplemented) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (b *badRequest) UnmarshalJSON(data []byte) error {
	return b.Err.UnmarshalJSON(data)
}
func (u *unauthorized) UnmarshalJSON(data []byte) error {
	return u.Err.UnmarshalJSON(data)
}
func (n *notSupported) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (t *timeout) UnmarshalJSON(data []byte) error {
	return t.Err.UnmarshalJSON(data)
}
func (b *badGateway) UnmarshalJSON(data []byte) error {
	return b.Err.UnmarshalJSON(data)
}
