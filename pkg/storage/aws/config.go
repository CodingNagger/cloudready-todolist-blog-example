package aws

// TodoConfig is the main configuration item for the AWS configuration
type TodoConfig struct {
	TableName string
	Region    string
	Endpoint  string
}
