package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"judgeMore/biz/service/model"
	"judgeMore/pkg/constants"
	"judgeMore/pkg/errno"
)

func IsAppealExist(ctx context.Context, result_Id string) (bool, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("result_id = ?", result_Id).
		First(&appeal).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return true, nil
}
func IsAppealExistByAppealId(ctx context.Context, appeal_id string) (bool, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		First(&appeal).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return true, nil
}

// 该函数调用前检验存在性
func QueryAppealById(ctx context.Context, appeal_id string) (*model.Appeal, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		First(&appeal).
		Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return buildAppeal(appeal), nil
}
func QueryAppealByUserId(ctx context.Context, userId string) ([]*model.Appeal, int64, error) {
	var appeal []*Appeal
	var count int64
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("user_id = ?", userId).
		Find(&appeal).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return buildAppealList(appeal), count, nil
}
func CreateAppeal(ctx context.Context, a *model.Appeal) (string, error) {
	appeal := &Appeal{
		ResultId:       a.ResultId,
		UserId:         a.UserId,
		AppealReason:   a.AppealReason,
		AttachmentPath: a.AttachmentPath,
		AppealCount:    1, //默认为1
		AppealType:     a.AppealType,
		Status:         "pending",
	}
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Create(appeal).
		Error
	if err != nil {
		return "", errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to insert appeal: %v", err)
	}
	return appeal.AppealId, nil
}
func DeleteAppealById(ctx context.Context, appeal_id string) error {
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		Delete(&Appeal{}).
		Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to delete appeal: %v", err)
	}
	return nil
}

func buildAppeal(data *Appeal) *model.Appeal {
	r := &model.Appeal{
		ResultId:       data.ResultId,
		AppealId:       data.AppealId,
		UserId:         data.UserId,
		AppealReason:   data.AppealReason,
		AttachmentPath: data.AttachmentPath,
		AppealCount:    data.AppealCount,
		AppealType:     data.AppealType,
		Status:         data.Status,
		HandleResult:   data.HandledResult,
		HandledBy:      data.HandledBy,
		UpdateAT:       data.UpdatedAt.Unix(),
		CreateAT:       data.CreatedAt.Unix(),
		DeleteAT:       0,
	}
	if data.HandledAt == nil {
		r.HandledAt = 0
	} else {
		r.HandledAt = data.HandledAt.Unix()
	}
	return r
}

func buildAppealList(data []*Appeal) []*model.Appeal {
	result := make([]*model.Appeal, 0)
	for i := range data {
		result = append(result, buildAppeal(data[i]))
	}
	return result
}
