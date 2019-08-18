package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gopkg.in/go-playground/validator.v8"
    "net/http"
    "reflect"
    "time"
)
//https://gin-gonic.com/docs/examples/custom-validators/

//Validator docs -
//https://godoc.org/gopkg.in/go-playground/validator.v8#Validate.RegisterStructValidation
//
//Struct level example -
//https://github.com/go-playground/validator/blob/v8.18.2/examples/struct-level/struct_level.go
//
//Validator release notes -
//https://github.com/go-playground/validator

//gin 官网的例子

type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
    CheckOut time.Time `form:"check_out" binding:"required,bookabledate" time_format:"2006-01-02"`
}

func bookableDate(
    v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
    field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
    if date, ok := field.Interface().(time.Time); ok {
        today := time.Now()

        if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
            return false
        }
    }
    return true
}

func main() {
    route := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        _ = v.RegisterValidation("bookabledate", bookableDate)
    }

    route.GET("/bookable", getBookable)
    _ = route.Run(":5000")
}

func getBookable(c *gin.Context) {
    var b Booking
    if err := c.ShouldBindWith(&b, binding.Query); err == nil {
        c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

//$ curl "localhost:8085/bookable?check_in=2018-04-16&check_out=2018-04-17"
//{"message":"Booking dates are valid!"}
//
//$ curl "localhost:8085/bookable?check_in=2018-03-08&check_out=2018-03-09"
//{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}


