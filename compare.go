package querystring

type DiffLine struct {
	Key      string
	Expected string
	Found    string
}

func Compare(reference, compared string, ignoredKeys []string) []DiffLine {
	diff := []DiffLine{}
	ref := CreateInstance("")
	ref.HydrateFromQueryString(reference)
	comp := CreateInstance("")
	comp.HydrateFromQueryString(compared)

	for key, _ := range ref.Data {
		if inSlice(ignoredKeys, key) == true {
			continue
		}
		if ref.Get(key) != comp.Get(key) {
			diff = append(diff, DiffLine{
				Key:      key,
				Expected: ref.Get(key),
				Found:    comp.Get(key),
			})
		}
	}

	return diff
}

func inSlice(a []string, match string) bool {
	for _, entry := range a {
		if entry == match {
			return true
		}
	}
	return false
}
