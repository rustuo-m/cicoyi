package r

type UserListReq struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	PageSize   int    `json:"pageSize"`
	PageNumber int    `json:"pageNumber"`
}
