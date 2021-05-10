package stringtable

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/stringtable"
)

// NewTable returns a new StringTable
func NewTable(stringData *lumps.TexDataStringData, stringTable *lumps.TexDataStringTable) *stringtable.StringTable {
	// Prepare texture lookup table
	return stringtable.NewFromExistingStringTableData(stringData.GetData(), stringTable.GetData())
}

// SortUnique builds a unique list of materials in a StringTable
// referenced by BSP TexInfo lump data.
func SortUnique(stringTable *stringtable.StringTable, texInfos *[]texinfo.TexInfo) []string {
	materialList := make([]string, 0)
	for _, ti := range *texInfos {
		target, _ := stringTable.FindString(int(ti.TexData))
		found := false
		for _, cur := range materialList {
			if cur == target {
				found = true
				break
			}
		}
		if !found {
			materialList = append(materialList, target)
		}
	}

	return materialList
}
