package log

type (
	IdentityInfo struct {
		Identity interface{} `json:"indentity,omitempty"`
	}
)

func (this *IdentityInfo) SetIdentity(i interface{}) {

	this.Identity = i
}
