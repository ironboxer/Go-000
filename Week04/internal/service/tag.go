package service

import (
	"context"

	v1 "Week04/api/tag/v1"

	"Week04/internal/biz"
	"github.com/pkg/errors"
)

type TagService struct {
	t *biz.TagUsecase
	v1.UnimplementedTagServer
}

func NewTagService(t *biz.TagUsecase) v1.TagServer {
	return &TagService{t: t}
}

func (t *TagService) GetTag(ctx context.Context, req *v1.TagRequest) (*v1.TagResponse, error) {
	tag, err := t.t.GetTag(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "Not Found")
	}
	resp := v1.TagResponse{Id: tag.ID, Name: tag.Name}
	return &resp, nil
}
