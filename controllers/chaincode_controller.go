package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hospital-rest/dtos"
	"hospital-rest/services"
	"hospital-rest/utils/apiResponses"
	"hospital-rest/utils/chaincodeConnection"
	"hospital-rest/utils/validation"
)

type ChaincodeController struct {
	FabricService services.FabricService
}

func CreateChaincodeController(fabricService services.FabricService) *ChaincodeController {
	return &ChaincodeController{
		FabricService: fabricService,
	}
}

func (c ChaincodeController) GetAllAssets(context *gin.Context) {
	pageString := context.Param("page")
	sizeString := context.Param("size")

	page, size, err := validation.ValidatePageAndSize(pageString, sizeString)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Please provide a valid page and size")
		return
	}

	filter := &dtos.GetAllAssetsRequest{}
	err = context.ShouldBindJSON(filter)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Please provide a valid object")
		return
	}

	filterEncoded, err := json.Marshal(filter)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Could not encode the request")
		return
	}

	answer := chaincodeConnection.ReadInFabric(
		"GetAllAssets",
		string(page),
		string(size),
		string(filterEncoded),
	)
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
	return
}

func (c ChaincodeController) PostAsset(context *gin.Context) {
	newAsset := &dtos.PostAssetRequest{}
	err := context.ShouldBindJSON(newAsset)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Please provide an object")
		return
	}

	err = validation.ValidateObject(newAsset)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Please provide a valid object")
		return
	}

	newAssetEncoded, err := json.Marshal(newAsset)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Could not encode the request")
		return
	}
	answer := chaincodeConnection.PostInFabric("PostAsset", string(newAssetEncoded))
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
}

func (c ChaincodeController) GetById(context *gin.Context) {
	idString := context.Param("id")
	answer := chaincodeConnection.ReadInFabric("GetAssetById", idString)
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
}

func (c ChaincodeController) PutById(context *gin.Context) {
	idString := context.Param("id")
	newAsset := &dtos.PostAssetRequest{}
	context.BindJSON(newAsset)

	encodedAsset, err := json.Marshal(newAsset)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Something went wrong encoding the object")
		return
	}

	answer := chaincodeConnection.PostInFabric("PutAssetById", idString, string(encodedAsset))
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
}

func (c ChaincodeController) DeleteById(context *gin.Context) {
	idString := context.Param("id")
	answer := chaincodeConnection.PostInFabric("DeleteAssetById", idString)
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
}
