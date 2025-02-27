package dto

type Risk struct {
	RiskId string
	State  string
	Title  string
	Desc   string
}

type GetRisksResponse struct {
	Risks []*Risk
}

type CreateRiskRequest struct {
	State string
	Title string
	Desc  string
}

type EmptyRequest struct {
}

type GetRiskByIdRequest struct {
	RiskId string
}
