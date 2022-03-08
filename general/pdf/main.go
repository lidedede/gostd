package main
import(
	"fmt"
	"rsc.io/pdf"
)
func main() {
	file, err := pdf.Open("/Users/lidedong/go/src/github.com/lidedede/gostd/general/pdf/test2222.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Page(2).Content())
}
