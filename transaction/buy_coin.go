package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for buy a coin paying another coin (owned by sender).
// CoinToBuy - Symbol of a coin to get. ValueToBuy - Amount of CoinToBuy to get.
// CoinToSell - Symbol of a coin to give.
// MaximumValueToSell - Maximum value of coins to sell.
type BuyCoinData struct {
	CoinToBuy          [10]byte
	ValueToBuy         *big.Int
	CoinToSell         [10]byte
	MaximumValueToSell *big.Int
}

func NewBuyCoinData() *BuyCoinData {
	return &BuyCoinData{}
}

func (d *BuyCoinData) SetCoinToSell(symbol string) *BuyCoinData {
	copy(d.CoinToSell[:], symbol)
	return d
}

func (d *BuyCoinData) SetCoinToBuy(symbol string) *BuyCoinData {
	copy(d.CoinToBuy[:], symbol)
	return d
}

func (d *BuyCoinData) SetValueToBuy(value *big.Int) *BuyCoinData {
	d.ValueToBuy = value
	return d
}

func (d *BuyCoinData) SetMaximumValueToSell(value *big.Int) *BuyCoinData {
	d.MaximumValueToSell = value
	return d
}

func (d *BuyCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *BuyCoinData) fee() Fee {
	return feeTypeBuyCoin
}
