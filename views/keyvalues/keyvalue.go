package keyvalues

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
)

type keyvalue struct {
	isEditable bool
	Key string
	Value string
}

func (kv *keyvalue) Render() {
	imgui.Text(fmt.Sprintf("%s : %s", kv.Key, kv.Value))
}

func newKeyValue(key string, value string, editable bool) keyvalue{
	return keyvalue{
		isEditable:editable,
		Key:key,
		Value:value,
	}
}
