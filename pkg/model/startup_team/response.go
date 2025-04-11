package startup_team

type ListStartupTeamMemberResponse struct {
	List  []*StartupTeamMember `json:"list"`
	Total int64                `json:"total"`
}
