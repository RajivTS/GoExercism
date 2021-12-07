package cars

const baseProduction = 221.0
const minutesInHour = 60

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed < 1 {
		return 0.0
	} else if speed < 5 {
		return 1.0
	} else if speed < 9 {
		return 0.9
	} else {
		return 0.77
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return baseProduction * float64(speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / minutesInHour)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	actualProduction := CalculateProductionRatePerHour(speed)
	if actualProduction > limit {
		return limit
	} else {
		return actualProduction
	}
}
