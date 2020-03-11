package src

import "time"

// DateFilter gathers all the records that are at most 1 year old
func DateFilter(data []InputRecord, start time.Time) []InputRecord {
	result := make([]InputRecord, 0)
	for _, record := range data {
		if record.Date.After(start) {
			result = append(result, record)
		}
	}
	return result
}

// AggregateAmount counts sum of all amounts for all records.
func AggregateAmount(data []InputRecord) map[string]map[string]float64 {
	aggregator := make(map[string]map[string]float64)
	for _, record := range data {

		if _, ok := aggregator[record.UserID]; !ok {
			aggregator[record.UserID] = make(map[string]float64)
		}

		aggregator[record.UserID][record.Currency] += record.Amount
	}
	return aggregator
}

// ProcessData turns input data to aggregated output data
func ProcessData(data []InputRecord) []OutputRecord {
	result := make([]OutputRecord, 0)
	aggregatedData := AggregateAmount(data)

	for UserID, v := range aggregatedData {

		for Currency, Amount := range v {
			result = append(result, newOutputRecord(UserID, Currency, Amount, 0))
		}
	}

	return result
}
