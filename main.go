package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"sample-app-backend/entities"
)

const (
	acknowledgeURL = "https://api.dev.app.thrivestack.ai/thrivestackWebhook/acknowledgeTenant"
)

func main() {
	http.HandleFunc("/accountMgmt", handleAccountMgmt)
	log.Println("listening on port :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleAccountMgmt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	var accountManagementReq entities.AccountManagement
	err := json.NewDecoder(r.Body).Decode(&accountManagementReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !validateRequest(accountManagementReq) {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	switch accountManagementReq.EventType {
	case "CREATE":
		simulateAccountCreation(accountManagementReq.CreateAccountRequest)
	case "JOIN":
		simulateAccountJoin(accountManagementReq.JoinAccountRequest)
	default:
		http.Error(w, "Invalid event type", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Account created/joined"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateRequest(req entities.AccountManagement) bool {
	if req.EventType == "" {
		return false
	}

	if req.EventType == "CREATE" && req.CreateAccountRequest.UserEmailId == "" {
		return false
	}

	if req.EventType == "JOIN" && req.JoinAccountRequest.UserEmailId == "" {
		return false
	}

	if req.EventType == "JOIN_AND_CREATE" && (req.JoinAccountRequest.UserEmailId == "" || req.CreateAccountRequest.UserEmailId == "") {
		return false
	}

	return true
}

func simulateAccountCreation(request entities.CreateAccountRequest) error {
	log.Println("Creating tenant...")
	go acknowledgeCreate(request)
	return nil
}

func simulateAccountJoin(request entities.JoinAccountRequest) error {
	log.Println("Joining tenant...")
	go acknowledgeJoin(request)
	return nil
}

func acknowledgeCreate(message entities.CreateAccountRequest) {
	accountId := "6dae52a1-20e0-41ff-a92b-b28463ebf3f6"
	accountName := "" // TODO: Fetch name from the request

	inputReq := entities.AccountManagementResponse{
		EventType: "account_created",
		CreateAccountResponse: entities.CreateAccountResponse{
			WorkflowDesignTimeId: message.WorkflowDesignTimeId,
			WorkflowRuntimeId:    message.WorkflowRuntimeId,
			AccountId:            accountId,
			AccountName:          accountName,
			UserEmailId:          message.UserEmailId,
			ThrivestackTenantId:  message.ThrivestackTenantId,
		},
	}

	jsonData, err := json.Marshal(inputReq)
	if err != nil {
		log.Println("Error marshalling input data for tenant ack API")
		return
	}

	body := bytes.NewReader(jsonData)
	httpReq, err := http.NewRequest(http.MethodPost, acknowledgeURL, body)
	if err != nil {
		log.Println("Error creating HTTP request for tenant ack API")
		return
	}

	token := fetchToken()
	headers := http.Header{}
	headers.Add("Authorization", token)
	httpReq.Header = headers
	_, err = http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Println("Error sending HTTP request to tenant ack API")
	} else {
		log.Println("Acknowledge tenant was successful")
	}

	log.Println("Create ack complete")
}

func acknowledgeJoin(message entities.JoinAccountRequest) {
	inputReq := entities.AccountManagementResponse{
		EventType: "account_added_user",
		JoinAccountResponse: entities.JoinAccountResponse{
			WorkflowDesignTimeId: message.WorkflowDesignTimeId,
			WorkflowRuntimeId:    message.WorkflowRuntimeId,
			AccountIds:           message.AccountIds,
			UserEmailId:          message.UserEmailId,
		},
	}

	jsonData, err := json.Marshal(inputReq)
	if err != nil {
		log.Println("Error marshalling input data for tenant ack API")
		return
	}

	body := bytes.NewReader(jsonData)
	httpReq, err := http.NewRequest(http.MethodPost, acknowledgeURL, body)
	if err != nil {
		log.Println("Error creating HTTP request for tenant ack API")
		return
	}

	token := fetchToken()
	headers := http.Header{}
	headers.Add("Authorization", token)
	httpReq.Header = headers
	_, err = http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Println("Error sending HTTP request to tenant ack API")
	} else {
		log.Println("Acknowledge join account was successful")
	}

	log.Println("Join ack complete")
}

func fetchToken() string {
	token := os.Getenv("THRIVESTACK_API_KEY")
	if token == "" {
		return "Bearer dummy-token"
	}
	return "Bearer " + token
}
