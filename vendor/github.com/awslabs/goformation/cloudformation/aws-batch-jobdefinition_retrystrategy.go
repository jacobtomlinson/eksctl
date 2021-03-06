package cloudformation

import (
	"encoding/json"
)

// AWSBatchJobDefinition_RetryStrategy AWS CloudFormation Resource (AWS::Batch::JobDefinition.RetryStrategy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html
type AWSBatchJobDefinition_RetryStrategy struct {

	// Attempts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html#cfn-batch-jobdefinition-retrystrategy-attempts
	Attempts *Value `json:"Attempts,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_RetryStrategy) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.RetryStrategy"
}

func (r *AWSBatchJobDefinition_RetryStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
