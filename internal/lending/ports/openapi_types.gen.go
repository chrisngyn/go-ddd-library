// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for PatronType.
const (
	Regular    PatronType = "Regular"
	Researcher PatronType = "Researcher"
)

// CheckedOut defines model for CheckedOut.
type CheckedOut struct {
	BookId          openapi_types.UUID `json:"bookId"`
	CheckedOutAt    time.Time          `json:"checkedOutAt"`
	LibraryBranchId openapi_types.UUID `json:"libraryBranchId"`
}

// Hold defines model for Hold.
type Hold struct {
	BookId          openapi_types.UUID `json:"bookId"`
	From            time.Time          `json:"from"`
	IsOpenEnded     bool               `json:"isOpenEnded"`
	LibraryBranchId openapi_types.UUID `json:"libraryBranchId"`
	Till            time.Time          `json:"till"`
}

// OverdueCheckout defines model for OverdueCheckout.
type OverdueCheckout struct {
	BookId          openapi_types.UUID `json:"bookId"`
	LibraryBranchId openapi_types.UUID `json:"libraryBranchId"`
}

// PatronProfile defines model for PatronProfile.
type PatronProfile struct {
	CheckedOuts      []CheckedOut       `json:"checkedOuts"`
	Holds            []Hold             `json:"holds"`
	OverdueCheckouts []OverdueCheckout  `json:"overdueCheckouts"`
	PatronId         openapi_types.UUID `json:"patronId"`
	PatronType       PatronType         `json:"patronType"`
}

// PatronType defines model for PatronType.
type PatronType string

// CheckoutJSONBody defines parameters for Checkout.
type CheckoutJSONBody struct {
	BookId openapi_types.UUID `json:"bookId"`
}

// CancelHoldJSONBody defines parameters for CancelHold.
type CancelHoldJSONBody struct {
	BookId openapi_types.UUID `json:"bookId"`
}

// PlaceHoldJSONBody defines parameters for PlaceHold.
type PlaceHoldJSONBody struct {
	BookId    openapi_types.UUID `json:"bookId"`
	NumOfDays int                `json:"numOfDays"`
}

// CheckoutJSONRequestBody defines body for Checkout for application/json ContentType.
type CheckoutJSONRequestBody CheckoutJSONBody

// CancelHoldJSONRequestBody defines body for CancelHold for application/json ContentType.
type CancelHoldJSONRequestBody CancelHoldJSONBody

// PlaceHoldJSONRequestBody defines body for PlaceHold for application/json ContentType.
type PlaceHoldJSONRequestBody PlaceHoldJSONBody
