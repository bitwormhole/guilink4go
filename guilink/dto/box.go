package dto

// Box 表示一个矩形的框
type Box struct {
	Base

	X      int `json:"x"`
	Y      int `json:"y"`
	Z      int `json:"z"`
	Width  int `json:"w"`
	Height int `json:"h"`
}
