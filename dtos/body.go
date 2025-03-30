package dtos

type GetAllAssetsRequest struct {
	Id            string `json:"id"`
	TypeForm      string `json:"type_form"`
	Description   string `json:"description"`
	InsertionType string `json:"insertion_type"`
	Hash          string `json:"hash"`
	Timestamp     string `json:"timestamp"`
}

type Filter struct {
	Ids            []string `json:"ids"`
	TypeForms      []string `json:"type_forms"`
	InsertionTypes []string `json:"insertion_types"`
	Hashs          []string `json:"hashs"`
}

type PostAssetRequest struct {
	Id            string `json:"id"`
	TypeForm      string `json:"type_form"`
	Description   string `json:"description"`
	InsertionType string `json:"insertion_type"`
	Hash          string `json:"hash"`
	Timestamp     string `json:"timestamp"`
}
