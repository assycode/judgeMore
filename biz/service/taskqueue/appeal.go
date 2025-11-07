package taskqueue

import (
	"context"
	"fmt"
	"judgeMore/biz/dal/mysql"
	"judgeMore/pkg/errno"
	"judgeMore/pkg/taskqueue"
)

func AddAppealToScoreTask(ctx context.Context, key, result_id, appeal_id, status string) {
	taskQueue.Add(key, taskqueue.QueueTask{Execute: func() error {
		return updateEventAppeal(ctx, result_id, appeal_id, status)
	}})
}
func updateEventAppeal(ctx context.Context, result_id, appeal_id, status string) error {
	err := mysql.UpdateResultAppealInfo(ctx, result_id, appeal_id, status)
	if err != nil {
		fmt.Printf("update appeal info failed, %s\n", err.Error())
		return errno.NewErrNo(errno.InternalDatabaseErrorCode,
			fmt.Sprintf("updateEventAppeal: failed: %v", err))
	}
	return nil
}
