package dispatcher

func Sort(width, height, length, mass float64) string {
	volume := width * height * length
	bulky := volume >= 1_000_000 || width >= 150 || height >= 150 || length >= 150
	heavy := mass >= 20

	if bulky && heavy {
		return "REJECTED"
	}
	if bulky || heavy {
		return "SPECIAL"
	}
	return "STANDARD"
}
