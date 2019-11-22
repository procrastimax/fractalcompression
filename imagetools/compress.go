package imagetools

import "image"

import "math"

//EncodingParams saves the encoding parameters for every range which shall be recreated
type EncodingParams struct {
	S  float64
	G  float64
	Dx int
	Dy int
	//T is the used transformation (0-7)
	T uint8
}

//FindBestMatchingDomains finds the best matching domain
func FindBestMatchingDomains(ranges [][]*image.Gray, domains [][]*image.Gray) [][]EncodingParams {
	var encodings = make([][]EncodingParams, len(ranges))
	for i := range ranges {
		encodings[i] = make([]EncodingParams, len(ranges[i]))

		for j := range ranges[i] {
			var e = math.MaxFloat64
			for a := range domains {
				for b := range domains[a] {
					// try out some transformations
					for z := 0; z < 8; z++ {
						var eTemp = math.MaxFloat64
						eTemp, s, g := CalcSquarredEuclideanDistance(ranges[i][j], TransformImage(domains[a][b], uint8(z)))
						if eTemp < e {
							e = eTemp
							encodings[i][j].S = s
							encodings[i][j].G = g
							encodings[i][j].Dx = a
							encodings[i][j].Dy = b
							encodings[i][j].T = uint8(z)
						}
					}
				}
			}
		}
	}
	return encodings
}
