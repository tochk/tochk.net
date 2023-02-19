package service

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (s *Service) GetArticlesByLanguage(language string) (res []datastruct.Articles, err error) {
	res, err = s.r.GetArticlesByLanguage(language)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetArticlesByLanguage")
	}

	articleIDs := make([]int, 0, len(res))
	for _, article := range res {
		articleIDs = append(articleIDs, article.ID)
	}

	articleTagsMap, err := s.GetArticlesTagsMap(articleIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetArticlesByLanguage")
	}

	articleTeamMembersMap, err := s.GetArticlesTeamMembersMap(articleIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetArticlesByLanguage")
	}

	for idx, article := range res {
		article.TagsIDs = articleTagsMap[article.ID]
		article.TeamMembersIDs = articleTeamMembersMap[article.ID]
		res[idx] = article
	}
	return res, nil
}

func (s *Service) GetArticleByID(id int) (res datastruct.Articles, err error) {
	res, err = s.r.GetArticleByID(id)
	if err != nil {
		return res, errors.Wrap(err, "service.GetArticleByID")
	}
	articleTagsMap, err := s.GetArticlesTagsMap([]int{res.ID})
	if err != nil {
		return res, errors.Wrap(err, "service.GetArticleByID")
	}

	articleTeamMembersMap, err := s.GetArticlesTeamMembersMap([]int{res.ID})
	if err != nil {
		return res, errors.Wrap(err, "service.GetArticleByID")
	}

	res.TagsIDs = articleTagsMap[res.ID]
	res.TeamMembersIDs = articleTeamMembersMap[res.ID]

	return res, errors.Wrap(err, "service.GetArticleByID")
}

func (s *Service) GetArticlesTagsMap(articleIDs []int) (res map[int][]int, err error) {
	articlesTags, err := s.r.GetArticlesTags(articleIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetArticlesTagsMap")
	}
	res = make(map[int][]int)
	for _, articleTag := range articlesTags {
		res[articleTag.ArticleID] = append(res[articleTag.ArticleID], articleTag.TagID)
	}
	return res, nil
}

func (s *Service) GetArticlesTeamMembersMap(articleIDs []int) (res map[int][]int, err error) {
	articlesTeamMembers, err := s.r.GetArticlesTeamMembers(articleIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetArticlesTeamMembersMap")
	}
	res = make(map[int][]int)
	for _, articleTeamMember := range articlesTeamMembers {
		res[articleTeamMember.ArticleID] = append(res[articleTeamMember.ArticleID], articleTeamMember.TeamMemberID)
	}
	return res, nil
}
