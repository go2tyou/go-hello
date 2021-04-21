package model

type student struct {
	Name string
	Score float64
}

// 结构体首字母小写，只能在本包使用 -> 工厂模式暴露
func NewStudent(name string, score float64) *student {
	return &student {
		Name : name,
		Score : score,
	}
}