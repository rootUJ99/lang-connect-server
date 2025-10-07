package language

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/rootUJ99/lang-connect-server/internal/repository"
)

type Service struct {
	querry 	*repository.Queries

}

func NewLanguageService(querry *repository.Queries)Service{
	return Service{
		querry: querry,
	}
}

func (s Service) CreateLanguageService (ctx context.Context, requestJson LangCodeJSON) (bool, error){
	byteMetadatJson, err :=json.Marshal(requestJson.Metadata)
	if err != nil {
		return false, err
	}
	languageRequest := repository.CreateLanguageParams{
			Name: requestJson.Name,
			Code: requestJson.Code,
			NativeName: requestJson.NativeName,
			Description: requestJson.Description,
			Script: requestJson.Script,
			Metadata: byteMetadatJson,
		}	
	result, err:= s.querry.CreateLanguage(ctx, languageRequest)
	if err != nil {
		return false, err
	}
	
	slog.Info(fmt.Sprintf("data added %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) UpdateLanguageService (ctx context.Context, requestJson LangCodeJSONWithEmpty, queryId int64) (bool, error){
	result, err := s.querry.GetLanguageById(context.Background(), int32(queryId))
	if err != nil {
		return false, err
	}
	updatedLanguage := repository.UpdateLanguageByIdParams{
		LanguageID: result.LanguageID,
		Code :result.Code,
		Name :result.Name,
		NativeName: result.NativeName,
		Script: result.Script,
		Description: result.Description,
		Metadata: result.Metadata,
	}	
	slog.Info("the jsonData", "", fmt.Sprintf("%s %s %s \n", updatedLanguage.Code, updatedLanguage.Name, updatedLanguage.Script ))
	if requestJson.Name != nil {
		updatedLanguage.Name = *requestJson.Name
	}
	if requestJson.Code != nil {
		updatedLanguage.Code = *requestJson.Code
	}
	if requestJson.NativeName != nil {
		updatedLanguage.NativeName = *requestJson.NativeName
	}
	if requestJson.Description != nil {
		updatedLanguage.Description = *requestJson.Description
	}
	if requestJson.Script != nil {
		updatedLanguage.Script = *requestJson.Script
	}

	result, err = s.querry.UpdateLanguageById(context.Background(), updatedLanguage)

	slog.Info(fmt.Sprintf("got updated value %s", updatedLanguage))

	if err != nil {
		return false, err
	}
	
	// slog.Info(fmt.Sprintf("data updated %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) DeleteLanguagesService(ctx context.Context, queryId int64) (bool, error){
	result, err := s.querry.DeleteLanguageById(context.Background(), int32(queryId))

	if err != nil {
		return false, err
	}

	slog.Info(fmt.Sprintf("data deleted %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
	return true, nil
}


func (s Service) ListLanguagesService(ctx context.Context) ([]repository.Language, error) {
	result, err:=s.querry.ListLanguage(context.Background())

	if err != nil {
		return nil, err
	}

	return result, err

}
