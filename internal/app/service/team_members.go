package service

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (s *Service) GetTeamMembersMap() (res map[int]datastruct.TeamMembers, err error) {
	teamMembers, err := s.r.GetAllTeamMembers()
	if err != nil {
		return nil, errors.Wrap(err, "service.GetTeamMembersMap")
	}
	res = make(map[int]datastruct.TeamMembers)
	for _, teamMember := range teamMembers {
		res[teamMember.ID] = teamMember
	}
	return res, nil
}

func (s *Service) GetTeamMembers() (res []datastruct.TeamMembers, err error) {
	teamMembers, err := s.r.GetAllTeamMembers()
	return teamMembers, errors.Wrap(err, "service.GetTeamMembers")
}
