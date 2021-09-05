package models

type AssetRequest struct {
	ID string `json:"id"`
}

type BlockRequest struct {
	Block uint32 `json:"block"`
}
