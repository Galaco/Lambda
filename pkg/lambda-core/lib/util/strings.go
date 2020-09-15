package util

// RemoveDuplicatesFromList
func RemoveDuplicatesFromList(list []string) (uniqueList []string) {
	for _, entry := range list {
		found := false
		for _, unique := range uniqueList {
			if entry == unique {
				found = true
				break
			}
		}
		if !found {
			uniqueList = append(uniqueList, entry)
		}
	}

	return uniqueList
}
