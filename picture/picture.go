package main

import ("golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
        var ret [][]uint8
	for i := 0; i < dy; i++ {
		var si []uint8
		for j := 0; j < dx; j++ {
		     si = append(si, uint8((j * i)))	
		}
		ret = append(ret, si)
        }
	return ret
	
}
func printSlice(s []uint8) {
	fmt.Println(s)
}


func main() {
	pic.Show(Pic)
	//for i := 0; i < 10; i++ {
		//printSlice(Pic(0, 0)[i])
	//}
}
