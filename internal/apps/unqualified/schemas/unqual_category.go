package schemas

var (
	ThkUnqualCategory = NewUnqualCategory("厚度")
	WidUnqualCategory = NewUnqualCategory("宽度")
	FdtUnqualCategory = NewUnqualCategory("终轧温度")
	CtUnqualCategory = NewUnqualCategory("卷取温度")
	ShapeUnqualCategory = NewUnqualCategory("板形")
	WdgUnqualCategory = NewUnqualCategory("楔形")
	CrnUnqualCategory = NewUnqualCategory("凸度")
	FlatnessUnqualCategory = NewUnqualCategory("平直度")
	SymFltUnqualCategory = NewUnqualCategory("对称平直度")
	AsymFltUnqualCategory = NewUnqualCategory("非对称平直度")
	LooperAngleUnqualCategory = NewUnqualCategory("活套角度")
	R2dtUnqualCategory = NewUnqualCategory("粗轧出口温度")
	CoilShapeUnqualCategory = NewUnqualCategory("卷形")
	RollBreakUnqualCategory = NewUnqualCategory("轧破甩尾")
)
type UnqualCategory struct {
	CategoryName string
}

func NewUnqualCategory(name string) *UnqualCategory {
	return &UnqualCategory{CategoryName: name}
}

func (s *UnqualCategory) Name() string {
	return s.CategoryName
}
