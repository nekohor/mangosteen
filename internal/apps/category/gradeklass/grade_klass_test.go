package gradeklass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetClassificationResultBySteelGrade(t *testing.T) {
	var res *SteelGradeClassificationResult
	var k *GradeKlass
	k = &GradeKlass{}
	res = k.GetClassificationResultBySteelGrade("M610L陈丹")
	assert.Equal(t, "商品材", res.Category1)
	assert.Equal(t, "高强商品材", res.Category2)
	assert.Equal(t, "大梁钢", res.Category3)

	res = k.GetClassificationResultBySteelGrade("MGW350D邵健")
	assert.Equal(t, "硅钢", res.Category1)
	assert.Equal(t, "高牌号无取向硅钢", res.Category2)

	res = k.GetClassificationResultBySteelGrade("M510L-张雅倩-P")
	assert.Equal(t, "酸洗基料", res.Category1)
	assert.Equal(t, "普通表面酸洗基料", res.Category2)
	assert.Equal(t, "高强酸洗基料", res.Category3)

	res = k.GetClassificationResultBySteelGrade("X42李轶伦")
	assert.Equal(t, "商品材", res.Category1)
	assert.Equal(t, "高强商品材", res.Category2)
	assert.Equal(t, "管线钢", res.Category3)

	res = k.GetClassificationResultBySteelGrade("SPHC-酸洗板-MZ")
	assert.Equal(t, "酸洗基料", res.Category1)
	assert.Equal(t, "高表面酸洗基料", res.Category2)
	assert.Equal(t, "", res.Category3)
}