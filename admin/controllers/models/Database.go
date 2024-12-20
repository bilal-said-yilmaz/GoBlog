package models

var Dsn = "postgres://postgres:Golang007@localhost:5432/blog_db?sslmode=disable"

/*type Post struct { // TODO: tek bir fonksiyonda halledilebilir mi ?
	gorm.Model
	Title, Slug, Description, Content, Picture string
	CategoryID                                 int
}

func (post Post) OpenDB(where ...interface{}) interface{} {
	Db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	Db.AutoMigrate(&post)

	Db.Create(&post)

	Db.First(&post, where...)

	var posts []Post

	Db.Find(&posts, where...)

	//Db.Model(&post).Update()

	Db.Model(&post).Updates(where)
	return post
}
*/
