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
			equalError(t, c.Err, err)
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
			equalError(t, c.Err, err)
		})
	}
}

func Test_CandleField_MarshalText(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Text        string
		Err         error
	}{
		"Invalid CandleField": {
			CandleField: 70,
			Err:         ErrInvalidCandleField,
		},
		"Successful CandleOpen marshal": {
			CandleField: CandleOpen,
			Text:        "open",
		},
		"Successful CandleHigh marshal": {
			CandleField: CandleHigh,
			Text:        "high",
		},
		"Successful CandleLow marshal": {
			CandleField: CandleLow,
			Text:        "low",
		},
		"Successful CandleClose marshal": {
			CandleField: CandleClose,
			Text:        "close",
		},
		"Successful CandleVolume marshal": {
			CandleField: CandleVolume,
			Text:        "volume",
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			res, err := c.CandleField.MarshalText()
			equalError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Text, string(res))
		})
	}
}

func Test_CandleField_UnmarshalJSON(t *testing.T) {
	cc := map[string]struct {
		Text   string
		Result CandleField
		Err    error
	}{
		"Invalid CandleField": {
			Text: "70",
			Err:  ErrInvalidCandleField,
		},
		"Successful CandleOpen unmarshal (long form)": {
			Text:   "open",
			Result: CandleOpen,
		},
		"Successful CandleOpen unmarshal (short form)": {
			Text:   "o",
			Result: CandleOpen,
		},
		"Successful CandleHigh unmarshal (long form)": {
			Text:   "high",
			Result: CandleHigh,
		},
		"Successful CandleHigh unmarshal (short form)": {
			Text:   "h",
			Result: CandleHigh,
		},
		"Successful CandleLow unmarshal (long form)": {
			Text:   "low",
			Result: CandleLow,
		},
		"Successful CandleLow unmarshal (short form)": {
			Text:   "l",
			Result: CandleLow,
		},
		"Successful CandleClose unmarshal (long form)": {
			Text:   "close",
			Result: CandleClose,
		},
		"Successful CandleClose unmarshal (short form)": {
			Text:   "c",
			Result: CandleClose,
		},
		"Successful CandleVolume unmarshal (long form)": {
			Text:   "volume",
			Result: CandleVolume,
		},
		"Successful CandleVolume unmarshal (short form)": {
			Text:   "v",
			Result: CandleVolume,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var cf CandleField
			err := cf.UnmarshalText([]byte(c.Text))
			equalError(t, c.Err, err)
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
		Last          string
		Ask           string
		Bid           string
		Change        string
		PercentChange string
		Volume        string
		Result        Ticker
		Err           error
	}{
		"Invalid Last": {
			Last:          "-",
			Ask:           "3",
			Bid:           "5",
			Change:        "2",
			PercentChange: "2",
			Err:           assert.AnError,
		},
		"Invalid Ask": {
			Last:          "1",
			Ask:           "-",
			Bid:           "5",
			Change:        "3",
			PercentChange: "2",
			Err:           assert.AnError,
		},
		"Invalid Bid": {
			Last:          "1",
			Ask:           "3",
			Bid:           "-",
			Change:        "2",
			PercentChange: "2",
			Err:           assert.AnError,
		},
		"Invalid Change": {
			Last:          "1",
			Ask:           "3",
			Bid:           "4",
			Change:        "-",
			PercentChange: "2",
			Err:           assert.AnError,
		},
		"Invalid PercentChange": {
			Last:          "1",
			Ask:           "3",
			Bid:           "4",
			Change:        "2",
			PercentChange: "-",
			Err:           assert.AnError,
		},
		"Invalid Volume": {
			Last:          "1",
			Ask:           "3",
			Bid:           "4",
			Change:        "2",
			PercentChange: "2",
			Volume:        "-",
			Err:           assert.AnError,
		},
		"Successful parse": {
			Last:          "1",
			Ask:           "3",
			Bid:           "5",
			Change:        "4",
			PercentChange: "2",
			Volume:        "1",
			Result: Ticker{
				Last:          decimal.NewFromInt(1),
				Ask:           decimal.NewFromInt(3),
				Bid:           decimal.NewFromInt(5),
				Change:        decimal.NewFromInt(4),
				PercentChange: decimal.NewFromInt(2),
				Volume:        decimal.NewFromInt(1),
			},
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			tr, err := ParseTicker(c.Last, c.Ask, c.Bid, c.Change, c.PercentChange, c.Volume)
			equalError(t, c.Err, err)
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
		"Successful TickerPercentChange validation": {
			TickerField: TickerPercentChange,
		},
		"Successful TickerVolume validation": {
			TickerField: TickerVolume,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			err := c.TickerField.Validate()
			equalError(t, c.Err, err)
		})
	}
}

func Test_TickerField_MarshalText(t *testing.T) {
	cc := map[string]struct {
		TickerField TickerField
		Text        string
		Err         error
	}{
		"Invalid TickerField": {
			TickerField: 70,
			Err:         ErrInvalidTickerField,
		},
		"Successful TickerLast marshal": {
			TickerField: TickerLast,
			Text:        "last",
		},
		"Successful TickerAsk marshal": {
			TickerField: TickerAsk,
			Text:        "ask",
		},
		"Successful TickerBid marshal": {
			TickerField: TickerBid,
			Text:        "bid",
		},
		"Successful TickerChange marshal": {
			TickerField: TickerChange,
			Text:        "change",
		},
		"Successful TickerPercentChange marshal": {
			TickerField: TickerPercentChange,
			Text:        "percent_change",
		},
		"Successful TickerVolume marshal": {
			TickerField: TickerVolume,
			Text:        "volume",
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			res, err := c.TickerField.MarshalText()
			equalError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Text, string(res))
		})
	}
}

