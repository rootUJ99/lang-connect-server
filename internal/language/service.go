package language 

import (
	"context"
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
	langWord := repository.CreateLanguageParams{
			Name: requestJson.Name,
			Code: requestJson.Code,
			NativeName: requestJson.NativeName,
			Description: requestJson.Description,
			Script: requestJson.Script,
		}	
	result, err:= s.querry.CreateLanguage(ctx, langWord)
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
	updatedWords := repository.UpdateLanguageByIdParams{LanguageID: result.LanguageID}	
	if (requestJson.Name != "") {
		updatedWords.Name = requestJson.Name
	} else { 
		updatedWords.Name = result.Name
	} 
	if (requestJson.Code != "") {
		updatedWords.Code = requestJson.Code
	} else {
		updatedWords.Code = result.Code
	} 

	result, err = s.querry.UpdateLanguageById(context.Background(), updatedWords)

	slog.Info(fmt.Sprintf("got updated value %s", result))

	if err != nil {
		return false, err
	}
	
	slog.Info(fmt.Sprintf("data updated %s, %s, %s", result.Code, result.Name, result.LanguageID))
	
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
