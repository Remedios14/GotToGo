package hello

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	if dy < 1 || dx < 1 {
		return nil
	}
	var ret [][]uint8
	for y := 0; y < dy; y++ {
		var input = make([]uint8, dx, dx)
		for x := 0; x < dx; x++ {
			input = append(input, uint8((x+y)/2))
		}
		ret = append(ret, input)
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
