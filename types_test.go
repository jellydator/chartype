package chartype

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_ParseCandle(t *testing.T) {
	cc := map[string]struct {
		Timestamp time.Time
		Open      string
		High      string
		Low       string
		Close     string
		Volume    string
		Result    Candle
		Err       error
	}{
		"Invalid Open": {
			Timestamp: time.Time{},
			Open:      "-",
			High:      "3",
			Low:       "5",
			Close:     "7",
			Volume:    "9",
			Err:       assert.AnError,
		},
		"Invalid High": {
			Timestamp: time.Time{},
			Open:      "1",
			High:      "-",
			Low:       "5",
			Close:     "7",
			Volume:    "9",
			Err:       assert.AnError,
		},
		"Invalid Low": {
			Timestamp: time.Time{},
			Open:      "1",
			High:      "3",
			Low:       "-",
			Close:     "7",
			Volume:    "9",
			Err:       assert.AnError,
		},
		"Invalid Close": {
			Timestamp: time.Time{},
			Open:      "1",
			High:      "3",
			Low:       "5",
			Close:     "-",
			Volume:    "9",
			Err:       assert.AnError,
		},
		"Invalid Volume": {
			Timestamp: time.Time{},
			Open:      "1",
			High:      "3",
			Low:       "5",
			Close:     "7",
			Volume:    "-",
			Err:       assert.AnError,
		},
		"Successful parse": {
			Timestamp: time.Time{},
			Open:      "1",
			High:      "3",
			Low:       "5",
			Close:     "7",
			Volume:    "9",
			Result: Candle{
				Timestamp: time.Time{},
				Open:      decimal.NewFromInt(1),
				High:      decimal.NewFromInt(3),
				Low:       decimal.NewFromInt(5),
				Close:     decimal.NewFromInt(7),
				Volume:    decimal.NewFromInt(9),
			},
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			res, err := ParseCandle(c.Timestamp, c.Open, c.High, c.Low, c.Close, c.Volume)
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, res)
		})
	}
}

func Test_CandleField_Validate(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Err         error
	}{
		"Invalid CandleField": {
			CandleField: 70,
			Err:         ErrInvalidCandleField,
		},
		"Successful CandleOpen validation": {
			CandleField: CandleOpen,
		},
		"Successful CandleHigh validation": {
			CandleField: CandleHigh,
		},
		"Successful CandleLow validation": {
			CandleField: CandleLow,
		},
		"Successful CandleClose validation": {
			CandleField: CandleClose,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			err := c.CandleField.Validate()
			AssertEqualError(t, c.Err, err)
		})
	}
}

func Test_CandleField_MarshalJSON(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		JSON        string
		Err         error
	}{
		"Invalid CandleField": {
			CandleField: 70,
			Err:         ErrInvalidCandleField,
		},
		"Successful CandleOpen marshal": {
			CandleField: CandleOpen,
			JSON:        `"open"`,
		},
		"Successful CandleHigh marshal": {
			CandleField: CandleHigh,
			JSON:        `"high"`,
		},
		"Successful CandleLow marshal": {
			CandleField: CandleLow,
			JSON:        `"low"`,
		},
		"Successful CandleClose marshal": {
			CandleField: CandleClose,
			JSON:        `"close"`,
		},
		"Successful CandleVolume marshal": {
			CandleField: CandleVolume,
			JSON:        `"volume"`,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			res, err := c.CandleField.MarshalJSON()
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.JSONEq(t, c.JSON, string(res))
		})
	}
}

