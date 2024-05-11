package entities

type AccountManagement struct {
	EventType            string               `json:"eventType"`
	CreateAccountRequest CreateAccountRequest `json:"createAccountRequest"`
	JoinAccountRequest   JoinAccountRequest   `json:"joinAccountRequest"`
}

type CreateAccountRequest struct {
	WorkflowDesignTimeId string `json:"workflowDesignTimeId,omitempty"`
	WorkflowRuntimeId    string `json:"workflowRuntimeId,omitempty"`
	UserEmailId          string `json:"userEmailId"`
	PricingPlanId        string `json:"pricingPlanId"`
	AppRoleId            string `json:"appRoleId"`
	ThrivestackTenantId  string `json:"thrivestackTenantId,omitempty"`
}

type JoinAccountRequest struct {
	WorkflowDesignTimeId string   `json:"workflowDesignTimeId,omitempty"`
	WorkflowRuntimeId    string   `json:"workflowRuntimeId,omitempty"`
	UserEmailId          string   `json:"userEmailId"`
	AccountIds           []string `json:"accountIds"`
}

type AccountManagementResponse struct {
	EventType             string                `json:"eventType"`
	CreateAccountResponse CreateAccountResponse `json:"createAccountResponse"`
	JoinAccountResponse   JoinAccountResponse   `json:"joinAccountResponse"`
}
type CreateAccountResponse struct {
	WorkflowDesignTimeId string `json:"workflowDesignTimeId,omitempty"`
	WorkflowRuntimeId    string `json:"workflowRuntimeId,omitempty"`
	AccountId            string `json:"accountId,omitempty"`
	AccountName          string `json:"accountName,omitempty"`
	UserEmailId          string `json:"userEmailId,omitempty"`
	ThrivestackTenantId  string `json:"thrivestackTenantId,omitempty"`
}
type JoinAccountResponse struct {
	WorkflowDesignTimeId string   `json:"workflowDesignTimeId,omitempty"`
	WorkflowRuntimeId    string   `json:"workflowRuntimeId,omitempty"`
	AccountIds           []string `json:"accountIds,omitempty"`
	UserEmailId          string   `json:"userEmailId,omitempty"`
	EventType            string   `json:"eventType"`
}
