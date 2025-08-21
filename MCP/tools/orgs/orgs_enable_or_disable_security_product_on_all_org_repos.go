package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Orgs_enable_or_disable_security_product_on_all_org_reposHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		orgVal, ok := args["org"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org"), nil
		}
		org, ok := orgVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org"), nil
		}
		security_productVal, ok := args["security_product"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: security_product"), nil
		}
		security_product, ok := security_productVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: security_product"), nil
		}
		enablementVal, ok := args["enablement"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: enablement"), nil
		}
		enablement, ok := enablementVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: enablement"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/orgs/%s/%s/%s%s", cfg.BaseURL, org, security_product, enablement, queryString)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateOrgs_enable_or_disable_security_product_on_all_org_reposTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_security_product_enablement",
		mcp.WithDescription("Enable or disable a security feature for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("security_product", mcp.Required(), mcp.Description("The security feature to enable or disable.")),
		mcp.WithString("enablement", mcp.Required(), mcp.Description("The action to take.\n\n`enable_all` means to enable the specified security feature for all repositories in the organization.\n`disable_all` means to disable the specified security feature for all repositories in the organization.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_enable_or_disable_security_product_on_all_org_reposHandler(cfg),
	}
}
