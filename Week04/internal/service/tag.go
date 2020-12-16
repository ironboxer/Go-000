package service

import (
	"context"
	"github.com/ironboxer/week04/api"
	"github.com/ironboxer/week04/internal/dao"
	"github.com/pkg/errors"
)

type TagService struct {
	dao dao.Dao
}

func NewTagService(dao dao.Dao) TagService {
	return TagService{dao: dao}
}

func (t *TagService) GetTag(ctx context.Context, req api.TagRequest) (*api.TagResponse, error) {
	tag, err := t.dao.GetTag(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "Not Found")
	}
	resp := api.TagResponse{Id: tag.ID, Name: tag.Name}
	return &resp, nil
}
