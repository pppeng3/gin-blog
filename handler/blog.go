package handler

import (
	"Blog/dal"
	"Blog/model"
	user_center "Blog/proto/user-center"
	"Blog/utils"
	"time"

	bl "Blog/db/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateBlog(c *gin.Context) {
	var req model.CreateBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	user := userInterface.(*user_center.User)
	blog := bl.Blog{
		UserID:         uint(user.Id),
		Title:          req.Title,
		Content:        req.Content,
		BlogCreateTime: time.Now(),
	}
	err := dal.CreateOrUpdateBlog(c, blog)
	if err != nil {
		logrus.Errorf("Createblog error: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
}

func UpdateBlog(c *gin.Context) {
	var req model.UpdateBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	user := userInterface.(*user_center.User)
	blog := bl.Blog{
		BlogID:         req.BlogID,
		UserID:         uint(user.Id),
		Title:          req.Title,
		Content:        req.Title,
		BlogUpdateTime: time.Now(),
	}
	err := dal.CreateOrUpdateBlog(c, blog)
	if err != nil {
		logrus.Errorf("Update error: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
}

func GetBlog(c *gin.Context) {
	var req model.GetBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params")
		utils.RetErr(c, err)
		return
	}
	resp, err := dal.GetBlogs(c, req, 1, 10)
	if err != nil {
		logrus.Errorf("Can not find any blogs: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, resp)
}

func DeleteBlog(c *gin.Context) {
	var req model.DeleteBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params")
		utils.RetErr(c, err)
		return
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	if err := dal.DeleteBlog(c, req); err != nil {
		logrus.Warnf("blog delete failed: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, req)
}

func LikeBlog(c *gin.Context) {
	var req model.LikeBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params")
		utils.RetErr(c, err)
		return
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	if err := dal.LikeBlog(c, req); err != nil {
		logrus.Warnf("blog delete failed: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, req)
}

func DislikeBlog(c *gin.Context) {
	var req model.DislikeBlogReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params")
		utils.RetErr(c, err)
		return
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	if err := dal.DislikeBlog(c, req); err != nil {
		logrus.Warnf("blog delete failed: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, req)
}
