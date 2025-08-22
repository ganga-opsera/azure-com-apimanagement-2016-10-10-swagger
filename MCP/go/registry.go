package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_regions "github.com/apimanagementclient/mcp-server/tools/regions"
	tools_policysnippets "github.com/apimanagementclient/mcp-server/tools/policysnippets"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_regions.CreateRegions_listbyserviceTool(cfg),
		tools_policysnippets.CreatePolicysnippets_listbyserviceTool(cfg),
	}
}