func Test_CandleField_UnmarshalJSON(t *testing.T) {
	cc := map[string]struct {
		JSON   string
		Result CandleField
		Err    error
	}{
		"Malformed JSON": {
			JSON: `{"70"`,
			Err:  assert.AnError,
		},
		"Invalid CandleField": {
			JSON: `"70"`,
			Err:  ErrInvalidCandleField,
		},
		"Successful CandleOpen unmarshal (long form)": {
			JSON:   `"open"`,
			Result: CandleOpen,
		},
		"Successful CandleOpen unmarshal (short form)": {
			JSON:   `"o"`,
			Result: CandleOpen,
		},
		"Successful CandleHigh unmarshal (long form)": {
			JSON:   `"high"`,
			Result: CandleHigh,
		},
		"Successful CandleHigh unmarshal (short form)": {
			JSON:   `"h"`,
			Result: CandleHigh,
		},
		"Successful CandleLow unmarshal (long form)": {
			JSON:   `"low"`,
			Result: CandleLow,
		},
		"Successful CandleLow unmarshal (short form)": {
			JSON:   `"low"`,
			Result: CandleLow,
		},
		"Successful CandleClose unmarshal (long form)": {
			JSON:   `"close"`,
			Result: CandleClose,
		},
		"Successful CandleClose unmarshal (short form)": {
			JSON:   `"c"`,
			Result: CandleClose,
		},
		"Successful CandleVolume unmarshal (long form)": {
			JSON:   `"volume"`,
			Result: CandleVolume,
		},
		"Successful CandleVolume unmarshal (short form)": {
			JSON:   `"v"`,
			Result: CandleVolume,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var cf CandleField
			err := cf.UnmarshalJSON([]byte(c.JSON))
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, cf)
		})
	}
}

func Test_CandleField_Extract(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Candle      Candle
		Result      decimal.Decimal
	}{
		"Invalid CandleField": {
			CandleField: 70,
			Candle: Candle{
				Open:   decimal.NewFromInt(30),
				High:   decimal.NewFromInt(30),
				Low:    decimal.NewFromInt(30),
				Close:  decimal.NewFromInt(30),
				Volume: decimal.NewFromInt(30),
			},
			Result: decimal.Zero,
		},
		"Successful Open extract": {
			CandleField: CandleOpen,
			Candle:      Candle{Open: decimal.NewFromInt(10)},
			Result:      decimal.NewFromInt(10),
		},
		"Successful High extract": {
			CandleField: CandleHigh,
			Candle:      Candle{High: decimal.NewFromInt(15)},
			Result:      decimal.NewFromInt(15),
		},
		"Successful Low extract": {
			CandleField: CandleLow,
			Candle:      Candle{Low: decimal.NewFromInt(20)},
			Result:      decimal.NewFromInt(20),
		},
		"Successful Close extract": {
			CandleField: CandleClose,
			Candle:      Candle{Close: decimal.NewFromInt(25)},
			Result:      decimal.NewFromInt(25),
		},
		"Successful Volume extract": {
			CandleField: CandleVolume,
			Candle:      Candle{Volume: decimal.NewFromInt(30)},
			Result:      decimal.NewFromInt(30),
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			v := c.CandleField.Extract(c.Candle)
			assert.Equal(t, c.Result, v)
		})
	}
}

func Test_FromCandles(t *testing.T) {
	cc := []Candle{
		{
			Open: decimal.NewFromInt(10),
		},
		{
			Open: decimal.NewFromInt(15),
		},
		{
			Open: decimal.NewFromInt(5),
		},
	}

	dd := FromCandles(cc, CandleOpen)

	res := []decimal.Decimal{
		decimal.NewFromInt(10),
		decimal.NewFromInt(15),
		decimal.NewFromInt(5),
	}

	assert.Equal(t, res, dd)
}

func Test_ParseTicker(t *testing.T) {
	cc := map[string]struct {
		Last   string
		Ask    string
		Bid    string
		Change string
		Result Ticker
		Err    error
	}{
		"Invalid Last": {
			Last:   "-",
			Ask:    "3",
			Bid:    "5",
			Change: "2",
			Err:    assert.AnError,
		},
		"Invalid Ask": {
			Last:   "1",
			Ask:    "-",
			Bid:    "5",
			Change: "3",
			Err:    assert.AnError,
		},
		"Invalid Bid": {
			Last:   "1",
			Ask:    "3",
			Bid:    "-",
			Change: "2",
			Err:    assert.AnError,
		},
		"Invalid Change": {
			Last:   "1",
			Ask:    "3",
			Bid:    "4",
			Change: "-",
			Err:    assert.AnError,
		},
		"Successful parse": {
			Last:   "1",
			Ask:    "3",
			Bid:    "5",
			Change: "4",
			Result: Ticker{
				Last:   decimal.NewFromInt(1),
				Ask:    decimal.NewFromInt(3),
				Bid:    decimal.NewFromInt(5),
				Change: decimal.NewFromInt(4),
			},
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			tr, err := ParseTicker(c.Last, c.Ask, c.Bid, c.Change)
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, tr)
		})
	}
}

