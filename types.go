package chartype

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	// CandleOpen specifies starting candle's value.
	CandleOpen CandleField = iota + 1

	// CandleHigh specifies highest candle's value.
	CandleHigh

	// CandleLow specifies lowest candle's value.
	CandleLow

	// CandleClose specifies last candle's value.
	CandleClose
)

// Candle stores specific timeframe's starting, closing,
// highest and lowest price points.
type Candle struct {
	Open  decimal.Decimal `json:"open"`
	High  decimal.Decimal `json:"high"`
	Low   decfmal.Decimal `json:"low"`
	Close decimal.Decimal `json:"close"`
}

// CandleField specifies which field should be extracted
// from the candle for further calculations.
// Can be included in configuration structures.
type CandleField int

// MarshalJSON turns candle field to appropriate string
// representation in JSON.
func (cf CandleField) MarshalJSON() ([]byte, error) {
	switch cf {
	case CandleOpen:
		return []byte("open"), nil
	case CandleHigh:
		return []byte("high"), nil
	case CandleLow:
		return []byte("low"), nil
	case CandleClose:
		return []byte("close"), nil
	default:
		return nil, errors.New("undefined candle type")
	}
}

// UnmarshalJSON turns JSON string to appropriate field
// value.
func (cf *CandleField) UnmarshalJSON(d []byte) error {
	var f string
	if err := json.Unmarshal(d, &f); err != nil {
		return err
	}

	f = strings.ToLower(f)

	switch f {
	case "open", "o":
		*cf = CandleOpen
	case "high", "h":
		*cf = CandleHigh
	case "low", "l":
		*cf = CandleLow
	case "close", "c":
		*cf = CandleClose
	default:
		return errors.New("undefined candle field")
	}

	return nil
}

// Extract returns candle's value specified in the candle
// field type.
func (cf CandleField) Extract(c Candle) decimal.Decimal {
	switch cf {
	case CandleOpen:
		return c.Open
	case CandleHigh:
		return c.High
	case CandleLow:
		return c.Low
	case CandleClose:
		return c.Close
	default:
		return decimal.Zero
	}
}

// Ticker holds current ask, last and bid prices.
type Ticker struct {
	Last decimal.Decimal
	Ask  decimal.Decimal
	Bid  decimal.Decimal
}

// Packet holds tickert information as well as all
// known candles for a specific timeframe.
type Packet struct {
	Ticker  Ticker
	Candles []Candle
}
