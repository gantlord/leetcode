package lc1436

func destCity(paths [][]string) string {
	originCities := map[string]bool{}
	for _, path := range paths {
		originCities[path[0]] = true
	}
	var destCity string
	for _, path := range paths {
		if exists := originCities[path[1]]; !exists {
			destCity = path[1]
		}
	}
	return destCity
}
