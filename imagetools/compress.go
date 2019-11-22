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
		encodings[i] = make([]EncodingParams, len(ranges[i]))

		for j := range ranges[i] {
			fmt.Printf("%d/%d von %d/%d\n", i+1, j+1, len(ranges), len(ranges[0]))
			var e = math.MaxFloat64
			for a := range domains {
				for b := range domains[a] {
					var eTemp = math.MaxFloat64
					eTemp, s, g := CalcSquarredEuclideanDistance(ranges[i][j], domains[a][b])
					if eTemp < e {
						e = eTemp
						encodings[i][j].S = s
						encodings[i][j].G = g
						encodings[i][j].Dx = a
						encodings[i][j].Dy = b
					}
				}
			}
		}
	}
	return encodings
}
