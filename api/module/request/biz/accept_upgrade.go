package reqbiz

import (
	"context"
	"shopbee/common"
)

func (biz *requestUpgradeBiz) AcceptRequestUpgrade(
	ctx context.Context,
	id int,
) error {
	if biz.requester.GetRole() != "admin" {
		return common.ErrNoPermission(nil)
	}

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{
		"user_id": id,
	})

	if err != nil {
		return common.ErrInternal(err)
	}

	if err := biz.store.AcceptRequestUpgrade(ctx, data); err != nil {
		return common.ErrInternal(nil)
	}

	return nil
}
