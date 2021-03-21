package utils

import "fmt"

func TypeCheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int, int64:
			fmt.Printf("%+v is int.\n", v)
		case float64:
			fmt.Printf("%+v is float64.\n", v)
		case float32:
			fmt.Printf("%+v is float32.\n", v)
		case string:
			fmt.Printf("%+v is string.\n", v)
		case bool:
			fmt.Printf("%+v is bool.\n", v)
		case nil:
			fmt.Printf("%+v is nil.\n", v)
		default:
			fmt.Printf("%T is unknown.\n", v)
		}
	}
}
