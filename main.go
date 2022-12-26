package validator

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Date string's validation
func IsValidDate(date string) bool {
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

// Email validation 
func IsValidEmail(v string) bool {
	_, err := mail.ParseAddress(v)
	return err == nil
}

// URL validation
func IsValidURL(v string) bool {
	_, err := url.ParseRequestURI(v)
	return err == nil
}

// IP validation
func IsValidIP(v string) bool {
	return net.ParseIP(v) != nil
}

// IPv4 validation
func IsValidIpV4(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && ip.To4() != nil
}


// Validation messages
var messages = map[string] string{
	"required": "This field is required",
	"email": "This field must be a valid email address",
	"date": "This field must be a valid date",
	"min": "This field must be at least %s",
	"max": "This field must be at most %s",
	"enum": "This field must be one of the following values: %s",
	"include": "This field must include one of the following values: %s",
	"eq": "This field must be equal to %s",
	"ne": "This field must not be equal to %s",
	"gt": "This field must be greater than %s",
	"gte": "This field must be greater than or equal to %s",
	"lt": "This field must be less than %s",
	"lte": "This field must be less than or equal to %s",
	"url": "This field must be a valid URL",
	"ip": "This field must be a valid IP address",
	"ipv4": "This field must be a valid IPv4 address",
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
						return messages["required"]
					}
				}else if v.Kind() == reflect.Int {
					if v.Int() == 0 {
						return messages["required"]
					}
				}else if v.Kind() == reflect.Float32 {
					if v.Float() == 0 {
						return messages["required"]
					}
				} else if v.Kind() == reflect.Float64 {
					if v.Float() == 0 {
						return messages["required"]
					}
				}else if v.Kind() == reflect.Slice {
					if v.Len() == 0 {
						return messages["required"]
					}
				}
			case "include": 
				if v.Kind() == reflect.String {
					if !strings.Contains(value, v.String()) {
						return fmt.Sprintf(messages["include"], value)
					}
				}
			case "min":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) < vv{
								return fmt.Sprintf(messages["min"], value)
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() < int64(vv){
								return fmt.Sprintf(messages["min"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return fmt.Sprintf(messages["min"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return fmt.Sprintf(messages["min"], value)
							}
						}
					}else if v.Kind() == reflect.Slice {
						if vv > 0 {
							if v.Len() < vv{
								return fmt.Sprintf(messages["min"], value)
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
								return fmt.Sprintf(messages["max"], value)
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() > int64(vv){
								return fmt.Sprintf(messages["max"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return fmt.Sprintf(messages["max"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return fmt.Sprintf(messages["max"], value)
							}
						}
					} else if v.Kind() == reflect.Slice {
						if vv > 0 {
							if v.Len() > vv{
								return fmt.Sprintf(messages["max"], value)
							}
						}
					}
				}
			case "eq": 
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() != int64(vv){
								return fmt.Sprintf(messages["eq"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() != float64(vv){
								return fmt.Sprintf(messages["eq"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() != float64(vv){
								return fmt.Sprintf(messages["eq"], value)
							}
						}
					}else if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) != vv{
								return fmt.Sprintf(messages["eq"], value)
							} else {
								if v.String() != value{
									return fmt.Sprintf(messages["eq"], value)
								}
							}
						}
					}
				}
			case "ne":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() == int64(vv){
								return fmt.Sprintf(messages["ne"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() == float64(vv){
								return fmt.Sprintf(messages["ne"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() == float64(vv){
								return fmt.Sprintf(messages["ne"], value)
							}
						}
					} else if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) == vv{
								return fmt.Sprintf(messages["ne"], value)
							} else {
								if v.String() == value{
									return fmt.Sprintf(messages["ne"], value)
								}
							}
						}
					}
				}
			case "gt":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() <= int64(vv){
								return fmt.Sprintf(messages["gt"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() <= float64(vv){
								return fmt.Sprintf(messages["gt"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() <= float64(vv){
								return fmt.Sprintf(messages["gt"], value)
							}
						}
					}
				}
			case "gte":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() < int64(vv){
								return fmt.Sprintf(messages["gte"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return fmt.Sprintf(messages["gte"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return fmt.Sprintf(messages["gte"], value)
							}
						}
					}
				}
			case "lt":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() >= int64(vv){
								return fmt.Sprintf(messages["lt"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() >= float64(vv){
								return fmt.Sprintf(messages["lt"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() >= float64(vv){
								return fmt.Sprintf(messages["lt"], value)
							}
						}
					}
				}
			case "lte":
				if value != "" {
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() > int64(vv){
								return fmt.Sprintf(messages["lte"], value)
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return fmt.Sprintf(messages["lte"], value)
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return fmt.Sprintf(messages["lte"], value)
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
							return fmt.Sprintf(messages["enum"], value)
						}
					}
				}
			case "email":
				if v.Kind() == reflect.String {
					if !IsValidEmail(v.String()) {
						return messages["email"]
					}
				}
			case "url":
				if v.Kind() == reflect.String {
					if !IsValidURL(v.String()) {
						return messages["url"]
					}
				}
			case "ip":
				if v.Kind() == reflect.String {
					if !IsValidIP(v.String()) {
						return messages["ip"]
					}
				}
			case "ipv4":
				if v.Kind() == reflect.String {
					if !IsValidIpV4(v.String()) {
						return messages["ipv4"]
					}
				}
			case "date":
				if v.Kind() == reflect.String {
					if !IsValidDate(v.String()) {
						return messages["date"]
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
		if name == "" {
			name = field.Name
		}
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

// Validate struct fields
// @param s interface{} struct
// @return bool, map[string]interface{}
func Validate(s interface{}) (bool, map[string]interface{}) {
	ut := reflect.TypeOf(s)
	uv := reflect.ValueOf(s)

	err := validateStruct(ut, uv)
	kind := reflect.TypeOf(err).Kind()
	if kind == reflect.Map {
		if len(err.(map[string]interface{})) > 0 {
			return false, err.(map[string]interface{})
		}
	}

	return true, nil
}
