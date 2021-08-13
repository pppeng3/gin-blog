package handler

import (
	"Blog/dal"
	dbm "Blog/db/model"
	"Blog/model"
	user_center "Blog/proto/user-center"
	"Blog/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context) {
	var req model.CreateCommentReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	user := userInterface.(*user_center.User)
	comment := dbm.Comment{
		Model:             gorm.Model{ID: req.CommentID},
		UserID:            uint(user.Id),
		BlogID:            req.BlogID,
		CommentContent:    req.CommentContent,
		CommentCreateTime: time.Now(),
	}
	if err := dal.CreateOrUpdateComment(c, comment); err != nil {
		logrus.Warnf("create comment failed: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
}

func UpdateComment(c *gin.Context) {
	var req model.UpdateCommentReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
	}
	userInterface, existed := c.Get("user")
	if !existed {
		logrus.Warnf("not login yet")
		utils.RetData(c, userInterface)
		return
	}
	user := userInterface.(*user_center.User)
	comment := dbm.Comment{
		Model:             gorm.Model{ID: req.CommentID},
		UserID:            uint(user.Id),
		BlogID:            req.BlogID,
		CommentContent:    req.CommentContent,
		CommentUpdateTime: time.Now(),
	}
	if err := dal.CreateOrUpdateComment(c, comment); err != nil {
		logrus.Warnf("create comment failed: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
}

func GetComment(c *gin.Context) {
	var req model.GetCommentReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	resp, err := dal.GetComments(c, req, 0, 20)
	if err != nil {
		logrus.Warnf("could not find comments: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, resp)
}

func DeleteComment(c *gin.Context) {
	var req model.DeleteCommentReq
	if err := c.ShouldBind(req); err != nil {
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
	if err := dal.DeleteComment(c, req); err != nil {
		logrus.Warnf("comment could not delete")
		utils.RetErr(c, err)
		return
	}

}
