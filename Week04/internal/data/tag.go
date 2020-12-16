package data

import "Week04/internal/biz"

var _ biz.TagRepo = new(tagRepo)

func NewTagRepo() biz.TagRepo {
	return &tagRepo{}
}

type tagRepo struct {
}

func (t *tagRepo) GetTag(id uint64) (*biz.Tag, error) {
	return &biz.Tag{ID: 1, Name: "Hello"}, nil
}