package chartype

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_NewCandle(t *testing.T) {
	time := time.Now()

	res := Candle{
		Timestamp: time,
	}

	d := NewCandle(time)

	assert.Equal(t, res, d)
}

func Test_Candle_Parse(t *testing.T) {
	cc := map[string]struct {
		Candle Candle
		Open   string
		High   string
		Low    string
		Close  string
		Volume string
		Result Candle
		Err    error
	}{
		"Successful parse": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "1",
			High:   "3",
			Low:    "5",
			Close:  "7",
			Volume: "9",
			Result: Candle{
				Timestamp: time.Time{},
				Open:      decimal.NewFromInt(1),
				High:      decimal.NewFromInt(3),
				Low:       decimal.NewFromInt(5),
				Close:     decimal.NewFromInt(7),
				Volume:    decimal.NewFromInt(9),
			},
		},
		"Invalid Open": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "-",
			High:   "3",
			Low:    "5",
			Close:  "7",
			Volume: "9",
			Err:    assert.AnError,
		},
		"Invalid High": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "1",
			High:   "-",
			Low:    "5",
			Close:  "7",
			Volume: "9",
			Err:    assert.AnError,
		},
		"Invalid Low": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "1",
			High:   "3",
			Low:    "-",
			Close:  "7",
			Volume: "9",
			Err:    assert.AnError,
		},
		"Invalid Close": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "1",
			High:   "3",
			Low:    "5",
			Close:  "-",
			Volume: "9",
			Err:    assert.AnError,
		},
		"Invalid Volume": {
			Candle: Candle{Timestamp: time.Time{}},
			Open:   "1",
			High:   "3",
			Low:    "5",
			Close:  "7",
			Volume: "-",
			Err:    assert.AnError,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			err := c.Candle.Parse(c.Open, c.High, c.Low, c.Close, c.Volume)
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, c.Candle)
		})
	}
}

func Test_CandleField_Validate(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Err         error
	}{
		"Successful CandleOpen validation": {
			CandleField: CandleOpen,
			Err:         nil,
		},
		"Successful CandleHigh validation": {
			CandleField: CandleHigh,
			Err:         nil,
		},
		"Successful CandleLow validation": {
			CandleField: CandleLow,
			Err:         nil,
		},
		"Successful CandleClose validation": {
			CandleField: CandleClose,
			Err:         nil,
		},
		"Invalid CandleField": {
			CandleField: 69,
			Err:         ErrInvalidCandleField,
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

func Test_CandleField_UnmarshalJSON(t *testing.T) {
	cc := map[string]struct {
		Data   string
		Result CandleField
		Err    error
	}{
		"Successful CandleOpen unmarshal": {
			Data:   `"open"`,
			Result: CandleOpen,
		},
		"Successful CandleHigh unmarshal": {
			Data:   `"high"`,
			Result: CandleHigh,
		},
		"Successful CandleLow unmarshal": {
			Data:   `"low"`,
			Result: CandleLow,
		},
		"Successful CandleClose unmarshal": {
			Data:   `"close"`,
			Result: CandleClose,
		},
		"Successful CandleVolume unmarshal": {
			Data:   `"volume"`,
			Result: CandleVolume,
		},
		"Malformed JSON": {
			Data: `{"69"`,
			Err:  assert.AnError,
		},
		"Invalid CandleField": {
			Data: `"69"`,
			Err:  ErrInvalidCandleField,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var cf CandleField
			err := cf.UnmarshalJSON([]byte(c.Data))
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.Equal(t, c.Result, cf)
		})
	}
}

func Test_CandleField_MarshalJSON(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Result      string
		Err         error
	}{
		"Successful CandleOpen marshal": {
			CandleField: CandleOpen,
			Result:      `"open"`,
		},
		"Successful CandleHigh marshal": {
			CandleField: CandleHigh,
			Result:      `"high"`,
		},
		"Successful CandleLow marshal": {
			CandleField: CandleLow,
			Result:      `"low"`,
		},
		"Successful CandleClose marshal": {
			CandleField: CandleClose,
			Result:      `"close"`,
		},
		"Successful CandleVolume marshal": {
			CandleField: CandleVolume,
			Result:      `"volume"`,
		},
		"Invalid CandleField": {
			CandleField: 69,
			Err:         ErrInvalidCandleField,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			d, err := c.CandleField.MarshalJSON()
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.JSONEq(t, c.Result, string(d))
		})
	}
}

