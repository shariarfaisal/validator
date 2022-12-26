package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Location struct {
	City string `json:"city" v:"required;min:3;max:10"`
	Street string `json:"street" v:"required"`
	Label string `json:"label" v:"required"`
}

type User struct {
	Name string `json:"name" v:"required;min:3;max:10;email"`
	Phone string `json:"phone"`
	Age int `json:"age" v:"min:18;max:60"`
	Address struct {
		City string `json:"city" v:"required"`
		Street string `json:"street" v:"required"`
	} `json:"address" v:"required"`
	Emails []string `json:"emails" v:"required;min:1;max:3"`
	Locations []Location `json:"locations" v:"required;min:1;max:3"`
	Status string `json:"status" v:"required;enum:active,inactive"`
	CreatedAt string `json:"created_at" v:"required;date"`
}

func (u *User) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err == nil {
		return true
	}
	_, err = time.Parse("2006/01/02", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02 15:04:05", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05Z", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05.000Z", date)
	if err == nil {
		return true
	}

	return false
}
 
func validateField (v reflect.Value, l string) interface{} {
	fields := make([]string, 2)
	keyValue := strings.Split(l, ":")

	if len(keyValue) == 2 {
		fields[0] = keyValue[0]
		fields[1] = keyValue[1]
	} else {
		fields[0] = keyValue[0]
	}
	key := fields[0]
	value := fields[1]

	if key != "" {
		switch key {
			case "required":
				if v.Kind() == reflect.String {
					if v.String() == "" {
						return "required"
					}
				}else if v.Kind() == reflect.Int {
					if v.Int() == 0 {
						return "required"
					}
				}else if v.Kind() == reflect.Slice {
					fmt.Println("Slice")
					if v.Len() == 0 {
						return "required"
					}
				}

			case "min":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) < vv{
								return "min " + value + " characters"
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() < int64(vv){
								return "min " + value
							}
						}
					} else if v.Kind() == reflect.Slice {
						if vv > 0 {
							if v.Len() < vv{
								return "min " + value
							}
						}
					}
				}
			case "max":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) > vv{
								return "max " + value + " characters"
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() > int64(vv){
								return "max " + value
							}
						}
					} else if v.Kind() == reflect.Slice {
						if vv > 0 {
							if v.Len() > vv{
								return "max " + value
							}
						}
					}
				}
			case "enum":
				if value != "" {
					if v.Kind() == reflect.String {
						enums := strings.Split(value, ",")
						if len(enums) > 0 {
							for _, e := range enums {
								if v.String() == e {
									return ""
								}
							}
							return "valid only " + value
						}
					}
				}
			case "email":
				if v.Kind() == reflect.String {
					_, err := mail.ParseAddress(v.String())
					if err != nil {
						return "invalid email"
					}
				}
			case "date":
				if v.Kind() == reflect.String {
					fmt.Println("Date", v.String())
					if !isValidDate(v.String()) {
						return "invalid date"
					}
				}
		}
	}

	return ""
}

func validateStruct(ut reflect.Type, uv reflect.Value) interface{} {
	
	errors := map[string]interface{}{}

	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)

		name := field.Tag.Get("json")
		tag := field.Tag.Get("v")
		v := uv.Field(i)

		tags := strings.Split(tag, ";")

		for _, l := range tags {
			if v.Kind() == reflect.Struct {
				err := validateStruct(v.Type(), v)
				if err != "" {
					errors[name] = err
					break
				}
			}else if v.Kind() == reflect.Slice {
				err := validateField(v, l)
				fmt.Println(name, l)
				if err != "" {
					errors[name] = err
					break
				}else if v.Len() > 0 {
					errors[name] = map[int]interface{}{}
					for i := 0; i < v.Len(); i++ {
						if v.Index(i).Kind() == reflect.Struct {
							err := validateStruct(v.Index(i).Type(), v.Index(i))
							if err != "" {
								errors[name].(map[int]interface{})[i] = err
							}
						}
					}
				}
			}else {
				err := validateField(v, l)
				if err != "" {
					errors[name] = err
					break
				}
			}
		}
	}

	return errors
}

func Validate(s interface{}) (bool, map[string]interface{}) {
	ut := reflect.TypeOf(s)
	uv := reflect.ValueOf(s)

	err := validateStruct(ut, uv)
	
	if reflect.TypeOf(err).NumField() > 0 {
		return true, err.(map[string]interface{})
	} else {
		return false, nil
	}
}
