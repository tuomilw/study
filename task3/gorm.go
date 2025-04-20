package task3

import "gorm.io/gorm"

type User struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	PostCounts int
	Posts      []Post `gorm:"foreignKey:UserId"`
}
type Post struct {
	Id            uint `gorm:"primaryKey;autoIncrement"`
	UserId        uint
	CommentCounts int
	State         string
	Content       string
	Comments      []Comment `gorm:"foreignKey:PostId"`
}
type Comment struct {
	Id      uint `gorm:"primaryKey;autoIncrement"`
	PostId  uint
	Content string
}

func GetUserPostAndComments(db *gorm.DB, id uint, posts []Post, comments []Comment) error {
	if err := db.Model(&User{}).Where("id=?", id).Association("posts").Find(&posts); err != nil {
		return err
	}
	if err := db.Model(&Post{}).Where("user_id=?", id).Association("comments").Find(&comments); err != nil {
		return err
	}
	return nil
}
func GetMaxCommentsPost(db *gorm.DB, post *Post) error {
	err := db.Model(&Post{}).Debug().Order("comment_counts desc").First(post)
	return err.Error
}

// 钩子函数
func (post *Post) AfterCreateByPost(db *gorm.DB) error {
	result := db.Model(&User{}).Where("id=?", post.UserId).Update("post_counts", gorm.Expr("post_counts+?", 1))
	return result.Error
}
func (comment *Comment) BeforeDeleteByComment(db *gorm.DB) error {
	var post Post
	if err := db.Model(&Post{}).Where("id=?", comment.PostId).Find(&post).Error; err != nil {
		return err
	}
	if post.CommentCounts == 0 {
		result := db.Model(&Post{}).Where("id=?", comment.PostId).Update("state", "无评论")
		return result.Error
	}
	return nil
}
