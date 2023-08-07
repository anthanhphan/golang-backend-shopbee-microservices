package reqmodel

type Filter struct {
	ReqStatus int `json:"req_status,omitempty" form:"request_status"`
}
