package validator

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) []string {
	var out []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			out = append(out, e.Field()+" failed on the '"+e.Tag()+"' tag")
		}
	} else {
		out = append(out, err.Error())
	}

	return out
}
