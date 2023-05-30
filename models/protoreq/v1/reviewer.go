package v1

// auth
type ReviewerSubmitRequest struct {
	UserId        int64 `json:"UserId,omitempty"`
	ApplicationId int64 `json:"ApplicationId,omitempty"`
	ReviewStatus  bool  `json:"ReviewStatus,omitempty"`
}
type ReviewerWithDrawRequest struct {
	UserId int64 `json:"UserId,omitempty"`
}
