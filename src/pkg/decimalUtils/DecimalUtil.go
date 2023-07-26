package decimalUtils

import (
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/shopspring/decimal"
)

type DecimalUtil struct {
	*decimal.Decimal
}

func NewDecimalFromString(str string) decimal.Decimal {
	res, err := decimal.NewFromString(str)
	if err != nil {
		log.Errorf("NewDecimalFromString error: %v", err)
		panic(err)
	}
	return res
}
