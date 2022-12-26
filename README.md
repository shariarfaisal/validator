# Go Struct Validator

### Installation

Use go get
```
go get -u github.com/shariarfaisal/validator
```

Import into your packages
```
import "github.com/shariarfaisal/validator"
```

### Tags 

|Tag | Description | 
| - | - |
| required | Field is not empty or undefined |
| email | Email validity |
| date | Date format validity |
| min | Minimum value for number and minimum length for string value | 
| max | Maximum value for number and maximum length for string value |
| enum | Is input data is enum |
| include | Is data contains or not in between given values | 
| eq | Is data equal |
| ne | Is data not eqal |
| gt | Is data greater than |
| gte | Is data greater than or equal |
| lt | Is data less than |
| lte| Is data less than or equal |
| url | Is valid url |
| ip | Is valid ip address |
| ipv4 | is valid ip address v4 | 

### Functions 
|Func| Arguments | Description| Return 
| - | - | - | - |
| IsValidDate | (***date*** string) | Check date's format validity | bool |
| IsValidEmail | (***email*** string) | Check email validity | bool |
| IsValidURL | (***url*** string) | Check url validity | bool |
| IsValidIP | (***ip*** string) | Check ip address validity | bool |
| IsValidIpV4 | (***ip*** string) | Check ip address v4 validity | bool |


### Example 
```
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shariarfaisal/validator"
)

type Address struct {
	Label string `json:"label" v:"enum:Home,Office"`
	Street string `json:"street" v:"required"`
	City string `json:"city" v:"required"`
}

type User struct {
	Name  string `json:"name" v:"required;min:3;max:20"`
	Age   int    `json:"age" v:"required;min:18;max:60"`
	Email string `json:"email" v:"required;email"`
	Addresses []Address `json:"addresses" v:"required;min:1;max:3"`
	DateOfBirth string `json:"dateOfBirth" v:"required;date"`
	Gender string `json:"gender" v:"enum:male,female"`
}

func main() {
	user := User{
		Name:  "test",
		Age:   17,
		Email: "example@gmail.com",
		Addresses: []Address{
			{
				Label: "Home",
			},
		},
		Gender: "mole",
	}

	// Validate the user
	ok, errors := validator.Validate(user)
	if !ok {
		fmt.Println(json.NewEncoder(os.Stdout).Encode(errors))
	}

	// Email validation using function
	if isValid := validator.IsValidEmail(user.Email); isValid {
		fmt.Println("Email is valid")
	}

}

```

### Console Outpur 
```
{
   "addresses":{
      "0":{
         "city":"This field is required",
         "street":"This field is required"
      }
   },
   "age":"This field must be at least 18",
   "dateOfBirth":"This field is required",
   "gender":"This field must be one of the following values: male,female"
}

Email is valid
```
