package user

import "github.com/google/uuid"

//UUID user UUID
type UUID = uuid.UUID

//NewUUID create a new user UUID
func NewUUID() UUID {
	return UUID(uuid.New())
}

//StringToUUID convert a string to an user UUID
func StringToUUID(s string) (UUID, error) {
	parseUuid, err := uuid.Parse(s)
	return UUID(parseUuid), err
}
