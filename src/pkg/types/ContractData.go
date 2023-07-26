package types

type ContractData struct {
	PledgeAmount          string `json:"pledge_amount"`
	PledgeTotalCalcAmount string `json:"pledge_total_calc_amount"`
	OwnerIsContract       bool   `json:"owner_is_contract"`
}
