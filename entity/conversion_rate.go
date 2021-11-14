package entity

type ConversionRate struct {
	CurrencyIDFrom int     `json:"currency_id_from"`
	CurrencyIDTo   int     `json:"currency_id_to"`
	Rate           float64 `json:"currency_rate"`
}

func (cr *ConversionRate) Valid() bool {
	return cr.CurrencyIDFrom > 0 &&
		cr.CurrencyIDTo > 0 &&
		cr.Rate > 0
}

func (cr *ConversionRate) Convert(v float64) float64 {
	return v * cr.Rate
}
