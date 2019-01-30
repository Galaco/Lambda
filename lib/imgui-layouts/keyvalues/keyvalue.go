package keyvalues

import "github.com/inkyblackness/imgui-go"

type keyValue struct {
	key string
	value string
	onChange func(string, string)
}

// Render draws this keyvalue as 2 columns.
func (kv *keyValue) Render() {
	imgui.BeginColumns(2)
	imgui.Text(kv.key)
	imgui.NextColumn()
	imgui.Text(kv.value)
	imgui.NextColumn()
}

// SetKey sets this key
func (kv *keyValue) SetKey(key string) {
	kv.key = key
	kv.onChange(kv.key, kv.value)
}

// SetValues sets this keys associated value
func (kv *keyValue) SetValue(value string) {
	kv.value = value
	kv.onChange(kv.key, kv.value)
}

// NewKeyValue returns a new keyvalue pair.
func NewKeyValue(key string, value string, onChange func(string, string)) *keyValue {
	return &keyValue{
		key: key,
		value: value,
		onChange: onChange,
	}
}
