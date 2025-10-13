package language

import (
	"context"
	"fmt"
	"log/slog"

	platform_postgres "github.com/rootUJ99/lang-connect-server/internal/platform/postgres"
)

type ILanguageRepository interface {
	ListLanguage() ([]platform_postgres.Language, error) 
	DeleteLanguageById(int32) (platform_postgres.Language, error) 
	CreateLanguageById(LangCodeJSON, []byte) (platform_postgres.Language, error)
	GetLanguageById(int32) (platform_postgres.Language, error) 
	UpdateLanguageById(platform_postgres.Language, LangCodeJSONWithEmpty) (platform_postgres.Language, error)
}


type langRepo struct {
	query *platform_postgres.Queries
}


func NewLangRepo(q *platform_postgres.Queries) ILanguageRepository{
	return langRepo {
		query: q,
	}	
}

func (l langRepo) ListLanguage() ([]platform_postgres.Language, error) {
	return l.query.ListLanguage(context.Background())
} 

func (l langRepo) DeleteLanguageById(id int32) (platform_postgres.Language, error) {
	return l.query.DeleteLanguageById(context.Background(), id)
} 

func (l langRepo) CreateLanguageById(requestJson LangCodeJSON, byteMetadata []byte) (platform_postgres.Language, error) {
	languageRequest := platform_postgres.CreateLanguageParams{
			Name: requestJson.Name,
			Code: requestJson.Code,
			NativeName: requestJson.NativeName,
			Description: requestJson.Description,
			Script: requestJson.Script,
			Metadata: byteMetadata,
		}	
	return l.query.CreateLanguage(context.Background(), languageRequest)

} 

func (l langRepo) GetLanguageById(id int32) (platform_postgres.Language, error) {
	return l.query.GetLanguageById(context.Background(), id)
}

func (l langRepo) UpdateLanguageById(result platform_postgres.Language, requestJson LangCodeJSONWithEmpty) (platform_postgres.Language, error) {

	updatedLanguage := platform_postgres.UpdateLanguageByIdParams{
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

	return  l.query.UpdateLanguageById(context.Background(), updatedLanguage)

}

