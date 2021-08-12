package dal

import (
	"Blog/db/model"
	"Blog/db/mysql"
	req "Blog/model"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateBlog(ctx context.Context, blog model.Blog) error {
	db := mysql.GetDB().WithContext(ctx)
	return db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&blog).Error
}

func GetBlogs(ctx context.Context, blogReq req.GetBlogReq, page, pageSize int) (blogs []model.Blog, err error) {
	db := mysql.GetDB().WithContext(ctx)
	if page <= 0 {
		page = 1
	}
	offest := (page - 1) * pageSize
	err = db.Where(&blogReq).Offset(offest).Find(&blogs).Limit(pageSize).Error
	return
}

func DeleteBlog(ctx context.Context, blogReq req.DeleteBlogReq) (err error) {
	db := mysql.GetDB().WithContext(ctx)
	return db.Delete(&blogReq).Error
}

func LikeBlog(ctx context.Context, blogReq req.LikeBlogReq) (err error) {
	db := mysql.GetDB().WithContext(ctx)
	return db.Model(&blogReq).UpdateColumn("like", gorm.Expr("like +?", 1)).Error
}

func DislikeBlog(ctx context.Context, blogReq req.DislikeBlogReq) (err error) {
	db := mysql.GetDB().WithContext(ctx)
	return db.Model(&blogReq).UpdateColumn("like", gorm.Expr("dislike +?", 1)).Error
}
