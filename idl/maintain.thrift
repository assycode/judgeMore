namespace go maintain
include "./model.thrift"

// 这个用于获取所有学院信息
struct QueryAllCollegeRequest{
      1: required i64 page_num,
      2: required i64 page_size,
}
struct QueryAllCollegeResponse{
     1: required model.BaseResp base,
     2: required model.CollegeList data,
}
// 获取学院的专业
struct QueryMajorByCollegeIdRequest{
    1: required i64 page_num,
    2: required i64 page_size,
    3: required i64 college_id,
}
struct QueryMajorByCollegeIdResponse{
     1: required model.BaseResp base,
     2: required model.MajorList data,
}
// 上传专业
struct UploadMajorRequest{
     1: required string major_name,
     2: required i64 college_id,
}
struct UploadMajorResponse{
     1: required model.BaseResp base,
     2: required i64 major_id,
}
// 上传学院
struct UploadCollegeRequest{
     1: required string college_name,
}
struct UploadCollegeResponse{
     1: required model.BaseResp base,
     2: required i64 college_id,
}


service maintainService{
     QueryAllCollegeResponse QueryCollege(1: QueryAllCollegeRequest req) (api.get = "/api/admin/colleges"),
     QueryMajorByCollegeIdResponse QueryMajorByCollegeId(1: QueryMajorByCollegeIdRequest req) (api.get = "/api/admin/majors"),
     UploadMajorResponse UploadMajor(1: UploadMajorRequest req) (api.post = "/api/admin/majors"),
     UploadCollegeResponse UploadCollege(1: UploadCollegeRequest req) (api.post = "/api/admin/colleges"),
}

