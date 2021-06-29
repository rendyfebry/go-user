package rest

type (
	// HealhcheckResponse ...
	HealhcheckResponse struct {
		Data string `json:"data"`
		Meta Meta   `json:"meta"`
	}
)
