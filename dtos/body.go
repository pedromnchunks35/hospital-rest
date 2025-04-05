package dtos

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GetAllAssetsRequest struct {
	Id            string    `json:"id"`
	TypeForm      string    `json:"type_form"`
	Description   string    `json:"description"`
	InsertionType string    `json:"insertion_type"`
	Hash          string    `json:"hash"`
	Timestamp     time.Time `json:"timestamp"`
}

type Filter struct {
	Ids            []string        `json:"ids"`
	TypeForms      []string        `json:"type_forms"`
	InsertionTypes []string        `json:"insertion_types"`
	Hashs          []string        `json:"hashs"`
	TimeFilter     TimestampFilter `json:"time_filter"`
}

type TimestampFilter struct {
	Min time.Time `json:"min"`
	Max time.Time `json:"max"`
}

type PostAssetRequest struct {
	Id            string    `json:"id"`
	TypeForm      string    `json:"type_form"`
	Description   string    `json:"description"`
	InsertionType string    `json:"insertion_type"`
	Hash          string    `json:"hash"`
	Timestamp     time.Time `json:"timestamp"`
}

type Asset struct {
	Id            string    `json:"id"`
	TypeForm      string    `json:"type_form"`
	Description   string    `json:"description"`
	InsertionType string    `json:"insertion_type"`
	Hash          string    `json:"hash"`
	Timestamp     time.Time `json:"timestamp"`
}

type KeyModification struct {
	TxId      string                 `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsDelete  bool                   `protobuf:"varint,4,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
}

type AssetHistory struct {
	TxId      string                 `json:"tx_id,omitempty"`
	Value     *Asset                 `json:"value,omitempty"`
	Timestamp *timestamppb.Timestamp `json:"timestamp,omitempty"`
	IsDelete  bool                   `json:"is_delete,omitempty"`
}
