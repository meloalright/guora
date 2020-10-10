package h

// Response struct
type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}
