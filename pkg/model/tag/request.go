package tag

import "ceres/pkg/router"

type ListRequest struct {
	Keyword  string   `form:"keyword"`
	Category Category `form:"category"`
	IsIndex  bool     `form:"isIndex"`
	Limit    int      `form:"limit"`
	Offset   int      `form:"offset"`
}

type TagListRequest struct {
	Ad   bool   `form:"ad"`
	Type string `form:"type"`
}

func (l ListRequest) Validate() error {
	if l.Limit <= 0 || l.Limit >= 100 {
		return router.ErrBadRequest.WithMsg("please input right limit")
	}
	if l.Offset < 0 {
		return router.ErrBadRequest.WithMsg("please input right offset")
	}
	return nil
}

func (r TagListRequest) Validate() error {
	if r.Type != "startup" && r.Type != "comer" {
		return router.ErrBadRequest.WithMsg("please input right type")
	}
	return nil
}
