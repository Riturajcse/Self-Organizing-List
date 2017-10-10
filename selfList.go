package selfOrganizingList

type List interface {
	Prepend(int)
	Append(int)
	Find(int) int
	PrintList()
}

type Strategy int

const (
	MTF       = Strategy(0)
	Transpose = Strategy(1)
	Counter   = Strategy(2)
)

func NewList(str Strategy) List {
	switch str {
	case MTF:
		return &mtfList{}
	case Transpose:
		return &transposeList{}
	case Counter:
		return &counterList{}
	}
	return nil
}
