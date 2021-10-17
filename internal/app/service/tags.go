package service

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (s *Service) GetTagsMap() (res map[int]datastruct.Tags, err error) {
	tags, err := s.r.GetAllTags()
	if err != nil {
		return nil, errors.Wrap(err, "service.GetTagsMap")
	}
	res = make(map[int]datastruct.Tags)
	for _, tag := range tags {
		res[tag.ID] = tag
	}
	return res, nil
}

func (s *Service) GetTopTags() (res []datastruct.Tags, err error) {
	//todo get top tags instead of all
	res, err = s.r.GetTopTags()
	return res, errors.Wrap(err, "service.GetTopTags")
}
