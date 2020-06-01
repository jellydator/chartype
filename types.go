package chartype

import (
	"encoding/json"
	"errors"
	"strings"
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
	Timestamp time.Time       `json:"timestamp"`
	Open      decimal.Decimal `json:"open"`
	High      decimal.Decimal `json:"high"`
	Low       decimal.Decimal `json:"low"`
	Close     decimal.Decimal `json:"close"`
	Volume    decimal.Decimal `json:"volume"`
}

// NewCandle instantiates a new candle with the timestamp provided.
func NewCandle(timestamp time.Time) Candle {
	return Candle{
		Timestamp: timestamp,
	}
}

// Parse parses provided string parameters into corresponding candle's fields.
func (c *Candle) Parse(opn, hgh, low, cls, vol string) error {
	o, err := decimal.NewFromString(opn)
	if err != nil {
		return err
	}

	h, err := decimal.NewFromString(hgh)
	if err != nil {
		return err
	}

	l, err := decimal.NewFromString(low)
	if err != nil {
		return err
	}

	clsD, err := decimal.NewFromString(cls)
	if err != nil {
		return err
	}

	v, err := decimal.NewFromString(vol)
	if err != nil {
		return err

	}

	(*c).Open = o
	(*c).High = h
	(*c).Low = l
	(*c).Close = clsD
	(*c).Volume = v

	return nil
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

// MarshalJSON turns candle field to appropriate string
// representation in JSON.
func (cf CandleField) MarshalJSON() ([]byte, error) {
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
		v = "volume"
	default:
		return nil, ErrInvalidCandleField
	}

	return json.Marshal(v)
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
	var res []decimal.Decimal
	for _, c := range cc {
		res = append(res, cf.Extract(c))
	}

	return res
}

const (
	// TickerLast specifies the last ticker value.
	TickerLast TickerField = iota + 1

	// TickerAsk specifies the ask ticker value.
	TickerAsk

	// TickerBid specifies the bid ticker value.
	TickerBid
)

var (
	// ErrInvalidTickerField is returned when ticker field
	// with invalid value is being used.
	ErrInvalidTickerField = errors.New("invalid ticker field")
)

// Ticker holds current ask, last and bid prices.
type Ticker struct {
	Last decimal.Decimal `json:"last"`
	Ask  decimal.Decimal `json:"ask"`
	Bid  decimal.Decimal `json:"bid"`
}

// ParseTicker parses provided string parameters into decimal type values,
// adds them into a new ticker instance and returns it.
func ParseTicker(lst, ask, bid string) (Ticker, error) {
	l, err := decimal.NewFromString(lst)
	if err != nil {
		return Ticker{}, err
	}

	a, err := decimal.NewFromString(ask)
	if err != nil {
		return Ticker{}, err
	}

	b, err := decimal.NewFromString(bid)
	if err != nil {
		return Ticker{}, err
	}

	return Ticker{Last: l, Ask: a, Bid: b}, nil
}

// TickerField specifies which field should be extracted
// from the ticker for further calculations.
// Can be included in configuration structures.
type TickerField int

// Validate checks whether the ticker field is one of
// supported field types or not.
func (tf TickerField) Validate() error {
	switch tf {
	case TickerLast, TickerAsk, TickerBid:
		return nil
	default:
		return ErrInvalidTickerField
	}
}

// MarshalJSON turns ticker field to appropriate string
// representation in JSON.
func (tf TickerField) MarshalJSON() ([]byte, error) {
	var v string
	switch tf {
	case TickerLast:
		v = "last"
	case TickerAsk:
		v = "ask"
	case TickerBid:
		v = "bid"
	default:
		return nil, ErrInvalidTickerField
	}

	return json.Marshal(v)
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
