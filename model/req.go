package model

import user_center "Blog/proto/user-center"

//user
type RegisterReq struct {
	//User struct {
	//	Username string `json:"username" form:"username"`
	//	Password string `json:"password" form:"password"`
	//}
	//UserExtra struct {
	//	Nickname string `json:"nickname" form:"nickname"`
	//	Email 	 string `json:"email" form:"email"`
	//	Phonenumber string `json:"phonenumber" form:"phonenumber"`
	//	Avatarurl string `json:"avatarurl" form:"avatarurl"`
	//}
	User      user_center.User
	Userextra user_center.UserExtra
}

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type DeleteReq struct {
	Id int32 `json:"id" form:"id"`
}

type GetUserInfoReq struct {
	Username string `json:"username" form:"username"`
}

//Blog
type CreateBlogReq struct {
	Content string `json:"content" form:"content"`
	Title   string `json:"title"   form:"string"`
}

type UpdateBlogReq struct {
	BlogID  uint
	Content string `json:"content" form:"content"`
	Title   string `json:"title"   form:"string"`
}

type GetBlogReq struct {
	Content string `json:"content" form:"content"`
	Title   string `json:"title"   form:"string"`
}

type DeleteBlogReq struct {
	BlogID uint `json:"blog_id" form:"blog_id"`
}

type LikeBlogReq struct {
	BlogID uint `json:"blog_id" form:"blog_id"`
}

type DislikeBlogReq struct {
	BlogID uint `json:"blog_id" form:"blog_id"`
}

//comment
type CreateCommentReq struct {
	CommentID      uint   `json:"comment_id" form:"comment_id"`
	BlogID         uint   `json:"blog_id" form:"blog_id"`
	CommentContent string `json:"comment_content" form:"comment_content"`
}

type UpdateCommentReq struct {
	CommentID      uint   `json:"comment_id" form:"comment_id"`
	BlogID         uint   `json:"blog_id" form:"blog_id"`
	CommentContent string `json:"comment_content" form:"comment_content"`
}

type GetCommentReq struct {
	CommentContent string `json:"comment_content" form:"comment_content"`
}

type DeleteCommentReq struct {
	CommentID uint `json:"comment_id" form:"comment_id"`
}
