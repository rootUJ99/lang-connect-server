package words

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/rootUJ99/lang-connect-server/internal/repository"
)

type Service struct {
	querry 	*lang_repository.Queries

}

func NewWordsService(querry *lang_repository.Queries)Service{
	return Service{
		querry: querry,
	}
}

func (s Service) CreateCommonWordService (ctx context.Context, requestJson LangCodeJSON) (bool, error){
	langWord := lang_repository.CreateCommonWordParams{
			Word: requestJson.Word,
			LangCode: requestJson.LangCode,
		}	
	result, err:= s.querry.CreateCommonWord(ctx, langWord)
	if err != nil {
		return false, err
	}
	
	slog.Info(fmt.Sprintf("data added %s, %s, %s", result.LangCode, result.Word, result.ID))
	
	return true, nil
}


func (s Service) UpdateCommonWordService (ctx context.Context, requestJson LangCodeJSONWithEmpty, queryId int64) (bool, error){
	result, err := s.querry.GetCommonWordById(context.Background(), int64(queryId))
	if err != nil {
		return false, err
	}
	updatedWords := lang_repository.UpdateCommonWordByIdParams{ID: result.ID}	
	if (requestJson.Word != "") {
		updatedWords.Word = requestJson.Word
	} else { 
		updatedWords.Word = result.Word
	} 
	if (requestJson.LangCode != "") {
		updatedWords.LangCode = requestJson.LangCode
	} else {
		updatedWords.LangCode = result.LangCode
	} 

	result, err = s.querry.UpdateCommonWordById(context.Background(), updatedWords)

	slog.Info(fmt.Sprintf("got updated value %s", result))

	if err != nil {
		return false, err
	}
	
	slog.Info(fmt.Sprintf("data updated %s, %s, %s", result.LangCode, result.Word, result.ID))
	
	return true, nil
}


func (s Service) DeleteCommonWordsService(ctx context.Context, queryId int64) (bool, error){
	result, err := s.querry.DeleteCommonWordById(context.Background(), queryId)

	if err != nil {
		return false, err
	}

	slog.Info(fmt.Sprintf("data deleted %s, %s, %s", result.LangCode, result.Word, result.ID))
	
	return true, nil
}


func (s Service) ListCommonWordsService(ctx context.Context) ([]lang_repository.CommonWord, error) {
	result, err:=s.querry.ListCommonWords(context.Background())

	if err != nil {
		return nil, err
	}

	return result, err

}
