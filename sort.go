package dispatcher

const (
	Standard = "STANDARD"
	Special  = "SPECIAL"
	Rejected = "REJECTED"

	VolThreshold  = 1_000_000
	DimThreshold  = 150
	MassThreshold = 20
)

func isBulky(width, height, length float64) bool {
	volume := width * height * length
	return volume >= VolThreshold || width >= DimThreshold || height >= DimThreshold || length >= DimThreshold
}

func isHeavy(mass float64) bool {
	return mass >= MassThreshold
}

// Sort dispatches packages into stacks based on physical properties.
//
// Stacks:
//   - STANDARD: not bulky or heavy
//   - SPECIAL: either bulky or heavy
//   - REJECTED: both bulky and heavy
//
// Definitions:
//   - Bulky: volume >= 1,000,000 cm³ or any dimension >= 150 cm
//   - Heavy: mass >= 20 kg
//
// Units are cm for dimensions and kg for mass.
func Sort(width, height, length, mass float64) string {
	bulky := isBulky(width, height, length)
	heavy := isHeavy(mass)

	if bulky && heavy {
		return Rejected
	}
	if bulky || heavy {
		return Special
	}
	return Standard
}
