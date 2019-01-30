package keyvalues

// View contains an updatable list of keyvalue pairs
type View struct {
	keyValues []*keyValue
}

// Render draws each keyvalue pair
func (view *View) Render() {
	for _,kv := range view.keyValues {
		kv.Render()
	}
}

// AddKeyValue appends a new KeyValue to the list
func (view *View) AddKeyValue(value *keyValue) {
	view.keyValues = append(view.keyValues, value)
}

// NewKeyValues returns a new view
func NewKeyValues() *View {
	return &View{
		keyValues: make([]*keyValue, 0),
	}
}