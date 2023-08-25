package userstorage

import "context"

type likeShop struct {
	UserId int `gorm:"column:user_id;"`
	ShopId int `gorm:"column:shop_id;"`
}

func (s *userMySql) LikeShop(
	ctx context.Context,
	userId int,
	shopId int,
) error {
	db := s.db.Table("shop_follows")

	like := likeShop{
		UserId: userId,
		ShopId: shopId,
	}

	if err := db.Create(&like).Error; err != nil {
		return err
	}

	return nil
}

func (s *userMySql) DislikeShop(
	ctx context.Context,
	userId int,
	shopId int,
) error {
	db := s.db.Table("shop_follows")

	like := likeShop{
		UserId: userId,
		ShopId: shopId,
	}

	if err := db.
		Where("user_id = ? AND shop_id = ?", userId, shopId).
		Delete(&like).Error; err != nil {
		return err
	}

	return nil
}

func (s *userMySql) IsLiked(
	ctx context.Context,
	userId int,
	shopId int,
) bool {
	db := s.db.Table("shop_follows")
	var like likeShop

	if err := db.
		Where("user_id = ? AND shop_id = ?", userId, shopId).
		First(&like).Error; err != nil {
		return false
	}

	return true
}

func (s *userMySql) CountLike(
	ctx context.Context,
	shopId int,
) int {
	db := s.db.Table("shop_follows")
	db = db.Where("shop_id = ?", shopId)

	var likeCount int64

	if err := db.Count(&likeCount).Error; err != nil {
		return 0
	}

	return int(likeCount)
}
