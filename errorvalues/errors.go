package errorvalues

import (
	svc "github.com/DecodeWorms/server-contract/errors"
)

var (
	PersonalIdDoesNotExistStatusCode      = 309
	EmailExistStatusCode                  = 411
	PhoneNumberExistStatusCode            = 413
	JerseyNumberExistError                = 415
	CreatingPlayerWithFieldInfoStatusCode = 500
	CreatingPlayerAddressStatusCode       = 502

	errorMessage = map[int]string{
		PersonalIdDoesNotExistStatusCode:      "error player personal id does not exist",
		EmailExistStatusCode:                  "error player email already exist",
		PhoneNumberExistStatusCode:            "error player phone number already exist",
		JerseyNumberExistError:                "error player jersey number already exist",
		CreatingPlayerWithFieldInfoStatusCode: "error creating player with field info",
		CreatingPlayerAddressStatusCode:       "error creating player address",
	}

	errorType = map[int]string{
		PersonalIdDoesNotExistStatusCode:      "error type of player personal id does not exist",
		EmailExistStatusCode:                  "error type of email already exist",
		PhoneNumberExistStatusCode:            "error type of phone number already exist",
		JerseyNumberExistError:                "error type jersey number already exist",
		CreatingPlayerWithFieldInfoStatusCode: "error type creating player with field info",
		CreatingPlayerAddressStatusCode:       "error type creating player address",
	}
)

func Type(code int) string {
	res, ok := errorType[code]
	if ok {
		return res
	}
	return "unknown error code"
}

func Message(code int) string {
	res, ok := errorMessage[code]
	if ok {
		return res
	}
	return "unknown error code"
}

func Format(code int) error {
	return svc.NewDerror(Message(code), Type(code), code)
}
