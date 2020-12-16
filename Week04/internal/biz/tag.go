package biz

type Tag struct {
	ID   uint64
	Name string
}

type TagRepo interface {
	GetTag(uint64) (*Tag, error)
}

func NewTagUsecase(repo TagRepo) *TagUsecase {
	return &TagUsecase{repo: repo}
}

type TagUsecase struct {
	repo TagRepo
}

func (t *TagUsecase) GetTag(id uint64) (*Tag, error) {
	tag, err := t.repo.GetTag(id)
	if err != nil {
		return nil, err
	}
	return &Tag{ID: tag.ID, Name: tag.Name}, nil
}
