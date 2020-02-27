package basic

// EveryPtr removes duplicates.
func EveryPtr() string {
	return `
// Every<FTYPE>Ptr returns true if supplied function returns logical true for every item in the list

func Every<FTYPE>Ptr(f func(*<TYPE>) bool, list []*<TYPE>) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
`
}
