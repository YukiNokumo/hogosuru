package htmlmapelement

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	baseobject.Eval(`m= document.createElement("map")`)
	m.Run()
}

func TestNew(t *testing.T) {

	if doc, err := document.New(); testingutils.AssertErr(t, err) {
		if b, err := New(doc); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "HTMLMapElement", b.ConstructName_())
		}

	}
}

func TestNewFromJSObject(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "m"); testingutils.AssertErr(t, err) {

		if b, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "HTMLMapElement", b.ConstructName_())
		}

	}
}

var getterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "Areas", "type": "constructnamechecking", "resultattempt": "HTMLCollection"},
	{"method": "Name", "resultattempt": ""},
}

func TestGetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "m"); testingutils.AssertErr(t, err) {

		if maphtml, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range getterAttempt {
				testingutils.InvokeCheck(t, maphtml, result)
			}

		}

	}
}

var setterAttempt []map[string]interface{} = []map[string]interface{}{
	{"method": "SetName", "args": []interface{}{"test"}, "gettermethod": "Name", "resultattempt": "test"},
}

func TestSetters(t *testing.T) {

	if obj, err := baseobject.Get(js.Global(), "m"); testingutils.AssertErr(t, err) {

		if maphtml, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			for _, result := range setterAttempt {

				testingutils.InvokeCheck(t, maphtml, result)

			}

		}

	}

}
