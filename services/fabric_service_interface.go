package services

type FabricService interface {
	ReadInFabric(operation string, args ...string) string
	PostInFabric(operation string, args ...string) string
}
