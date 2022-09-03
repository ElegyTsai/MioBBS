package post

type Post struct {
	CommonPost
	HotScore int
	Title    string
	Tag      string
	Nails    [][]byte
}
