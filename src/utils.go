package src

import "time"

func DateFilter(data []InputRecord, start time.Time) ([]InputRecord, error) {
	result := make([]InputRecord, 0);
	for _, record := range data {
		if record.Date.After(start) {
			result = append(result, record);
		}
	}
	return result, nil
}

func aggregateAmount(data []InputRecord) map[string]map[string]float64 {
	aggregator := make(map[string]map[string]float64)
	for _, record := range data {

		if _, ok := aggregator[record.UserID]; !ok {
			aggregator[record.UserID] = make(map[string]float64)
		}

		aggregator[record.UserID][record.Currency] += record.Amount
	}
	return aggregator
}

func ProcessData(data []InputRecord) []OutputRecord {
	result := make([]OutputRecord, 0);
	aggregatedData := aggregateAmount(data)

	for userId, v := range aggregatedData {

		for curr, amount := range v {
			result = append(result, newOutputRecord(userId, curr, amount, 0))
		}
	}

	return result
}
