package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
	"xiaolehuigo/accountserver/util"
	"xiaolehuigo/accountserver/util/log"
)

const DefaultMemory = 32 * 1024 * 1024

var timeFormat = "2006-01-02 15:04:05.000"

func LoggerMiddlerware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log := log.GetLogrus()
		url := context.Request.URL
		method := context.Request.Method
		path := url.Path
		host := context.Request.Host
		clientUserAgent := context.Request.UserAgent()
		clientIp := GetIP(context)

		var requestBody interface{}
		requestBody = GetRequestBody(context)
		requestHeaders := GetHeaders(context.Request.Header)
		start := util.GetNowTime()
		context.Next()
		stop := time.Since(start)
		requestTime := int(math.Ceil(float64(stop.Nanoseconds()) / 1000.0))
		statusCode := context.Writer.Status()

		dataLength := context.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			"hostname":    host,
			"statusCode":  statusCode,
			"requestTime": requestTime, // time to process
			"clientIP":    clientIp,
			"headers":     requestHeaders,
			"method":      method,
			"path":        path,
			"dataLength":  dataLength,
			"userAgent":   clientUserAgent,
			"body":        requestBody,
		}).WithTime(start)

		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("currentTime:[%s] success requestTime[%dms]", util.GetNowTime().Format(timeFormat), requestTime)
			fmt.Println(statusCode)
			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}

// GetHeaders ...
func GetHeaders(head http.Header) map[string]string {
	hdr := make(map[string]string, len(head))
	for k, v := range head {
		hdr[k] = v[0]
	}
	return hdr
}

// GetIP ...
func GetIP(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

// GetMultiPartFormValue ...
func GetMultiPartFormValue(c *http.Request) interface{} {
	var requestBody interface{}

	multipartForm := make(map[string]interface{})
	if err := c.ParseMultipartForm(DefaultMemory); err != nil {
		// handle error
	}
	if c.MultipartForm != nil {
		for key, values := range c.MultipartForm.Value {
			multipartForm[key] = strings.Join(values, "")
		}

		for key, file := range c.MultipartForm.File {
			for k, f := range file {
				formKey := fmt.Sprintf("%s%d", key, k)
				multipartForm[formKey] = map[string]interface{}{"filename": f.Filename, "size": f.Size}
			}
		}

		if len(multipartForm) > 0 {
			requestBody = multipartForm
		}
	}
	return requestBody
}

// GetFormBody ...
func GetFormBody(c *http.Request) interface{} {
	var requestBody interface{}

	form := make(map[string]string)
	if err := c.ParseForm(); err != nil {
		// handle error
	}
	for key, values := range c.PostForm {
		form[key] = strings.Join(values, "")
	}
	if len(form) > 0 {
		requestBody = form
	}

	return requestBody
}

// GetRequestBody ...
func GetRequestBody(c *gin.Context) interface{} {
	//multiPartFormValue := GetMultiPartFormValue(c.Request)
	//if multiPartFormValue != nil {
	//	return multiPartFormValue
	//}
	//
	//formBody := GetFormBody(c.Request)
	//if formBody != nil {
	//	return formBody
	//}

	method := c.Request.Method
	if method == "GET" {
		return nil
	}
	contentType := c.ContentType()
	body := c.Request.Body
	var model interface{}
	bodyContent, err := ioutil.ReadAll(body)
	if err != nil {
		return model
	}
	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyContent))

	switch contentType {
	case binding.MIMEJSON:
		json.Unmarshal(bodyContent, &model)
		return model
	default:
		model = string(bodyContent)
		return model
	}
}
