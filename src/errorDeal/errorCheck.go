package errorDeal

import "fmt"

type ErrorCheck struct {
}

func (errorCheck *ErrorCheck) CommonCheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}


