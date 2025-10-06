package language

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/rootUJ99/lang-connect-server/pkg/jsonutil"
)
type Handler struct {
	service Service	

}

func NewLanguageHandler(service Service) Handler{
	return Handler{
		service : service,
	}
}

func (h Handler)CreateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	requestJson, err:= jsonutil.JsonDecode[LangCodeJSON](r)
	if (err != nil) {
		slog.Error("Error parsing the json body", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
		
	}
	
	_, err = h.service.CreateLanguageService(context.Background(), requestJson)

	if (err != nil) {
		slog.Error("Service Error", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}
	err = jsonutil.SendJson(w, 200, GenericJSONMesage{ Message: "Success"})

	if (err != nil) {
		slog.Error("Error parsing response", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

}

func (h Handler) UpdateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()

	requestJson, err:= jsonutil.JsonDecode[LangCodeJSONWithEmpty](r)
	if (err != nil) {
		slog.Error("Error parsing the json body", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
		
	}

	id := queryParam.Get("id")


	if (id == "") {
		slog.Error("Please provide querry param", "error", fmt.Errorf("id is empty"))
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
		
	}

	int_id, err :=  strconv.ParseInt(id, 10, 64) 
	if (err != nil) {
		slog.Error("Error parising querry param", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

	_, err = h.service.UpdateLanguageService(context.Background(), requestJson, int_id)

	if (err != nil) {
		slog.Error("Service Error", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}
	
	err = jsonutil.SendJson(w, 200, GenericJSONMesage{Message: fmt.Sprintf("%s", queryParam)})

	if (err != nil) {
		slog.Error("Error parsing response", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}
	
}

func (h Handler) DeleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	id := queryParam.Get("id")
	slog.Info(fmt.Sprintf("query param -> %s", id))


	if (id == "") {
		slog.Error("Please provide querry param", "error", fmt.Errorf("id is empty"))
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
		
	}

	int_id, err :=  strconv.ParseInt(id, 10, 64) 
	if (err != nil) {
		slog.Error("Error parising querry param", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

	_, err = h.service.DeleteLanguagesService(context.Background(), int_id)

	if (err != nil) {
		slog.Error("Service Error", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

	err = jsonutil.SendJson(w, 200, GenericJSONMesage{Message: fmt.Sprintf("%s", queryParam)})

	if (err != nil) {
		slog.Error("Error parsing response", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

}


func (h Handler) ListLanguageHandler(w http.ResponseWriter, r *http.Request) {

	result, err := h.service.ListLanguagesService(context.Background())

	if (err != nil){
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}
		
	var langJsonList  []LangCodeJSON

	for _, ele := range(result) {
		langJsonList =  append(langJsonList, LangCodeJSON{ Name: ele.Name, Code: ele.Code, NativeName: ele.NativeName})	
	}
	
	restJson := make(map[string][]LangCodeJSON)
	restJson["result"] = langJsonList
	err = jsonutil.SendJson(w, 200, restJson)
	if (err != nil) {
		slog.Error("Error parsing response", "error", err)
		jsonutil.SendJson(w, 400, GenericJSONMesage{ Message: "Something went wrong"})
		return
	}

}



