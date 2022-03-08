package a

func (a *StructA) SayNum() string {
	if a.elem.SayNum() < 100 {
		return "小于100"
	}
	return "大于等于100"
}
