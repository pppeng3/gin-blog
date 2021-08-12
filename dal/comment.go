package dal

import (
	"Blog/db/model"
	"Blog/db/mysql"
	req "Blog/model"
	"context"

	"gorm.io/gorm/clause"
)

func CreateOrUpdateComment(ctx context.Context, comment model.Comment) (err error) {
	db := mysql.GetDB().WithContext(ctx)
	return db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&comment).Error
}

func GetComments(ctx context.Context, commentReq req.GetCommentReq, page, pageSize int) (comments []model.Comment, err error) {
	db := mysql.GetDB().WithContext(ctx)
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	err = db.Where(&commentReq).Offset(offset).Find(&comments).Limit(pageSize).Error
	return comments, err
}

func DeleteComment(ctx context.Context, commentReq req.DeleteCommentReq) (err error) {
	db := mysql.GetDB().WithContext(ctx)
	err = db.Delete(&commentReq).Error
	return err
}
