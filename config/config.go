package config

const (
	OperatorName      string = "splunk-forwarder-operator"
	OperatorNamespace string = "splunk-forwarder-operator"

	SplunkAuthSecretName string = "splunk-auth" // #nosec G101 -- This is a false positive
)
