package api

import (
	"net/url"
	"os"
	"strings"
)

/*
Clopud DB for MySQL 로그 리스트 확인 API
*/
func GetDbServerLogList() ReturnValue {
	params := map[string]string{
		"regionCode":                 "FKR",
		"logType":                    "GENERAL",
		"cloudMysqlServerInstanceNo": os.Getenv("MySQL_INSTANCE_NUM"),
		"responseFormatType":         "json",
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("?")

	for key, value := range params {
		queryBuilder.WriteString(url.QueryEscape(key))
		queryBuilder.WriteString("=")
		queryBuilder.WriteString(url.QueryEscape(value))
		queryBuilder.WriteString("&")
	}

	query := strings.TrimSuffix(queryBuilder.String(), "&")

	method := "GET"
	baseUrl := "https://fin-ncloud.apigw.fin-ntruss.com"
	uri := "/vmysql/v2/getDbServerLogList" + query

	param := ReturnValue{method, baseUrl, uri}

	return param
}
