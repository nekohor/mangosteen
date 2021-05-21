package schemas

var (
	HeadThkUnqualType = NewUnqualType("头厚不合")
	BodyThkUnqualType = NewUnqualType("身厚不合")
	ThkUnqualType = NewUnqualType("厚度不合")

	WidUnqualType = NewUnqualType("宽度不合")

	WdgUnqualType = NewUnqualType("楔形不合")
	CrnUnqualType = NewUnqualType("凸度不合")

	FdtUnqualType = NewUnqualType("FDT不合")
	CtUnqualType = NewUnqualType("CT不合")

	SymFltUnqualType = NewUnqualType("对称平直不合")
	AsymFltUnqualType = NewUnqualType("非对平直不合")

	CoilShapeUnqualType = NewUnqualType("卷形不合")
)
type UnqualType struct {
	TypeName string
}

func NewUnqualType(typeName string) *UnqualType {
	return &UnqualType{TypeName: typeName}
}

func (t *UnqualType) Name() string {
	return t.TypeName
}