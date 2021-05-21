package schemas

var (
	ThkFspUnqualCategory = NewFspUnqualCategory("厚度")
	WidFspUnqualCategory = NewFspUnqualCategory("宽度")
	FdtFspUnqualCategory = NewFspUnqualCategory("FDT")
	CtFspUnqualCategory = NewFspUnqualCategory("CT")
	WdgFspUnqualCategory = NewFspUnqualCategory("楔形")
	CrnFspUnqualCategory = NewFspUnqualCategory("凸度")
	SymFltFspUnqualCategory = NewFspUnqualCategory("对称平直度")
	AsymFltFspUnqualCategory = NewFspUnqualCategory("非对称平直度")
)
type FspUnqualCategory struct {
	CategoryName string
}

func NewFspUnqualCategory(name string) *FspUnqualCategory {
	return &FspUnqualCategory{CategoryName: name}
}

func (s *FspUnqualCategory) Name() string {
	return s.CategoryName
}
