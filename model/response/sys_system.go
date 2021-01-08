package response

import "github.com/cxpgo/ginf/model/config"

type SysConfigResponse struct {
	Config config.Config `json:"config"`
}
