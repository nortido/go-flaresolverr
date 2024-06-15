package flaresolverr

import (
	"github.com/google/uuid"
)

type Response struct {
	Status         string            `json:"status"`
	Message        string            `json:"message"`
	StartTimestamp int64             `json:"startTimestamp"`
	EndTimestamp   int64             `json:"endTimestamp"`
	Version        string            `json:"version"`
	Session        string            `json:"session"`
	Sessions       []uuid.UUID       `json:"sessions"`
	Solution       *ResponseSolution `json:"solution"`
}

type ResponseSolution struct {
	URL     string `json:"url"`
	Status  int    `json:"status"`
	Headers struct {
		Status              string `json:"status"`
		Date                string `json:"date"`
		ContentType         string `json:"content-type"`
		Expires             string `json:"expires"`
		CacheControl        string `json:"cache-control"`
		Pragma              string `json:"pragma"`
		XFrameOptions       string `json:"x-frame-options"`
		XContentTypeOptions string `json:"x-content-type-options"`
		CfCacheStatus       string `json:"cf-cache-status"`
		ExpectCt            string `json:"expect-ct"`
		ReportTo            string `json:"report-to"`
		Nel                 string `json:"nel"`
		Server              string `json:"server"`
		CfRay               string `json:"cf-ray"`
		ContentEncoding     string `json:"content-encoding"`
		AltSvc              string `json:"alt-svc"`
	} `json:"headers"`
	Response string `json:"response"`
	Cookies  Cookies `json:"cookies"`
	UserAgent string `json:"userAgent"`
}
