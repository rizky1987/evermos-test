package helper

func NotFoundValidationForSearching(err error ) []string {

	var errResults []string
	if err != nil && !IsNotFoundErrorValidation(err.Error()) {

		errResults = append(errResults, err.Error())
	}

	return errResults
}
