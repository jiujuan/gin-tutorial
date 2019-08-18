package middle

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "math"
    "os"
    "time"
)

//https://github.com/toorop/gin-logrus
var timeFormat = "02/Jan/2006:15:04:05 -0700"

func Logger(logger logrus.FieldLogger) gin.HandlerFunc  {
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "unknow"
    }

    return func(c *gin.Context) {
        path := c.Request.URL.Path
        start := time.Now()
        c.Next()
        stop := time.Since(start)
        latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
        statusCode := c.Writer.Status()
        clientIP := c.ClientIP()
        clientUserAgent := c.Request.UserAgent()
        referer := c.Request.Referer()
        dataLength := c.Writer.Size()
        if dataLength < 0 {
            dataLength = 0
        }

        entry := logger.WithFields(logrus.Fields{
            "hostname":   hostname,
            "statusCode": statusCode,
            "latency":    latency, // time to process
            "clientIP":   clientIP,
            "method":     c.Request.Method,
            "path":       path,
            "referer":    referer,
            "dataLength": dataLength,
            "userAgent":  clientUserAgent,
        })

        if len(c.Errors) > 0 {
            entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
        } else {
            msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
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
