package models

type QualityInspectionAccount struct {
	Id  uint
	CoilId  string

	//order_thick = Column(String(80))
	//order_width = Column(String(80))
	//steel_grade = Column(String(80))
	//next_process = Column(String(80))
	//act_weight = Column(String(80))
	//process_defect = Column(String(80))
	//process_defect_desc = Column(String(80))

	CoilDefect string
	CoilDefectDesc string
	SurfaceDefect string
	SurfaceDefectDesc string

	//treatment = Column(String(80))

	BlockState string

	//block_man = Column(String(80))
	//slab_grade = Column(String(80))
	//surface_feedback_grade = Column(String(80))
	//coil_quality_grade = Column(String(80))
	//shape_quality_grade = Column(String(80))
	//convertor_id = Column(String(80))
	//slab_id = Column(String(80))

	Month uint
}

func (*QualityInspectionAccount) TableName() string {
	return "quality_inspection_accounts"
}