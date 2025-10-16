package bs_eth

import (
	"math/big"
)

type Receipt struct { //12
	Type              uint8    `json:"type,omitempty"`
	PostState         string   `json:"root"`
	Status            uint64   `json:"status"`
	CumulativeGasUsed uint64   `json:"cumulative_gas_used" gencodec:"required"`
	Bloom             string   `json:"logs_bloom"         gencodec:"required"`
	Logs              []*Log   `json:"logs"              gencodec:"required"`
	TxHash            string   `json:"transaction_hash"   gencodec:"required"`
	ContractAddress   string   `json:"contract_address"`
	GasUsed           uint64   `json:"gas_used"           gencodec:"required"`
	BlockHash         string   `json:"block_hash,omitempty"`
	BlockNumber       *big.Int `json:"block_number,omitempty"`
	TransactionIndex  uint     `json:"transaction_index"`
	To                string   `json:"to"`
	From              string   `json:"from"`
	EffectiveGasPrice string   `json:"effective_gas_price"`
}

/* 14
   "blockHash": "0x0b71e27419dd0dd04f551c573c65bd6cacfe760b0e3d1fa12566115ed0e320fb",
   "blockNumber": "0x6f25",
   "contractAddress": "0x97077686617a8f4478863c23b73b7387ba35a802",
   "cumulativeGasUsed": "0x4e1e3",
   "effectiveGasPrice": "0x3b9aca07",
   "from": "0x30e938b0630c02f394d17925fdb5fb046f70d452",
   "gasUsed": "0x4e1e3",
   "logs": [],
   "logsBloom": "",
   "status": "0x1",
   "to": null,
   "transactionHash": "",
   "transactionIndex": "",
   "type": ""
*/
