package gorm

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func ResetUserTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&User{}) {
		db.Debug().Migrator().DropTable(&User{})
	}

	db.Debug().AutoMigrate(&User{})

	for i := 0; i < 20; i++ {
		db.Debug().Create(&User{
			UserName: "user" + strconv.Itoa(i+1),
		})
	}
}

func ResetPostTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&Post{}) {
		db.Debug().Migrator().DropTable(&Post{})
	}

	db.Debug().AutoMigrate(&Post{})

	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			db.Debug().Create(&Post{
				Title:   "title" + strconv.Itoa(j+1),
				Content: "content" + strconv.Itoa(j+1),
				UserID:  uint(i + 1),
			})
		}
	}
}

func ResetCommentTable(db *gorm.DB) {
	if db.Debug().Migrator().HasTable(&Comment{}) {
		db.Debug().Migrator().DropTable(&Comment{})
	}

	db.Debug().AutoMigrate(&Comment{})

	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				db.Debug().Create(&Comment{
					Content: "content" + strconv.Itoa(k+1),
					PostID:  uint(j + 1),
				})
			}
		}
	}

	db.Debug().Model(&Comment{}).Where("post_id = 1").Update("post_id", 2)
}

func QueryUserPosts(db *gorm.DB, user_id uint) []Post {
	var user User
	db.Debug().Preload("Posts").Preload("Posts.Comments").Where("id = ?", user_id).First(&user)

	return user.Posts
}

func QueryMaxCommentsPost(db *gorm.DB) Post {
	type PostCommentResult struct {
		CommentCount uint
		PostID       uint
	}
	var result PostCommentResult
	db.Debug().Table("comments").Select("COUNT(id) AS comment_count, post_id").Group("post_id").Order("comment_count DESC").Limit(1).Scan(&result)

	var post Post
	db.Debug().Model(&Post{}).Where("ID = ?", result.PostID).Take(&post)

	return post
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Debug().Raw("UPDATE users SET post_count = post_count + 1 WHERE id = ?", p.UserID)
	return
}

func CreateUserPostWithHook(db *gorm.DB, title string, content string, user_id uint) {
	tx := db.Begin()
	tx.Debug().Create(&Post{
		Title:   title,
		Content: content,
		UserID:  user_id,
	})
	tx.Commit()
}

func UpdatePostCallback(db *gorm.DB) {
	fmt.Println("UpdatePostCallback Call............................")
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	tx.Debug().Where("id = ?", c.PostID).Find(&post)
	if post.ID != 0 && post.ID == c.PostID && len(post.Comments) == 0 {
		post.Status = "无评论"
		tx.Save(&post)
	}
	return
}

func DeleteCommentWithHook(db *gorm.DB, id uint) {
	tx := db.Begin()
	tx.Debug().Where("id = ?", id).Delete(&Comment{})
	tx.Commit()
}
