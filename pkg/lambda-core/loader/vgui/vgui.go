package vgui

import (
	"errors"
	"fmt"
	keyvalues "github.com/galaco/KeyValues"
	"github.com/galaco/Lambda/pkg/lambda-core/vgui"
	"github.com/golang-source-engine/filesystem"
	"log"
)

// LoadVGUI
func LoadVGUI(fs *filesystem.FileSystem, resourceName string) (*vgui.Panel, error) {
	stream, err := fs.GetFile(fmt.Sprintf("/resource/%s.res", resourceName))
	if err != nil {
		return nil, err
	}
	kvReader := keyvalues.NewReader(stream)
	kv, err := kvReader.Read()
	if err != nil {
		return nil, err
	}
	log.Println(kv.Type())
	if !kv.HasChildren() {
		return nil, errors.New("empty vgui resource keyvalues")
	}

	children, err := kv.Children()
	if err != nil {
		return nil, err
	}
	panel := recursiveBuildVGUITree(children[0], nil)

	return panel, nil
}

func recursiveBuildVGUITree(node *keyvalues.KeyValue, parent *vgui.Panel) *vgui.Panel {
	if node.HasChildren() {
		children, _ := node.Children()
		var isPanel bool
		// If a child also has children, then it should be a panel
		for _, c := range children {
			if c.HasChildren() {
				isPanel = true
			}
		}

		if isPanel {
			var p *vgui.Panel
			if parent == nil {
				// @TODO this needs proper construction rules
				parent = &vgui.Panel{}
				p = parent
			} else {
				p = parent.NewChildPanel(0, 0, 640, 480, true)
			}

			for _, c := range children {
				childPanel := recursiveBuildVGUITree(c, p)

				if childPanel != nil {
					p.AddChild(childPanel)
				}
			}
		} else {
			for _, c := range children {
				if c.Key() == "label" {
					if text, err := c.AsString(); err == nil {
						parent.AddElement(vgui.NewButton(text))
					}
				}
			}
			return nil
		}
	}
	return parent
}