func Test_TickerField_Validate(t *testing.T) {
	cc := map[string]struct {
		TickerField TickerField
		Err         error
	}{
		"Invalid TickerField": {
			TickerField: 70,
			Err:         ErrInvalidTickerField,
		},
		"Successful TickerLast validation": {
			TickerField: TickerLast,
		},
		"Successful TickerAsk validation": {
			TickerField: TickerAsk,
		},
		"Successful TickerBid validation": {
			TickerField: TickerBid,
		},
		"Successful TickerChange validation": {
			TickerField: TickerChange,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			err := c.TickerField.Validate()
			AssertEqualError(t, c.Err, err)
		})
	}

}

func Test_TickerField_MarshalJSON(t *testing.T) {
	cc := map[string]struct {
		TickerField TickerField
		JSON        string
		Err         error
	}{
		"Invalid TickerField": {
			TickerField: 70,
			Err:         ErrInvalidTickerField,
		},
		"Successful TickerLast marshal": {
			TickerField: TickerLast,
			JSON:        `"last"`,
		},
		"Successful TickerAsk marshal": {
			TickerField: TickerAsk,
			JSON:        `"ask"`,
		},
		"Successful TickerChange marshal": {
			TickerField: TickerChange,
			JSON:        `"change"`,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			res, err := c.TickerField.MarshalJSON()
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.JSONEq(t, c.JSON, string(res))
		})
	}
}

func Test_TickerField_UnmarshalJSON(t *testing.T) {
	cc := map[string]struct {
		JSON   string
		Result TickerField
		Err    error
	}{
		"Malformed JSON": {
			JSON: `{"70"`,
			Err:  assert.AnError,
		},
		"Invalid TickerField": {
			JSON: `"70"`,
			Err:  ErrInvalidTickerField,
		},
		"Successful TickerLast unmarshal (long form)": {
			JSON:   `"last"`,
			Result: TickerLast,
		},
		"Successful TickerLast unmarshal (short form)": {
			JSON:   `"l"`,
			Result: TickerLast,
		},
		"Successful TickerAsk unmarshal  (long form)": {
			JSON:   `"ask"`,
			Result: TickerAsk,
		},
		"Successful TickerAsk unmarshal  (short form)": {
			JSON:   `"a"`,
			Result: TickerAsk,
		},
		"Successful TickerBid unmarshal  (long form)": {
			JSON:   `"bid"`,
			Result: TickerBid,
		},
		"Successful TickerBid unmarshal  (short form)": {
			JSON:   `"b"`,
			Result: TickerBid,
		},
		"Successful TickerChange unmarshal  (long form)": {
			JSON:   `"change"`,
			Result: TickerChange,
		},
		"Successful TickerChange unmarshal  (short form)": {
			JSON:   `"c"`,
			Result: TickerChange,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var tf TickerField
			err := tf.UnmarshalJSON([]byte(c.JSON))
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, tf)
		})
	}
}

func Test_TickerField_Extract(t *testing.T) {
	cc := map[string]struct {
		TickerField TickerField
		Ticker      Ticker
		Result      decimal.Decimal
	}{
		"Invalid CandleField": {
			TickerField: 70,
			Ticker: Ticker{
				Last: decimal.NewFromInt(30),
				Ask:  decimal.NewFromInt(30),
				Bid:  decimal.NewFromInt(30),
			},
			Result: decimal.Zero,
		},
		"Successful Last extract": {
			TickerField: TickerLast,
			Ticker:      Ticker{Last: decimal.NewFromInt(10)},
			Result:      decimal.NewFromInt(10),
		},
		"Successful Ask extract": {
			TickerField: TickerAsk,
			Ticker:      Ticker{Ask: decimal.NewFromInt(15)},
			Result:      decimal.NewFromInt(15),
		},
		"Successful Bid extract": {
			TickerField: TickerBid,
			Ticker:      Ticker{Bid: decimal.NewFromInt(20)},
			Result:      decimal.NewFromInt(20),
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			v := c.TickerField.Extract(c.Ticker)
			assert.Equal(t, c.Result, v)
		})
	}
}
