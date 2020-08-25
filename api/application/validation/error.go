package validation

import "google.golang.org/genproto/googleapis/rpc/errdetails"

func ConvertToBadRequestDetails(
	errors map[string]string) []*errdetails.BadRequest_FieldViolation {

	details := make([]*errdetails.BadRequest_FieldViolation, 0)
	for field, msg := range errors {
		details = append(details, &errdetails.BadRequest_FieldViolation{
			Field:       field,
			Description: msg,
		})
	}
	return details
}
