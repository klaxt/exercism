package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// rate * success / 100
	return float64(productionRate) * successRate / float64(100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

const (
	BatchSize = 10
	BatchCost = 95_000
	CarCost   = 10_000
)

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	batches := carsCount / BatchSize
	cars := carsCount % BatchSize
	return uint(batches*BatchCost + cars*CarCost)
}
