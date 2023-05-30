package v1

// auth
type ReviewerSubmitRequest struct {
	UserId        int64 `json:"UserId" validate:"required,gte=1"`
	ApplicationId int64 `json:"ApplicationId" validate:"required,gte=1"`
	ReviewStatus  bool  `json:"ReviewStatus" validate:"required"`
}
type ReviewerWithDrawRequest struct {
	UserId int64 `json:"UserId" validate:"required,gte=1"`
}
