package inputs


type ValidateUserInput struct {
	UserID uint64
	ValidationCode string `json:"validationCode"`
}
