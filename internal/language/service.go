package language

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/rootUJ99/lang-connect-server/internal/platform/postgres"
)

type Service struct {
	querry 	*platform_postgres.Queries
	langRepo ILanguageRepository
}

func NewLanguageService(langRepo ILanguageRepository )Service{
	return Service{
		langRepo: langRepo,
	}
}

func (s Service) CreateLanguageService (ctx context.Context, requestJson LangCodeJSON) (bool, error){
	byteMetadatJson, err :=json.Marshal(requestJson.Metadata)
	if err != nil {
		return false, err
	}

	result, err := s.langRepo.CreateLanguageById(requestJson, byteMetadatJson)

	if err != nil {
		return false, err
	}
	
	slog.Info(fmt.Sprintf("data added %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) UpdateLanguageService (ctx context.Context, requestJson LangCodeJSONWithEmpty, queryId int64) (bool, error){
	result, err := s.langRepo.GetLanguageById(int32(queryId))
	if err != nil {
		return false, err
	}

	result, err = s.langRepo.UpdateLanguageById(result, requestJson)

	slog.Info(fmt.Sprintf("got updated value %s", result))

	if err != nil {
		return false, err
	}
	
	// slog.Info(fmt.Sprintf("data updated %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) DeleteLanguagesService(ctx context.Context, queryId int64) (bool, error){
	result, err := s.langRepo.DeleteLanguageById(int32(queryId))

	if err != nil {
		return false, err
	}

	slog.Info(fmt.Sprintf("data deleted %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) ListLanguagesService(ctx context.Context) ([]platform_postgres.Language, error) {
	result, err := s.langRepo.ListLanguage()

	if err != nil {
		return nil, err
	}

	return result, err
}
