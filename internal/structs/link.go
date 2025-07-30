package structs

type Link struct {
	UserID     int64
	Link       string
	LastAuthor string
	LastText   string
	LastStars  int
}

func NewLink() *Link {
	return &Link{}
}
