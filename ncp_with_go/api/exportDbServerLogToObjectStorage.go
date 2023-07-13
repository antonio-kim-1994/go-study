package api

import (
	"net/url"
	"os"
	"strings"
	"time"
)

/*
Clopud DB for MySQL 로그 보관용 API
*/
func ExportDbServerLogToObjectStorage() ReturnValue {
	// Get File Name
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayStr := yesterday.Format("20060102")
	fileName := "live-mysql-002-y35.log." + yesterdayStr

	params := map[string]string{
		"regionCode":                 "FKR",
		"logType":                    "GENERAL",
		"fileName":                   fileName,
		"bucketName":                 "mysql",
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
	uri := "/vmysql/v2/exportDbServerLogToObjectStorage" + query

	param := ReturnValue{method, baseUrl, uri}

	return param
}
