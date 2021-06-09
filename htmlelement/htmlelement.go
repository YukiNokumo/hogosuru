package htmlelement

// https://developer.mozilla.org/fr/docs/Web/API/HTMLElement

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/element"
)

var singleton sync.Once

var htmlelementinterface js.Value

//HtmlInputElement struct
type HtmlElement struct {
	element.Element
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmlelementinterface, err = js.Global().GetWithErr("HTMLElement"); err != nil {
			htmlelementinterface = js.Null()
		}

	})

	return htmlelementinterface
}

func NewFromJSObject(obj js.Value) (HtmlElement, error) {
	var h HtmlElement
	var err error
	if ai := GetInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			h.BaseObject = h.SetObject(obj)

		} else {
			err = ErrNotAnHtmlElement
		}

	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromElement(elem element.Element) (HtmlElement, error) {
	var h HtmlElement
	var err error

	if ai := GetInterface(); !ai.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(ai) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}
func (h HtmlElement) getAttributeString(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {

		valueStr = obj.String()
	}

	return valueStr, err

}

func (h HtmlElement) getAttributeBool(attribute string) (bool, error) {

	var err error
	var obj js.Value
	var ret bool

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeBoolean {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrObjectNotBool
		}
	}

	return ret, err
}

func (h HtmlElement) getAttributeInt(attribute string) (int, error) {

	var err error
	var obj js.Value
	var result int

	if obj, err = h.JSObject().GetWithErr(attribute); err == nil {
		if obj.Type() == js.TypeNumber {
			result = obj.Int()
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}

	return result, err
}

func (h HtmlElement) setAttributeBool(attribute string, value bool) error {

	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (h HtmlElement) setAttributeString(attribute string, value string) error {

	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (h HtmlElement) setAttributeInt(attribute string, value int) error {

	return h.JSObject().SetWithErr(attribute, js.ValueOf(value))
}

func (h HtmlElement) AccessKey() (string, error) {
	return h.getAttributeString("accessKey")
}

func (h HtmlElement) SetAccessKey(value string) error {
	return h.setAttributeString("accessKey", value)
}

func (h HtmlElement) AccessKeyLabel() (string, error) {
	return h.getAttributeString("accessKeyLabel")
}

func (h HtmlElement) SetAccessKeyLabel(value string) error {
	return h.setAttributeString("accessKeyLabel", value)
}

func (h HtmlElement) ClassName() (string, error) {
	return h.getAttributeString("accessKeyLabel")
}

func (h HtmlElement) SetClassName(classname string) error {
	return h.setAttributeString("accessKeyLabel", classname)
}

func (h HtmlElement) ContentEditable() (string, error) {
	return h.getAttributeString("contentEditable")
}

func (h HtmlElement) SetContentEditable(value string) error {
	return h.setAttributeString("contentEditable", value)
}

func (h HtmlElement) IsContentEditable() (bool, error) {
	return h.getAttributeBool("isContentEditable")
}

func (h HtmlElement) Dataset(name string) (interface{}, error) {
	var err error
	var obj, objv js.Value
	var ret interface{}

	if obj, err = h.JSObject().GetWithErr("dataset"); err == nil {
		if objv, err = obj.GetWithErr(name); err == nil {
			ret = baseobject.GoValue(objv)
		}

	}

	return ret, err

}

func (h HtmlElement) SetDataset(name string, value interface{}) error {

	var err error
	var obj js.Value

	if obj, err = h.JSObject().GetWithErr("dataset"); err == nil {
		err = obj.SetWithErr(name, js.ValueOf(value))

	}
	return err
}

func (h HtmlElement) Dir() (string, error) {
	return h.getAttributeString("dir")
}

func (h HtmlElement) Hidden() (bool, error) {
	return h.getAttributeBool("hidden")
}

func (h HtmlElement) SetHidden(value bool) error {
	return h.setAttributeBool("hidden", value)
}

func (h HtmlElement) SetDir(value string) error {
	return h.setAttributeString("dir", value)
}

func (h HtmlElement) Id() (string, error) {
	return h.getAttributeString("id")
}

func (h HtmlElement) SetId(value string) error {
	return h.setAttributeString("id", value)
}

func (h HtmlElement) Lang() (string, error) {
	return h.getAttributeString("lang")
}

func (h HtmlElement) SetLang(value string) error {
	return h.setAttributeString("lang", value)
}

func (h HtmlElement) OffsetHeight() (int, error) {
	return h.getAttributeInt("offsetHeight")
}

func (h HtmlElement) SetOffsetHeight(value int) error {
	return h.setAttributeInt("offsetHeight", value)
}

func (h HtmlElement) OffsetLeft() (int, error) {
	return h.getAttributeInt("offsetLeft")
}

func (h HtmlElement) SetOffsetLeft(value int) error {
	return h.setAttributeInt("offsetLeft", value)
}

func (h HtmlElement) OffsetParent() (baseobject.BaseObject, error) {
	var err error
	var obj js.Value
	var ret baseobject.BaseObject

	if obj, err = h.JSObject().GetWithErr("offsetParent"); err == nil {
		if !obj.IsNull() {
			ret, err = baseobject.NewFromJSObject(obj)
		} else {
			err = baseobject.ErrNotAnObject
		}

	}
	return ret, err
}

func (h HtmlElement) OffsetTop() (int, error) {
	return h.getAttributeInt("offsetTop")
}

func (h HtmlElement) SetOffsetTop(value int) error {
	return h.setAttributeInt("offsetTop", value)
}

func (h HtmlElement) OffsetWidth() (int, error) {
	return h.getAttributeInt("offsetWidth")
}

func (h HtmlElement) SetOffsetWidth(value int) error {
	return h.setAttributeInt("offsetWidth", value)
}

func (h HtmlElement) Title() (string, error) {
	return h.getAttributeString("title")
}

func (h HtmlElement) SetTitle(value string) error {
	return h.setAttributeString("title", value)
}

func (h HtmlElement) Blur() error {
	_, err := h.BaseObject.JSObject().CallWithErr("blur")
	return err
}

func (h HtmlElement) Click() error {
	_, err := h.BaseObject.JSObject().CallWithErr("click")
	return err
}

func (h HtmlElement) Focus() error {
	_, err := h.BaseObject.JSObject().CallWithErr("focus")
	return err
}
