package gradeklass


var (
	EmptySteelGradeCategory = NewSteelGradeCategory("")

	AgreementProduct = NewSteelGradeCategory("带出品")

	PicklePlate = NewSteelGradeCategory("酸洗基料")
	HighSurfacePicklePlate =   NewSteelGradeCategory("高表面酸洗基料")
	NormalSurfacePicklePlate =   NewSteelGradeCategory("普通表面酸洗基料")
	HighStrengthPicklePlate = NewSteelGradeCategory("高强酸洗基料")

	ColdBase = NewSteelGradeCategory("冷轧基料")
	HighStrengthColdBase = NewSteelGradeCategory("高强冷轧基料")
	NormalStrengthColdBase = NewSteelGradeCategory("普通冷轧基料")

	Silicon = NewSteelGradeCategory("硅钢")
	OrientedSilicon = NewSteelGradeCategory("取向硅钢")
	NonOrientedSilicon = NewSteelGradeCategory("无取向硅钢")
	HighGradeNonOrientedSilicon = NewSteelGradeCategory("高牌号无取向硅钢")
	MidLowGradeNonOrientedSilicon = NewSteelGradeCategory("中低牌号无取向硅钢")

	CommercialSteel = NewSteelGradeCategory("商品材")
	HighStrengthCommercialSteel = NewSteelGradeCategory("高强商品材")
	NormalStrengthCommercialSteel = NewSteelGradeCategory("普通商品材")

	AxleSteel = NewSteelGradeCategory("桥壳钢")
	BeamSteel = NewSteelGradeCategory("大梁钢")
	WheelSteel = NewSteelGradeCategory("车轮钢")
	PipelineSteel = NewSteelGradeCategory("管线钢")
	MechanicalMixingSteel = NewSteelGradeCategory("机械搅拌钢")

)

type SteelGradeCategory struct {
	name string
}

func (s *SteelGradeCategory) Name() string {
	return s.name
}

func (s *SteelGradeCategory) SetName(name string) {
	s.name = name
}

func NewSteelGradeCategory(name string) *SteelGradeCategory {
	c := &SteelGradeCategory{}
	c.SetName(name)
	return c
}