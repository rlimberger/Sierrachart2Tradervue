package dtc2tv

import (
	"Sierrachart2Tradervue/dtc"
	"Sierrachart2Tradervue/tradervue"
	"strconv"
	"time"
)

func ExecutionFromDTCOrderFill(fill dtc.HistoricalOrderFillResponse) (tradervue.Execution, error) {

	isoTime := time.Unix(fill.DateTime, 0).Format(time.RFC3339)

	quantity := strconv.Itoa(int(fill.Quantity))
	if fill.BuySell == dtc.BuySellEnum_SELL {
		quantity = "-" + quantity
	}

	price := strconv.FormatFloat(fill.Price, 'f', 2, 64)

	return tradervue.Execution{
		Datetime: isoTime,
		Symbol:   fill.Symbol,
		Quantity: quantity,
		Price:    price,
	}, nil
}
