package godraft

import (
	"net/http"
	"strconv"

	"github.com/gothing/draft"
)

func InvalidCaseBody(key, errType string, value interface{}) map[string]map[string]interface{} {
	r := make(map[string]map[string]interface{})
	r[key] = map[string]interface{}{
		"value": value,
		"error": errType,
	}
	return r
}

var statusMap = map[int]draft.StatusType{
	http.StatusProcessing:           draft.Status.Processing,                 // 102 processing
	http.StatusOK:                   draft.Status.OK,                         // 200 ok
	http.StatusAccepted:             draft.Status.Accepted,                   // 202 accepted
	http.StatusNonAuthoritativeInfo: draft.Status.NonAuthoritative,           // 203 non_authoritative
	http.StatusPartialContent:       draft.Status.Partial,                    // 206 partial
	http.StatusMovedPermanently:     draft.Status.Move,                       // 301 move
	http.StatusFound:                draft.Status.Found,                      // 302
	http.StatusNotModified:          draft.Status.NotModified,                // 304 notmodified
	http.StatusBadRequest:           draft.Status.Invalid,                    // 400 invalid
	http.StatusPaymentRequired:      draft.Status.PaymentRequired,            // 402 payment_required
	http.StatusForbidden:            draft.Status.Denied,                     // 403 denied
	http.StatusNotFound:             draft.Status.NotFound,                   // 404 notfound
	http.StatusNotAcceptable:        draft.Status.Unacceptable,               // 406 unacceptable
	http.StatusRequestTimeout:       draft.Status.Timeout,                    // 408 timeout
	http.StatusConflict:             draft.Status.Conflict,                   // 409 conflict
	http.StatusExpectationFailed:    draft.Status.ExpectationFailed,          // 417 expectation_failed
	http.StatusUnprocessableEntity:  draft.Status.Unprocessable,              // 422 unprocessable
	http.StatusLocked:               draft.Status.Locked,                     // 423 locked
	http.StatusFailedDependency:     draft.Status.FailedDependency,           // 424 failed_dependency
	http.StatusUpgradeRequired:      draft.Status.UpgradeRequired,            // 426 upgrade_required
	http.StatusTooManyRequests:      draft.Status.ManyRequests,               // 429 many_requests
	449:                             draft.Status.RetryWith,                  // 449 retry_with
	451:                             draft.Status.UnavailableForLegalReasons, // 451 unavailable_for_legal_reasons
	http.StatusInternalServerError:  draft.Status.Fail,                       // 500 fail
	http.StatusNotImplemented:       draft.Status.NotImplemented,             // 501 not_implemented
	http.StatusServiceUnavailable:   draft.Status.Unavaliable,                // 503 unavailable
	http.StatusInsufficientStorage:  draft.Status.Insufficient,               // 507 insufficient
}

func HTTPStatus(status int) draft.StatusType {
	st, ok := statusMap[status]
	if !ok {
		return draft.StatusType(strconv.Itoa(status))
	}
	return st
}
