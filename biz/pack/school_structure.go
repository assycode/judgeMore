package pack

import (
	resp "judgeMore/biz/model/model"
	"judgeMore/biz/service/model"
)

func College(data *model.College) *resp.College {
	return &resp.College{
		CollegeId:   data.CollegeId,
		CollegeName: data.CollegeName,
	}
}
func CollegeList(data []*model.College, total int64) *resp.CollegeList {
	r := make([]*resp.College, 0)
	for _, v := range data {
		r = append(r, College(v))
	}
	return &resp.CollegeList{
		Item:  r,
		Total: total,
	}
}

func Major(data *model.Major) *resp.Major {
	return &resp.Major{
		MajorId:   data.MajorId,
		MajorName: data.MajorName,
		CollegeId: data.CollegeId,
	}
}
func MajorList(data []*model.Major, total int64) *resp.MajorList {
	r := make([]*resp.Major, 0)
	for _, v := range data {
		r = append(r, Major(v))
	}
	return &resp.MajorList{
		Item:  r,
		Total: total,
	}
}
