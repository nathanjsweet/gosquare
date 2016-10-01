package gosquare

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

// Provides a business's account information, such as its name and associated email address.
//
// Required permissions:  MERCHANT_PROFILE_READ
func RetrieveBusiness(token string) (*Merchant, error) {
	v := new(Merchant)
	_, err := squareRequest("GET", "/v1/me", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides details for a business's locations, including their IDs.
//
// The account_capabilities array returned in each Merchant object indicates which
// account capabilities the location has enabled. For example, if this array does not include the
// value CREDIT_CARD_PROCESSING, the location cannot currently process credit cards with
// Square.
//
// Required permissions:  MERCHANT_PROFILE_READ
func ListLocations(token string) ([]*Merchant, *NextRequest, error) {
	v := make([]*Merchant)
	nr, err := squareRequest("GET", "/v1/me/locations", token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateEmployeeReqObject struct {
	// The employee's first name.
	FirstName string `json:"first_name"`
	// The employee's last name.
	LastName string `json:"last_name"`
	// An optional second ID to associate the employee with an entity in another system.
	ExternalId string `json:"external_id"`
	// The ids of the employee's associated roles. Currently, you can specify only one
	// or zero roles per employee.Default value: []
	RoleIds []string `json:"role_ids"`
	// An optional email address to associate with the employee.Note that you cannot edit an existing employee's email address with the Connect API.
	// You can only set its initial value when creating an employee.
	Email string `json:"email"`
}

// Creates an employee for a business.
//
// Required permissions:  EMPLOYEES_WRITE
func CreateEmployee(token string, reqObj *CreateEmployeeReqObject) (*Employee, error) {
	v := new(Employee)
	_, err := squareRequest("POST", "/v1/me/employees", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employees.
//
// You can filter the results returned by this endpoint by exactly one of the following
// fields:
//
// `order`:
// The order in which employees are listed in the response, based on their
// created_at field.Default value: ASC
//
// `beginUpdatedAt`:
// If filtering results by their updated_at field, the beginning of the requested
// reporting period, in ISO
// 8601 format.
//
// `endUpdatedAt`:
// If filtering results by there updated_at field, the end of the requested
// reporting period, in ISO
// 8601 format.
//
// `beginCreatedAt`:
// If filtering results by their created_at field, the beginning of the requested
// reporting period, in ISO
// 8601 format.
//
// `endCreatedAt`:
// If filtering results by their created_at field, the end of the requested
// reporting period, in ISO
// 8601 format.
//
// `status`:
// If provided, the endpoint returns only employee entities with the specified
// status (ACTIVE or INACTIVE).
//
// `externalId`:
// If provided, the endpoint returns only employee entities with the specified
// external_id.
//
// `limit`:
// The maximum number of employee entities to return in a single response. This value
// cannot exceed 200.This value is always an integer.Default value: 100
func ListEmployees(token string, order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalId string, limit int) ([]*Employee, *NextRequest, error) {
	v := make([]*Employee)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/me/employees?order=%s&begin_updated_at=%s&end_updated_at=%s&begin_created_at=%s&end_created_at=%s&status=%s&external_id=%s&limit=%d", order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalId, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for a single employee.
//
// Required permissions:  EMPLOYEES_READ
func RetrieveEmployee(token, employeeId string) (*Employee, error) {
	v := new(Employee)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/me/employees/%s", employeeId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateEmployeeReqObject struct {
	// The employee's first name.
	FirstName string `json:"first_name"`
	// The employee's last name.
	LastName string `json:"last_name"`
	// An optional second ID to associate the employee with an entity in another system.
	ExternalId string `json:"external_id"`
	// The employee's associated roles. Currently, you can specify only one or zero roles per
	// employee.
	RoleIds []string `json:"role_ids"`
}

// Modifies the details of an employee.
//
// Required permissions:  EMPLOYEES_WRITE
func UpdateEmployee(token, employeeId string, reqObj *UpdateEmployeeReqObject) (*Employee, error) {
	v := new(Employee)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/me/employees/%s", employeeId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateRoleReqObject struct {
	// The role's name.
	Name string `json:"name"`
	// The role's permissions.
	Permissions []string `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the
	// values indicated in permissions.Default value: false
	IsOwner bool `json:"is_owner"`
}

// Creates an employee role you can then assign to employees.
//
// Required permissions:  EMPLOYEES_WRITE
func CreateRole(token string, reqObj *CreateRoleReqObject) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	_, err := squareRequest("POST", "/v1/me/roles", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employee roles.
//
// Required permissions:  EMPLOYEES_READ
//
// `order`:
// The order in which roles are listed in the response.Default value: ASC
//
// `limit`:
// The maximum number of employee entities to return in a single response. This value
// cannot exceed 200.This value is always an integer.Default value: 100
func ListRoles(token, order string, limit int) ([]*EmployeeRole, *NextRequest, error) {
	v := make([]*EmployeeRole)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/me/roles?order=%s&limit=%d", order, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for a single employee role.
//
// Required permissions:  EMPLOYEES_READ
func RetrieveRole(token, roleId string) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/me/roles/%s", roleId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateRoleReqObject struct {
	// The role's name.
	Name string `json:"name"`
	// The role's permissions.
	Permissions []string `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the
	// values indicated in permissions.
	IsOwner bool `json:"is_owner"`
}

// Modifies the details of an employee role.
//
// Required permissions:  EMPLOYEES_WRITE
func UpdateRole(token, roleId string, reqObj *UpdateRoleReqObject) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/me/roles/%s", roleId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateTimecardReqObject struct {
	// The employee to create a timecard for.
	EmployeeId string `json:"employee_id"`
	// The clock-in time for the timecard, in ISO 8601 format.Default value: The current time.
	ClockinTime string `json:"clockin_time"`
	// The clock-out time for the timecard, in ISO 8601 format.
	// Provide this value only if importing timecard information from another system.
	ClockoutTime string `json:"clockout_time"`
	// The ID of the location the employee clocked in from, if any.
	ClockinLocationId string `json:"clockin_location_id"`
	// The ID of the location the employee clocked out from. Provide this value only if
	// importing timecard information from another system.If you provide this value, you must also provide a value for clockout_time.
	ClockoutLocationId string `json:"clockout_location_id"`
}

// Creates a timecard for an employee. Each timecard corresponds to a single shift.
//
// This endpoint automatically creates an API_CREATE event for the new timecard.
//
// Required permissions:  TIMECARDS_WRITE
func CreateTimecard(token string, reqObj *CreateTimecardReqObject) (*Timecard, error) {
	v := new(Timecard)
	_, err := squareRequest("POST", "/v1/me/timecards", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employee timecards.
//
// You can filter the results returned by this endpoint by exactly one of the following
// fields:
//
// `order`:
// The order in which timecards are listed in the response, based on their
// created_at field.Default value: ASC
//
// `employeeId`:
// If provided, the endpoint returns only timecards for the employee with the specified
// ID.
//
// `beginClockinTime`:
// If filtering results by their clockin_time field, the beginning of the requested
// reporting period, in ISO
// 8601 format.
//
// `endClockinTime`:
// If filtering results by their clockin_time field, the end of the requested
// reporting period, in ISO
// 8601 format.
//
// `beginClockoutTime`:
// If filtering results by their clockout_time field, the beginning of the
// requested reporting period, in ISO 8601 format.
//
// `endClockoutTime`:
// If filtering results by their clockout_time field, the end of the requested
// reporting period, in ISO
// 8601 format.
//
// `beginUpdatedAt`:
// If filtering results by their updated_at field, the beginning of the requested
// reporting period, in ISO
// 8601 format.
//
// `endUpdatedAt`:
// If filtering results by their updated_at field, the end of the requested
// reporting period, in ISO
// 8601 format.
//
// `deleted`:
// If true, only deleted timecards are returned. If false, only valid
// timecards are returned.If you don't provide this parameter, both valid and deleted timecards are returned.
//
// `limit`:
// The maximum number of timecards to return in a single response. This value cannot
// exceed 200.This value is always an integer.
func ListTimecards(token, order, employeeId, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt string, deleted bool, limit int) ([]*Timecard, *NextRequest, error) {
	v := make([]*Timecard)
	nr, err := squareRequest("GET",
		fmt.Sprintf("/v1/me/timecards?order=%s&employee_id=%s&begin_clockin_time=%s&end_clockin_time=%s&begin_clockout_time=%s&end_clockout_time=%s&begin_updated_at=%s&end_updated_at=%s&deleted=%t&limit=%d",
			order, employeeId, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt, deleted, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Currently, only approved merchants can manage their employees with Square. Unapproved merchants cannot use employee management features you include in your application.
// Provides the details for a single timecard.
// Required permissions: TIMECARDS_READ
func RetrieveTimecard(token, timecardId string) (*Timecard, error) {
	v := new(Timecard)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/me/timecards/%s", timecardId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateTimecardReqObject struct {
	// The clock-in time for the timecard, in ISO 8601 format.
	ClockinTime string `json:"clockin_time"`
	// The clock-out time for the timecard, in ISO 8601 format.
	ClockoutTime string `json:"clockout_time"`
	// The ID of the location the employee clocked in from, if any.
	ClockinLocationId string `json:"clockin_location_id"`
	// The ID of the location the employee clocked out from, if any.
	ClockoutLocationId string `json:"clockout_location_id"`
}

// Modifies a timecard's details. This creates an API_EDIT event for the timecard. You
// can view a timecard's event history with the List Timecard
// Events endpoint.
//
// Required permissions:  TIMECARDS_WRITE
func UpdateTimecard(token, timecardId string, reqObj *UpdateTimecardReqObject) (*Timecard, error) {
	v := new(Timecard)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/me/timecards/%s", timecardId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func DeleteTimecard(token, timecardId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/me/timecards/%s", timecardId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all events associated with a particular timecard.
//
// Required permissions:  TIMECARDS_READ
func ListTimecardEvents(token, timecardId string) ([]*TimecardEvent, *NextRequest, error) {
	v := make([]*TimecardEvent)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/me/timecards/%s/events", timecardId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for all of a location's cash drawer shifts during a date range. The date
// range you specify cannot exceed 90 days.
//
// CashDrawerShift objects returned by this endpoint do not include the
// events field, which lists the events that occurred during the shift. To get a
// particular shift's events, use the Retrieve Cash Drawer
// Shift endpoint.
//
// Required permissions:  PAYMENTS_READ
//
// `beginTime`:
// The beginning of the requested reporting period, in ISO 8601 format.Default value: The current time minus 90 days.
//
// `endTime`:
// The beginning of the requested reporting period, in ISO 8601 format.Default value: The current time.
//
// `order`:
// The order in which cash drawer shifts are listed in the response, based on their
// created_at field.Default value: ASC
func ListCashDrawerShifts(token, locationId, beginTime, endTime, order string) ([]*CashDrawerShift, *NextRequest, error) {
	v := make([]*CashDrawerShift)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts?begin_time=%s&end_time=%s&order=%s", locationId, beginTime, endTime, order), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for a single cash drawer shift, including all events that occurred
// during the shift.
//
// Required permissions:  PAYMENTS_READ
func RetrieveCashDrawerShift(token, locationId, shiftId string) (*CashDrawerShift, error) {
	v := new(CashDrawerShift)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts/%s", locationId, shiftId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all payments taken by a merchant or any of the merchant's
// mobile staff during a date range.
// Date ranges cannot exceed one year in length. See Date ranges
// for details of inclusive and exclusive dates.
//
// Required permissions:  PAYMENTS_READ
//
// `beginTime`:
// The beginning of the requested reporting period, in ISO 8601 format.If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint
// returns an error.Default value: The current time minus one year.
//
// `endTime`:
// The end of the requested reporting period, in ISO 8601 format.If this value is more than one year greater than begin_time, this endpoint
// returns an error.Default value: The current time.
//
// `order`:
// The order in which payments are listed in the response.Default value: ASC
//
// `limit`:
// The maximum number of payments to return in a single response. This value cannot exceed
// 200.This value is always an integer.Default value: 100
func ListPayments(token, locationId, beginTime, endTime, order string, limit int) ([]*Payment, *NextRequest, error) {
	v := make([]*Payment)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/payments?begin_time=%s&end_time=%s&order=%s&limit=%d", locationId, beginTime, endTime, order, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides comprehensive information for a single payment.
//
// Required permissions:  PAYMENTS_READ
func RetrievePayment(token, locationId, paymentId string) (*Payment, error) {
	v := new(Payment)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/payments/%s", locationId, paymentId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all deposits and withdrawals initiated by Square to a
// merchant's bank account during a date range. Date ranges cannot exceed one year in length. See
// Date ranges for details of inclusive and exclusive
// dates.
//
// Settlement objects returned by this endpoint do not include the entries
// field, which lists the transactions that contribute to the total of the settlement. To get a
// particular settlement's entries, use the Retrieve
// Settlement endpoint.
//
// Square initiates its regular deposits to merchant bank accounts on the schedule indicated on
// this page. A deposit initiated by
// Square on a given day is usually not provided by this endpoint before 10 p.m. PST the same
// day.
//
// Square does not know when an initiated settlement completes, only whether it has
// failed. A completed settlement is typically reflected in a merchant's bank account
// within three business days, but in exceptional cases it might take longer.
//
// Required permissions:  SETTLEMENTS_READ
//
// `beginTime`:
// The beginning of the requested reporting period, in ISO 8601 format.If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint
// returns an error.Default value: The current time minus one year.
//
// `endTime`:
// The end of the requested reporting period, in ISO 8601 format.If this value is more than one year greater than begin_time, this endpoint
// returns an error.Default value: The current time.
//
// `order`:
// The order in which settlements are listed in the response.Default value: ASC
//
// `limit`:
// The maximum number of settlements to return in a single response. This value cannot
// exceed 200.This value is always an integer.Default value: 100
//
// `status`:
// Provide this parameter to retrieve only settlements with a particular status
// (SENT or FAILED).
func ListSettlements(token, locationId, beginTime, endTime, order string, limit int, status string) ([]*Settlement, *NextRequest, error) {
	v := make([]*Settlement)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/settlements?begin_time=%s&end_time=%s&order=%s&limit=%d&status=%s", locationId, beginTime, endTime, order, limit, status), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides comprehensive information for a single settlement, including the entries that
// contribute to the settlement's total.
//
// See SettlementEntry.Type for
// descriptions of the types of entries that compose a settlement.
//
// Required permissions:  SETTLEMENTS_READ
func RetrieveSettlement(token, locationId, settlementId string) (*Settlement, error) {
	v := new(Settlement)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/settlements/%s", locationId, settlementId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateRefundReqObject struct {
	// The ID of the payment to refund.If you're creating a PARTIAL refund for a split tender payment, instead provide
	// the id of the particular tender you want to refund. See Split Tender Payments for details.
	PaymentId string `json:"payment_id"`
	// The type of refund (FULL or PARTIAL).
	Type string `json:"type"`
	// The reason for the refund.
	Reason string `json:"reason"`
	// The amount of money to refund. Required only for PARTIAL refunds.The value of amount must be negative.
	RefundedMoney Money `json:"refunded_money"`
	// An optional key to ensure idempotence if you issue the same PARTIAL refund
	// request more than once.If you attempt to issue a partial refund and you aren't sure whether your request
	// succeeded, you can safely repeat your request with the same
	// request_idempotence_key. If you want to issue another partial refund for
	// the same payment, you must use a request_idempotence_key that is unique among
	// refunds you have issued for the payment.
	RequestIdempotenceKey string `json:"request_idempotence_key"`
}

// Issues a refund for a previously processed payment. You must issue a refund within 60 days of
// the associated payment. See this
// article for more information on refund behavior.
//
// Issuing a refund for a card payment is not reversible. To develop against this
// endpoint, you can create fake cash payments in Square Register and refund them.
//
// You can issue either full refunds or partial refunds. If you issue a partial refund, you must
// specify the amount of money to refund.
//
// Required permissions:  PAYMENTS_WRITE
func CreateRefund(token, locationId string, reqObj *CreateRefundReqObject) (*Refund, error) {
	v := new(Refund)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/refunds", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for all refunds initiated by a merchant or any of the merchant's mobile staff during a date range. Date
// ranges cannot exceed one year in length. See Date ranges for
// details of inclusive and exclusive dates.
//
// Required permissions:  PAYMENTS_READ
//
// `beginTime`:
// The beginning of the requested reporting period, in ISO 8601 format.If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint
// returns an error.Default value: The current time minus one year.
//
// `endTime`:
// The end of the requested reporting period, in ISO 8601 format.If this value is more than one year greater than begin_time, this endpoint
// returns an error.Default value: The current time.
//
// `order`:
// The order in which refunds are listed in the response.Default value: ASC
//
// `limit`:
// The maximum number of refunds to return in a single response. This value cannot exceed
// 200.This value is always an integer.Default value: 100
func ListRefunds(token, locationId, beginTime, endTime, order string, limit int) ([]*Refund, *NextRequest, error) {
	v := make([]*Refund)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/refunds?begin_time=%s&end_time=%s&order=%s&limit=%d", locationId, beginTime, endTime, order, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

//
// `limit`:
// The maximum number of orders to return in a single response. This value cannot exceed
// 200.This value is always an integer.Default value: 100
//
// `order`:
// Indicates whether orders are listed in chronological (ASC) or
// reverse-chronological (DESC) order.Default value: ASC
func ListOrders(token, locationId string, limit int, order string) ([]*Order, *NextRequest, error) {
	v := make([]*Order)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/orders?limit=%d&order=%s", locationId, limit, order), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

func RetrieveOrder(token, locationId, orderId string) (*Order, error) {
	v := new(Order)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/orders/%s", locationId, orderId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateOrderReqObject struct {
	// The action to perform on the order (COMPLETE, CANCEL, or
	// REFUND).
	Action string `json:"action"`
	// The tracking number of the shipment associated with the order. Only valid if
	// action is COMPLETE.
	ShippedTrackingNumber string `json:"shipped_tracking_number"`
	// A merchant-specified note about the completion of the order. Only valid if
	// action is COMPLETE.
	CompletedNote string `json:"completed_note"`
	// A merchant-specified note about the refunding of the order. Only valid if action
	// is REFUND.
	RefundedNote string `json:"refunded_note"`
	// A merchant-specified note about the canceling of the order. Only valid if action
	// is CANCEL.
	CanceledNote string `json:"canceled_note"`
}

func UpdateOrder(token, locationId, orderId string, reqObj *UpdateOrderReqObject) (*Order, error) {
	v := new(Order)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/orders/%s", locationId, orderId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides non-confidential details for all of a location's associated bank accounts. This
// endpoint does not provide full bank account numbers, and there is no way to obtain a
// full bank account number with the Connect API.
//
// Required permissions:  BANK_ACCOUNTS_READ
func ListBankAccounts(token, locationId string) ([]*BankAccount, *NextRequest, error) {
	v := make([]*BankAccount)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides non-confidential details for a merchant's associated bank account. This endpoint
// does not provide full bank account numbers, and there is no way to obtain a full bank
// account number with the Connect API.
//
// Required permissions:  BANK_ACCOUNTS_READ
func RetrieveBankAccount(token, locationId, bankAccountId string) (*BankAccount, error) {
	v := new(BankAccount)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts/%s", locationId, bankAccountId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateItemReqObject struct {
	// The item's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The item's name.
	Name string `json:"name"`
	// The item's description.
	Description string `json:"description"`
	// The ID of the item's category, if any.
	CategoryId string `json:"category_id"`
	// The color of the item's display label in Square Register.Default value: 9da2a6
	Color string `json:"color"`
	// The text of the item's display label in Square Register. Only up to the first five
	// characters of the string are used.Default value: The first two characters of the item's name.
	Abbreviation string `json:"abbreviation"`
	// Indicates whether the item is viewable from the merchant's online store (PUBLIC)
	// or PRIVATE.Default value: PUBLIC
	Visibility string `json:"visibility"`
	// If true, the item can be added to shipping orders from the merchant's online
	// store.Default value: false
	AvailableOnline bool `json:"available_online"`
	// If true, the item can be added to pickup orders from the merchant's online
	// store.Default value: false
	AvailableForPickup bool `json:"available_for_pickup"`
	// The item's variations. You must specify at least one variation.
	Variations []ItemVariation `json:"variations"`
}

// Creates an item and at least one variation for it.
//
// Required permissions:  ITEMS_WRITE
func CreateItem(token, locationId string, reqObj *CreateItemReqObject) (*Item, error) {
	v := new(Item)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/items", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a location's items.
//
// Required permissions:  ITEMS_READ
func ListItems(token, locationId string) ([]*Item, *NextRequest, error) {
	v := make([]*Item)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/items", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for a single item, including associated modifier lists and fees.
//
// Required permissions:  ITEMS_READ
func RetrieveItem(token, locationId, itemId string) (*Item, error) {
	v := new(Item)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/items/%s", locationId, itemId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateItemReqObject struct {
	// The item's name.
	Name string `json:"name"`
	// The item's description.
	Description string `json:"description"`
	// The ID of the item's category, if any.If you provide the empty string for this value, any existing category association is
	// removed from the item.
	CategoryId string `json:"category_id"`
	// The color of the item's display label in Square Register.
	Color string `json:"color"`
	// The text of the item's display label in Square Register. Only up to the first five
	// characters of the string are used.
	Abbreviation string `json:"abbreviation"`
	// Indicates whether the item is viewable from the merchant's online store (PUBLIC) or
	// PRIVATE.
	Visibility string `json:"visibility"`
	// If true, the item can be purchased from the merchant's online store.
	AvailableOnline bool `json:"available_online"`
	// If true, the item can be added to pickup orders from the merchant's online
	// store.
	AvailableForPickup bool `json:"available_for_pickup"`
}

// Modifies the core details of an existing item.
//
// If you want to modify an item's variations, use the Update Variation endpoint instead.
//
// If you want to add or remove a modifier list from an item, use the Apply Modifier List and Remove Modifier List endpoints instead.
//
// If you want to add or remove a fee from an item, use the Apply Fee and Remove Fee endpoints
// instead.
//
// Required permissions:  ITEMS_WRITE
func UpdateItem(token, locationId, itemId string, reqObj *UpdateItemReqObject) (*Item, error) {
	v := new(Item)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/items/%s", locationId, itemId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item and all item variations associated with it.
//
// Required permissions:  ITEMS_WRITE
func DeleteItem(token, locationId, itemId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s", locationId, itemId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Uploads a JPEG or PNG image and sets it as the master image for an item. See this article for recommended
// specifications for item images.
//
// If you upload an image for an item that already has a master image, the new image replaces
// the existing one.
//
// Important: Requests to this endpoint use the Content-Type: multipart/form-data
// header instead of Content-Type: application/json. It's recommended that you use an HTTP
// library (such as Requests for
// Python) that simplifies the process for sending multipart/form-data requests.
//
// The example request body shown assumes that you've set your request's multipart boundary to
// BOUNDARY in your Content-Type header, like so:
//
// Content-Type: multipart/form-data; boundary=BOUNDARY
//
// Note that some HTTP libraries set your request's multipart boundary for you.
//
// Required permissions:  ITEMS_WRITE
func UploadItemImage(token, locationId, itemId, imageName, imageMime string, body io.Reader) (*ItemImage, error) {
	v := new(ItemImage)
	b := bytes.NewBuffer(make([]byte))
	bw := multipart.NewWriter(b)
	boundary := bw.Boundary()
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", imageMime)
	mh.Set("Content-Disposition", fmt.Sprtinf(`form-data; name="image_data"; filename="%s"`, imageName))
	pw, err := bw.CreatePart(mh)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(pw, body)
	err = bw.Close()
	if err != nil {
		return nil, err
	}
	_, err = baseSquareRequest("POST", fmt.Sprintf("/v1/%s/items/%s/image", locationId, itemId), token, fmt.Sprintf("multipart/form-data; boundary=%s", boundary), b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateVariationReqObject struct {
	// The variation's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The item variation's name.
	Name string `json:"name"`
	// Indicates whether the item variation's price is fixed or determined at the time of
	// sale.Default value: FIXED_PRICING
	PricingType string `json:"pricing_type"`
	// The item variation's price, if any.
	PriceMoney Money `json:"price_money"`
	// The item variation's SKU, if any.
	Sku string `json:"sku"`
	// If true, inventory tracking is active for the variation.Default value: false
	TrackInventory bool `json:"track_inventory"`
	// Indicates whether the item variation displays an alert when its inventory quantity is
	// less than or equal to its inventory_alert_threshold.Default value: NONE
	InventoryAlertType string `json:"inventory_alert_type"`
	// If the inventory quantity for the variation is less than or equal to this value and
	// inventory_alert_type is LOW_QUANTITY, the variation displays an alert in
	// the merchant dashboard.This value is always an integer.Default value: 0
	InventoryAlertThreshold int `json:"inventory_alert_threshold"`
	// Arbitrary metadata to associate with the variation. Cannot exceed 255 characters.
	UserData string `json:"user_data"`
}

// Creates an item variation for an existing item.
//
// Required permissions:  ITEMS_WRITE
func CreateVariation(token, locationId, itemId string, reqObj *CreateVariationReqObject) (*ItemVariation, error) {
	v := new(ItemVariation)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/items/%s/variations", locationId, itemId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateVariationReqObject struct {
	// The item variation's name.
	Name string `json:"name"`
	// Indicates whether the item variation's price is fixed or determined at the time of
	// sale.Default value: FIXED_PRICING
	PricingType string `json:"pricing_type"`
	// The item variation's price, if any.
	PriceMoney Money `json:"price_money"`
	// The item variation's SKU, if any.
	Sku string `json:"sku"`
	// If true, inventory tracking is active for the variation.
	TrackInventory bool `json:"track_inventory"`
	// Indicates whether the item variation displays an alert when its inventory quantity goes
	// below its inventory_alert_threshold.
	InventoryAlertType string `json:"inventory_alert_type"`
	// If the inventory quantity for the variation is below this value and
	// inventory_alert_type is LOW_QUANTITY, the variation displays an alert in
	// the merchant dashboard.
	InventoryAlertThreshold int `json:"inventory_alert_threshold"`
	// Arbitrary metadata to associate with the variation. Cannot exceed 255 characters.
	UserData string `json:"user_data"`
}

// Modifies the details of an existing item variation.
//
// Required permissions:  ITEMS_WRITE
func UpdateVariation(token, locationId, itemId, variationId string, reqObj *UpdateVariationReqObject) (*ItemVariation, error) {
	v := new(ItemVariation)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationId, itemId, variationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item variation from an item.
//
// Every item must have at least one varation. This endpoint returns an error if you attempt to
// delete an item's only variation.
//
// Required permissions:  ITEMS_WRITE
func DeleteVariation(token, locationId, itemId, variationId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationId, itemId, variationId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides inventory information for all of a merchant's inventory-enabled item variations.
//
// See Managing inventory to learn how to enable an item
// variation for inventory tracking.
//
// Required permissions:  ITEMS_READ
//
// `limit`:
// The maximum number of inventory entries to return in a single response. This value
// cannot exceed 1000.This value is always an integer.Default value: 1000
func ListInventory(token, locationId string, limit int) ([]*InventoryEntry, *NextRequest, error) {
	v := make([]*InventoryEntry)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/inventory?limit=%d", locationId, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type AdjustInventoryReqObject struct {
	// The number to adjust the variation's quantity by.This value must be negative if adjustment_type is SALE, and it must be
	// positive if adjustment_type is RECEIVE_STOCK.
	QuantityDelta int `json:"quantity_delta"`
	// The reason for the inventory adjustment.
	AdjustmentType string `json:"adjustment_type"`
	// A note about the inventory adjustment.
	Memo string `json:"memo"`
}

// Adjusts an item variation's current available inventory.
//
// See Managing inventory to learn how to enable an item
// variation for inventory tracking.
//
// Required permissions:  ITEMS_WRITE
func AdjustInventory(token, locationId, variationId string, reqObj *AdjustInventoryReqObject) (*InventoryEntry, error) {
	v := new(InventoryEntry)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/inventory/%s", locationId, variationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateModifierListReqObject struct {
	// The modifier list's ID. Must be unique among all entity IDs ever provided on behalf of
	// the merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The modifier list's name.
	Name string `json:"name"`
	// Indicates whether multiple options from the modifier list can be applied to a single
	// item.Default value: SINGLE
	SelectionType string `json:"selection_type"`
	// The options included in the modifier list. You must include at least one modifier
	// option.
	ModifierOptions []ModifierOption `json:"modifier_options"`
}

// Creates an item modifier list and at least one modifier option for it.
//
// Required permissions:  ITEMS_WRITE
func CreateModifierList(token, locationId string, reqObj *CreateModifierListReqObject) (*ModifierList, error) {
	v := new(ModifierList)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's modifier lists.
//
// Required permissions:  ITEMS_READ
func ListModifierLists(token, locationId string) ([]*ModifierList, *NextRequest, error) {
	v := make([]*ModifierList)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides the details for a single modifier list.
//
// Required permissions:  ITEMS_READ
func RetrieveModifierList(token, locationId, modifierListId string) (*ModifierList, error) {
	v := new(ModifierList)
	_, err := squareRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationId, modifierListId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateModifierListReqObject struct {
	// The modifier list's name.
	Name string `json:"name"`
	// Indicates whether multiple options from the modifier list can be applied to a single
	// item.
	SelectionType string `json:"selection_type"`
}

// Modifies the details of an existing item modifier list.
//
// If you want to modify the details of a single modifier option, use the Update Modifier Option endpoint instead.
//
// Required permissions:  ITEMS_WRITE
func UpdateModifierList(token, locationId, modifierListId string, reqObj *UpdateModifierListReqObject) (*ModifierList, error) {
	v := new(ModifierList)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationId, modifierListId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item modifier list and all modifier options associated with it.
//
// Required permissions:  ITEMS_WRITE
func DeleteModifierList(token, locationId, modifierListId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationId, modifierListId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Associates a modifier list with an item, meaning modifier options from the list can be
// applied to the item.
//
// Required permissions:  ITEMS_WRITE
func ApplyModifierList(token, locationId, itemId, modifierListId string) (*Item, error) {
	v := new(Item)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationId, itemId, modifierListId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Removes a modifier list association from an item, meaning modifier options from the list can
// no longer be applied to the item.
//
// Required permissions:  ITEMS_WRITE
func RemoveModifierList(token, locationId, itemId, modifierListId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationId, itemId, modifierListId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateModifierOptionReqObject struct {
	// The modifier option's ID. Must be unique among all entity IDs ever provided on behalf
	// of the merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The modifier option's name.
	Name string `json:"name"`
	// The modifier option's price.
	PriceMoney Money `json:"price_money"`
	// If true, the modifier option is the default option in a modifier list for which
	// selection_type is SINGLE.Default value: false
	OnByDefault bool `json:"on_by_default"`
}

// Creates an item modifier option and adds it to a modifier list.
//
// Required permissions:  ITEMS_WRITE
func CreateModifierOption(token, locationId, modifierListId string, reqObj *CreateModifierOptionReqObject) (*ModifierOption, error) {
	v := new(ModifierOption)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options", locationId, modifierListId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateModifierOptionReqObject struct {
	// The modifier option's name.
	Name string `json:"name"`
	// The modifier option's price.
	PriceMoney Money `json:"price_money"`
	// If true, the modifier option is the default option in a modifier list for which
	// selection_type is SINGLE.
	OnByDefault bool `json:"on_by_default"`
}

// Modifies the details of an existing item modifier option.
//
// Required permissions:  ITEMS_WRITE
func UpdateModifierOption(token, locationId, modifierListId, modifierOptionId string, reqObj *UpdateModifierOptionReqObject) (*ModifierOption, error) {
	v := new(ModifierOption)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationId, modifierListId, modifierOptionId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item modifier option from a modifier list.
//
// Every modifier list must have at least one option. This endpoint returns an error if you
// attempt to delete a modifier list's only option.
//
// Required permissions:  ITEMS_WRITE
func DeleteModifierOption(token, locationId, modifierListId, modifierOptionId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationId, modifierListId, modifierOptionId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateCategoryReqObject struct {
	// The category's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The category's name.
	Name string `json:"name"`
}

// Creates an item category.
//
// To add or remove an item from a category, use the Update
// Item endpoint.
//
// Required permissions:  ITEMS_WRITE
func CreateCategory(token, locationId string, reqObj *CreateCategoryReqObject) (*Category, error) {
	v := new(Category)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/categories", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's item categories.
//
// Required permissions:  ITEMS_READ
func ListCategories(token, locationId string) ([]*Category, *NextRequest, error) {
	v := make([]*Category)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/categories", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateCategoryReqObject struct {
	// The new name of the category.
	Name string `json:"name"`
}

// Modifies the details of an existing item category.
//
// To add or remove an item from a category, use the Update
// Item endpoint.
//
// Required permissions:  ITEMS_WRITE
func UpdateCategory(token, locationId, categoryId string, reqObj *UpdateCategoryReqObject) (*Category, error) {
	v := new(Category)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/categories/%s", locationId, categoryId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item category.
//
// Required permissions:  ITEMS_WRITE
func DeleteCategory(token, locationId, categoryId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/categories/%s", locationId, categoryId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateDiscountReqObject struct {
	// The discount's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The discount's name.
	Name string `json:"name"`
	// The rate of the discount, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%. Specify a rate of 0 if discount_type
	// is VARIABLE_PERCENTAGE.Do not include this field for amount-based discounts.
	Rate string `json:"rate"`
	// The amount of the discount. Specify an amount of 0 if discount_type is
	// VARIABLE_AMOUNT.Do not include this field for rate-based discounts.
	AmountMoney Money `json:"amount_money"`
	// Indicates whether the discount is a FIXED value or entered at the time of
	// sale.Default value: FIXED
	DiscountType string `json:"discount_type"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount
	// to a payment.Default value: false
	PinRequired bool `json:"pin_required"`
	// The color of the discount's display label in Square Register.Default value: 9da2a6
	Color string `json:"color"`
}

// Creates a discount.
//
// Required permissions:  ITEMS_WRITE
func CreateDiscount(token, locationId string, reqObj *CreateDiscountReqObject) (*Discount, error) {
	v := new(Discount)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/discounts", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's discounts.
//
// Required permissions:  ITEMS_READ
func ListDiscounts(token, locationId string) ([]*Discount, *NextRequest, error) {
	v := make([]*Discount)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/discounts", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateDiscountReqObject struct {
	// The discount's name.
	Name string `json:"name"`
	// The rate of the discount, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%. Specify a rate of 0 if discount_type
	// is VARIABLE_PERCENTAGE.Do not include this field for amount-based discounts.
	Rate string `json:"rate"`
	// The amount of the discount. Specify an amount of 0 if discount_type is
	// VARIABLE_AMOUNT.Do not include this field for rate-based discounts.
	AmountMoney Money `json:"amount_money"`
	// Indicates whether the discount is a FIXED value or entered at the time of
	// sale.Default value: FIXED
	DiscountType string `json:"discount_type"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount
	// to a payment.Default value: false
	PinRequired bool `json:"pin_required"`
	// The color of the discount's display label in Square Register.Default value: 9da2a6
	Color string `json:"color"`
}

// Modifies the details of an existing discount.
//
// Required permissions:  ITEMS_WRITE
func UpdateDiscount(token, locationId, discountId string, reqObj *UpdateDiscountReqObject) (*Discount, error) {
	v := new(Discount)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/discounts/%s", locationId, discountId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing discount.
//
// Required permissions:  ITEMS_WRITE
func DeleteDiscount(token, locationId, discountId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/discounts/%s", locationId, discountId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateFeeReqObject struct {
	// The fee's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The fee's name.
	Name string `json:"name"`
	// The rate of the fee, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate"`
	// Whether the fee is calculated based on a payment's subtotal or total.Default value: FEE_SUBTOTAL_PHASE
	CalculationPhase string `json:"calculation_phase"`
	// The type of adjustment the fee applies to a payment. Currently, this value is
	// TAX for all fees.Default value: TAX
	AdjustmentType sring `json:"adjustment_type"`
	// If true, the fee applies to custom amounts entered into Square Register that are
	// not associated with a particular item.Default value: true
	AppliesToCustomAmounts bool `json:"applies_to_custom_amounts"`
	// If true, the fee is applied to payments. If false, it isn't.Default value: true
	Enabled bool `json:"enabled"`
	// Whether the fee is ADDITIVE or INCLUSIVE.Default value: ADDITIVE
	InclusionType string `json:"inclusion_type"`
}

// Creates a fee (tax).
//
// Required permissions:  ITEMS_WRITE
func CreateFee(token, locationId string, reqObj *CreateFeeReqObject) (*Fee, error) {
	v := new(Fee)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/fees", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's fees (taxes).
//
// Required permissions:  ITEMS_READ
func ListFees(token, locationId string) ([]*Fee, *NextRequest, error) {
	v := make([]*Fee)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/fees", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateFeeReqObject struct {
	// The fee's name.
	Name string `json:"name"`
	// The rate of the fee, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate"`
	// Whether the fee is calculated based on a payment's subtotal or total.Default value: FEE_SUBTOTAL_PHASE
	CalculationPhase string `json:"calculation_phase"`
	// The type of adjustment the fee applies to a payment. Currently, this value is
	// TAX for all fees.Default value: TAX
	AdjustmentType string `json:"adjustment_type"`
	// If true, the fee applies to custom amounts entered into Square Register that are
	// not associated with a particular item.Default value: true
	AppliesToCustomAmounts bool `json:"applies_to_custom_amounts"`
	// If true, the fee is applied to all appropriate items. If false, the fee
	// is not applied at all.Default value: true
	Enabled bool `json:"enabled"`
	// Whether the fee is ADDITIVE or INCLUSIVE.Default value: ADDITIVE
	InclusionType string `json:"inclusion_type"`
}

// Modifies the details of an existing fee (tax).
//
// Required permissions:  ITEMS_WRITE
func UpdateFee(token, locationId, feeId string, reqObj *UpdateFeeReqObject) (*Fee, error) {
	v := new(Fee)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/fees/%s", locationId, feeId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing fee (tax).
//
// Required permissions:  ITEMS_WRITE
func DeleteFee(token, locationId, feeId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/fees/%s", locationId, feeId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Associates a fee with an item, meaning the fee is automatically applied to the item in Square
// Register.
//
// Required permissions:  ITEMS_WRITE
func ApplyFee(token, locationId, itemId, feeId string) (*Item, error) {
	v := new(Item)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationId, itemId, feeId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Removes a fee assocation from an item, meaning the fee is no longer automatically applied to
// the item in Square Register.
//
// Required permissions:  ITEMS_WRITE
func RemoveFee(token, locationId, itemId, feeId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationId, itemId, feeId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreatePageReqObject struct {
	// The page's ID. Must be unique among all entity IDs ever provided on behalf of the
	// merchant. You can never reuse an ID. This value can include alphanumeric
	// characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id string `json:"id"`
	// The page's name.
	Name string `json:"name"`
	// The page's position in the list of pages. Must be an integer between 0 and
	// 4, inclusive.The endpoint returns an error if you specify a page_index that another page is
	// already using.
	PageIndex int `json:"page_index"`
}

// Creates a Favorites page in Square Register.
//
// A merchant can have up to five pages, each of which has a page_index between 0
// and 4, inclusive.
//
// After you create a page, you can set the values of its cells with the Update Cell endpoint. A page doesn't appear in Square Register
// unless at least one of its cells has an assigned value.
//
// Required permissions:  ITEMS_WRITE
func CreatePage(token, locationId string, reqObj *CreatePageReqObject) (*Page, error) {
	v := new(Page)
	_, err := squareRequest("POST", fmt.Sprintf("/v1/%s/pages", locationId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's Favorites pages in Square Register.
//
// Required permissions:  ITEMS_READ
func ListPages(token, locationId string) ([]*Page, *NextRequest, error) {
	v := make([]*Page)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/pages", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdatePageReqObject struct {
	// The page's name.
	Name string `json:"name"`
	// The page's position in the merchant's list of pages. Must be an integer between
	// 0 and 4, inclusive.The page's index is not updated if another page already exists at the specified
	// index.
	PageIndex int `json:"page_index"`
}

// Modifies the details of a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
func UpdatePage(token, locationId, pageId string, reqObj *UpdatePageReqObject) (*Page, error) {
	v := new(Page)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s", locationId, pageId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing Favorites page and all of its cells.
//
// Required permissions:  ITEMS_WRITE
func DeletePage(token, locationId, pageId string) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s", locationId, pageId), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateCellReqObject struct {
	// The row of the cell to update. Always an integer between 0 and 4,
	// inclusive. Row 0 is the top row.
	Row int `json:"row"`
	// The column of the cell to update. Always an integer between 0 and 4,
	// inclusive. Column 0 is the leftmost row.
	Column int `json:"column"`
	// The type of entity represented in the cell (ITEM, DISCOUNT,
	// CATEGORY, or PLACEHOLDER).
	ObjectType string `json:"object_type"`
	// The unique identifier of the entity to represent in the cell. Do not include if the
	// cell's object_type is PLACEHOLDER.
	ObjectId string `json:"object_id"`
	// For a cell with an object_type of PLACEHOLDER, indicates the cell's
	// behavior.
	PlaceholderType string `json:"placeholder_type"`
}

// Modifies a cell of a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
func UpdateCell(token, locationId, pageId string, reqObj *UpdateCellReqObject) (*PageCell, error) {
	v := new(PageCell)
	_, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s/cells", locationId, pageId), token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes a cell from a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
//
// `row`:
// The row of the cell to clear. Always an integer between 0 and 4,
// inclusive. Row 0 is the top row.
//
// `column`:
// The column of the cell to clear. Always an integer between 0 and 4,
// inclusive. Column 0 is the leftmost column.
func DeleteCell(token, locationId, pageId string, row number, column number) error {
	_, err := squareRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s/cells?row=%s&column=%s", locationId, pageId, row, column), token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type SubmitBatchReqObject struct {
	// The requests to perform.
	Requests []BatchRequest `json:"requests"`
}

// Lets you batch multiple requests to other Connect API endpoints into a single request. This
// endpoint's response is an array that contains the response for each batched request.
//
// You don't need to provide an access token in the header of your request to the Submit Batch
// endpoint. Instead, you provide an access_token parameter for each request included in
// the batch.
//
// Note the following when using the Submit Batch endpoint:
func SubmitBatch(token string, reqObj *SubmitBatchReqObject) ([]*BatchResponse, *NextRequest, error) {
	if len(reqObj.Requests) > 30 {
		return nil, nil, fmt.Errorf("You cannot submit more than 30 requests to `/v1/batch`")
	}
	v := make([]*BatchResponse)
	nr, err := squareRequest("POST", "/v1/batch", token, reqObj, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Lists which types of events trigger webhook notifications for a particular location.
func ListWebhooks(token, locationId string) ([]*WebhookEventType, *NextRequest, error) {
	v := make([]*WebhookEventType)
	nr, err := squareRequest("GET", fmt.Sprintf("/v1/%s/webhooks", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Sets which types of events trigger webhook notifications for a location.
//
// Simply provide a JSON array of the event types you want notifications for in your request
// body (see Example Requests below).
func UpdateWebhooks(token, locationId string) ([]*WebhookEventType, *NextRequest, error) {
	v := make([]*WebhookEventType)
	nr, err := squareRequest("PUT", fmt.Sprintf("/v1/%s/webhooks", locationId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Lists subscriptions that have been created for an application. You can look up subscription
// information for a single merchant by providing the merchant_id parameter to this
// endpoint.
//
// Subscription objects returned by this endpoint do not include the fees
// field, which lists the subscription's payment history. To get a particular subscription's
// payment history, use the Retrieve Subscription
// endpoint.
//
// Important: The Authorization header you provide to this endpoint must have the
// following format:
//
// `merchantId`:
// If you provide this parameter, the endpoint returns only subscription information for
// the specified merchant.You can get a merchant's ID with the Retrieve
// Merchant endpoint.
//
// `limit`:
// The maximum number of subscriptions to return in a single response. This value cannot
// exceed 200.Default value: 100
func ListSubscriptions(token, clientId, merchantId string, limit int) ([]*Subscription, *NextRequest, error) {
	v := make([]*Subscription)
	nr, err := squareRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions?merchant_id=%s&limit=%d", clientId, merchantId, limit), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

// Provides comprehensive information for a single subscription, including its payment
// history.
//
// Important: The Authorization header you provide to this endpoint must have the
// following format:
func RetrieveSubscription(token, clientId, subscriptionId string) (*Subscription, error) {
	v := new(Subscription)
	_, err := squareRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions/%s", clientId, subscriptionId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides information for all of an application's subscription plans.
//
// Important: The Authorization header you provide to this endpoint must have the
// following format:
func ListSubscriptionPlans(token, clientId string) ([]*SubscriptionPlan, *NextRequest, error) {
	v := make([]*SubscriptionPlan)
	nr, err := squareRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans", clientId), token, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, nr, nil
}

func RetrieveSubscriptionPlan(token, clientId, planId string) (*SubscriptionPlan, error) {
	v := new(SubscriptionPlan)
	_, err := squareRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans/%s", clientId, planId), token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
