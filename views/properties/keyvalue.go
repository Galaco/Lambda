package properties

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
)

type keyValue struct {
	isEditable bool
	Key        string
	Value      string
}

func (kv *keyValue) Render() {
	imgui.Text(fmt.Sprintf("%s : %s", kv.Key, kv.Value))
}

func newKeyValue(key string, value string, editable bool) keyValue {
	return keyValue{
		isEditable: editable,
		Key:        key,
		Value:      value,
	}
}
