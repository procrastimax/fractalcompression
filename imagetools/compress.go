package imagetools

import (
	"fmt"
	"image"
	"math"
)

//EncodingParams saves the encoding parameters for every range which shall be recreated
type EncodingParams struct {
	S  float64
	G  float64
	Dx int
	Dy int
}

//FindBestMatchingDomains finds the best matching domain
func FindBestMatchingDomains(ranges [][]*image.Gray, domains [][]*image.Gray) [][]EncodingParams {
	var encodings = make([][]EncodingParams, len(ranges))
	for i := range ranges {
		fmt.Printf("%d from %d\n", i+1, len(ranges))
		encodings[i] = make([]EncodingParams, len(ranges[i]))

		for j := range ranges[i] {
			var e = math.MaxFloat64
			for a := range domains {
				for b := range domains[a] {
					var eTemp = math.MaxFloat64
					eTemp = CalcSquarredEuclideanDistance(ranges[i][j], domains[a][b])
					if eTemp == 0 {
						encodings[i][j].S = 1
						encodings[i][j].G = 0
						encodings[i][j].Dx = a
						encodings[i][j].Dy = b
						break
					} else if eTemp < e {
						e = eTemp
						encodings[i][j].S = 1
						encodings[i][j].G = 0
						encodings[i][j].Dx = a
						encodings[i][j].Dy = b
					}
				}
			}
		}
	}
	fmt.Println("Optimize s and g for all found domains...")
	//now iterate over all encodings and optimize s and g for it
	for i := range ranges {
		for j := range ranges[i] {
			s, g := CalcContrastAndBrightness(ranges[i][j], domains[encodings[i][j].Dx][encodings[i][j].Dy])
			encodings[i][j].S = s
			encodings[i][j].G = g
		}
	}

	return encodings
}
