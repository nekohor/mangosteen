package schemas


var (
	Level2AssuringUnqualSource = NewUnqualSource("二级")
	FspUnqualSource = NewUnqualSource("FSP")
	QmsUnqualSource = NewUnqualSource("QMS")
	InspectionUnqualSource = NewUnqualSource("质检")
)
type UnqualSource struct {
	SourceName string
}

func NewUnqualSource(sourceName string) *UnqualSource {
	return &UnqualSource{SourceName: sourceName}
}

func (s *UnqualSource) Name() string {
	return s.SourceName
}


