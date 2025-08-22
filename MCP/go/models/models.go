package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RegionListResult represents the RegionListResult schema from the OpenAPI specification
type RegionListResult struct {
	Value []RegionContract `json:"value,omitempty"` // Lists of Regions.
}

// ErrorBodyContract represents the ErrorBodyContract schema from the OpenAPI specification
type ErrorBodyContract struct {
	Code string `json:"code,omitempty"` // Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Details []ErrorFieldContract `json:"details,omitempty"` // The list of invalid fields send in request, in case of validation error.
	Message string `json:"message,omitempty"` // Human-readable representation of the error.
}

// ErrorFieldContract represents the ErrorFieldContract schema from the OpenAPI specification
type ErrorFieldContract struct {
	Message string `json:"message,omitempty"` // Human-readable representation of property-level error.
	Target string `json:"target,omitempty"` // Property name.
	Code string `json:"code,omitempty"` // Property level error code.
}

// PolicySnippetContract represents the PolicySnippetContract schema from the OpenAPI specification
type PolicySnippetContract struct {
	Content string `json:"content,omitempty"` // Snippet content.
	Name string `json:"name,omitempty"` // Snippet name.
	Scope string `json:"scope,omitempty"` // Snippet scope.
	Tooltip string `json:"toolTip,omitempty"` // Snippet toolTip.
}

// PolicySnippetsCollection represents the PolicySnippetsCollection schema from the OpenAPI specification
type PolicySnippetsCollection struct {
	Value []PolicySnippetContract `json:"value,omitempty"` // Policy snippet value.
}

// RegionContract represents the RegionContract schema from the OpenAPI specification
type RegionContract struct {
	Name string `json:"name"` // Region name.
	Ismasterregion bool `json:"isMasterRegion"` // whether Region is the master region.
}
