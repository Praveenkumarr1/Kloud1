package main

import (
    "fmt"
    "reflect"
    "regexp"
    "gopkg.in/go-playground/validator.v8"
)
type Person struct {
    ID       string `json:"id" validate:"required,min=36,max=36"`
    Password string `json:"id" validate:"upwd"`
}
var (
    validate       = validator.New(&validator.Config{TagName: "validate"})
    passwordString = "^[a-zA-Z0-9]$" // regex that compiles
    passwordRegex  = regexp.MustCompile(passwordString)
)
func main() {

    validate.RegisterValidation("validpasswd", func(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
        return passwordRegex.MatchString(field.String())
    })
    validate.RegisterAliasValidation("upwd", "min=8,max=20,validpasswd")
var errMsg string
	 person := Person{
        ID:       "36",
        Password: "!!!!!!!!!!!!!!",
    }

    if errs := validate.Struct(person); errs != nil {
        if err := errs.(validator.ValidationErrors)["Person.Password"]; err != nil {
            if err.Tag == "upwd" {
                errMsg = "invalid password!"
 }} }
fmt.Println(errMsg)
}