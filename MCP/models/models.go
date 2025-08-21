package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Number int `json:"number,omitempty"` // The security alert number.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Closed_at string `json:"closed_at"`
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Description string `json:"description"`
	Id float64 `json:"id"`
	Public bool `json:"public"`
	Short_description string `json:"short_description"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Deleted_at string `json:"deleted_at"`
	Number int `json:"number"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"` // The global node ID of the installation.
	Id int `json:"id"` // The ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Context string `json:"context"`
	Required bool `json:"required,omitempty"`
	Id int `json:"id"`
	State string `json:"state"`
	Target_url string `json:"target_url"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Comments_url string `json:"comments_url"`
	Html_url string `json:"html_url"`
	Commit map[string]interface{} `json:"commit"`
	Parents []map[string]interface{} `json:"parents"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Uniques int `json:"uniques"`
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Node_id string `json:"node_id"`
	Url string `json:"url"` // The URL to the workflow run.
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Id int `json:"id"` // The ID of the workflow run.
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Path string `json:"path"` // The full path of the workflow
	Pull_requests []GeneratedType `json:"pull_requests"`
	Updated_at string `json:"updated_at"`
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
	Event string `json:"event"`
	Status string `json:"status"`
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Html_url string `json:"html_url"`
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Head_branch string `json:"head_branch"`
	Created_at string `json:"created_at"`
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Conclusion string `json:"conclusion"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User string `json:"user,omitempty"`
	Id string `json:"id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Comments int `json:"comments,omitempty"`
	Git_pull_url string `json:"git_pull_url,omitempty"`
	History []GeneratedType `json:"history,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Description string `json:"description,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Truncated bool `json:"truncated,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
	Url string `json:"url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Forks_url string `json:"forks_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Public bool `json:"public,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Created_at string `json:"created_at"`
	Public_gists int `json:"public_gists"`
	Public_repos int `json:"public_repos"`
	Members_url string `json:"members_url"`
	Blog string `json:"blog,omitempty"`
	Issues_url string `json:"issues_url"`
	Updated_at string `json:"updated_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Billing_email string `json:"billing_email,omitempty"`
	Company string `json:"company,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Email string `json:"email,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Is_verified bool `json:"is_verified,omitempty"`
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Followers int `json:"followers"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Name string `json:"name,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Repos_url string `json:"repos_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Url string `json:"url"`
	Location string `json:"location,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Login string `json:"login"`
	TypeField string `json:"type"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Collaborators int `json:"collaborators,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Events_url string `json:"events_url"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Following int `json:"following"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tags_url string `json:"tags_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Stargazers_url string `json:"stargazers_url"`
	Topics []string `json:"topics,omitempty"`
	Merges_url string `json:"merges_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Svn_url string `json:"svn_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Teams_url string `json:"teams_url"`
	Assignees_url string `json:"assignees_url"`
	Contents_url string `json:"contents_url"`
	Html_url string `json:"html_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Role_name string `json:"role_name,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Has_pages bool `json:"has_pages"`
	Subscribers_url string `json:"subscribers_url"`
	Open_issues int `json:"open_issues"`
	Trees_url string `json:"trees_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Forks_count int `json:"forks_count"`
	Git_refs_url string `json:"git_refs_url"`
	Git_url string `json:"git_url"`
	Watchers_count int `json:"watchers_count"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Master_branch string `json:"master_branch,omitempty"`
	Contributors_url string `json:"contributors_url"`
	License GeneratedType `json:"license"` // License Simple
	Id int `json:"id"` // Unique identifier of the repository
	Archived bool `json:"archived"` // Whether the repository is archived.
	Forks int `json:"forks"`
	Hooks_url string `json:"hooks_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Branches_url string `json:"branches_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Archive_url string `json:"archive_url"`
	Full_name string `json:"full_name"`
	Clone_url string `json:"clone_url"`
	Forks_url string `json:"forks_url"`
	Statuses_url string `json:"statuses_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Node_id string `json:"node_id"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Comments_url string `json:"comments_url"`
	Labels_url string `json:"labels_url"`
	Mirror_url string `json:"mirror_url"`
	Open_issues_count int `json:"open_issues_count"`
	Language string `json:"language"`
	Ssh_url string `json:"ssh_url"`
	Stargazers_count int `json:"stargazers_count"`
	Issue_events_url string `json:"issue_events_url"`
	Updated_at string `json:"updated_at"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Blobs_url string `json:"blobs_url"`
	Size int `json:"size"`
	Releases_url string `json:"releases_url"`
	Pushed_at string `json:"pushed_at"`
	Compare_url string `json:"compare_url"`
	Description string `json:"description"`
	Subscription_url string `json:"subscription_url"`
	Notifications_url string `json:"notifications_url"`
	Commits_url string `json:"commits_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Pulls_url string `json:"pulls_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Fork bool `json:"fork"`
	Keys_url string `json:"keys_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Name string `json:"name"` // The name of the repository.
	Collaborators_url string `json:"collaborators_url"`
	Homepage string `json:"homepage"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Issues_url string `json:"issues_url"`
	Network_count int `json:"network_count,omitempty"`
	Watchers int `json:"watchers"`
	Git_commits_url string `json:"git_commits_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Created_at string `json:"created_at"`
	Languages_url string `json:"languages_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Downloads_url string `json:"downloads_url"`
	Git_tags_url string `json:"git_tags_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Account GeneratedType `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Url string `json:"url"`
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Ping_url string `json:"ping_url"`
	Test_url string `json:"test_url"`
	Config map[string]interface{} `json:"config"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Created_at string `json:"created_at"`
	Last_response GeneratedType `json:"last_response"`
	TypeField string `json:"type"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the webhook.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	User map[string]interface{} `json:"user,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Node_id string `json:"node_id"`
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Archive_download_url string `json:"archive_download_url"`
	Created_at string `json:"created_at"`
	Expires_at string `json:"expires_at"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the artifact.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Before string `json:"before"`
	After string `json:"after"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
	Event string `json:"event,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Reason string `json:"reason"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	End_column int `json:"end_column"`
	Start_line int `json:"start_line"`
	Title string `json:"title"`
	Message string `json:"message"`
	Start_column int `json:"start_column"`
	Blob_href string `json:"blob_href"`
	End_line int `json:"end_line"`
	Raw_details string `json:"raw_details"`
	Annotation_level string `json:"annotation_level"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Url string `json:"url"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Size int `json:"size"`
	Download_url string `json:"download_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
	Id int `json:"id"`
	Import_url string `json:"import_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url"`
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Environment string `json:"environment"` // Name for the target deployment environment.
	Id int `json:"id"` // Unique identifier of the deployment
	Original_environment string `json:"original_environment,omitempty"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Has_free_trial bool `json:"has_free_trial"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	State string `json:"state"`
	Url string `json:"url"`
	Bullets []string `json:"bullets"`
	Unit_name string `json:"unit_name"`
	Accounts_url string `json:"accounts_url"`
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"`
	Number int `json:"number"`
	Price_model string `json:"price_model"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Target_type string `json:"target_type"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Account map[string]interface{} `json:"account"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository"` // A git repository
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 20 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/reference/repos#commits) to fetch additional commits. This limit is applied to timeline events only and isn't applied to webhook deliveries.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Created bool `json:"created"` // Whether this push created the `ref`.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Base_ref string `json:"base_ref"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Head_commit map[string]interface{} `json:"head_commit"`
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Milestone map[string]interface{} `json:"milestone"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer map[string]interface{} `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester map[string]interface{} `json:"requester"`
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Single_file_name string `json:"single_file_name"`
	Id int `json:"id"` // The ID of the installation.
	Html_url string `json:"html_url"`
	Contact_email string `json:"contact_email,omitempty"`
	Events []string `json:"events"`
	Target_type string `json:"target_type"`
	Account interface{} `json:"account"`
	App_slug string `json:"app_slug"`
	App_id int `json:"app_id"`
	Created_at string `json:"created_at"`
	Suspended_at string `json:"suspended_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Updated_at string `json:"updated_at"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Access_tokens_url string `json:"access_tokens_url"`
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
	Id int `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Html_url string `json:"html_url"`
	Site_admin bool `json:"site_admin"`
	Gravatar_id string `json:"gravatar_id"`
	Gists_url string `json:"gists_url"`
	Repos_url string `json:"repos_url"`
	Login string `json:"login"`
	Email string `json:"email,omitempty"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	Received_events_url string `json:"received_events_url"`
	Id int `json:"id"`
	Starred_url string `json:"starred_url"`
	TypeField string `json:"type"`
	Following_url string `json:"following_url"`
	Name string `json:"name,omitempty"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author map[string]interface{} `json:"author"`
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Tree_id string `json:"tree_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Number int `json:"number"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Columns_url string `json:"columns_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Permissions map[string]interface{} `json:"permissions"`
	Node_id string `json:"node_id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Owner_url string `json:"owner_url"`
	Body string `json:"body"`
	State string `json:"state"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
	Avatar_url string `json:"avatar_url"`
	Name string `json:"name"` // The name of the enterprise.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Description string `json:"description"`
	Target_url string `json:"target_url"`
	Context string `json:"context"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester map[string]interface{} `json:"requester"`
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Assignee map[string]interface{} `json:"assignee"`
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Commit map[string]interface{} `json:"commit"`
	Updated_at string `json:"updated_at"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Id int `json:"id"` // The unique identifier of the status.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sha string `json:"sha"` // The Commit SHA.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Avatar_url string `json:"avatar_url,omitempty"`
	Created_at string `json:"created_at"`
	Target_url string `json:"target_url"` // The optional link added to the status.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Description string `json:"description"` // The optional human-readable description added to the status.
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Context string `json:"context"`
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Guid string `json:"guid"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Archive_url string `json:"archive_url,omitempty"`
	Exclude_metadata bool `json:"exclude_metadata"`
	State string `json:"state"`
	Url string `json:"url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Lock_repositories bool `json:"lock_repositories"`
	Node_id string `json:"node_id"`
	Exclude []interface{} `json:"exclude,omitempty"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Exclude_releases bool `json:"exclude_releases"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // URL for the team
	Description string `json:"description"` // Description of the team
	Id int `json:"id"` // Unique identifier of the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Html_url string `json:"html_url"`
	Repositories_url string `json:"repositories_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Git_url string `json:"git_url"`
	Language string `json:"language,omitempty"`
	Name string `json:"name"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Html_url string `json:"html_url"`
	Line_numbers []string `json:"line_numbers,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Target string `json:"target"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author map[string]interface{} `json:"author"`
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Tree_id string `json:"tree_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Statuses_url string `json:"statuses_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Deletions int `json:"deletions"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Additions int `json:"additions"`
	Labels []map[string]interface{} `json:"labels"`
	Commits_url string `json:"commits_url"`
	Url string `json:"url"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Review_comments_url string `json:"review_comments_url"`
	Links map[string]interface{} `json:"_links"`
	Merged_at string `json:"merged_at"`
	Changed_files int `json:"changed_files"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Title string `json:"title"` // The title of the pull request.
	User GeneratedType `json:"user"` // A GitHub user.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Created_at string `json:"created_at"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Comments int `json:"comments"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Patch_url string `json:"patch_url"`
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Issue_url string `json:"issue_url"`
	Locked bool `json:"locked"`
	Review_comments int `json:"review_comments"`
	Review_comment_url string `json:"review_comment_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Body string `json:"body"`
	Closed_at string `json:"closed_at"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Html_url string `json:"html_url"`
	Commits int `json:"commits"`
	Base map[string]interface{} `json:"base"`
	Updated_at string `json:"updated_at"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
	Comments_url string `json:"comments_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Merged bool `json:"merged"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Head map[string]interface{} `json:"head"`
	Mergeable_state string `json:"mergeable_state"`
	Diff_url string `json:"diff_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Member map[string]interface{} `json:"member"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Action string `json:"action"`
	Sender map[string]interface{} `json:"sender"`
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Path string `json:"path"` // The file path in the repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
	User GeneratedType `json:"user"` // Collaborator
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Zipball_url string `json:"zipball_url"`
	Mentions_count int `json:"mentions_count,omitempty"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Assets_url string `json:"assets_url"`
	Body_text string `json:"body_text,omitempty"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Html_url string `json:"html_url"`
	Published_at string `json:"published_at"`
	Upload_url string `json:"upload_url"`
	Assets []GeneratedType `json:"assets"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Body string `json:"body,omitempty"`
	Id int `json:"id"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user map[string]interface{} `json:"blocked_user"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Updated_at string `json:"updated_at"`
	Location string `json:"location"` // The Azure region where this codespace is located.
	Recent_folders []string `json:"recent_folders"`
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Url string `json:"url"` // API URL for this codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Id int `json:"id"`
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	State string `json:"state"` // State of this codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Created_at string `json:"created_at"`
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Name string `json:"name"` // Automatically generated name of this codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
	Payload map[string]interface{} `json:"payload"`
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Labels_url string `json:"labels_url"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Number int `json:"number"` // The number of the milestone.
	Due_on string `json:"due_on"`
	State string `json:"state"` // The state of the milestone.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Url string `json:"url"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Open_issues int `json:"open_issues"`
	Title string `json:"title"` // The title of the milestone.
	Closed_issues int `json:"closed_issues"`
	Closed_at string `json:"closed_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/reference/repos#list-github-pages-builds) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Details interface{} `json:"details"`
	TypeField string `json:"type"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	State_reason string `json:"state_reason,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Url string `json:"url"`
	Protection_rules []interface{} `json:"protection_rules,omitempty"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Name string `json:"name"` // The name of the environment.
	Html_url string `json:"html_url"`
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Id int `json:"id"` // The id of the environment.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the package version.
	License string `json:"license,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Package_html_url string `json:"package_html_url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the package version.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role_name string `json:"role_name"`
	Followers_url string `json:"followers_url"`
	TypeField string `json:"type"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Repos_url string `json:"repos_url"`
	Node_id string `json:"node_id"`
	Starred_url string `json:"starred_url"`
	Site_admin bool `json:"site_admin"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	Url string `json:"url"`
	Name string `json:"name,omitempty"`
	Id int `json:"id"`
	Organizations_url string `json:"organizations_url"`
	Email string `json:"email,omitempty"`
	Html_url string `json:"html_url"`
	Gravatar_id string `json:"gravatar_id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Avatar_url string `json:"avatar_url"`
	Received_events_url string `json:"received_events_url"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Member map[string]interface{} `json:"member"`
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Sender map[string]interface{} `json:"sender"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Files []GeneratedType `json:"files,omitempty"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Committer GeneratedType `json:"committer"` // A GitHub user.
	Node_id string `json:"node_id"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Slug string `json:"slug"`
	Members_count int `json:"members_count"`
	Name string `json:"name"` // Name of the team
	Html_url string `json:"html_url"`
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Members_url string `json:"members_url"`
	Created_at string `json:"created_at"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Repositories_url string `json:"repositories_url"`
	Id int `json:"id"` // Unique identifier of the team
	Updated_at string `json:"updated_at"`
	Organization GeneratedType `json:"organization"` // Team Organization
	Url string `json:"url"` // URL for the team
	Repos_count int `json:"repos_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Location GeneratedType `json:"location"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding"`
	Html_url string `json:"html_url"`
	Target string `json:"target,omitempty"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Path string `json:"path"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Action string `json:"action"`
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Details_url string `json:"details_url"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Name string `json:"name"` // The name of the check.
	Started_at string `json:"started_at"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Node_id string `json:"node_id"`
	Completed_at string `json:"completed_at"`
	Id int `json:"id"` // The id of the check.
	Pull_requests []GeneratedType `json:"pull_requests"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	External_id string `json:"external_id"`
	Output map[string]interface{} `json:"output"`
	Url string `json:"url"`
	Conclusion string `json:"conclusion"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Html_url string `json:"html_url"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Clone_url string `json:"clone_url"`
	Topics []string `json:"topics,omitempty"`
	Git_url string `json:"git_url"`
	Name string `json:"name"` // The name of the repository.
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Open_issues_count int `json:"open_issues_count"`
	Stargazers_url string `json:"stargazers_url"`
	Fork bool `json:"fork"`
	Subscription_url string `json:"subscription_url"`
	Svn_url string `json:"svn_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Forks_count int `json:"forks_count"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Issue_events_url string `json:"issue_events_url"`
	Homepage string `json:"homepage"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Events_url string `json:"events_url"`
	Network_count int `json:"network_count,omitempty"`
	Releases_url string `json:"releases_url"`
	Contents_url string `json:"contents_url"`
	Full_name string `json:"full_name"`
	Trees_url string `json:"trees_url"`
	Updated_at string `json:"updated_at"`
	Watchers_count int `json:"watchers_count"`
	Forks int `json:"forks"`
	Watchers int `json:"watchers"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Milestones_url string `json:"milestones_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Comments_url string `json:"comments_url"`
	Has_pages bool `json:"has_pages"`
	Tags_url string `json:"tags_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Language string `json:"language"`
	Ssh_url string `json:"ssh_url"`
	Git_commits_url string `json:"git_commits_url"`
	License GeneratedType `json:"license"` // License Simple
	Compare_url string `json:"compare_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	Downloads_url string `json:"downloads_url"`
	Hooks_url string `json:"hooks_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Node_id string `json:"node_id"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Statuses_url string `json:"statuses_url"`
	Url string `json:"url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Master_branch string `json:"master_branch,omitempty"`
	Merges_url string `json:"merges_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Subscribers_url string `json:"subscribers_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Assignees_url string `json:"assignees_url"`
	Description string `json:"description"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Notifications_url string `json:"notifications_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Contributors_url string `json:"contributors_url"`
	Created_at string `json:"created_at"`
	Archive_url string `json:"archive_url"`
	Stargazers_count int `json:"stargazers_count"`
	Git_tags_url string `json:"git_tags_url"`
	Pushed_at string `json:"pushed_at"`
	Collaborators_url string `json:"collaborators_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Keys_url string `json:"keys_url"`
	Mirror_url string `json:"mirror_url"`
	Labels_url string `json:"labels_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Forks_url string `json:"forks_url"`
	Git_refs_url string `json:"git_refs_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Issues_url string `json:"issues_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Languages_url string `json:"languages_url"`
	Commits_url string `json:"commits_url"`
	Branches_url string `json:"branches_url"`
	Blobs_url string `json:"blobs_url"`
	Teams_url string `json:"teams_url"`
	Html_url string `json:"html_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Pulls_url string `json:"pulls_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Open_issues int `json:"open_issues"`
	Deployments_url string `json:"deployments_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the issue comment
	Reactions GeneratedType `json:"reactions,omitempty"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the issue comment
	Issue_url string `json:"issue_url"`
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Followers_url string `json:"followers_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Email string `json:"email,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Name string `json:"name,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Contributions int `json:"contributions"`
	Login string `json:"login,omitempty"`
	TypeField string `json:"type"`
	Id int `json:"id,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Url string `json:"url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Html_url string `json:"html_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_url string `json:"pull_request_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body string `json:"body"` // The text of the review.
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the review
	Links map[string]interface{} `json:"_links"`
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Url string `json:"url"`
	Conditions []string `json:"conditions"`
	Description string `json:"description"`
	Key string `json:"key"`
	Limitations []string `json:"limitations"`
	Node_id string `json:"node_id"`
	Implementation string `json:"implementation"`
	Permissions []string `json:"permissions"`
	Spdx_id string `json:"spdx_id"`
	Body string `json:"body"`
	Featured bool `json:"featured"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Collaborators int `json:"collaborators,omitempty"`
	Login string `json:"login"`
	Events_url string `json:"events_url"`
	Public_repos int `json:"public_repos"`
	Followers_url string `json:"followers_url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Repos_url string `json:"repos_url"`
	Received_events_url string `json:"received_events_url"`
	Blog string `json:"blog"`
	Starred_url string `json:"starred_url"`
	Email string `json:"email"`
	Hireable bool `json:"hireable"`
	Public_gists int `json:"public_gists"`
	Following_url string `json:"following_url"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Site_admin bool `json:"site_admin"`
	Subscriptions_url string `json:"subscriptions_url"`
	Url string `json:"url"`
	Id int `json:"id"`
	Private_gists int `json:"private_gists,omitempty"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Company string `json:"company"`
	Bio string `json:"bio"`
	Updated_at string `json:"updated_at"`
	Name string `json:"name"`
	TypeField string `json:"type"`
	Location string `json:"location"`
	Gists_url string `json:"gists_url"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Following int `json:"following"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Followers int `json:"followers"`
	Plan map[string]interface{} `json:"plan,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Count int `json:"count"`
	Referrer string `json:"referrer"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [commit comment](https://docs.github.com/rest/reference/repos#get-a-commit-comment) resource.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Insecure_ssl GeneratedType `json:"insecure_ssl,omitempty"`
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uniques int `json:"uniques"`
	Views []Traffic `json:"views"`
	Count int `json:"count"`
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Blocked_user map[string]interface{} `json:"blocked_user"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Column_name string `json:"column_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Url string `json:"url"`
	Members_url string `json:"members_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Slug string `json:"slug"`
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Permission string `json:"permission"`
	Privacy string `json:"privacy,omitempty"`
	Name string `json:"name"`
	Repositories_url string `json:"repositories_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscriptions_url string `json:"subscriptions_url"`
	Public_gists int `json:"public_gists"`
	Disk_usage int `json:"disk_usage"`
	Private_gists int `json:"private_gists"`
	Business_plus bool `json:"business_plus,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Received_events_url string `json:"received_events_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Company string `json:"company"`
	Created_at string `json:"created_at"`
	Email string `json:"email"`
	Collaborators int `json:"collaborators"`
	Starred_url string `json:"starred_url"`
	Location string `json:"location"`
	Site_admin bool `json:"site_admin"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Blog string `json:"blog"`
	Html_url string `json:"html_url"`
	Public_repos int `json:"public_repos"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Login string `json:"login"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Following int `json:"following"`
	Repos_url string `json:"repos_url"`
	Events_url string `json:"events_url"`
	Bio string `json:"bio"`
	Gists_url string `json:"gists_url"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Followers_url string `json:"followers_url"`
	Following_url string `json:"following_url"`
	Updated_at string `json:"updated_at"`
	Followers int `json:"followers"`
	TypeField string `json:"type"`
	Owned_private_repos int `json:"owned_private_repos"`
	Total_private_repos int `json:"total_private_repos"`
	Name string `json:"name"`
	Hireable bool `json:"hireable"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion map[string]interface{} `json:"discussion"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
	Payload string `json:"payload"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Ping_url string `json:"ping_url"`
	Updated_at string `json:"updated_at"`
	Active bool `json:"active"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Events []string `json:"events"`
	Name string `json:"name"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Config map[string]interface{} `json:"config"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_group map[string]interface{} `json:"merge_group"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The action that was performed.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Behind_by int `json:"behind_by"`
	Patch_url string `json:"patch_url"`
	Permalink_url string `json:"permalink_url"`
	Status string `json:"status"`
	Ahead_by int `json:"ahead_by"`
	Commits []Commit `json:"commits"`
	Diff_url string `json:"diff_url"`
	Html_url string `json:"html_url"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Total_commits int `json:"total_commits"`
	Url string `json:"url"`
	Base_commit Commit `json:"base_commit"` // Commit
	Files []GeneratedType `json:"files,omitempty"`
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Timeline_url string `json:"timeline_url,omitempty"`
	User map[string]interface{} `json:"user"`
	Repository_url string `json:"repository_url"`
	Id int `json:"id"`
	Title string `json:"title"`
	Updated_at string `json:"updated_at"`
	Comments int `json:"comments"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Category map[string]interface{} `json:"category"`
	Locked bool `json:"locked"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Node_id string `json:"node_id"`
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Created_at string `json:"created_at"`
	Number int `json:"number"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Active_lock_reason string `json:"active_lock_reason"`
	Answer_html_url string `json:"answer_html_url"`
	Body string `json:"body"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Url string `json:"url"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Changes map[string]interface{} `json:"changes"`
	Action string `json:"action"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Team_count int `json:"team_count"`
	Created_at string `json:"created_at"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Failed_at string `json:"failed_at,omitempty"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Node_id string `json:"node_id"`
	Email string `json:"email"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Login string `json:"login"`
	Id int `json:"id"`
	Role string `json:"role"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	User_url string `json:"user_url"`
	Links map[string]interface{} `json:"_links"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Timeline_url string `json:"timeline_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Position int `json:"position"`
	Commit_id string `json:"commit_id"`
	Line int `json:"line"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Source_repository GeneratedType `json:"source_repository"` // Minimal Repository
	Badge_url string `json:"badge_url"`
	Id int `json:"id"`
	Path string `json:"path"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Html_url string `json:"html_url,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Commit_sha string `json:"commit_sha,omitempty"`
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parents []map[string]interface{} `json:"parents"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Event string `json:"event,omitempty"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Tree map[string]interface{} `json:"tree"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Sha string `json:"sha"` // SHA for the commit
	Verification map[string]interface{} `json:"verification"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Client_secret string `json:"client_secret,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Html_url string `json:"html_url"`
	Pem string `json:"pem,omitempty"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	External_url string `json:"external_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Events []string `json:"events"` // The list of events for the GitHub app
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Name string `json:"name"` // The name of the GitHub app
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body,omitempty"`
	State string `json:"state"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Score float64 `json:"score"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Labels_url string `json:"labels_url"`
	Events_url string `json:"events_url"`
	Node_id string `json:"node_id"`
	Labels []map[string]interface{} `json:"labels"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Locked bool `json:"locked"`
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	State_reason string `json:"state_reason,omitempty"`
	Draft bool `json:"draft,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Comments int `json:"comments"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Number int `json:"number"`
	Repository_url string `json:"repository_url"`
	Html_url string `json:"html_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Updated_at string `json:"updated_at"`
	Comments_url string `json:"comments_url"`
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name pattern that branches must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	Id int `json:"id,omitempty"` // The unique identifier of the branch policy.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Repository_announcement_banners string `json:"repository_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for a repository.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects beta (where available).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // Unique identifier of the webhook delivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Event string `json:"event"` // The event that triggered the delivery.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	State string `json:"state"` // The state of the Dependabot alert.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Number int `json:"number"` // The pull request number.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // The repository's current description.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the repository invitation.
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Url string `json:"url"` // URL for the repository invitation
	Created_at string `json:"created_at"`
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Node_id string `json:"node_id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. **Note**: The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Topics []string `json:"topics,omitempty"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Statuses_url string `json:"statuses_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Fork bool `json:"fork"`
	Language string `json:"language"`
	Teams_url string `json:"teams_url"`
	Pushed_at string `json:"pushed_at"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Master_branch string `json:"master_branch,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Milestones_url string `json:"milestones_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Ssh_url string `json:"ssh_url"`
	Git_tags_url string `json:"git_tags_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Subscribers_url string `json:"subscribers_url"`
	Open_issues int `json:"open_issues"`
	Stargazers_url string `json:"stargazers_url"`
	Trees_url string `json:"trees_url"`
	Deployments_url string `json:"deployments_url"`
	Forks int `json:"forks"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Downloads_url string `json:"downloads_url"`
	Watchers_count int `json:"watchers_count"`
	Node_id string `json:"node_id"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Languages_url string `json:"languages_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Releases_url string `json:"releases_url"`
	Subscription_url string `json:"subscription_url"`
	Merges_url string `json:"merges_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Issues_url string `json:"issues_url"`
	Keys_url string `json:"keys_url"`
	Issue_events_url string `json:"issue_events_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Svn_url string `json:"svn_url"`
	Blobs_url string `json:"blobs_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Notifications_url string `json:"notifications_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Updated_at string `json:"updated_at"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Has_pages bool `json:"has_pages"`
	Compare_url string `json:"compare_url"`
	Pulls_url string `json:"pulls_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Html_url string `json:"html_url"`
	Commits_url string `json:"commits_url"`
	Contributors_url string `json:"contributors_url"`
	Hooks_url string `json:"hooks_url"`
	Assignees_url string `json:"assignees_url"`
	License GeneratedType `json:"license"` // License Simple
	Labels_url string `json:"labels_url"`
	Forks_count int `json:"forks_count"`
	Contents_url string `json:"contents_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Full_name string `json:"full_name"`
	Watchers int `json:"watchers"`
	Description string `json:"description"`
	Tags_url string `json:"tags_url"`
	Url string `json:"url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Events_url string `json:"events_url"`
	Mirror_url string `json:"mirror_url"`
	Created_at string `json:"created_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Network_count int `json:"network_count,omitempty"`
	Name string `json:"name"` // The name of the repository.
	Archive_url string `json:"archive_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Forks_url string `json:"forks_url"`
	Comments_url string `json:"comments_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Clone_url string `json:"clone_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Git_refs_url string `json:"git_refs_url"`
	Branches_url string `json:"branches_url"`
	Git_url string `json:"git_url"`
	Stargazers_count int `json:"stargazers_count"`
	Homepage string `json:"homepage"`
	Open_issues_count int `json:"open_issues_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	DefaultField bool `json:"default"`
	Score float64 `json:"score"`
	Color string `json:"color"`
	Description string `json:"description"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Created_at string `json:"created_at"`
	Discussion_url string `json:"discussion_url"`
	Html_url string `json:"html_url"`
	Last_edited_at string `json:"last_edited_at"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_html string `json:"body_html"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body string `json:"body"` // The main text of the comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Issue_url string `json:"issue_url"`
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"`
	Patch_url string `json:"patch_url"`
	Updated_at string `json:"updated_at"`
	Review_comments_url string `json:"review_comments_url"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Review_comment_url string `json:"review_comment_url"`
	Number int `json:"number"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Diff_url string `json:"diff_url"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Title string `json:"title"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments_url string `json:"comments_url"`
	Closed_at string `json:"closed_at"`
	Html_url string `json:"html_url"`
	Merged_at string `json:"merged_at"`
	State string `json:"state"`
	Commits_url string `json:"commits_url"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Locked bool `json:"locked"`
	Head map[string]interface{} `json:"head"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Base map[string]interface{} `json:"base"`
	Labels []map[string]interface{} `json:"labels"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository GeneratedType `json:"repository"` // Full Repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone,omitempty"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Hook map[string]interface{} `json:"hook"` // The modified webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Event string `json:"event"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Old_answer map[string]interface{} `json:"old_answer"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Token_last_eight string `json:"token_last_eight"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Note_url string `json:"note_url"`
	Token string `json:"token"`
	Created_at string `json:"created_at"`
	Note string `json:"note"`
	Id int `json:"id"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	Url string `json:"url"`
	App map[string]interface{} `json:"app"`
	Installation GeneratedType `json:"installation,omitempty"`
	Fingerprint string `json:"fingerprint"`
	Expires_at string `json:"expires_at"`
	Hashed_token string `json:"hashed_token"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Number int `json:"number,omitempty"` // The security alert number.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reactions GeneratedType `json:"reactions,omitempty"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Links map[string]interface{} `json:"_links"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Pull_request_review_id int `json:"pull_request_review_id"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Original_position int `json:"original_position"`
	Position int `json:"position"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Path string `json:"path"`
	Node_id string `json:"node_id"`
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Url string `json:"url"`
	Diff_hunk string `json:"diff_hunk"`
	Original_commit_id string `json:"original_commit_id"`
	Pull_request_url string `json:"pull_request_url"`
	Created_at string `json:"created_at"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	User GeneratedType `json:"user"` // A GitHub user.
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Closed_issues int `json:"closed_issues"`
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	Labels_url string `json:"labels_url"`
	Number int `json:"number"` // The number of the milestone.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Open_issues int `json:"open_issues"`
	Due_on string `json:"due_on"`
	Description string `json:"description"`
	State string `json:"state"` // The state of the milestone.
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Title string `json:"title"` // The title of the milestone.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Status string `json:"status"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
	Base_branch string `json:"base_branch,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
	Change_status map[string]interface{} `json:"change_status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
	Commit_id string `json:"commit_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Commit_id string `json:"commit_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Milestone map[string]interface{} `json:"milestone"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Events []string `json:"events"` // The list of events for the GitHub app
	External_url string `json:"external_url"`
	Pem string `json:"pem,omitempty"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Client_secret string `json:"client_secret,omitempty"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Week int `json:"week"`
	Days []int `json:"days"`
	Total int `json:"total"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Node_id string `json:"node_id"`
	Verification map[string]interface{} `json:"verification"`
	Tree map[string]interface{} `json:"tree"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id,omitempty"`
	Content_node_id string `json:"content_node_id"`
	Id float64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Archived_at string `json:"archived_at"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Project_node_id string `json:"project_node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Twitter_username string `json:"twitter_username,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Billing_email string `json:"billing_email,omitempty"`
	Html_url string `json:"html_url"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Issues_url string `json:"issues_url"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Repos_url string `json:"repos_url"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Plan map[string]interface{} `json:"plan,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Public_members_url string `json:"public_members_url"`
	Members_url string `json:"members_url"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Private_gists int `json:"private_gists,omitempty"`
	Company string `json:"company,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Email string `json:"email,omitempty"`
	Followers int `json:"followers"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Following int `json:"following"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Events_url string `json:"events_url"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Url string `json:"url"`
	Blog string `json:"blog,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	TypeField string `json:"type"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Disk_usage int `json:"disk_usage,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Location string `json:"location,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // Whether dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Public_repos int `json:"public_repos"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Public_gists int `json:"public_gists"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // Ref at which the workflow file will be selected
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	State string `json:"state"` // State of the required workflow
	Path string `json:"path"` // Path of the workflow file
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Scope string `json:"scope"` // Scope of the required workflow
	Created_at string `json:"created_at"`
	Name string `json:"name"` // Name present in the workflow file
	Updated_at string `json:"updated_at"`
	Id float64 `json:"id"` // Unique identifier for a required workflow
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Id int64 `json:"id"`
	Name string `json:"name"` // The name of the label.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protection_url string `json:"protection_url,omitempty"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issues_url string `json:"issues_url"`
	Repos_url string `json:"repos_url"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Description string `json:"description"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Hooks_url string `json:"hooks_url"`
	Login string `json:"login"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Login string `json:"login"`
	Site_admin bool `json:"site_admin"`
	Html_url string `json:"html_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Url string `json:"url"`
	Received_events_url string `json:"received_events_url"`
	Gravatar_id string `json:"gravatar_id"`
	Following_url string `json:"following_url"`
	Gists_url string `json:"gists_url"`
	Starred_url string `json:"starred_url"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Node_id string `json:"node_id"`
	TypeField string `json:"type"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Rule GeneratedType `json:"rule"`
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType `json:"tool"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Number int `json:"number"` // The security alert number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reactions GeneratedType `json:"reactions,omitempty"`
	Url string `json:"url"` // URL for the pull request review comment
	Body string `json:"body"` // The text of the comment.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	User GeneratedType `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"` // The ID of the pull request review comment.
	Position int `json:"position"` // The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Content_type string `json:"content_type"`
	Id int `json:"id"`
	Name string `json:"name"` // The file name of the asset.
	State string `json:"state"` // State of the release asset.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Url string `json:"url"`
	Browser_download_url string `json:"browser_download_url"`
	Download_count int `json:"download_count"`
	Label string `json:"label"`
	Size int `json:"size"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate GeneratedType `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository_url string `json:"repository_url"`
	Sha string `json:"sha"`
	Statuses_url string `json:"statuses_url"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Task string `json:"task"` // Parameter to specify a task to execute
	Updated_at string `json:"updated_at"`
	Payload interface{} `json:"payload"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Environment string `json:"environment"` // Name for the target deployment environment.
	Id int `json:"id"` // Unique identifier of the deployment
	Original_environment string `json:"original_environment,omitempty"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"`
	Description string `json:"description"`
	Name string `json:"name"`
	Released string `json:"released"`
	Featured bool `json:"featured"`
	Updated_at string `json:"updated_at"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Short_description string `json:"short_description"`
	Created_at string `json:"created_at"`
	Curated bool `json:"curated"`
	Logo_url string `json:"logo_url,omitempty"`
	Created_by string `json:"created_by"`
	Repository_count int `json:"repository_count,omitempty"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Followers_url string `json:"followers_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Notifications_url string `json:"notifications_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Issue_search_url string `json:"issue_search_url"`
	Commit_search_url string `json:"commit_search_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Repository_search_url string `json:"repository_search_url"`
	Emails_url string `json:"emails_url"`
	Current_user_url string `json:"current_user_url"`
	Hub_url string `json:"hub_url"`
	Issues_url string `json:"issues_url"`
	Public_gists_url string `json:"public_gists_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Events_url string `json:"events_url"`
	Authorizations_url string `json:"authorizations_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Label_search_url string `json:"label_search_url"`
	Feeds_url string `json:"feeds_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Organization_url string `json:"organization_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Keys_url string `json:"keys_url"`
	Repository_url string `json:"repository_url"`
	Emojis_url string `json:"emojis_url"`
	Gists_url string `json:"gists_url"`
	User_search_url string `json:"user_search_url"`
	Following_url string `json:"following_url"`
	User_url string `json:"user_url"`
	Starred_url string `json:"starred_url"`
	Code_search_url string `json:"code_search_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Statuses_url string `json:"statuses_url"`
	Git_tags_url string `json:"git_tags_url"`
	Has_pages bool `json:"has_pages"`
	Description string `json:"description"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Size int `json:"size"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Watchers int `json:"watchers"`
	Watchers_count int `json:"watchers_count"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Has_downloads bool `json:"has_downloads"`
	Issue_comment_url string `json:"issue_comment_url"`
	Branches_url string `json:"branches_url"`
	Git_commits_url string `json:"git_commits_url"`
	Pushed_at string `json:"pushed_at"`
	Git_refs_url string `json:"git_refs_url"`
	Topics []string `json:"topics,omitempty"`
	Score float64 `json:"score"`
	Stargazers_count int `json:"stargazers_count"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Mirror_url string `json:"mirror_url"`
	Homepage string `json:"homepage"`
	Fork bool `json:"fork"`
	Commits_url string `json:"commits_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Languages_url string `json:"languages_url"`
	Keys_url string `json:"keys_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Has_wiki bool `json:"has_wiki"`
	Contents_url string `json:"contents_url"`
	Tags_url string `json:"tags_url"`
	Subscription_url string `json:"subscription_url"`
	Deployments_url string `json:"deployments_url"`
	Merges_url string `json:"merges_url"`
	Has_issues bool `json:"has_issues"`
	Hooks_url string `json:"hooks_url"`
	Open_issues_count int `json:"open_issues_count"`
	Downloads_url string `json:"downloads_url"`
	Updated_at string `json:"updated_at"`
	Forks int `json:"forks"`
	Node_id string `json:"node_id"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Milestones_url string `json:"milestones_url"`
	Created_at string `json:"created_at"`
	Compare_url string `json:"compare_url"`
	License GeneratedType `json:"license"` // License Simple
	Labels_url string `json:"labels_url"`
	Forks_count int `json:"forks_count"`
	Pulls_url string `json:"pulls_url"`
	Events_url string `json:"events_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Clone_url string `json:"clone_url"`
	Blobs_url string `json:"blobs_url"`
	Default_branch string `json:"default_branch"`
	Open_issues int `json:"open_issues"`
	Ssh_url string `json:"ssh_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Archive_url string `json:"archive_url"`
	Is_template bool `json:"is_template,omitempty"`
	Name string `json:"name"`
	Collaborators_url string `json:"collaborators_url"`
	Stargazers_url string `json:"stargazers_url"`
	Teams_url string `json:"teams_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Releases_url string `json:"releases_url"`
	Trees_url string `json:"trees_url"`
	Forks_url string `json:"forks_url"`
	Archived bool `json:"archived"`
	Url string `json:"url"`
	Full_name string `json:"full_name"`
	Has_projects bool `json:"has_projects"`
	Assignees_url string `json:"assignees_url"`
	Git_url string `json:"git_url"`
	Subscribers_url string `json:"subscribers_url"`
	Issues_url string `json:"issues_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Language string `json:"language"`
	Comments_url string `json:"comments_url"`
	Contributors_url string `json:"contributors_url"`
	Id int `json:"id"`
	Private bool `json:"private"`
	Html_url string `json:"html_url"`
	Notifications_url string `json:"notifications_url"`
	Svn_url string `json:"svn_url"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Subject map[string]interface{} `json:"subject"`
	Subscription_url string `json:"subscription_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Id string `json:"id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Unread bool `json:"unread"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Columns_url string `json:"columns_url"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Name string `json:"name"` // Name of the project
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Owner_url string `json:"owner_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Number int `json:"number"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Description string `json:"description"` // The repository description.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Name string `json:"name"` // The name of the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Private bool `json:"private"` // Whether the repository is private.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Id int `json:"id"` // A unique identifier of the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Check_run_url string `json:"check_run_url"`
	Url string `json:"url"`
	Run_url string `json:"run_url"`
	Name string `json:"name"` // The name of the job.
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Node_id string `json:"node_id"`
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Html_url string `json:"html_url"`
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Id int `json:"id"` // The id of the job.
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Content map[string]interface{} `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The security alert number.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType `json:"tool"`
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	State string `json:"state"` // State of a code scanning alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Rule GeneratedType `json:"rule"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deletions int `json:"deletions"`
	Sha string `json:"sha"`
	Changes int `json:"changes"`
	Filename string `json:"filename"`
	Patch string `json:"patch,omitempty"`
	Status string `json:"status"`
	Additions int `json:"additions"`
	Blob_url string `json:"blob_url"`
	Contents_url string `json:"contents_url"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Raw_url string `json:"raw_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Last_edited_at string `json:"last_edited_at"`
	Comments_count int `json:"comments_count"`
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Url string `json:"url"`
	Team_url string `json:"team_url"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Body string `json:"body"` // The main text of the discussion.
	Comments_url string `json:"comments_url"`
	Body_html string `json:"body_html"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization administrators.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Title string `json:"title"` // The title of the discussion.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscribers_url string `json:"subscribers_url"`
	Fork bool `json:"fork"`
	Commits_url string `json:"commits_url"`
	Svn_url string `json:"svn_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Git_url string `json:"git_url"`
	Archived bool `json:"archived"`
	Hooks_url string `json:"hooks_url"`
	Git_refs_url string `json:"git_refs_url"`
	Has_wiki bool `json:"has_wiki"`
	Master_branch string `json:"master_branch,omitempty"`
	Has_discussions bool `json:"has_discussions"`
	Description string `json:"description"`
	Branches_url string `json:"branches_url"`
	Comments_url string `json:"comments_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_issues bool `json:"has_issues"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Blobs_url string `json:"blobs_url"`
	Assignees_url string `json:"assignees_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Name string `json:"name"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Url string `json:"url"`
	Has_pages bool `json:"has_pages"`
	Labels_url string `json:"labels_url"`
	Full_name string `json:"full_name"`
	Network_count int `json:"network_count"`
	Open_issues_count int `json:"open_issues_count"`
	Issues_url string `json:"issues_url"`
	Has_projects bool `json:"has_projects"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Html_url string `json:"html_url"`
	Trees_url string `json:"trees_url"`
	Archive_url string `json:"archive_url"`
	Homepage string `json:"homepage"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Keys_url string `json:"keys_url"`
	Contributors_url string `json:"contributors_url"`
	Watchers int `json:"watchers"`
	Node_id string `json:"node_id"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Clone_url string `json:"clone_url"`
	Releases_url string `json:"releases_url"`
	Mirror_url string `json:"mirror_url"`
	Default_branch string `json:"default_branch"`
	Deployments_url string `json:"deployments_url"`
	Id int `json:"id"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Statuses_url string `json:"statuses_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Tags_url string `json:"tags_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	License GeneratedType `json:"license"` // License Simple
	Created_at string `json:"created_at"`
	Events_url string `json:"events_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Stargazers_count int `json:"stargazers_count"`
	Open_issues int `json:"open_issues"`
	Forks_url string `json:"forks_url"`
	Merges_url string `json:"merges_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Private bool `json:"private"`
	Compare_url string `json:"compare_url"`
	Watchers_count int `json:"watchers_count"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Milestones_url string `json:"milestones_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Has_downloads bool `json:"has_downloads"`
	Contents_url string `json:"contents_url"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Stargazers_url string `json:"stargazers_url"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Language string `json:"language"`
	Topics []string `json:"topics,omitempty"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Forks_count int `json:"forks_count"`
	Subscribers_count int `json:"subscribers_count"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Languages_url string `json:"languages_url"`
	Ssh_url string `json:"ssh_url"`
	Teams_url string `json:"teams_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Forks int `json:"forks"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Notifications_url string `json:"notifications_url"`
	Downloads_url string `json:"downloads_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Label map[string]interface{} `json:"label"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Status string `json:"status,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Conclusion string `json:"conclusion,omitempty"`
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	After string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Created_at string `json:"created_at,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Event string `json:"event"`
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Forks []interface{} `json:"forks,omitempty"`
	Node_id string `json:"node_id"`
	Comments int `json:"comments"`
	Git_pull_url string `json:"git_pull_url"`
	Html_url string `json:"html_url"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Forks_url string `json:"forks_url"`
	Truncated bool `json:"truncated,omitempty"`
	Created_at string `json:"created_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	History []interface{} `json:"history,omitempty"`
	Commits_url string `json:"commits_url"`
	Public bool `json:"public"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Git_push_url string `json:"git_push_url"`
	Files map[string]interface{} `json:"files"`
	Id string `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Created_at string `json:"created_at"`
	Issue_url string `json:"issue_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"` // Contents of the issue comment
	User GeneratedType `json:"user"` // A GitHub user.
	Url string `json:"url"` // URL for the issue comment
	Event string `json:"event"`
	Body_html string `json:"body_html,omitempty"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the issue comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Before string `json:"before"`
	Id int `json:"id"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Conclusion string `json:"conclusion"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Created_at string `json:"created_at"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	After string `json:"after"`
	Node_id string `json:"node_id"`
	Check_runs_url string `json:"check_runs_url"`
	Updated_at string `json:"updated_at"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Head_branch string `json:"head_branch"`
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	State string `json:"state"` // State of this codespace.
	Id int `json:"id"`
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Location string `json:"location"` // The Azure region where this codespace is located.
	Repository GeneratedType `json:"repository"` // Full Repository
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Url string `json:"url"` // API URL for this codespace.
	Updated_at string `json:"updated_at"`
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Created_at string `json:"created_at"`
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Recent_folders []string `json:"recent_folders"`
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pulls_url string `json:"pulls_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Downloads_url string `json:"downloads_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Tags_url string `json:"tags_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Comments_url string `json:"comments_url"`
	Archive_url string `json:"archive_url"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	License map[string]interface{} `json:"license,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Compare_url string `json:"compare_url"`
	Keys_url string `json:"keys_url"`
	Git_url string `json:"git_url,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Fork bool `json:"fork"`
	Teams_url string `json:"teams_url"`
	Forks_url string `json:"forks_url"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Languages_url string `json:"languages_url"`
	Created_at string `json:"created_at,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Language string `json:"language,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Trees_url string `json:"trees_url"`
	Stargazers_url string `json:"stargazers_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Name string `json:"name"`
	Network_count int `json:"network_count,omitempty"`
	Branches_url string `json:"branches_url"`
	Labels_url string `json:"labels_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Contents_url string `json:"contents_url"`
	Milestones_url string `json:"milestones_url"`
	Disabled bool `json:"disabled,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Hooks_url string `json:"hooks_url"`
	Archived bool `json:"archived,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Forks int `json:"forks,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Blobs_url string `json:"blobs_url"`
	Visibility string `json:"visibility,omitempty"`
	Url string `json:"url"`
	Notifications_url string `json:"notifications_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Issues_url string `json:"issues_url"`
	Releases_url string `json:"releases_url"`
	Description string `json:"description"`
	Merges_url string `json:"merges_url"`
	Private bool `json:"private"`
	Has_issues bool `json:"has_issues,omitempty"`
	Html_url string `json:"html_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Has_projects bool `json:"has_projects,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Full_name string `json:"full_name"`
	Commits_url string `json:"commits_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Events_url string `json:"events_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Collaborators_url string `json:"collaborators_url"`
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Status string `json:"status"` // The status of the most recent build of the Page.
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Source GeneratedType `json:"source,omitempty"`
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Url string `json:"url"` // The API address for accessing this Page resource.
	Cname string `json:"cname"` // The Pages site's custom domain
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/reference/deployments#list-deployment-statuses).
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Workflow map[string]interface{} `json:"workflow,omitempty"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Added_by string `json:"added_by,omitempty"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Verified bool `json:"verified"`
	Last_used string `json:"last_used,omitempty"`
	Title string `json:"title"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Rename map[string]interface{} `json:"rename"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Full_description string `json:"full_description,omitempty"` // description of the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Encoding string `json:"encoding"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Links map[string]interface{} `json:"_links"`
	Content string `json:"content"`
	Download_url string `json:"download_url"`
	License GeneratedType `json:"license"` // License Simple
	Name string `json:"name"`
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Detail string `json:"detail,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // Name of the team
	Repositories_url string `json:"repositories_url"`
	Url string `json:"url"` // URL for the team
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Slug string `json:"slug"`
	Description string `json:"description"` // Description of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_sign bool `json:"can_sign"`
	Expires_at string `json:"expires_at"`
	Key_id string `json:"key_id"`
	Raw_key string `json:"raw_key"`
	Can_certify bool `json:"can_certify"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Id int `json:"id"`
	Public_key string `json:"public_key"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Created_at string `json:"created_at"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Primary_key_id int `json:"primary_key_id"`
	Emails []map[string]interface{} `json:"emails"`
	Revoked bool `json:"revoked"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Effective_date string `json:"effective_date"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Version string `json:"version"`
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Lock_reason string `json:"lock_reason,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Event string `json:"event"`
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"`
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Comments_url string `json:"comments_url"`
	Body_text string `json:"body_text,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Body string `json:"body,omitempty"` // Contents of the issue
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	User GeneratedType `json:"user"` // A GitHub user.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the issue
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Title string `json:"title"` // Title of the issue
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Labels_url string `json:"labels_url"`
	Locked bool `json:"locked"`
	Repository_url string `json:"repository_url"`
	Comments int `json:"comments"`
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Draft bool `json:"draft,omitempty"`
	Html_url string `json:"html_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Updated_at string `json:"updated_at"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Events_url string `json:"events_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Comments int `json:"comments"`
	Updated_at string `json:"updated_at"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Html_url string `json:"html_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Closed_at string `json:"closed_at"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Url string `json:"url"` // URL for the issue
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Events_url string `json:"events_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Locked bool `json:"locked"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Created_at string `json:"created_at"`
	Title string `json:"title"` // Title of the issue
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Labels_url string `json:"labels_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Comments_url string `json:"comments_url"`
	Repository_url string `json:"repository_url"`
	Draft bool `json:"draft,omitempty"`
	Node_id string `json:"node_id"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Body_text string `json:"body_text,omitempty"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Label map[string]interface{} `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rule GeneratedType `json:"rule"`
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Tool GeneratedType `json:"tool"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	State string `json:"state"` // State of a code scanning alert.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
	Strict bool `json:"strict"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_requests []GeneratedType `json:"pull_requests"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Completed_at string `json:"completed_at"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	External_id string `json:"external_id"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Url string `json:"url"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Started_at string `json:"started_at"`
	Details_url string `json:"details_url"`
	Conclusion string `json:"conclusion"`
	Output map[string]interface{} `json:"output"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Id int `json:"id"` // The id of the check.
	Name string `json:"name"` // The name of the check.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Version_count int `json:"version_count"` // The number of versions of the package.
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the package.
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the package.
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Visibility string `json:"visibility"`
	Package_type string `json:"package_type"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the review
	Node_id string `json:"node_id"`
	Pull_request_url string `json:"pull_request_url"`
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the review.
	State string `json:"state"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_text string `json:"body_text,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Score float64 `json:"score"`
	Gists_url string `json:"gists_url"`
	Received_events_url string `json:"received_events_url"`
	Organizations_url string `json:"organizations_url"`
	Hireable bool `json:"hireable,omitempty"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Id int `json:"id"`
	Company string `json:"company,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Repos_url string `json:"repos_url"`
	Bio string `json:"bio,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Starred_url string `json:"starred_url"`
	Following_url string `json:"following_url"`
	Email string `json:"email,omitempty"`
	Location string `json:"location,omitempty"`
	Node_id string `json:"node_id"`
	Subscriptions_url string `json:"subscriptions_url"`
	Followers int `json:"followers,omitempty"`
	Following int `json:"following,omitempty"`
	Blog string `json:"blog,omitempty"`
	Events_url string `json:"events_url"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Public_gists int `json:"public_gists,omitempty"`
	Public_repos int `json:"public_repos,omitempty"`
	Suspended_at string `json:"suspended_at,omitempty"`
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Email string `json:"email,omitempty"`
	Url string `json:"url"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
	Name string `json:"name,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Login string `json:"login"`
	Following_url string `json:"following_url"`
	TypeField string `json:"type"`
	Subscriptions_url string `json:"subscriptions_url"`
	Role_name string `json:"role_name"`
	Events_url string `json:"events_url"`
	Gists_url string `json:"gists_url"`
	Node_id string `json:"node_id"`
	Received_events_url string `json:"received_events_url"`
	Starred_url string `json:"starred_url"`
	Gravatar_id string `json:"gravatar_id"`
	Html_url string `json:"html_url"`
	Organizations_url string `json:"organizations_url"`
	Id int `json:"id"`
	Repos_url string `json:"repos_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Url string `json:"url,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Name string `json:"name,omitempty"`
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The project card's ID
	Project_id string `json:"project_id,omitempty"`
	Column_name string `json:"column_name,omitempty"`
	Column_url string `json:"column_url"`
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Content_url string `json:"content_url,omitempty"`
	Created_at string `json:"created_at"`
	Note string `json:"note"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user-to-server access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	ErrorField string `json:"error"`
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Id int `json:"id"` // Unique identifier for this analysis.
	Tool GeneratedType `json:"tool"`
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Deletable bool `json:"deletable"`
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Url string `json:"url"` // The REST API URL of the analysis resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Review_id int `json:"review_id"`
	State string `json:"state"`
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Lock_reason string `json:"lock_reason"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/reference/repos#get-a-repository) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Reason string `json:"reason"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	All []int `json:"all"`
	Owner []int `json:"owner"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Svc_root string `json:"svc_root,omitempty"`
	Status string `json:"status"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Vcs string `json:"vcs"`
	Authors_count int `json:"authors_count,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Message string `json:"message,omitempty"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Push_percent int `json:"push_percent,omitempty"`
	Repository_url string `json:"repository_url"`
	Svn_root string `json:"svn_root,omitempty"`
	Error_message string `json:"error_message,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Status_text string `json:"status_text,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Authors_url string `json:"authors_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Pull_request map[string]interface{} `json:"pull_request"`
	Reason string `json:"reason,omitempty"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester map[string]interface{} `json:"requester,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Field1 int `json:"+1"`
	Laugh int `json:"laugh"`
	Total_count int `json:"total_count"`
	Confused int `json:"confused"`
	Eyes int `json:"eyes"`
	Heart int `json:"heart"`
	Rocket int `json:"rocket"`
	Field1 int `json:"-1"`
	Hooray int `json:"hooray"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Number int `json:"number"` // The security alert number.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	State string `json:"state"` // The state of the Dependabot alert.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Language string `json:"language"` // The language of the CodeQL database.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Name string `json:"name"` // The name of the CodeQL database.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Id int `json:"id"` // The ID of the CodeQL database.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Stargazers_url string `json:"stargazers_url"`
	Trees_url string `json:"trees_url"`
	Language string `json:"language,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Languages_url string `json:"languages_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Branches_url string `json:"branches_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Url string `json:"url"`
	Role_name string `json:"role_name,omitempty"`
	Merges_url string `json:"merges_url"`
	Statuses_url string `json:"statuses_url"`
	Name string `json:"name"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Deployments_url string `json:"deployments_url"`
	Fork bool `json:"fork"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Has_pages bool `json:"has_pages,omitempty"`
	Id int `json:"id"`
	Archived bool `json:"archived,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	License map[string]interface{} `json:"license,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Events_url string `json:"events_url"`
	Forks_url string `json:"forks_url"`
	Contents_url string `json:"contents_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Forks int `json:"forks,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Contributors_url string `json:"contributors_url"`
	Disabled bool `json:"disabled,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Description string `json:"description"`
	Pulls_url string `json:"pulls_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Issues_url string `json:"issues_url"`
	Full_name string `json:"full_name"`
	Milestones_url string `json:"milestones_url"`
	Node_id string `json:"node_id"`
	Comments_url string `json:"comments_url"`
	Labels_url string `json:"labels_url"`
	Git_refs_url string `json:"git_refs_url"`
	Private bool `json:"private"`
	Tags_url string `json:"tags_url"`
	Compare_url string `json:"compare_url"`
	Commits_url string `json:"commits_url"`
	Releases_url string `json:"releases_url"`
	Git_commits_url string `json:"git_commits_url"`
	Keys_url string `json:"keys_url"`
	Downloads_url string `json:"downloads_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Teams_url string `json:"teams_url"`
	Collaborators_url string `json:"collaborators_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Git_tags_url string `json:"git_tags_url"`
	Archive_url string `json:"archive_url"`
	Subscription_url string `json:"subscription_url"`
	Forks_count int `json:"forks_count,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Severity string `json:"severity"` // The severity of the advisory.
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Git []string `json:"git,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Api []string `json:"api,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Web []string `json:"web,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Target_url string `json:"target_url"` // Deprecated: the URL to associate with this status.
	Repository_url string `json:"repository_url"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	State string `json:"state"` // The state of the status.
	Deployment_url string `json:"deployment_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Description string `json:"description"` // A short description of the status.
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Effective_date string `json:"effective_date"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion interface{} `json:"discussion"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
	Comment string `json:"comment"` // The comment submitted with the deployment review
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Os string `json:"os"` // The Operating System of the runner.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Id int `json:"id"` // The id of the runner.
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repos only. `organization` level access allows sharing across the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Users []GeneratedType `json:"users"`
	Teams []Team `json:"teams"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Id int `json:"id"` // Unique identifier of the delivery.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Status string `json:"status"` // Description of the status of the attempted delivery
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Response map[string]interface{} `json:"response"`
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Request map[string]interface{} `json:"request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Download_url string `json:"download_url"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Path string `json:"path"`
	Git_url string `json:"git_url"`
	Submodule_git_url string `json:"submodule_git_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Url string `json:"url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Path string `json:"path"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	State string `json:"state"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Badge_url string `json:"badge_url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
	Total_count int `json:"total_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}
