package service

import "github.com/lttzzlll/week02/dao"

// TagService handles biz for Tag Model
type TagService struct {
}

// GetTagByID returns Tag by id
func (s *TagService) GetTag(id uint64) (*dao.Tag, error) {
	return dao.GetTag(id)
}
