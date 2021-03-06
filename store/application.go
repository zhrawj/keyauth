package store

import (
	"github.com/defineiot/keyauth/dao/application"
	"github.com/defineiot/keyauth/dao/user"
	"github.com/defineiot/keyauth/internal/exception"
)

// CreateApplication todo
func (s *Store) CreateApplication(app *application.Application) error {
	_, err := s.dao.User.GetUser(user.UserID, app.UserID)
	if err != nil {
		return err
	}

	if err := s.dao.Application.CreateApplication(app); err != nil {
		return err
	}

	return nil
}

// DeleteApplication todo
func (s *Store) DeleteApplication(id string) error {
	var err error

	cacheKey := "app_" + id

	app, err := s.dao.Application.GetApplication(id)
	if err != nil {
		return err
	}

	err = s.dao.Application.DeleteApplication(app.ID)
	if err != nil {
		return err
	}
	if s.isCache {
		if !s.cache.Delete(cacheKey) {
			s.log.Debug("delete app from cache failed, key: %s", cacheKey)
		}
		s.log.Debug("delete app from cache success, key: %s", cacheKey)
	}

	return nil
}

// ListUserApps todo
func (s *Store) ListUserApps(userID string) ([]*application.Application, error) {
	_, err := s.dao.User.GetUser(user.UserID, userID)
	if err != nil {
		return nil, err
	}

	apps, err := s.dao.Application.ListUserApplications(userID)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetUserApp todo
func (s *Store) GetUserApp(appid string) (*application.Application, error) {
	var err error

	app := new(application.Application)
	cacheKey := "app_" + appid

	if s.isCache {
		if s.cache.Get(cacheKey, app) {
			s.log.Debug("get app from cache key: %s", cacheKey)
			return app, nil
		}
		s.log.Debug("get app from cache failed, key: %s", cacheKey)
	}

	app, err = s.dao.Application.GetApplication(appid)
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, exception.NewBadRequest("app %s not found", appid)
	}

	if s.isCache {
		if !s.cache.Set(cacheKey, app, s.ttl) {
			s.log.Debug("set app cache failed, key: %s", cacheKey)
		}
		s.log.Debug("set app cache ok, key: %s", cacheKey)
	}

	return app, nil
}
