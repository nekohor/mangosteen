package gradeklass

type GradeKlass struct {
}

func (k *GradeKlass) GetClassificationResultBySteelGrade(steelGrade string) *SteelGradeClassificationResult {
	res := &SteelGradeClassificationResult{}
	res.SteelGrade = steelGrade
	res.Category1 = k.GetCategory1(steelGrade).Name()
	res.Category2 = k.GetCategory2(steelGrade).Name()
	res.Category3 = k.GetCategory3(steelGrade).Name()

	return res
}

func (k *GradeKlass) GetCategory1(steelGrade string) *SteelGradeCategory {
	if k.IsPicklePlate(steelGrade) {
		return PicklePlate
	} else if k.IsSilicon(steelGrade) {
		return Silicon
	} else if k.IsColdBase(steelGrade) {
		return ColdBase
	} else {
		return CommercialSteel
	}
}

func (k *GradeKlass) GetCategory2(steelGrade string) *SteelGradeCategory {
	if k.IsPicklePlate(steelGrade) {
		if k.IsHighSurfacePicklePlate(steelGrade) {
			return HighSurfacePicklePlate
		} else {
			return NormalSurfacePicklePlate
		}
	} else if k.IsSilicon(steelGrade) {
		if k.IsOrientedSilicon(steelGrade) {
			return OrientedSilicon
		} else if k.IsHighGradeNonOrientedSilicon(steelGrade) {
			return HighGradeNonOrientedSilicon
		} else {
			return MidLowGradeNonOrientedSilicon
		}
	} else if k.IsColdBase(steelGrade) {
		if k.IsHighStrengthColdBase(steelGrade) {
			return HighStrengthColdBase
		} else {
			return NormalStrengthColdBase
		}
	} else {
		if k.IsHighStrengthCommercialSteel(steelGrade) {
			return HighStrengthCommercialSteel
		} else {
			return NormalStrengthCommercialSteel
		}
	}
}

func (k *GradeKlass) GetCategory3(steelGrade string) *SteelGradeCategory {
	if k.IsHighStrengthPicklePlate(steelGrade) {
		return HighStrengthPicklePlate
	}

	if k.IsAxleSteel(steelGrade) {
		return AxleSteel
	}

	if k.IsBeamSteel(steelGrade) {
		return BeamSteel
	}

	if k.IsWheelSteel(steelGrade) {
		return WheelSteel
	}

	if k.IsPipelineSteel(steelGrade) {
		return PipelineSteel
	}

	if k.IsMechanicalMixingSteel(steelGrade) {
		return MechanicalMixingSteel
	}

	return EmptySteelGradeCategory
}
