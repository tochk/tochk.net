package service

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (s *Service) GetProjectsByLanguage(language string) (res []datastruct.Projects, err error) {
	res, err = s.r.GetProjectsByLanguage(language)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetProjectsByLanguage")
	}

	projectIDs := make([]int, 0, len(res))
	for _, project := range res {
		projectIDs = append(projectIDs, project.ID)
	}

	projectTagsMap, err := s.GetProjectsTagsMap(projectIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetProjectsByLanguage")
	}

	projectTeamMembersMap, err := s.GetProjectsTeamMembersMap(projectIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetProjectsByLanguage")
	}

	for idx, project := range res {
		project.ImageURL = s.cnf.CDNBaseURL + project.ImageURL
		project.TagsIDs = projectTagsMap[project.ID]
		project.TeamMembersIDs = projectTeamMembersMap[project.ID]
		res[idx] = project
	}
	return res, nil
}

func (s *Service) GetProjectByID(id int) (res datastruct.Projects, err error) {
	res, err = s.r.GetProjectByID(id)
	if err != nil {
		return res, errors.Wrap(err, "service.GetProjectByID")
	}

	projectTagsMap, err := s.GetProjectsTagsMap([]int{res.ID})
	if err != nil {
		return res, errors.Wrap(err, "service.GetProjectByID")
	}

	projectTeamMembersMap, err := s.GetProjectsTeamMembersMap([]int{res.ID})
	if err != nil {
		return res, errors.Wrap(err, "service.GetProjectByID")
	}

	res.ImageURL = s.cnf.CDNBaseURL + res.ImageURL
	res.TagsIDs = projectTagsMap[res.ID]
	res.TeamMembersIDs = projectTeamMembersMap[res.ID]

	return res, errors.Wrap(err, "service.GetProjectByID")
}

func (s *Service) GetProjectsTagsMap(projectIDs []int) (res map[int][]int, err error) {
	projectsTags, err := s.r.GetProjectsTags(projectIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetProjectsTagsMap")
	}
	res = make(map[int][]int)
	for _, projectTag := range projectsTags {
		res[projectTag.ProjectID] = append(res[projectTag.ProjectID], projectTag.TagID)
	}
	return res, nil
}

func (s *Service) GetProjectsTeamMembersMap(projectIDs []int) (res map[int][]int, err error) {
	projectsTeamMembers, err := s.r.GetProjectsTeamMembers(projectIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetProjectsTeamMembersMap")
	}
	res = make(map[int][]int)
	for _, projectTeamMember := range projectsTeamMembers {
		res[projectTeamMember.ProjectID] = append(res[projectTeamMember.ProjectID], projectTeamMember.TeamMemberID)
	}
	return res, nil
}
