package models

type User struct {
	Base
	Username   string
	Email      string
	Phone      string
	Password   string
	ProfilePic string
	CoverPhoto string
	Posts      []Post    `gorm:"foreignKey:UserID"`
	Comments   []Comment `gorm:"foreignKey:UserID"`
}

type Post struct {
	Base
	UserID      uint64
	Description string
	PostMedias  []PostMedia `gorm:"foreignKey:PostID"`
	Comments    []Comment   `gorm:"foreignKey:PostID"`
}

type PostMedia struct {
	Base
	PostID    uint64
	MediaType string
	FileName  string
}

type Comment struct {
	Base
	PostID  uint64
	UserID  uint64
	Comment string
}

type Like struct {
	Base
	ItemID   uint64
	UserID   uint64
	Category string
}

type CommentReply struct {
	Base
	CommentID uint64
	UserID    uint64
	Reply     string
}

func ReturnListOfTables() []interface{} {
	return []interface{}{
		&User{}, &Post{}, &PostMedia{}, &Comment{},
		&CommentReply{}, &Like{}}
}
