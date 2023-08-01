package userstorage

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

func (s *userMySql) ListDataWithCondition(
	context context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.Role != "" {
			db = db.Where("role = ?", f.Role)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))
	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)

		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
