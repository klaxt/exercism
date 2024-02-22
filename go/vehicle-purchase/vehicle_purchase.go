package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// TODO another autoformat annoyance - why do cases get moved back
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	message := "%s is clearly the better choice."
	// ternary is not supported https://go.dev/doc/faq#Does_Go_have_a_ternary_form
	if option1 < option2 {
		return fmt.Sprintf(message, option1)
	} else {
		return fmt.Sprintf(message, option2)
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percent float64 = 0
	if age < 3 {
		percent = .8
	} else if age < 10 {
		percent = .7
	} else {
		percent = .5
	}
	return originalPrice * float64(percent)
}
