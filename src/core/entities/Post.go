package entities

type Post struct {
	NickName string `json:"nickname"`
	Content  string `json:"content"`
}

func NewPost(NickName, Content string) Post {
	return Post{NickName, Content}
}

func (p Post) GetNickLength() int {
	return len(p.NickName)
}

func (p Post) GetContentLength() int {
	return len(p.Content)
}