func Test_TickerField_UnmarshalText(t *testing.T) {
	cc := map[string]struct {
		Text   string
		Result TickerField
		Err    error
	}{
		"Invalid TickerField": {
			Text: "70",
			Err:  ErrInvalidTickerField,
		},
		"Successful TickerLast unmarshal (long form)": {
			Text:   "last",
			Result: TickerLast,
		},
		"Successful TickerLast unmarshal (short form)": {
			Text:   "l",
			Result: TickerLast,
		},
		"Successful TickerAsk unmarshal  (long form)": {
			Text:   "ask",
			Result: TickerAsk,
		},
		"Successful TickerAsk unmarshal  (short form)": {
			Text:   "a",
			Result: TickerAsk,
		},
		"Successful TickerBid unmarshal  (long form)": {
			Text:   "bid",
			Result: TickerBid,
		},
		"Successful TickerBid unmarshal  (short form)": {
			Text:   "b",
			Result: TickerBid,
		},
		"Successful TickerChange unmarshal  (long form)": {
			Text:   "change",
			Result: TickerChange,
		},
		"Successful TickerChange unmarshal  (short form)": {
			Text:   "c",
			Result: TickerChange,
		},
		"Successful TickerPercentChange unmarshal  (long form)": {
			Text:   "percent_change",
			Result: TickerPercentChange,
		},
		"Successful TickerPercentChange unmarshal  (short form)": {
			Text:   "pc",
			Result: TickerPercentChange,
		},
		"Successful Volume unmarshal  (long form)": {
			Text:   "volume",
			Result: TickerVolume,
		},
		"Successful Volume unmarshal  (short form)": {
			Text:   "v",
			Result: TickerVolume,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var tf TickerField
			err := tf.UnmarshalText([]byte(c.Text))
			equalError(t, c.Err, err)
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
		"Successful Change extract": {
			TickerField: TickerChange,
			Ticker:      Ticker{Change: decimal.NewFromInt(220)},
			Result:      decimal.NewFromInt(220),
		},
		"Successful PercentChange extract": {
			TickerField: TickerPercentChange,
			Ticker:      Ticker{PercentChange: decimal.NewFromInt(203)},
			Result:      decimal.NewFromInt(203),
		},
		"Successful Volume extract": {
			TickerField: TickerVolume,
			Ticker:      Ticker{Volume: decimal.NewFromInt(201)},
			Result:      decimal.NewFromInt(201),
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
