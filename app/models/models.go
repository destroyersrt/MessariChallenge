package models

type AssetPoolsRequest struct {
	ID string `json:"id"`
}

type TimeStamp struct {
	Lower uint64 `json:"lower"`
	Upper uint64 `json:"upper"`
}

type AssetVolumeRequest struct {
	ID        string    `json:"id"`
	Timestamp TimeStamp `json:"timestamp"`
}

type BlockRequest struct {
	Block uint32 `json:"block"`
}