func Test_CandleField_Extract(t *testing.T) {
	cc := map[string]struct {
		CandleField CandleField
		Candle      Candle
		Result      decimal.Decimal
	}{
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
		"Invalid CandleField": {
			CandleField: 69,
			Candle: Candle{
				Open:   decimal.NewFromInt(30),
				High:   decimal.NewFromInt(30),
				Low:    decimal.NewFromInt(30),
				Close:  decimal.NewFromInt(30),
				Volume: decimal.NewFromInt(30),
			},
			Result: decimal.Zero,
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
		Result Ticker
		Err    error
	}{
		"Successful parse": {
			Last: "1",
			Ask:  "3",
			Bid:  "5",
			Result: Ticker{
				Last: decimal.NewFromInt(1),
				Ask:  decimal.NewFromInt(3),
				Bid:  decimal.NewFromInt(5),
			},
		},
		"Invalid Last": {
			Last: "-",
			Ask:  "3",
			Bid:  "5",
			Err:  assert.AnError,
		},
		"Invalid Ask": {
			Last: "1",
			Ask:  "-",
			Bid:  "5",
			Err:  assert.AnError,
		},
		"Invalid Bid": {
			Last: "1",
			Ask:  "3",
			Bid:  "-",
			Err:  assert.AnError,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			tr, err := ParseTicker(c.Last, c.Ask, c.Bid)
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
		"Successful TickerLast validation": {
			TickerField: TickerLast,
			Err:         nil,
		},
		"Successful TickerAsk validation": {
			TickerField: TickerAsk,
			Err:         nil,
		},
		"Successful TickerBid validation": {
			TickerField: TickerBid,
			Err:         nil,
		},
		"Invalid TickerField": {
			TickerField: 69,
			Err:         ErrInvalidTickerField,
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
		Result      string
		Err         error
	}{
		"Successful TickerLast marshal": {
			TickerField: TickerLast,
			Result:      `"last"`,
		},
		"Successful TickerAsk marshal": {
			TickerField: TickerAsk,
			Result:      `"ask"`,
		},
		"Successful TickerBid marshal": {
			TickerField: TickerBid,
			Result:      `"bid"`,
		},
		"Invalid TickerField": {
			TickerField: 69,
			Err:         ErrInvalidTickerField,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			d, err := c.TickerField.MarshalJSON()
			AssertEqualError(t, c.Err, err)
			if err != nil {
				return
			}

			assert.JSONEq(t, c.Result, string(d))
		})
	}
}

func Test_TickerField_UnmarshalJSON(t *testing.T) {
	cc := map[string]struct {
		Data   string
		Result TickerField
		Err    error
	}{
		"Successful TickerLast unmarshal (long form)": {
			Data:   `"last"`,
			Result: TickerLast,
		},
		"Successful TickerLast unmarshal (short form)": {
			Data:   `"l"`,
			Result: TickerLast,
		},
		"Successful TickerAsk unmarshal  (long form)": {
			Data:   `"ask"`,
			Result: TickerAsk,
		},
		"Successful TickerAsk unmarshal  (short form)": {
			Data:   `"a"`,
			Result: TickerAsk,
		},
		"Successful TickerBid unmarshal  (long form)": {
			Data:   `"bid"`,
			Result: TickerBid,
		},
		"Successful TickerBid unmarshal  (short form)": {
			Data:   `"b"`,
			Result: TickerBid,
		},
		"Malformed JSON": {
			Data: `{"69"`,
			Err:  assert.AnError,
		},
		"Invalid TickerField": {
			Data: `"69"`,
			Err:  ErrInvalidTickerField,
		},
	}

	for cn, c := range cc {
		c := c

		t.Run(cn, func(t *testing.T) {
			t.Parallel()

			var tf TickerField
			err := tf.UnmarshalJSON([]byte(c.Data))
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
		"Invalid CandleField": {
			TickerField: 69,
			Ticker: Ticker{
				Last: decimal.NewFromInt(30),
				Ask:  decimal.NewFromInt(30),
				Bid:  decimal.NewFromInt(30),
			},
			Result: decimal.Zero,
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