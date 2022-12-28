# Go Struct Validator

### Introduction
- The purpose of this project is to create error messages with proper structure that are easy for both users and developers to understand and use. 
- Our goal is to provide error messages that include the name of the property in response and are simple to use in frontend applications, specifically in forms to display error messages for specific fields. 
- This will help improve the user experience by making it easier for users to understand and resolve errors that may occur while using the system.

### Installation

Use go get
```
go get -u github.com/shariarfaisal/validator
```

Import into your project
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
	Label string `json:"label" v:"enum=Home,Office"`
	Street string `json:"street" v:"required"`
	City string `json:"city" v:"required"`
}

type User struct {
	Name  string `json:"name" v:"required;min=3;max=20"`
	Age   int    `json:"age" v:"required;min=18;max=60"`
	Email string `json:"email" v:"required;email"`
	Addresses []Address `json:"addresses" v:"required;min=1;max=3"`
	DateOfBirth string `json:"dateOfBirth" v:"required;date"`
	Gender string `json:"gender" v:"enum=male,female"`
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

### Contributing
We welcome contributions to this project! If you have an idea for a new feature or have found a bug, please open an issue on GitHub. If you would like to contribute code, please follow these guidelines:

- Fork the repository and create a new branch for your changes.
- Make your changes, including appropriate tests.
- Submit a pull request for review.
