package validate

import (
	"fmt"

	"github.com/mt-sre/addon-metadata-operator/pkg/validators"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

func Validate(mb *utils.MetaBundle) []error {
	errs := []error{}
	validators := GetAllValidators()

	printHeading()

	for _, validator := range validators {
		fmt.Printf("\r%s\t\t", validator.Description)
		success, err := validator.Runner(mb)
		if err != nil && !success {
			errs = append(errs, err)
			printErrorMessage(validator.Description)
		} else if !success {
			printFailureMessage(validator.Description)
		} else {
			printSuccessMessage(validator.Description)
		}
		fmt.Println()
	}
	return errs
}

func GetAllValidators() []utils.Validator {
	return []utils.Validator{
		{
			Description: "Ensure defaultChannel is present in list of channels",
			Runner:      validators.ValidateDefaultChannel,
		},
		{
			Description: "Ensure `label` to follow the format api.openshift.com/addon-<operator-name>",
			Runner:      validators.ValidateAddonLabel,
		},
	}
}
