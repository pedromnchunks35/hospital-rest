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

func (c ChaincodeController) GetAssetHistoryById(context *gin.Context) {
	idString := context.Param("id")
	answerString, gateway := chaincodeConnection.ReadInFabric("GetHistoryAssetById", idString)
	defer gateway.Close()

	assetHistory := &dtos.KeyModification{}
	err := json.Unmarshal([]byte(answerString), assetHistory)
	if err != nil {
		apiResponses.ErrorResponse(context, answerString)
		return
	}

	apiResponses.SuccessResponse(context, assetHistory, "The fabric network retrieved an answer")
	return
}

func (c ChaincodeController) GetAllAssets(context *gin.Context) {
	pageString := context.Param("page")
	sizeString := context.Param("size")

	_, _, err := validation.ValidatePageAndSize(pageString, sizeString)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Please provide a valid page and size")
		return
	}

	filter := &dtos.Filter{}
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

	answer, gateway := chaincodeConnection.ReadInFabric(
		"GetAllAssets",
		pageString,
		sizeString,
		string(filterEncoded),
	)
	defer gateway.Close()

	assets := &[]dtos.PostAssetRequest{}
	err = json.Unmarshal([]byte(answer), assets)
	if err != nil {
		apiResponses.ErrorResponse(context, answer)
		return
	}

	apiResponses.SuccessResponse(context, assets, "The fabric network retrieved an answer")
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

	answer, gateway := chaincodeConnection.PostInFabric("CreateAsset", string(newAssetEncoded))
	defer gateway.Close()

	asset := &dtos.PostAssetRequest{}
	err = json.Unmarshal([]byte(answer), asset)
	if err != nil {
		apiResponses.ErrorResponse(context, answer)
		return
	}

	apiResponses.SuccessResponse(context, asset, "The fabric network retrieved an answer")
	return
}

func (c ChaincodeController) GetById(context *gin.Context) {
	idString := context.Param("id")
	answer, gateway := chaincodeConnection.ReadInFabric("GetAssetById", idString)
	defer gateway.Close()

	asset := &dtos.PostAssetRequest{}
	err := json.Unmarshal([]byte(answer), asset)
	if err != nil {
		apiResponses.ErrorResponse(context, answer)
		return
	}

	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
	return
}

func (c ChaincodeController) PatchById(context *gin.Context) {
	idString := context.Param("id")
	newAsset := &dtos.PostAssetRequest{}
	context.BindJSON(newAsset)

	encodedAsset, err := json.Marshal(newAsset)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, "Something went wrong encoding the object")
		return
	}

	answer, gateway := chaincodeConnection.PostInFabric("PatchAsset", string(encodedAsset), idString)
	defer gateway.Close()

	asset := &dtos.PostAssetRequest{}
	err = json.Unmarshal([]byte(answer), asset)
	if err != nil {
		apiResponses.ErrorResponse(context, answer)
		return
	}

	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
	return
}

func (c ChaincodeController) DeleteById(context *gin.Context) {
	idString := context.Param("id")
	answer, gateway := chaincodeConnection.PostInFabric("DeleteAssetById", idString)
	defer gateway.Close()
	apiResponses.SuccessResponse(context, answer, "The fabric network retrieved an answer")
	return
}
