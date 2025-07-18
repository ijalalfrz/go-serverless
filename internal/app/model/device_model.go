package model

type Device struct {
	PK          string `dynamodbav:"PK"`
	ID          string `dynamodbav:"id"`
	DeviceModel string `dynamodbav:"deviceModel"`
	Name        string `dynamodbav:"name"`
	Note        string `dynamodbav:"note"`
	Serial      string `dynamodbav:"serial"`
}
