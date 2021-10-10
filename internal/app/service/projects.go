package service

import "tochkru-golang/internal/app/datastruct"

func (s *Service) GetProjectsByLanguage(language string) (res []datastruct.Projects, err error) {
	res, err = s.r.GetProjectsByLanguage(language)
	if err != nil {
		return nil, err
	}

	//todo wrap cdn url

	return res, err
}
