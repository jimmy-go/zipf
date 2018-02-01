package zipf

// ByCountAsc implement sort interface
type ByCountAsc []Term

func (a ByCountAsc) Len() int           { return len(a) }
func (a ByCountAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCountAsc) Less(i, j int) bool { return a[i].Count > a[j].Count }
