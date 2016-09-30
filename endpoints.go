package gosquare

// Provides a business's account information, such as its name and associated email address.
//
// Required permissions:  MERCHANT_PROFILE_READ
func RetrieveBusiness(token string) (*Merchant, error) {
	v := new(Merchant)
	err := v1Request("GET", "/v1/me", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides details for a business's locations, including their IDs.
//
// The account_capabilities array returned in each Merchant object indicates which
//       account capabilities the location has enabled. For example, if this array does not include the
//       value CREDIT_CARD_PROCESSING, the location cannot currently process credit cards with
//       Square.
//
// Required permissions:  MERCHANT_PROFILE_READ
func ListLocations(token string) ([]*Merchant, error) {
	v := new(Merchant)
	err := v1Request("GET", "/v1/me/locations", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateEmployeeReqObject struct {
	// The employee's first name.
	FirstName []string `json:"first_name"`
	// The employee's last name.
	LastName []string `json:"last_name"`
	// An optional second ID to associate the employee with an entity in another system.
	ExternalId []string `json:"external_id"`
	// The ids of the employee's associated roles. Currently, you can specify only one
	//             or zero roles per employee.Default value: []
	RoleIds []string `json:"role_ids"`
	// An optional email address to associate with the employee.Note that you cannot edit an existing employee's email address with the Connect API.
	//             You can only set its initial value when creating an employee.
	Email []string `json:"email"`
}

// Creates an employee for a business.
//
// Required permissions:  EMPLOYEES_WRITE
func CreateEmployee(token string, reqObj CreateEmployeeReqObject) (*Employee, error) {
	v := new(Employee)
	err := v1Request("POST", "/v1/me/employees", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employees.
//
// You can filter the results returned by this endpoint by exactly one of the following
//       fields:
func ListEmployees(token string) ([]*Employee, error) {
	v := new([]*Employee)
	err := v1Request("GET", "/v1/me/employees", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for a single employee.
//
// Required permissions:  EMPLOYEES_READ
func RetrieveEmployee(token, EmployeeId string) (*Employee, error) {
	v := new(Employee)
	err := v1Request("GET", "/v1/me/employees/"+EmployeeId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateEmployeeReqObject struct {
	// The ID of the employee to modify.
	EmployeeId []string `json:"employee_id"`
	// The employee's first name.
	FirstName []string `json:"first_name"`
	// The employee's last name.
	LastName []string `json:"last_name"`
	// An optional second ID to associate the employee with an entity in another system.
	ExternalId []string `json:"external_id"`
	// The employee's associated roles. Currently, you can specify only one or zero roles per
	//             employee.
	RoleIds []string `json:"role_ids"`
}

// Modifies the details of an employee.
//
// Required permissions:  EMPLOYEES_WRITE
func UpdateEmployee(token, EmployeeId string, reqObj UpdateEmployeeReqObject) (*Employee, error) {
	v := new(Employee)
	err := v1Request("PUT", "/v1/me/employees/"+EmployeeId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateRoleReqObject struct {
	// The role's name.
	Name []string `json:"name"`
	// The role's permissions.
	Permissions []string `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the
	//             values indicated in permissions.Default value: false
	IsOwner []bool `json:"is_owner"`
}

// Creates an employee role you can then assign to employees.
//
// Required permissions:  EMPLOYEES_WRITE
func CreateRole(token string, reqObj CreateRoleReqObject) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	err := v1Request("POST", "/v1/me/roles", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employee roles.
//
// Required permissions:  EMPLOYEES_READ
func ListRoles(token string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/me/roles", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for a single employee role.
//
// Required permissions:  EMPLOYEES_READ
func RetrieveRole(token, RoleId string) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	err := v1Request("GET", "/v1/me/roles/"+RoleId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateRoleReqObject struct {
	// The ID of the role to modify.
	RoleId []string `json:"role_id"`
	// The role's name.
	Name []string `json:"name"`
	// The role's permissions.
	Permissions []EmployeeRole.Permission `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the
	//             values indicated in permissions.
	IsOwner []bool `json:"is_owner"`
}

// Modifies the details of an employee role.
//
// Required permissions:  EMPLOYEES_WRITE
func UpdateRole(token, RoleId string, reqObj UpdateRoleReqObject) (*EmployeeRole, error) {
	v := new(EmployeeRole)
	err := v1Request("PUT", "/v1/me/roles/"+RoleId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateTimecardReqObject struct {
	// The employee to create a timecard for.
	EmployeeId []string `json:"employee_id"`
	// The clock-in time for the timecard, in ISO 8601 format.Default value: The current time.
	ClockinTime []string `json:"clockin_time"`
	// The clock-out time for the timecard, in ISO 8601 format.
	//             Provide this value only if importing timecard information from another system.
	ClockoutTime []string `json:"clockout_time"`
	// The ID of the location the employee clocked in from, if any.
	ClockinLocationId []string `json:"clockin_location_id"`
	// The ID of the location the employee clocked out from. Provide this value only if
	//             importing timecard information from another system.If you provide this value, you must also provide a value for clockout_time.
	ClockoutLocationId []string `json:"clockout_location_id"`
}

// Creates a timecard for an employee. Each timecard corresponds to a single shift.
//
// This endpoint automatically creates an API_CREATE event for the new timecard.
//
// Required permissions:  TIMECARDS_WRITE
func CreateTimecard(token string, reqObj CreateTimecardReqObject) (*Timecard, error) {
	v := new(Timecard)
	err := v1Request("POST", "/v1/me/timecards", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a business's employee timecards.
//
// You can filter the results returned by this endpoint by exactly one of the following
//       fields:
func ListTimecards(token string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/me/timecards", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func RetrieveTimecard(token, TimecardId string) (*Timecard, error) {
	v := new(Timecard)
	err := v1Request("GET", "/v1/me/timecards/"+TimecardId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateTimecardReqObject struct {
	// The ID of the timecard to modify.
	TimecardId []string `json:"timecard_id"`
	// The clock-in time for the timecard, in ISO 8601 format.
	ClockinTime []string `json:"clockin_time"`
	// The clock-out time for the timecard, in ISO 8601 format.
	ClockoutTime []string `json:"clockout_time"`
	// The ID of the location the employee clocked in from, if any.
	ClockinLocationId []string `json:"clockin_location_id"`
	// The ID of the location the employee clocked out from, if any.
	ClockoutLocationId []string `json:"clockout_location_id"`
}

// Modifies a timecard's details. This creates an API_EDIT event for the timecard. You
//       can view a timecard's event history with the List Timecard
//         Events endpoint.
//
// Required permissions:  TIMECARDS_WRITE
func UpdateTimecard(token, TimecardId string, reqObj UpdateTimecardReqObject) (*Timecard, error) {
	v := new(Timecard)
	err := v1Request("PUT", "/v1/me/timecards/"+TimecardId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func DeleteTimecard(token, TimecardId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/me/timecards/"+TimecardId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all events associated with a particular timecard.
//
// Required permissions:  TIMECARDS_READ
func ListTimecardEvents(token, TimecardId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/me/timecards/"+TimecardId+"/events", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for all of a location's cash drawer shifts during a date range. The date
//       range you specify cannot exceed 90 days.
//
// CashDrawerShift objects returned by this endpoint do not include the
//         events field, which lists the events that occurred during the shift. To get a
//       particular shift's events, use the Retrieve Cash Drawer
//         Shift endpoint.
//
// Required permissions:  PAYMENTS_READ
func ListCashDrawerShifts(token, LocationId string) ([]*CashDrawerShift, error) {
	v := new(CashDrawerShift)
	err := v1Request("GET", "/v1/"+LocationId+"/cash-drawer-shifts", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for a single cash drawer shift, including all events that occurred
//       during the shift.
//
// Required permissions:  PAYMENTS_READ
func RetrieveCashDrawerShift(token, LocationId, ShiftId string) (*CashDrawerShift, error) {
	v := new(CashDrawerShift)
	err := v1Request("GET", "/v1/"+LocationId+"/cash-drawer-shifts/"+ShiftId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all payments taken by a merchant or any of the merchant's
//         mobile staff during a date range.
//       Date ranges cannot exceed one year in length. See Date ranges
//       for details of inclusive and exclusive dates.
//
// Required permissions:  PAYMENTS_READ
func ListPayments(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/payments", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides comprehensive information for a single payment.
//
// Required permissions:  PAYMENTS_READ
func RetrievePayment(token, LocationId, PaymentId string) (*Payment, error) {
	v := new(Payment)
	err := v1Request("GET", "/v1/"+LocationId+"/payments/"+PaymentId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all deposits and withdrawals initiated by Square to a
//       merchant's bank account during a date range. Date ranges cannot exceed one year in length. See
//         Date ranges for details of inclusive and exclusive
//       dates.
//
// Settlement objects returned by this endpoint do not include the entries
//       field, which lists the transactions that contribute to the total of the settlement. To get a
//       particular settlement's entries, use the Retrieve
//         Settlement endpoint.
//
// Square initiates its regular deposits to merchant bank accounts on the schedule indicated on
//         this page. A deposit initiated by
//       Square on a given day is usually not provided by this endpoint before 10 p.m. PST the same
//       day.
//
// Square does not know when an initiated settlement completes, only whether it has
//         failed. A completed settlement is typically reflected in a merchant's bank account
//       within three business days, but in exceptional cases it might take longer.
//
// Required permissions:  SETTLEMENTS_READ
func ListSettlements(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/settlements", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides comprehensive information for a single settlement, including the entries that
//       contribute to the settlement's total.
//
// See SettlementEntry.Type for
//       descriptions of the types of entries that compose a settlement.
//
// Required permissions:  SETTLEMENTS_READ
func RetrieveSettlement(token, LocationId, SettlementId string) (*Settlement, error) {
	v := new(Settlement)
	err := v1Request("GET", "/v1/"+LocationId+"/settlements/"+SettlementId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateRefundReqObject struct {
	// The ID of the original payment's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the payment to refund.If you're creating a PARTIAL refund for a split tender payment, instead provide
	//             the id of the particular tender you want to refund. See Split Tender Payments for details.
	PaymentId []string `json:"payment_id"`
	// The type of refund (FULL or PARTIAL).
	Type []Refund.Type `json:"type"`
	// The reason for the refund.
	Reason []string `json:"reason"`
	// The amount of money to refund. Required only for PARTIAL refunds.The value of amount must be negative.
	RefundedMoney []Money `json:"refunded_money"`
	// An optional key to ensure idempotence if you issue the same PARTIAL refund
	//             request more than once.If you attempt to issue a partial refund and you aren't sure whether your request
	//             succeeded, you can safely repeat your request with the same
	//               request_idempotence_key. If you want to issue another partial refund for
	//             the same payment, you must use a request_idempotence_key that is unique among
	//             refunds you have issued for the payment.
	RequestIdempotenceKey []string `json:"request_idempotence_key"`
}

// Issues a refund for a previously processed payment. You must issue a refund within 60 days of
//       the associated payment. See this
//         article for more information on refund behavior.
//
// Issuing a refund for a card payment is not reversible. To develop against this
//       endpoint, you can create fake cash payments in Square Register and refund them.
//
// You can issue either full refunds or partial refunds. If you issue a partial refund, you must
//       specify the amount of money to refund.
//
// Required permissions:  PAYMENTS_WRITE
func CreateRefund(token, LocationId string, reqObj CreateRefundReqObject) (*Refund, error) {
	v := new(Refund)
	err := v1Request("POST", "/v1/"+LocationId+"/refunds", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for all refunds initiated by a merchant or any of the merchant's mobile staff during a date range. Date
//       ranges cannot exceed one year in length. See Date ranges for
//       details of inclusive and exclusive dates.
//
// Required permissions:  PAYMENTS_READ
func ListRefunds(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/refunds", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ListOrders(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/orders", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func RetrieveOrder(token, LocationId, OrderId string) (*Order, error) {
	v := new(Order)
	err := v1Request("GET", "/v1/"+LocationId+"/orders/"+OrderId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateOrderReqObject struct {
	// The ID of the order's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the order to modify.
	OrderId []string `json:"order_id"`
	// The action to perform on the order (COMPLETE, CANCEL, or
	//             REFUND).
	Action []Order.Action `json:"action"`
	// The tracking number of the shipment associated with the order. Only valid if
	//               action is COMPLETE.
	ShippedTrackingNumber []string `json:"shipped_tracking_number"`
	// A merchant-specified note about the completion of the order. Only valid if
	//               action is COMPLETE.
	CompletedNote []string `json:"completed_note"`
	// A merchant-specified note about the refunding of the order. Only valid if action
	//             is REFUND.
	RefundedNote []string `json:"refunded_note"`
	// A merchant-specified note about the canceling of the order. Only valid if action
	//             is CANCEL.
	CanceledNote []string `json:"canceled_note"`
}

func UpdateOrder(token, LocationId, OrderId string, reqObj UpdateOrderReqObject) (*Order, error) {
	v := new(Order)
	err := v1Request("PUT", "/v1/"+LocationId+"/orders/"+OrderId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides non-confidential details for all of a location's associated bank accounts. This
//       endpoint does not provide full bank account numbers, and there is no way to obtain a
//       full bank account number with the Connect API.
//
// Required permissions:  BANK_ACCOUNTS_READ
func ListBankAccounts(token, LocationId string) ([]*BankAccount, error) {
	v := new(BankAccount)
	err := v1Request("GET", "/v1/"+LocationId+"/bank-accounts", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides non-confidential details for a merchant's associated bank account. This endpoint
//       does not provide full bank account numbers, and there is no way to obtain a full bank
//       account number with the Connect API.
//
// Required permissions:  BANK_ACCOUNTS_READ
func RetrieveBankAccount(token, LocationId, BankAccountId string) (*BankAccount, error) {
	v := new(BankAccount)
	err := v1Request("GET", "/v1/"+LocationId+"/bank-accounts/"+BankAccountId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateItemReqObject struct {
	// The ID of the location to create an item for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The item's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The item's name.
	Name []string `json:"name"`
	// The item's description.
	Description []string `json:"description"`
	// The ID of the item's category, if any.
	CategoryId []string `json:"category_id"`
	// The color of the item's display label in Square Register.Default value: 9da2a6
	Color []Item.Color `json:"color"`
	// The text of the item's display label in Square Register. Only up to the first five
	//             characters of the string are used.Default value: The first two characters of the item's name.
	Abbreviation []string `json:"abbreviation"`
	// Indicates whether the item is viewable from the merchant's online store (PUBLIC)
	//             or PRIVATE.Default value: PUBLIC
	Visibility []Item.Visibility `json:"visibility"`
	// If true, the item can be added to shipping orders from the merchant's online
	//             store.Default value: false
	AvailableOnline []bool `json:"available_online"`
	// If true, the item can be added to pickup orders from the merchant's online
	//             store.Default value: false
	AvailableForPickup []bool `json:"available_for_pickup"`
	// The item's variations. You must specify at least one variation.
	Variations []ItemVariation `json:"variations"`
}

// Creates an item and at least one variation for it.
//
// Required permissions:  ITEMS_WRITE
func CreateItem(token, LocationId string, reqObj CreateItemReqObject) (*Item, error) {
	v := new(Item)
	err := v1Request("POST", "/v1/"+LocationId+"/items", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides summary information for all of a location's items.
//
// Required permissions:  ITEMS_READ
func ListItems(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/items", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for a single item, including associated modifier lists and fees.
//
// Required permissions:  ITEMS_READ
func RetrieveItem(token, LocationId, ItemId string) (*Item, error) {
	v := new(Item)
	err := v1Request("GET", "/v1/"+LocationId+"/items/"+ItemId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateItemReqObject struct {
	// The ID of the item's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the item to modify.
	ItemId []string `json:"item_id"`
	// The item's name.
	Name []string `json:"name"`
	// The item's description.
	Description []string `json:"description"`
	// The ID of the item's category, if any.If you provide the empty string for this value, any existing category association is
	//             removed from the item.
	CategoryId []string `json:"category_id"`
	// The color of the item's display label in Square Register.
	Color []Item.Color `json:"color"`
	// The text of the item's display label in Square Register. Only up to the first five
	//             characters of the string are used.
	Abbreviation []string `json:"abbreviation"`
	// Indicates whether the item is viewable from the merchant's online store (PUBLIC) or
	//               PRIVATE.
	Visibility []Item.Visibility `json:"visibility"`
	// If true, the item can be purchased from the merchant's online store.
	AvailableOnline []bool `json:"available_online"`
	// If true, the item can be added to pickup orders from the merchant's online
	//             store.
	AvailableForPickup []bool `json:"available_for_pickup"`
}

// Modifies the core details of an existing item.
//
// If you want to modify an item's variations, use the Update Variation endpoint instead.
//
// If you want to add or remove a modifier list from an item, use the Apply Modifier List and Remove Modifier List endpoints instead.
//
// If you want to add or remove a fee from an item, use the Apply Fee and Remove Fee endpoints
//       instead.
//
// Required permissions:  ITEMS_WRITE
func UpdateItem(token, LocationId, ItemId string, reqObj UpdateItemReqObject) (*Item, error) {
	v := new(Item)
	err := v1Request("PUT", "/v1/"+LocationId+"/items/"+ItemId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item and all item variations associated with it.
//
// Required permissions:  ITEMS_WRITE
func DeleteItem(token, LocationId, ItemId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/items/"+ItemId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UploadItemImageReqObject struct {
	// The ID of the item's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the item to associate the image with.
	ItemId []string `json:"item_id"`
	// The image's binary data.
	ImageData []data `json:"image_data"`
}

// Uploads a JPEG or PNG image and sets it as the master image for an item. See this article for recommended
//       specifications for item images.
//
// If you upload an image for an item that already has a master image, the new image replaces
//       the existing one.
//
// Important: Requests to this endpoint use the Content-Type: multipart/form-data
//       header instead of Content-Type: application/json. It's recommended that you use an HTTP
//       library (such as Requests for
//       Python) that simplifies the process for sending multipart/form-data requests.
//
// The example request body shown assumes that you've set your request's multipart boundary to
//         BOUNDARY in your Content-Type header, like so:
//
// Content-Type: multipart/form-data; boundary=BOUNDARY
//
// Note that some HTTP libraries set your request's multipart boundary for you.
//
// Required permissions:  ITEMS_WRITE
func UploadItemImage(token, LocationId, ItemId string, reqObj UploadItemImageReqObject) (*ItemImage, error) {
	v := new(ItemImage)
	err := v1Request("POST", "/v1/"+LocationId+"/items/"+ItemId+"/image", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateVariationReqObject struct {
	// The ID of the item's associated location.Get a business's locations with the List Locations
	//             endpoint.
	MerchantId []string `json:"merchant_id"`
	// The ID of the item the variation applies to.
	ItemId []string `json:"item_id"`
	// The variation's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The item variation's name.
	Name []string `json:"name"`
	// Indicates whether the item variation's price is fixed or determined at the time of
	//             sale.Default value: FIXED_PRICING
	PricingType []ItemVariation.PricingType `json:"pricing_type"`
	// The item variation's price, if any.
	PriceMoney []Money `json:"price_money"`
	// The item variation's SKU, if any.
	Sku []string `json:"sku"`
	// If true, inventory tracking is active for the variation.Default value: false
	TrackInventory []bool `json:"track_inventory"`
	// Indicates whether the item variation displays an alert when its inventory quantity is
	//             less than or equal to its inventory_alert_threshold.Default value: NONE
	InventoryAlertType []InventoryAlertType `json:"inventory_alert_type"`
	// If the inventory quantity for the variation is less than or equal to this value and
	//               inventory_alert_type is LOW_QUANTITY, the variation displays an alert in
	//             the merchant dashboard.This value is always an integer.Default value: 0
	InventoryAlertThreshold []number `json:"inventory_alert_threshold"`
	// Arbitrary metadata to associate with the variation. Cannot exceed 255 characters.
	UserData []string `json:"user_data"`
}

// Creates an item variation for an existing item.
//
// Required permissions:  ITEMS_WRITE
func CreateVariation(token, LocationId, ItemId string, reqObj CreateVariationReqObject) (*ItemVariation, error) {
	v := new(ItemVariation)
	err := v1Request("POST", "/v1/"+LocationId+"/items/"+ItemId+"/variations", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateVariationReqObject struct {
	// The ID of the variation's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the item the variation applies to.
	ItemId []string `json:"item_id"`
	// The ID of the variation to modify.
	VariationId []string `json:"variation_id"`
	// The item variation's name.
	Name []string `json:"name"`
	// Indicates whether the item variation's price is fixed or determined at the time of
	//             sale.Default value: FIXED_PRICING
	PricingType []ItemVariation.PricingType `json:"pricing_type"`
	// The item variation's price, if any.
	PriceMoney []Money `json:"price_money"`
	// The item variation's SKU, if any.
	Sku []string `json:"sku"`
	// If true, inventory tracking is active for the variation.
	TrackInventory []bool `json:"track_inventory"`
	// Indicates whether the item variation displays an alert when its inventory quantity goes
	//             below its inventory_alert_threshold.
	InventoryAlertType []InventoryAlertType `json:"inventory_alert_type"`
	// If the inventory quantity for the variation is below this value and
	//               inventory_alert_type is LOW_QUANTITY, the variation displays an alert in
	//             the merchant dashboard.
	InventoryAlertThreshold []number `json:"inventory_alert_threshold"`
	// Arbitrary metadata to associate with the variation. Cannot exceed 255 characters.
	UserData []string `json:"user_data"`
}

// Modifies the details of an existing item variation.
//
// Required permissions:  ITEMS_WRITE
func UpdateVariation(token, LocationId, ItemId, VariationId string, reqObj UpdateVariationReqObject) (*ItemVariation, error) {
	v := new(ItemVariation)
	err := v1Request("PUT", "/v1/"+LocationId+"/items/"+ItemId+"/variations/"+VariationId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item variation from an item.
//
// Every item must have at least one varation. This endpoint returns an error if you attempt to
//       delete an item's only variation.
//
// Required permissions:  ITEMS_WRITE
func DeleteVariation(token, LocationId, ItemId, VariationId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/items/"+ItemId+"/variations/"+VariationId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides inventory information for all of a merchant's inventory-enabled item variations.
//
// See Managing inventory to learn how to enable an item
//       variation for inventory tracking.
//
// Required permissions:  ITEMS_READ
func ListInventory(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/inventory", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type AdjustInventoryReqObject struct {
	// The ID of the variation's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the variation to adjust inventory information for.
	VariationId []string `json:"variation_id"`
	// The number to adjust the variation's quantity by.This value must be negative if adjustment_type is SALE, and it must be
	//             positive if adjustment_type is RECEIVE_STOCK.
	QuantityDelta []number `json:"quantity_delta"`
	// The reason for the inventory adjustment.
	AdjustmentType []InventoryAdjustmentType `json:"adjustment_type"`
	// A note about the inventory adjustment.
	Memo []string `json:"memo"`
}

// Adjusts an item variation's current available inventory.
//
// See Managing inventory to learn how to enable an item
//       variation for inventory tracking.
//
// Required permissions:  ITEMS_WRITE
func AdjustInventory(token, LocationId, VariationId string, reqObj AdjustInventoryReqObject) (*InventoryEntry, error) {
	v := new(InventoryEntry)
	err := v1Request("POST", "/v1/"+LocationId+"/inventory/"+VariationId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateModifierListReqObject struct {
	// The ID of the location to create a modifier list for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The modifier list's ID. Must be unique among all entity IDs ever provided on behalf of
	//             the merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The modifier list's name.
	Name []string `json:"name"`
	// Indicates whether multiple options from the modifier list can be applied to a single
	//             item.Default value: SINGLE
	SelectionType []ModifierList.SelectionType `json:"selection_type"`
	// The options included in the modifier list. You must include at least one modifier
	//             option.
	ModifierOptions []ModifierOption `json:"modifier_options"`
}

// Creates an item modifier list and at least one modifier option for it.
//
// Required permissions:  ITEMS_WRITE
func CreateModifierList(token, LocationId string, reqObj CreateModifierListReqObject) (*ModifierList, error) {
	v := new(ModifierList)
	err := v1Request("POST", "/v1/"+LocationId+"/modifier-lists", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's modifier lists.
//
// Required permissions:  ITEMS_READ
func ListModifierLists(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/modifier-lists", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides the details for a single modifier list.
//
// Required permissions:  ITEMS_READ
func RetrieveModifierList(token, LocationId, ModifierListId string) (*ModifierList, error) {
	v := new(ModifierList)
	err := v1Request("GET", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateModifierListReqObject struct {
	// The ID of the modifier list's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the modifier list to edit.
	ModifierListId []string `json:"modifier_list_id"`
	// The modifier list's name.
	Name []string `json:"name"`
	// Indicates whether multiple options from the modifier list can be applied to a single
	//             item.
	SelectionType []ModifierList.SelectionType `json:"selection_type"`
}

// Modifies the details of an existing item modifier list.
//
// If you want to modify the details of a single modifier option, use the Update Modifier Option endpoint instead.
//
// Required permissions:  ITEMS_WRITE
func UpdateModifierList(token, LocationId, ModifierListId string, reqObj UpdateModifierListReqObject) (*ModifierList, error) {
	v := new(ModifierList)
	err := v1Request("PUT", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item modifier list and all modifier options associated with it.
//
// Required permissions:  ITEMS_WRITE
func DeleteModifierList(token, LocationId, ModifierListId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type ApplyModifierListReqObject struct {
	// The ID of the modifier list's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the item to add the modifier list to.
	ItemId []string `json:"item_id"`
	// The ID of the modifier list to apply.
	ModifierListId []string `json:"modifier_list_id"`
}

// Associates a modifier list with an item, meaning modifier options from the list can be
//       applied to the item.
//
// Required permissions:  ITEMS_WRITE
func ApplyModifierList(token, LocationId, ItemId, ModifierListId string, reqObj ApplyModifierListReqObject) (*Item, error) {
	v := new(Item)
	err := v1Request("PUT", "/v1/"+LocationId+"/items/"+ItemId+"/modifier-lists/"+ModifierListId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Removes a modifier list association from an item, meaning modifier options from the list can
//       no longer be applied to the item.
//
// Required permissions:  ITEMS_WRITE
func RemoveModifierList(token, LocationId, ItemId, ModifierListId string) (*Item, error) {
	v := new(Item)
	err := v1Request("DELETE", "/v1/"+LocationId+"/items/"+ItemId+"/modifier-lists/"+ModifierListId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateModifierOptionReqObject struct {
	// The ID of the modifier list's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the modifier list to add the option to.
	ModifierListId []string `json:"modifier_list_id"`
	// The modifier option's ID. Must be unique among all entity IDs ever provided on behalf
	//             of the merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The modifier option's name.
	Name []string `json:"name"`
	// The modifier option's price.
	PriceMoney []Money `json:"price_money"`
	// If true, the modifier option is the default option in a modifier list for which
	//               selection_type is SINGLE.Default value: false
	OnByDefault []bool `json:"on_by_default"`
}

// Creates an item modifier option and adds it to a modifier list.
//
// Required permissions:  ITEMS_WRITE
func CreateModifierOption(token, LocationId, ModifierListId string, reqObj CreateModifierOptionReqObject) (*ModifierOption, error) {
	v := new(ModifierOption)
	err := v1Request("POST", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId+"/modifier-options", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateModifierOptionReqObject struct {
	// The ID of the modifier option's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the modifier list that contains the option to edit.
	ModifierListId []string `json:"modifier_list_id"`
	// The ID of the modifier option to edit.
	ModifierOptionId []string `json:"modifier_option_id"`
	// The modifier option's name.
	Name []string `json:"name"`
	// The modifier option's price.
	PriceMoney []Money `json:"price_money"`
	// If true, the modifier option is the default option in a modifier list for which
	//               selection_type is SINGLE.
	OnByDefault []bool `json:"on_by_default"`
}

// Modifies the details of an existing item modifier option.
//
// Required permissions:  ITEMS_WRITE
func UpdateModifierOption(token, LocationId, ModifierListId, ModifierOptionId string, reqObj UpdateModifierOptionReqObject) (*ModifierOption, error) {
	v := new(ModifierOption)
	err := v1Request("PUT", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId+"/modifier-options/"+ModifierOptionId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item modifier option from a modifier list.
//
// Every modifier list must have at least one option. This endpoint returns an error if you
//       attempt to delete a modifier list's only option.
//
// Required permissions:  ITEMS_WRITE
func DeleteModifierOption(token, LocationId, ModifierListId, ModifierOptionId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/modifier-lists/"+ModifierListId+"/modifier-options/"+ModifierOptionId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateCategoryReqObject struct {
	// The ID of the location to create a category for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The category's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The category's name.
	Name []string `json:"name"`
}

// Creates an item category.
//
// To add or remove an item from a category, use the Update
//         Item endpoint.
//
// Required permissions:  ITEMS_WRITE
func CreateCategory(token, LocationId string, reqObj CreateCategoryReqObject) (*Category, error) {
	v := new(Category)
	err := v1Request("POST", "/v1/"+LocationId+"/categories", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's item categories.
//
// Required permissions:  ITEMS_READ
func ListCategories(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/categories", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateCategoryReqObject struct {
	// The ID of the category's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the category to edit.
	CategoryId []string `json:"category_id"`
	// The new name of the category.
	Name []string `json:"name"`
}

// Modifies the details of an existing item category.
//
// To add or remove an item from a category, use the Update
//         Item endpoint.
//
// Required permissions:  ITEMS_WRITE
func UpdateCategory(token, LocationId, CategoryId string, reqObj UpdateCategoryReqObject) (*Category, error) {
	v := new(Category)
	err := v1Request("PUT", "/v1/"+LocationId+"/categories/"+CategoryId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing item category.
//
// Required permissions:  ITEMS_WRITE
func DeleteCategory(token, LocationId, CategoryId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/categories/"+CategoryId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateDiscountReqObject struct {
	// The ID of the location to create a discount for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The discount's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The discount's name.
	Name []string `json:"name"`
	// The rate of the discount, as a string representation of a decimal number. A value of
	//             0.07 corresponds to a rate of 7%. Specify a rate of 0 if discount_type
	//             is VARIABLE_PERCENTAGE.Do not include this field for amount-based discounts.
	Rate []string `json:"rate"`
	// The amount of the discount. Specify an amount of 0 if discount_type is
	//             VARIABLE_AMOUNT.Do not include this field for rate-based discounts.
	AmountMoney []Money `json:"amount_money"`
	// Indicates whether the discount is a FIXED value or entered at the time of
	//             sale.Default value: FIXED
	DiscountType []Discount.Type `json:"discount_type"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount
	//             to a payment.Default value: false
	PinRequired []bool `json:"pin_required"`
	// The color of the discount's display label in Square Register.Default value: 9da2a6
	Color []Item.Color `json:"color"`
}

// Creates a discount.
//
// Required permissions:  ITEMS_WRITE
func CreateDiscount(token, LocationId string, reqObj CreateDiscountReqObject) (*Discount, error) {
	v := new(Discount)
	err := v1Request("POST", "/v1/"+LocationId+"/discounts", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's discounts.
//
// Required permissions:  ITEMS_READ
func ListDiscounts(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/discounts", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateDiscountReqObject struct {
	// The ID of the discount's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the discount to edit.
	DiscountId []string `json:"discount_id"`
	// The discount's name.
	Name []string `json:"name"`
	// The rate of the discount, as a string representation of a decimal number. A value of
	//             0.07 corresponds to a rate of 7%. Specify a rate of 0 if discount_type
	//             is VARIABLE_PERCENTAGE.Do not include this field for amount-based discounts.
	Rate []string `json:"rate"`
	// The amount of the discount. Specify an amount of 0 if discount_type is
	//             VARIABLE_AMOUNT.Do not include this field for rate-based discounts.
	AmountMoney []Money `json:"amount_money"`
	// Indicates whether the discount is a FIXED value or entered at the time of
	//             sale.Default value: FIXED
	DiscountType []Discount.Type `json:"discount_type"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount
	//             to a payment.Default value: false
	PinRequired []bool `json:"pin_required"`
	// The color of the discount's display label in Square Register.Default value: 9da2a6
	Color []Item.Color `json:"color"`
}

// Modifies the details of an existing discount.
//
// Required permissions:  ITEMS_WRITE
func UpdateDiscount(token, LocationId, DiscountId string, reqObj UpdateDiscountReqObject) (*Discount, error) {
	v := new(Discount)
	err := v1Request("PUT", "/v1/"+LocationId+"/discounts/"+DiscountId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing discount.
//
// Required permissions:  ITEMS_WRITE
func DeleteDiscount(token, LocationId, DiscountId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/discounts/"+DiscountId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreateFeeReqObject struct {
	// The ID of the location to create a fee for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The fee's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The fee's name.
	Name []string `json:"name"`
	// The rate of the fee, as a string representation of a decimal number. A value of
	//               0.07 corresponds to a rate of 7%.
	Rate []string `json:"rate"`
	// Whether the fee is calculated based on a payment's subtotal or total.Default value: FEE_SUBTOTAL_PHASE
	CalculationPhase []Fee.CalculationPhase `json:"calculation_phase"`
	// The type of adjustment the fee applies to a payment. Currently, this value is
	//               TAX for all fees.Default value: TAX
	AdjustmentType []Fee.AdjustmentType `json:"adjustment_type"`
	// If true, the fee applies to custom amounts entered into Square Register that are
	//             not associated with a particular item.Default value: true
	AppliesToCustomAmounts []bool `json:"applies_to_custom_amounts"`
	// If true, the fee is applied to payments. If false, it isn't.Default value: true
	Enabled []bool `json:"enabled"`
	// Whether the fee is ADDITIVE or INCLUSIVE.Default value: ADDITIVE
	InclusionType []Fee.InclusionType `json:"inclusion_type"`
}

// Creates a fee (tax).
//
// Required permissions:  ITEMS_WRITE
func CreateFee(token, LocationId string, reqObj CreateFeeReqObject) (*Fee, error) {
	v := new(Fee)
	err := v1Request("POST", "/v1/"+LocationId+"/fees", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's fees (taxes).
//
// Required permissions:  ITEMS_READ
func ListFees(token, LocationId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/v1/"+LocationId+"/fees", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateFeeReqObject struct {
	// The ID of the fee's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the fee to edit.
	FeeId []string `json:"fee_id"`
	// The fee's name.
	Name []string `json:"name"`
	// The rate of the fee, as a string representation of a decimal number. A value of
	//               0.07 corresponds to a rate of 7%.
	Rate []string `json:"rate"`
	// Whether the fee is calculated based on a payment's subtotal or total.Default value: FEE_SUBTOTAL_PHASE
	CalculationPhase []Fee.CalculationPhase `json:"calculation_phase"`
	// The type of adjustment the fee applies to a payment. Currently, this value is
	//               TAX for all fees.Default value: TAX
	AdjustmentType []Fee.AdjustmentType `json:"adjustment_type"`
	// If true, the fee applies to custom amounts entered into Square Register that are
	//             not associated with a particular item.Default value: true
	AppliesToCustomAmounts []bool `json:"applies_to_custom_amounts"`
	// If true, the fee is applied to all appropriate items. If false, the fee
	//             is not applied at all.Default value: true
	Enabled []bool `json:"enabled"`
	// Whether the fee is ADDITIVE or INCLUSIVE.Default value: ADDITIVE
	InclusionType []Fee.InclusionType `json:"inclusion_type"`
}

// Modifies the details of an existing fee (tax).
//
// Required permissions:  ITEMS_WRITE
func UpdateFee(token, LocationId, FeeId string, reqObj UpdateFeeReqObject) (*Fee, error) {
	v := new(Fee)
	err := v1Request("PUT", "/v1/"+LocationId+"/fees/"+FeeId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing fee (tax).
//
// Required permissions:  ITEMS_WRITE
func DeleteFee(token, LocationId, FeeId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/fees/"+FeeId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type ApplyFeeReqObject struct {
	// The ID of the fee's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the item to add the fee to.
	ItemId []string `json:"item_id"`
	// The ID of the fee to apply.
	FeeId []string `json:"fee_id"`
}

// Associates a fee with an item, meaning the fee is automatically applied to the item in Square
//       Register.
//
// Required permissions:  ITEMS_WRITE
func ApplyFee(token, LocationId, ItemId, FeeId string, reqObj ApplyFeeReqObject) (*Item, error) {
	v := new(Item)
	err := v1Request("PUT", "/v1/"+LocationId+"/items/"+ItemId+"/fees/"+FeeId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Removes a fee assocation from an item, meaning the fee is no longer automatically applied to
//       the item in Square Register.
//
// Required permissions:  ITEMS_WRITE
func RemoveFee(token, LocationId, ItemId, FeeId string) (*Item, error) {
	v := new(Item)
	err := v1Request("DELETE", "/v1/"+LocationId+"/items/"+ItemId+"/fees/"+FeeId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type CreatePageReqObject struct {
	// The ID of the location to create a Favorites page for.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The page's ID. Must be unique among all entity IDs ever provided on behalf of the
	//             merchant. You can never reuse an ID. This value can include alphanumeric
	//             characters, dashes (-), and underscores (_).If you don't provide this value, an ID is generated by Square.
	Id []string `json:"id"`
	// The page's name.
	Name []string `json:"name"`
	// The page's position in the list of pages. Must be an integer between 0 and
	//               4, inclusive.The endpoint returns an error if you specify a page_index that another page is
	//             already using.
	PageIndex []number `json:"page_index"`
}

// Creates a Favorites page in Square Register.
//
// A merchant can have up to five pages, each of which has a page_index between 0
//       and 4, inclusive.
//
// After you create a page, you can set the values of its cells with the Update Cell endpoint. A page doesn't appear in Square Register
//       unless at least one of its cells has an assigned value.
//
// Required permissions:  ITEMS_WRITE
func CreatePage(token, LocationId string, reqObj CreatePageReqObject) (*Page, error) {
	v := new(Page)
	err := v1Request("POST", "/v1/"+LocationId+"/pages", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists all of a location's Favorites pages in Square Register.
//
// Required permissions:  ITEMS_READ
func ListPages(token, LocationId string) ([]*Page, error) {
	v := new(Page)
	err := v1Request("GET", "/v1/"+LocationId+"/pages", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdatePageReqObject struct {
	// The ID of the Favorites page's associated location.Get a business's locations with the List Locations
	//             endpoint..
	LocationId []string `json:"location_id"`
	// The ID of the page to modify.
	PageId []string `json:"page_id"`
	// The page's name.
	Name []string `json:"name"`
	// The page's position in the merchant's list of pages. Must be an integer between
	//               0 and 4, inclusive.The page's index is not updated if another page already exists at the specified
	//             index.
	PageIndex []number `json:"page_index"`
}

// Modifies the details of a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
func UpdatePage(token, LocationId, PageId string, reqObj UpdatePageReqObject) (*Page, error) {
	v := new(Page)
	err := v1Request("PUT", "/v1/"+LocationId+"/pages/"+PageId, token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes an existing Favorites page and all of its cells.
//
// Required permissions:  ITEMS_WRITE
func DeletePage(token, LocationId, PageId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/pages/"+PageId, token, nil, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// For POST and PUT endpoints, you provide request parameters as JSON in your request's body.
type UpdateCellReqObject struct {
	// The ID of the Favorites page's associated location.Get a business's locations with the List Locations
	//             endpoint.
	LocationId []string `json:"location_id"`
	// The ID of the page the cell belongs to.
	PageId []string `json:"page_id"`
	// The row of the cell to update. Always an integer between 0 and 4,
	//             inclusive. Row 0 is the top row.
	Row []number `json:"row"`
	// The column of the cell to update. Always an integer between 0 and 4,
	//             inclusive. Column 0 is the leftmost row.
	Column []number `json:"column"`
	// The type of entity represented in the cell (ITEM, DISCOUNT,
	//               CATEGORY, or PLACEHOLDER).
	ObjectType []PageCell.Type `json:"object_type"`
	// The unique identifier of the entity to represent in the cell. Do not include if the
	//             cell's object_type is PLACEHOLDER.
	ObjectId []string `json:"object_id"`
	// For a cell with an object_type of PLACEHOLDER, indicates the cell's
	//             behavior.
	PlaceholderType []PageCell.PlaceholderType `json:"placeholder_type"`
}

// Modifies a cell of a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
func UpdateCell(token, LocationId, PageId string, reqObj UpdateCellReqObject) (*PageCell, error) {
	v := new(PageCell)
	err := v1Request("PUT", "/v1/"+LocationId+"/pages/"+PageId+"/cells", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Deletes a cell from a Favorites page in Square Register.
//
// Required permissions:  ITEMS_WRITE
func DeleteCell(token, LocationId, PageId string) error {
	v := new()
	err := v1Request("DELETE", "/v1/"+LocationId+"/pages/"+PageId+"/cells", token, nil, nil)
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
//       endpoint's response is an array that contains the response for each batched request.
//
// You don't need to provide an access token in the header of your request to the Submit Batch
//       endpoint. Instead, you provide an access_token parameter for each request included in
//       the batch.
//
// Note the following when using the Submit Batch endpoint:
func SubmitBatch(token string, reqObj SubmitBatchReqObject) ([]*BatchResponse, error) {
	v := new(BatchResponse)
	err := v1Request("POST", "/v1/batch", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists which types of events trigger webhook notifications for a particular location.
func ListWebhooks(token, LocationId string) ([]*WebhookEventType, error) {
	v := new(WebhookEventType)
	err := v1Request("GET", "/v1/"+LocationId+"/webhooks", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type UpdateWebhooksReqObject struct {
	// The ID of the location to update webhook preferences for.
	LocationId []string `json:"location_id"`
}

// Sets which types of events trigger webhook notifications for a location.
//
// Simply provide a JSON array of the event types you want notifications for in your request
//       body (see Example Requests below).
func UpdateWebhooks(token, LocationId string, reqObj UpdateWebhooksReqObject) ([]*WebhookEventType, error) {
	v := new(WebhookEventType)
	err := v1Request("PUT", "/v1/"+LocationId+"/webhooks", token, reqObj, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Lists subscriptions that have been created for an application. You can look up subscription
//       information for a single merchant by providing the merchant_id parameter to this
//       endpoint.
//
// Subscription objects returned by this endpoint do not include the fees
//       field, which lists the subscription's payment history. To get a particular subscription's
//       payment history, use the Retrieve Subscription
//       endpoint.
//
// Important: The Authorization header you provide to this endpoint must have the
//       following format:
func ListSubscriptions(token, ClientId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/oauth2/clients/"+ClientId+"/subscriptions", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides comprehensive information for a single subscription, including its payment
//       history.
//
// Important: The Authorization header you provide to this endpoint must have the
//       following format:
func RetrieveSubscription(token, ClientId, SubscriptionId string) (*Subscription, error) {
	v := new(Subscription)
	err := v1Request("GET", "/oauth2/clients/"+ClientId+"/subscriptions/"+SubscriptionId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Provides information for all of an application's subscription plans.
//
// Important: The Authorization header you provide to this endpoint must have the
//       following format:
func ListSubscriptionPlans(token, ClientId string) (*paginate, error) {
	v := new(paginate)
	err := v1Request("GET", "/oauth2/clients/"+ClientId+"/plans", token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func RetrieveSubscriptionPlan(token, ClientId, PlanId string) (*SubscriptionPlan, error) {
	v := new(SubscriptionPlan)
	err := v1Request("GET", "/oauth2/clients/"+ClientId+"/plans/"+PlanId, token, nil, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
