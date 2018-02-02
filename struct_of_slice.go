package main

import (
	"fmt"
)

type person struct {
	name     string
	lastName string
	age      int
}

type people struct {
	numberOfpeople []person
}

func main() {
	fmt.Println("Hello, playground")
	saminder := people{}

	saminder = people{numberOfpeople: []person{person{name: "Saminder", lastName: "Chahal", age: 25}}}

	fmt.Println("%v\n", saminder)

}

/* REST API deep structs 
Expected output:
{
	"data": {
		"percentagcalculations": [
			{
				"param1": "value1",
				"moreSubparams": {
					"sub_param_1": {"critical": "80", "warning": "50"},
					"sub_param_2": {"critical": "90", "warning": "60"},
					"sub_param_3": {"critical": "80", "warning": "50"},
					"sub_param_4": {"critical": "90", "warning": "60"}
				}
			},
			{
				"param1": "value2",
				"moreSubparams": {
					"sub_param_1": {"critical": "80", "warning": "50"},
					"sub_param_2": {"critical": "90", "warning": "60"},
					"sub_param_3": {"critical": "80", "warning": "50"},
					"sub_param_4": {"critical": "90", "warning": "60"}
				}
			}
		]
	}
}


----------------------------------------------------------------------------------------------------------------------
                                      SAMPLE CODE
----------------------------------------------------------------------------------------------------------------------


/ mainKeyForBackendService holds all the data structures for the
// percentagcalculations node in above mentioned JSON
type mainKeyForBackendService struct {
	perCentageCalculations []individualPercentageData `json:"percentagcalculations"`
}

// individualPercentageData holds the data structures for a single param1's
// threshold information (each JSON object in the list of above
// mentioned JSON)
type individualPercentageData struct {
	param1                        string                      `json:"param1"`
	moreSubparams moreSubparams `json:"moreSubparams,omitempty"`
}

// moreSubparams holds the threshold values of the droppedBytesPriorityHigh,
// droppedBytesPriorityMedium, sub_param_3 and sub_param_4 parameters
// in above mentioned JSON
type moreSubparams struct {
	sub_param_1   thresholds `json:"sub_param_1,omitempty"`
	sub_param_2 thresholds `json:"sub_param_2,omitempty"`
	sub_param_3           thresholds `json:"sub_param_3,omitempty"`
	sub_param_4          thresholds `json:"sub_param_4,omitempty"`
}

// thresholds holds the actual percentage value of the critical and
// warning values for each parameter of moreSubparams node
// in above mentioned JSON
type thresholds struct {
	Critical string `json:"critical,omitempty"`
	Warning  string `json:"warning,omitempty"`
}



----------------------- HardCode Values / Learn Struct declarations -----------------------------------------------------
param1thresholds := mainKeyForBackendService{}

	param1thresholds = mainKeyForBackendService{
		perCentageCalculations: []individualPercentageData{
			individualPercentageData{
				param1: "value1",
				moreSubparams: moreSubparams{
					sub_param_1:   thresholds{Critical: "80", Warning: "50"},
					sub_param_2: thresholds{Critical: "90", Warning: "60"},
					sub_param_3:           thresholds{Critical: "80", Warning: "50"},
					sub_param_4:          thresholds{Critical: "90", Warning: "60"},
				,
		}}

*/
