package json

// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/JSON

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var jsoninterface js.Value

//Json  struct
type Json struct {
	object.Object
}

type JsonFrom interface {
	Json_() Json
}

func (i Json) Json_() Json {
	return i
}

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if jsoninterface, err = js.Global().GetWithErr("JSON"); err != nil {
			jsoninterface = js.Null()
		}
		baseobject.Register(jsoninterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return jsoninterface
}

func Parse(data string) (Json, error) {

	var jsonObject js.Value
	var err error
	if jsoni := GetInterface(); !jsoni.IsNull() {

		if jsonObject, err = jsoni.CallWithErr("parse", data); err != nil {
			return Json{}, err
		} else {
			return NewFromJSObject(jsonObject)
		}

	} else {
		err = ErrNotImplemented
	}

	return Json{}, err
}

func NewFromJSObject(obj js.Value) (Json, error) {
	var j Json

	if jsoni := GetInterface(); !jsoni.IsNull() {

		j.BaseObject = j.SetObject(obj)
		return j, nil

	}

	return j, ErrNotAJson

}

func extractJsonFromObject(jsobj js.Value) interface{} {
	var retvalue interface{}

	if obj, err := object.NewFromJSObject(jsobj); err == nil {

		if ok, err := array.IsArray(obj.BaseObject); ok && err == nil {
			var array []interface{}
			keys, _ := obj.Values()
			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if obj1, ok := vkey.(baseobject.ObjectFrom); !ok {
					array = append(array, vkey)

				} else {
					array = append(array, extractJsonFromObject(obj1.JSObject()))
				}

			}

			retvalue = array

		} else {

			json := make(map[string]interface{})
			keys, _ := obj.Keys()

			itkeys, _ := keys.Values()

			for _, vkey, err := itkeys.Next(); err == nil; _, vkey, err = itkeys.Next() {

				if key, ok := vkey.(string); ok {
					if value, err := jsobj.GetWithErr(key); err == nil {
						i := baseobject.GoValue(value)
						if obj1, ok := i.(baseobject.ObjectFrom); !ok {
							json[key] = i
						} else {
							json[key] = extractJsonFromObject(obj1.JSObject())

						}

					}

				}

			}
			retvalue = json
		}

	}

	return retvalue

}

func (j Json) Map() interface{} {

	return extractJsonFromObject(j.JSObject())

}

func (j Json) Stringify() (string, error) {

	var stringObject js.Value
	var err error
	if jsoni := GetInterface(); !jsoni.IsNull() {

		if stringObject, err = jsoni.CallWithErr("stringify", j.JSObject()); err != nil {
			return "", err
		} else {

			return stringObject.String(), nil
		}

	} else {
		err = ErrNotImplemented
	}

	return "", err

}
