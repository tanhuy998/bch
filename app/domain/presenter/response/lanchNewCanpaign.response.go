package responsePresenter

import "github.com/google/uuid"

type (
	ILaunchNewCampaignResponsePresenter interface {
		SetMessage(string)
		SetData(*Data)
		GetData() *Data
	}

	Data struct {
		CreatedUUID *uuid.UUID `json:"createdUUID"`
	}

	LaunchNewCampaignResponse struct {
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}
)

func (this *LaunchNewCampaignResponse) SetMessage(msg string) {

	this.Message = msg
}

func (this *LaunchNewCampaignResponse) SetData(dt *Data) {

	this.Data = *dt
}
func (this *LaunchNewCampaignResponse) GetData() *Data {

	return &this.Data
}
