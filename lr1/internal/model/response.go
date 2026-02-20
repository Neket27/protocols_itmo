package model

type WorkResult struct {
	Value      int    `json:"value"`
	StrMessage string `json:"strMessage"`
}

type ServerResponse struct {
	WorkResult WorkResult `json:"workResult"`
}

func (r *ServerResponse) IsValid() bool {
	return r.WorkResult.Value != 0 || r.WorkResult.StrMessage != ""
}
