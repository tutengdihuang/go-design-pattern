package csdn_example

import "testing"

func TestChainofResponsibilityPattern(t *testing.T) {
	logger := testChainofResponsibilityPattern()

	logger.LogMessage(INFO, "This is an information.", nil)

	logger.LogMessage(DEBUG, "This is a debug level information.", nil)

	logger.LogMessage(ERROR, "This is an error information.", nil)
}
