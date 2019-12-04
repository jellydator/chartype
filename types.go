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
	Low   decimal.Decimal `json:"low"`
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
		return nil, errors.New("undefined candle field")
	}
}

// UnmarshalJSON turns JSON string to appropriate candle
// field value.
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

// Extract returns candle's value as specified in the candle
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

const (
	// TickerLast specifies the last ticker value.
	TickerLast TickerField = iota + 1

	// TickerAsk specifies the ask ticker value.
	TickerAsk

	// TickerBid specifies the bid ticker value.
	TickerBid
)

// Ticker holds current ask, last and bid prices.
type Ticker struct {
	Last decimal.Decimal
	Ask  decimal.Decimal
	Bid  decimal.Decimal
}

// TickerField specifies which field should be extracted
// from the ticker for further calculations.
// Can be included in configuration structures.
type TickerField int

// MarshalJSON turns ticker field to appropriate string
// representation in JSON.
func (tf TickerField) MarshalJSON() ([]byte, error) {
	switch tf {
	case TickerLast:
		return []byte("last"), nil
	case TickerAsk:
		return []byte("ask"), nil
	case TickerBid:
		return []byte("bid"), nil
	default:
		return nil, errors.New("undefined ticker field")
	}
}

// UnmarshalJSON turns JSON string to appropriate ticker
// field value.
func (tf *TickerField) UnmarshalJSON(d []byte) error {
	var t string
	if err := json.Unmarshal(d, &t); err != nil {
		return err
	}

	t = strings.ToLower(t)

	switch t {
	case "last", "l":
		*tf = TickerLast
	case "ask", "a":
		*tf = TickerAsk
	case "bid", "b":
		*tf = TickerBid
	default:
		return errors.New("undefined ticker field")
	}

	return nil
}

// Extract returns ticker's value as specified in the ticker
// field type.
func (tf TickerField) Extract(t Ticker) decimal.Decimal {
	switch tf {
	case TickerLast:
		return t.Last
	case TickerAsk:
		return t.Ask
	case TickerBid:
		return t.Bid
	default:
		return decimal.Zero
	}
}

// Packet holds ticker information as well as all
// known candles for a specific timeframe.
type Packet struct {
	Ticker  Ticker
	Candles []Candle
}
