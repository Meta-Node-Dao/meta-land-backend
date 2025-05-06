package chain

type ChainListResponse struct {
	List []ChainBasicResponse `json:"list"`
}

type ChainBasicResponse struct {
	Chain
	ChainContracts []ChainContract `json:"chain_contracts" gorm:"foreignKey:ChainID;references:ChainID"`
	ChainEndpoints []ChainEndpoint `json:"chain_endpoints" gorm:"foreignKey:ChainID;references:ChainID"`
}
