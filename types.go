// Package chartype provides types and functions for convenient work with
// market data structures.
package chartype

import (
	"errors"
	"time"

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

	// CandleVolume specifies candle's volume value.
	CandleVolume
)

var (
	// ErrInvalidCandleField is returned when candle field
	// with invalid value is being used.
	ErrInvalidCandleField = errors.New("invalid candle field")
)

// Candle stores specific timeframe's starting, closing,
// highest and lowest price points.
type Candle struct {
	Timestamp time.Time       `json:"timestamp" db:"timestamp"`
	Open      decimal.Decimal `json:"open" db:"open"`
	High      decimal.Decimal `json:"high" db:"high"`
	Low       decimal.Decimal `json:"low" db:"low"`
	Close     decimal.Decimal `json:"close" db:"close"`
	Volume    decimal.Decimal `json:"volume" db:"volume"`
}

// ParseCandle parses provided string parameters into newly created candle's fields
// and returns it.
func ParseCandle(t time.Time, os, hs, ls, cs, vs string) (Candle, error) {
	o, err := decimal.NewFromString(os)
	if err != nil {
		return Candle{}, err
	}

	h, err := decimal.NewFromString(hs)
	if err != nil {
		return Candle{}, err
	}

	l, err := decimal.NewFromString(ls)
	if err != nil {
		return Candle{}, err
	}

	c, err := decimal.NewFromString(cs)
	if err != nil {
		return Candle{}, err
	}

	v, err := decimal.NewFromString(vs)
	if err != nil {
		return Candle{}, err
	}

	return Candle{Timestamp: t, Open: o, High: h, Low: l, Close: c, Volume: v}, nil
}

// CandleField specifies which field should be extracted
// from the candle for further calculations.
// Can be included in configuration structures.
type CandleField int

// Validate checks whether the candle field is one of
// supported field types or not.
func (cf CandleField) Validate() error {
	switch cf {
	case CandleOpen, CandleHigh, CandleLow, CandleClose, CandleVolume:
		return nil
	default:
		return ErrInvalidCandleField
	}
}

// MarshalText turns candle field to appropriate string
// representation.
func (cf CandleField) MarshalText() ([]byte, error) {
	var v string

	switch cf {
	case CandleOpen:
		v = "open"
	case CandleHigh:
		v = "high"
	case CandleLow:
		v = "low"
	case CandleClose:
		v = "close"
	case CandleVolume:
		v = "volume" //nolint:goconst // we need to be explicit about these fields
	default:
		return nil, ErrInvalidCandleField
	}

	return []byte(v), nil
}

// UnmarshalText turns string to appropriate candle
// field value.
func (cf *CandleField) UnmarshalText(d []byte) error {
	switch string(d) {
	case "open", "o":
		*cf = CandleOpen
	case "high", "h":
		*cf = CandleHigh
	case "low", "l":
		*cf = CandleLow
	case "close", "c":
		*cf = CandleClose
	case "volume", "v":
		*cf = CandleVolume
	default:
		return ErrInvalidCandleField
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
	case CandleVolume:
		return c.Volume
	default:
		return decimal.Zero
	}
}

// FromCandles extracts specific candle fields from all provided candles
// and puts them in plain number slice.
func FromCandles(cc []Candle, cf CandleField) []decimal.Decimal {
	res := make([]decimal.Decimal, len(cc))
	for i, c := range cc {
		res[i] = cf.Extract(c)
	}

	return res
}

const (
	// TickerLast specifies last ticker value.
	TickerLast TickerField = iota + 1

	// TickerAsk specifies ask ticker value.
	TickerAsk

	// TickerBid specifies bid ticker value.
	TickerBid

	// TickerChange specifies 24 hours price units change in
	// last ticker price.
	TickerChange

	// TickerPercentChange specifies 24 hour price percent change in
	// last ticker price.
	TickerPercentChange

	// TickerVolume specifies volume ticker value.
	TickerVolume
)

var (
	// ErrInvalidTickerField is returned when ticker field
	// with invalid value is being used.
	ErrInvalidTickerField = errors.New("invalid ticker field")
)

// Ticker holds current ask, last and bid prices.
type Ticker struct {
	Last          decimal.Decimal `json:"last"`
	Ask           decimal.Decimal `json:"ask"`
	Bid           decimal.Decimal `json:"bid"`
	Change        decimal.Decimal `json:"change"`
	PercentChange decimal.Decimal `json:"percent_change"`
	Volume        decimal.Decimal `json:"volume"`
}

// ParseTicker parses provided string parameters into decimal type values,
// adds them into a new ticker instance and returns it.
func ParseTicker(ls, as, bs, cs, pcs, vs string) (Ticker, error) {
	l, err := decimal.NewFromString(ls)
	if err != nil {
		return Ticker{}, err
	}

	a, err := decimal.NewFromString(as)
	if err != nil {
		return Ticker{}, err
	}

	b, err := decimal.NewFromString(bs)
	if err != nil {
		return Ticker{}, err
	}

	c, err := decimal.NewFromString(cs)
	if err != nil {
		return Ticker{}, err
	}

	pc, err := decimal.NewFromString(pcs)
	if err != nil {
		return Ticker{}, err
	}

	v, err := decimal.NewFromString(vs)
	if err != nil {
		return Ticker{}, err
	}

	return Ticker{Last: l, Ask: a, Bid: b, Change: c, PercentChange: pc, Volume: v}, nil
}

// TickerField specifies which field should be extracted
// from the ticker for further calculations.
// Can be included in configuration structures.
type TickerField int

// Validate checks whether the ticker field is one of
// supported field types or not.
func (tf TickerField) Validate() error {
	switch tf {
	case TickerLast, TickerAsk, TickerBid, TickerChange, TickerPercentChange, TickerVolume:
		return nil
	default:
		return ErrInvalidTickerField
	}
}

// MarshalText turns ticker field to appropriate string
// representation.
func (tf TickerField) MarshalText() ([]byte, error) {
	var v string

	switch tf {
	case TickerLast:
		v = "last"
	case TickerAsk:
		v = "ask"
	case TickerBid:
		v = "bid"
	case TickerChange:
		v = "change"
	case TickerPercentChange:
		v = "percent_change"
	case TickerVolume:
		v = "volume"
	default:
		return nil, ErrInvalidTickerField
	}

	return []byte(v), nil
}

// UnmarshalText turns string to appropriate ticker
// field value.
func (tf *TickerField) UnmarshalText(d []byte) error {
	switch string(d) {
	case "last", "l":
		*tf = TickerLast
	case "ask", "a":
		*tf = TickerAsk
	case "bid", "b":
		*tf = TickerBid
	case "change", "c":
		*tf = TickerChange
	case "percent_change", "pc":
		*tf = TickerPercentChange
	case "volume", "v":
		*tf = TickerVolume
	default:
		return ErrInvalidTickerField
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
	case TickerChange:
		return t.Change
	case TickerPercentChange:
		return t.PercentChange
	case TickerVolume:
		return t.Volume
	default:
		return decimal.Zero
	}
}

// Packet holds ticker information as well as all
// known candles for a specific timeframe.
type Packet struct {
	Ticker  Ticker   `json:"ticker"`
	Candles []Candle `json:"candles"`
}
