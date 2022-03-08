package slice

//slice数据结构
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}

//与数组区别：数组传值，slice传指针

//初始化
var slice []int //nil 切片 需要进行初始化之后才能用

//make创建
//slice := make([]int, 3)  //len 3, cap 3
//slice := make([]int, 3, 5)  //len 3， cap 5

//快速创建
//#创建切片
//slice := []int{1,2,3} //len 3, cap 3
//#创建数组
//slice := [10]int{1,2,3}  //大小为10的数组
//#只初始化某个索引值
//s := []int{4:1} //len 5, cap 5 ,value [0 0 0 0 1]

func main() {

}