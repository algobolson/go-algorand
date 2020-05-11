// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/transactions/dryrun)
	TransactionDryRun(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionDryRun converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionDryRun(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionDryRun(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.POST("/v2/transactions/dryrun", wrapper.TransactionDryRun, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
