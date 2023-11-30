package entities

type Post struct {
	NickName string `json:"nickname"`
	Content  string `json:"content"`
}

func NewPost(NickName, Content string) Post {
	return Post{NickName, Content}
}
