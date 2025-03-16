package dtos

type GetAllAssetsRequest struct {
	Id            string `json:"id"`
	TypeForm      string `json:"type_form"`
	Description   string `json:"description"`
	Timestamp     string `json:"timestamp"`
	InsertionType string `json:"insertion_type"`
	Hash          string `json:"hash"`
}

type PostAssetRequest struct {
	Id            string `json:"id"`
	TypeForm      string `json:"type_form"`
	Description   string `json:"description"`
	Timestamp     string `json:"timestamp"`
	InsertionType string `json:"insertion_type"`
	Hash          string `json:"hash"`
}
