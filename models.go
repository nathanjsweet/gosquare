package gosquare

// Represents a merchant's bank account.
type BankAccount struct {
	// The bank account's Square-issued ID.
	ID string `json:"id"`
	// The Square-issued ID of the merchant associated with the bank account.
	MerchantID string `json:"merchant_id"`
	// The name of the bank that manages the account.
	BankName string `json:"bank_name"`
	// The name associated with the bank account.
	Name string `json:"name"`
	// The bank account's type (for example, savings or checking).
	Type string `json:"type"`
	// The bank account's routing number.
	RoutingNumber string `json:"routing_number"`
	// The last few digits of the bank account number.
	AccountNumberSuffix string `json:"account_numberSuffix"`
	// The currency code of the currency associated with the bank account, in ISO 4217
	// format. For example, the currency code for US dollars is USD.
	CurrencyCode string `json:"currency_code"`
}

// Represents a single request included in a call to the Submit Batch endpoint.
type BatchRequest struct {
	// The HTTP method of the request (DELETE, GET, POST, or PUT).
	Method string `json:"method"`
	// The path of the request's endpoint, relative to
	// https://connect.squareup.com.For example, this value is
	// /v1/MERCHANT_ID/payments for the List Payments
	// endpoint (with the proper merchant ID).For GET and DELETE requests,
	// include any request parameters in a query string appended to the path (for example,
	// /v1/MERCHANT_ID/payments?order=DESC).
	RelativePath string `json:"relative_path"`
	// The access token to use for the request. This can be a personal access token or an
	// access token generated with the OAuth API.
	AccessToken string `json:"access_token"`
	// The body of the request, if any. Include parameters for POST and PUT
	// requests here.
	Body interface{} `json:"body"`
	// An optional identifier for the request, returned in the request's corresponding
	// BatchResponse.
	RequestID string `json:"request_id"`
}

// Represents the response for a request included in a call to the Submit Batch endpoint.
type BatchResponse struct {
	// The response's HTTP status code.
	StatusCode int `json:"status_code"`
	// Contains any important headers for the response, indexed by header name. For example,
	// if the response includes a pagination header, the
	// header's value is available from headers["Link"].
	Headers interface{} `json:"headers"`
	// The body of the response, if any.
	Body interface{} `json:"body"`
	// The value you provided for request_id in the corresponding BatchRequest, if any.
	RequestID string `json:"request_id"`
}

// Represents geographic coordinates.
type Coordinates struct {
	// The latitude coordinate, in degrees.
	Latitude float64 `json:"latitude"`
	// The longitude coordinate, in degrees.
	Longitude float64 `json:"longitude"`
}

// Represents an event (such as a payment or refund) that involved opening the cash drawer
// during a cash drawer shift.
type CashDrawerEvent struct {
	// The event's unique ID.
	ID string `json:"id"`
	// The ID of the employee that created the event.
	EmployeeID string `json:"employee_id"`
	// The type of event that occurred, such as CASH_TENDER_PAYMENT or
	// CASH_TENDER_REFUND.
	EventType string `json:"event_type"`
	// The amount of money that was added to or removed from the cash drawer because of the
	// event. This value can be positive (for added money) or negative (for removed money).
	// event_money Money `json:"event_money"`
	// The time when the event occurred, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// An optional description of the event, entered by the employee that created it.
	Description string `json:"description"`
}

// Represents all cash drawer activity that takes place during a single cash drawer shift.
type CashDrawerShift struct {
	// The shift's unique ID.
	ID string `json:"id"`
	// The shift's current state (OPEN, ENDED, or CLOSED).
	CashDrawerState string `json:"cash_drawer_state"`
	// The time when the shift began, in ISO 8601 format.
	OpenedAt string `json:"opened_at"`
	// The time when the shift ended, in ISO 8601 format.
	EndedAt string `json:"ended_at"`
	// The time when the shift was closed, in ISO 8601 format.
	ClosedAt string `json:"closed_at"`
	// The IDs of all employees that were logged into Square Register at some point during the
	// cash drawer shift.
	EmployeeIDs []string `json:"employee_ids"`
	// The ID of the employee that started the cash drawer shift.
	OpeningEmployeeID string `json:"opening_employee_id"`
	// The ID of the employee that ended the cash drawer shift.
	EndingEmployeeID string `json:"ending_employee_id"`
	// The ID of the employee that closed the cash drawer shift by auditing the cash drawer's
	// contents.
	ClosingEmployeeID string `json:"closing_employee_id"`
	// An optional description of the shift, entered by the employee that ended it.
	Description string `json:"description"`
	// The amount of money in the cash drawer at the start of the shift.
	StartingCashMoney Money `json:"starting_cash_money"`
	// The amount of money added to the cash drawer from cash payments.
	CashPaymentMoney Money `json:"cash_payment_money"`
	// The amount of money removed from the cash drawer from cash refunds. This value is
	// always negative or zero.
	CashRefundsMoney Money `json:"cash_refunds_money"`
	// The amount of money added to the cash drawer for reasons other than cash payments.
	CashPaidInMoney Money `json:"cash_paid_in_money"`
	// The amount of money removed from the cash drawer for reasons other than cash
	// refunds.
	CashPaidOutMoney Money `json:"cash_paid_out_money"`
	// The amount of money that should be in the cash drawer at the end of the shift, based on
	// the shift's other money amounts.
	ExpectedCashMoney Money `json:"expected_cash_money"`
	// The amount of money found in the cash drawer at the end of the shift by an auditing
	// employee.
	ClosedCashMoney Money `json:"closed_cash_money"`
	// The device running Square Register that was connected to the cash drawer.
	Device Device `json:"device"`
	// All of the events (payments, refunds, and so on) that involved the cash drawer during
	// the shift.
	Events []CashDrawerEvent `json:"events"`
}

// Represents an item category.
type Category struct {
	// The category's unique ID.
	ID string `json:"id"`
	// The category's name.
	Name string `json:"name"`
}

// Represents a device running Square Register.
type Device struct {
	// The device's merchant-specified name.
	Name string `json:"name"`
	// The device's Square-issued ID.
	ID string `json:"id"`
}

// Represents a discount that can be applied to a payment. A discount can be either a
// percentage or a flat amount. You can determine a particular discount's type by checking
// which of rate or amount_money is included in the object.
type Discount struct {
	// The discount's unique ID.
	ID string `json:"id"`
	// The discount's name.
	Name string `json:"name"`
	// The rate of the discount, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%. This rate is 0 if discount_type
	// is VARIABLE_PERCENTAGE.This field is not included for amount-based discounts.
	Rate string `json:"rate"`
	// The amount of the discount. This amount is 0 if discount_type is
	// VARIABLE_AMOUNT.This field is not included for rate-based discounts.
	AmountMoney Money `json:"amount_money"`
	// Indicates whether the discount is a FIXED value or entered at the time of
	// sale.
	DiscountType string `json:"discount_type"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount
	// to a payment.
	PinRequired bool `json:"pin_required"`
	// The color of the discount's display label in Square Register, if not the default color.
	// The default color is 9da2a6.
	Color string `json:"color"`
}

// Represents one of a business's employees.
type Employee struct {
	// The employee's unique ID.
	ID string `json:"id"`
	// The employee's first name.
	FirstName string `json:"first_name"`
	// The employee's last name.
	LastName string `json:"last_name"`
	// The ids of the employee's associated roles. Currently, you can specify only one
	// or zero roles per employee.
	RoleIDs []string `json:"role_ids"`
	// The IDs of the locations the employee is allowed to clock in at.
	AuthorizedLocationIDs []string `json:"authorized_location_ids"`
	// The employee's email address.You cannot edit this value with the Connect API.
	// You can only set its initial value
	// when creating an employee with the Create Employee endpoint.
	Email string `json:"email"`
	// Whether the employee is ACTIVE or INACTIVE. Inactive employees cannot
	// sign in to Square Register.Merchants update this field from the Square Dashboard.
	// You cannot modify it with the Connect API.
	Status string `json:"status"`
	// An ID the merchant can set to associate the employee with an entity in another
	// system.You cannot set this value with the Connect API.
	ExternalID string `json:"external_id"`
	// The time when the employee entity was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The time when the employee entity was most recently updated, in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
}

// Represents a role that can be assigned to one or more employees. An employee's role
// indicates which permissions they have.
type EmployeeRole struct {
	// The role's unique ID.
	ID string `json:"id"`
	// The role's merchant-defined name.
	Name string `json:"name"`
	// The permissions that the role has been granted.
	Permissions []string `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the
	// values indicated in permissions.
	IsOwner bool `json:"is_owner"`
	// The time when the role was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The time when the role was most recently updated, in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
}

// Represents a tax or other fee that can be applied to a payment.
type Fee struct {
	// The fee's unique ID.
	ID string `json:"id"`
	// The fee's name.
	Name string `json:"name"`
	// The rate of the fee, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate"`
	// Forthcoming.
	CalculationPhase string `json:"calculation_phase"`
	// The type of adjustment the fee applies to a payment. Currently, this value is
	// TAX for all fees.
	AdjustmentType string `json:"adjustment_type"`
	// If true, the fee applies to custom amounts entered into Square Register that are
	// not associated with a particular item.
	AppliesToCustomAmounts bool `json:"applies_to_custom_amounts"`
	// If true, the fee is applied to all appropriate items. If false, the fee
	// is not applied at all.
	Enabled bool `json:"enabled"`
	// Whether the fee is ADDITIVE or INCLUSIVE.
	InclusionType string `json:"inclusion_type"`
	// In countries with multiple classifications for sales taxes, indicates which
	// classification the fee falls under. Currently relevant only to Canadian merchants.
	Type string `json:"type"`
}

// A generic representation of a physical address.
type GlobalAddress struct {
	// The first line of the address.Fields that start with address_line provide
	// the address's most specific details,
	// like street number, street name, and building name. They do not provide less
	// specific details like city, state/province, or country (these details are provided in
	// other fields).
	AddressLine1 string `json:"address_line_1"`
	// The second line of the address, if any.
	AddressLine2 string `json:"address_line_2"`
	// The third line of the address, if any.
	AddressLine3 string `json:"address_line_3"`
	// The fourth line of the address, if any.
	AddressLine4 string `json:"address_line_4"`
	// The fifth line of the address, if any.
	AddressLine5 string `json:"address_line_5"`
	// The city or town of the address.
	Locality string `json:"locality"`
	// A civil region within the address's locality, if any.
	Sublocality string `json:"sublocality"`
	// A civil region within the address's sublocality, if any.
	Sublocality1 string `json:"sublocality_1"`
	// A civil region within the address's sublocality_1, if any.
	Sublocality2 string `json:"sublocality_2"`
	// A civil region within the address's sublocality_2, if any.
	Sublocality3 string `json:"sublocality_3"`
	// A civil region within the address's sublocality_3, if any.
	Sublocality4 string `json:"sublocality_4"`
	// A civil region within the address's sublocality_4, if any.
	Sublocality5 string `json:"sublocality_5"`
	// A civil entity within the address's country. In the United States, this is the
	// state.
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1"`
	// A civil entity within the address's administrative_district_level_1, if any. In
	// the United States, this is the county.
	AdministrativeDistrictLevel2 string `json:"administrative_district_level_2"`
	// A civil entity within the address's administrative_district_level_2, if any.
	AdministrativeDistrictLevel3 string `json:"administrative_district_level_3"`
	// The address's postal code.
	PostalCode string `json:"postal_code"`
	// The address's country, in ISO 3166-1-alpha-2 format.
	CountryCode string `json:"country_code"`
	// The coordinates of the address.
	AddressCoordinates Coordinates `json:"address_coordinates"`
}

// Represents inventory information for one of a merchant's item variations.
type InventoryEntry struct {
	// The variation that the entry corresponds to.
	VariationID string `json:"variation_id"`
	// The current available quantity of the item variation.
	QuantityOnHand int `json:"quantity_on_hand"`
}

// Represents a merchant's item.
type Item struct {
	// The item's unique ID.
	ID string `json:"id"`
	// The item's name.
	Name string `json:"name"`
	// The item's description, if any.
	Description string `json:"description"`
	// The item's type. This value is NORMAL for almost all items.
	Type string `json:"type"`
	// The text of the item's display label in Square Register. This value is present only if
	// an abbreviation other than the default has been set.
	Abbreviation string `json:"abbreviation"`
	// The color of the item's display label in Square Register, if not the default color.
	// The default color is 9da2a6.
	Color string `json:"color"`
	// Indicates whether the item is viewable in the merchant's online store (PUBLIC)
	// or PRIVATE.
	Visibility string `json:"visibility"`
	// If true, the item is available for purchase from the merchant's online
	// store.
	AvailableOnline bool `json:"available_online"`
	// The item's master image, if any.
	MasterImage ItemImage `json:"master_image"`
	// The category the item belongs to, if any.
	Category Category `json:"category"`
	// The item's variations.
	Variations []ItemVariation `json:"variations"`
	// The modifier lists that apply to the item, if any.
	ModifierLists []ModifierList `json:"modifier_lists"`
	// The fees that apply to the item, if any.
	Fees []Fee `json:"fees"`
	// Deprecated. This field is not used.
	Taxable bool `json:"taxable"`
}

// Represents an image of an item.
type ItemImage struct {
	// The image's unique ID.
	ID string `json:"id"`
	// The image's publicly accessible URL.
	Url string `json:"url"`
}

// Represents a variation of an Item. Every item has
// at least one variation.
type ItemVariation struct {
	// The item variation's unique ID.
	ID string `json:"id"`
	// The item variation's name.
	Name string `json:"name"`
	// The ID of the variation's associated item.
	ItemID string `json:"item_id"`
	// Indicates the variation's list position when displayed in Square Register and the
	// merchant dashboard. If more than one variation for the same item has the same
	// ordinal value, those variations are displayed in alphabetical order.
	// An item's variation with the lowest ordinal value is displayed first.
	Ordinal int `json:"ordinal"`
	// Indicates whether the item variation's price is fixed or determined at the time of
	// sale.
	PricingType string `json:"pricing_type"`
	// The item variation's price, if any.
	PriceMoney Money `json:"price_money"`
	// The item variation's SKU, if any.
	Sku string `json:"sku"`
	// If true, inventory tracking is active for the variation.
	TrackInventory bool `json:"track_inventory"`
	// Indicates whether the item variation displays an alert when its inventory quantity is
	// less than or equal to its inventory_alert_threshold.
	InventoryAlertType string `json:"inventory_alert_type"`
	// If the inventory quantity for the variation is less than or equal to this value and
	// inventory_alert_type is LOW_QUANTITY, the variation displays an alert in
	// the merchant dashboard.This value is always an integer.
	InventoryAlertThreshold int `json:"inventory_alert_threshold"`
	// Arbitrary metadata associated with the variation. Cannot exceed 255 characters.
	UserData string `json:"user_data"`
}

// Represents a Square merchant account.
type Merchant struct {
	// The merchant account's unique identifier.
	ID string `json:"id"`
	// The name associated with the merchant account.
	Name string `json:"name"`
	// The email address associated with the merchant account.
	Email string `json:"email"`
	// Indicates whether the merchant account corresponds to a single-location account
	// (LOCATION) or a business account (BUSINESS). This value is almost always
	// LOCATION. See Multi-Location
	// Overview for more information.
	AccountType string `json:"account_type"`
	// Capabilities that are enabled for the merchant's Square account. Capabilities that are
	// not listed in this array are not enabled for the account. Currently there is only one
	// capability, CREDIT_CARD_PROCESSING.
	AccountCapabilities []string `json:"account_capabilities"`
	// The country associated with the merchant account, in ISO 3166-1-alpha-2
	// format.
	CountryCode string `json:"country_code"`
	// The language associated with the merchant account, in BCP 47 format.
	LanguageCode string `json:"language_code"`
	// The currency associated with the merchant account, in ISO 4217
	// format. For example, the currency code for US dollars is USD.
	CurrencyCode string `json:"currency_code"`
	// The name of the merchant's business.
	BusinessName string `json:"business_name"`
	// The address of the merchant's business.
	BusinessAddress GlobalAddress `json:"business_address"`
	// The phone number of the merchant's business.
	BusinessPhone PhoneNumber `json:"business_phone"`
	// The type of business operated by the merchant.
	BusinessType string `json:"business_type"`
	// The merchant's shipping address.
	ShippingAddress GlobalAddress `json:"shipping_address"`
	// Additional information for a single-location account specified by its associated
	// business account, if it has one.Never included in Merchant objects with the account_type
	// BUSINESS.
	LocationDetails MerchantLocationDetails `json:"location_details"`
	// The URL of the merchant's online store.
	MarketUrl string `json:"market_url"`
}

// Represents additional details for a single-location account as specified by its parent
// business.
type MerchantLocationDetails struct {
	// The nickname assigned to the single-location account by the parent business. This value
	// appears in the parent business's multi-location dashboard.
	Nickname string `json:"nickname"`
}

// Represents an item modifier list.
type ModifierList struct {
	// The modifier list's unique ID.
	ID string `json:"id"`
	// The modifier list's name.
	Name string `json:"name"`
	// Indicates whether MULTIPLE options or a SINGLE option from the modifier
	// list can be applied to a single item.
	SelectionType string `json:"selection_type"`
	// The options included in the modifier list.
	ModifierOptions []string `json:"modifier_options"`
}

// Represents an item modifier option.
type ModifierOption struct {
	// The modifier option's unique ID.
	ID string `json:"id"`
	// The modifier option's name.
	Name string `json:"name"`
	// The modifier option's price.
	PriceMoney Money `json:"price_money"`
	// If true, the modifier option is the default option in a modifier list for which
	// selection_type is SINGLE.
	OnByDefault bool `json:"on_by_default"`
	// Indicates the modifier option's list position when displayed in Square Register and the
	// merchant dashboard. If more than one modifier option in the same modifier list has the
	//same ordinal value, those options are displayed in alphabetical order.
	// A modifier list's option with the lowest ordinal value is displayed first.
	Ordinal int `json:"ordinal"`
	// The ID of the modifier list the option belongs to.
	ModifierListID string `json:"modifier_list_id"`
}

// Represents an amount of money. When you provide this object in a request,
// currency_code must match the currency associated with the merchant's Square account.
type Money struct {
	// The amount of money, in the smallest unit of the applicable currency. For US dollars,
	// this value is in cents.This value is always an integer.
	Amount int `json:"amount"`
	// The type of currency involved in the current payment, in ISO 4217 format. For example, the
	// currency code for US dollars is USD.
	CurrencyCode string `json:"currency_code"`
}

// Represents an order from a merchant's online store.
type Order struct {
	// The order's unique identifier.
	ID string `json:"id"`
	// The order's current state, such as OPEN or COMPLETED.
	State string `json:"state"`
	// The email address of the order's buyer.
	BuyerEmail string `json:"buyer_email"`
	// The name of the order's buyer.
	RecipientName string `json:"recipient_name"`
	// The phone number to use for the order's delivery.
	RecipientPhoneNumber string `json:"recipient_phone_number"`
	// The address to ship the order to.
	ShippingAddress GlobalAddress `json:"shipping_address"`
	// The amount of all items purchased in the order, before taxes and shipping.
	SubtotalMoney Money `json:"subtotal_money"`
	// The shipping cost for the order.
	TotalShippingMoney Money `json:"total_shipping_money"`
	// The total of all taxes applied to the order.
	TotalTaxMoney Money `json:"total_tax_money"`
	// The total cost of the order.
	TotalPriceMoney Money `json:"total_price_money"`
	// The total of all discounts applied to the order.
	TotalDiscountMoney Money `json:"total_discount_money"`
	// The time when the order was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The time when the order was last modified, in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	// The time when the order expires if no action is taken, in ISO 8601 format.
	ExpiresAt string `json:"expires_at"`
	// The unique identifier of the payment associated with the order.
	PaymentID string `json:"payment_id"`
	// A note provided by the buyer when the order was created, if any.
	BuyerNote string `json:"buyer_note"`
	// A note provided by the merchant when the order's state was set to COMPLETED, if any.
	CompletedNote string `json:"completed_note"`
	// A note provided by the merchant when the order's state was set to REFUNDED, if any.
	RefundedNote string `json:"refunded_note"`
	// A note provided by the merchant when the order's state was set to CANCELED, if any.
	CanceledNote string `json:"canceled_note"`
	// The tender used to pay for the order.
	Tender Tender `json:"tender"`
	// The history of actions associated with the order.
	OrderHistory []OrderHistoryEntry `json:"order_history"`
	// The promo code provided by the buyer, if any.
	PromoCode string `json:"promo_code"`
	// For Bitcoin transactions, the address that the buyer sent Bitcoin to.
	BtcReceiveAddress string `json:"btc_receive_address"`
	// For Bitcoin transactions, the price of the buyer's order in satoshi (100 million
	// satoshi equals 1 BTC).
	BtcPriceSatoshi int `json:"btc_price_satoshi"`
}

// Represents a prior action performed on an online store order.
type OrderHistoryEntry struct {
	// The type of action performed on the order.
	Action string `json:"action"`
	// The time when the action was performed, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
}

// Represents a Favorites page in the iPad version of Square Register.
type Page struct {
	// The page's unique identifier.
	ID string `json:"id"`
	// The page's name, if any.
	Name string `json:"name"`
	// The page's position in the merchant's list of pages. Always an integer between 0 and 4,
	// inclusive.
	PageIndex int `json:"page_index"`
	// The cells included on the page.
	Cells []PageCell `json:"cells"`
}

// Represents a cell of a Page.
type PageCell struct {
	// The unique identifier of the page the cell is included on.
	PageID string `json:"page_id"`
	// The row of the cell. Always an integer between 0 and 4, inclusive.
	Row int `json:"row"`
	// The column of the cell. Always an integer between 0 and 4, inclusive.
	Column int `json:"column"`
	// The type of entity represented in the cell (ITEM, DISCOUNT,
	// CATEGORY, or PLACEHOLDER).
	ObjectType string `json:"object_type"`
	// The unique identifier of the entity represented in the cell. Not present for cells with
	// an object_type of PLACEHOLDER.
	ObjectID string `json:"object_id"`
	// For a cell with an object_type of PLACEHOLDER, this value indicates the cell's
	// special behavior.
	PlaceholderType string `json:"placeholder_type"`
}

// Represents a payment taken by a Square merchant.
type Payment struct {
	// The payment's unique identifier.
	ID string `json:"id"`
	// The unique identifier of the merchant that took the payment.
	MerchantID string `json:"merchant_id"`
	// The time when the payment was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The unique identifier of the Square account that took the payment.
	// This value can differ from merchant_id if the merchant has mobile staff.
	CreatorID string `json:"creator_id"`
	// The device that took the payment.
	Device Device `json:"device"`
	// The URL of the payment's detail page in the merchant dashboard.
	// The merchant must be signed in to the merchant dashboard to view this page.
	PaymentUrl string `json:"payment_url"`
	// The URL of the receipt for the payment.Note that for split tender payments, this URL
	// corresponds to the receipt for the first tender listed in the payment's
	// tender field. Each Tender object has its own receipt_url field you can use
	// to get the other receipts associated with a split tender payment.
	ReceiptUrl string `json:"receipt_url"`
	// The sum of all inclusive taxes associated with the payment.
	InclusiveTaxMoney Money `json:"inclusive_tax_money"`
	// The sum of all additive taxes associated with the payment.
	AdditiveTaxMoney Money `json:"additive_tax_money"`
	// The total of all taxes applied to the payment.
	// This is always the sum of inclusive_tax_money and additive_tax_money.
	TaxMoney Money `json:"tax_money"`
	// The total of all tips applied to the payment.
	TipMoney Money `json:"tip_money"`
	// The total of all discounts applied to the payment.This value is always 0 or negative.
	DiscountMoney Money `json:"discount_money"`
	// The total amount of money collected from the buyer for the payment.
	TotalCollectedMoney Money `json:"total_collected_money"`
	// The total of all processing fees collected by Square for the payment.
	// This value is always 0 or negative.
	ProcessingFeeMoney Money `json:"processing_fee_money"`
	// The amount to be deposited into the merchant's bank account for the payment.
	// This is always the sum of total_collected_money and processing_fee_money
	// (note that processing_fee_money is always negative or 0).
	NetTotalMoney Money `json:"net_total_money"`
	// The total of all refunds applied to the payment.
	RefundedMoney Money `json:"refunded_money"`
	// All of the inclusive taxes associated with the payment.
	InclusiveTax []PaymentTax `json:"inclusive_tax"`
	// All of the additive taxes associated with the payment.
	AdditiveTax []PaymentTax `json:"additive_tax"`
	// The form(s) of tender provided by the buyer for the payment.
	Tender []Tender `json:"tender"`
	// All of the refunds applied to the payment.
	Refunds []Refund `json:"refunds"`
	// The items purchased in the payment.
	Itemizations []PaymentItemization `json:"itemizations"`
}

// Represents a discount applied to an itemization in a payment.
type PaymentDiscount struct {
	// The discount's name.
	Name string `json:"name"`
	// The amount of money that this discount adds to the payment (note that this value is
	// always negative or zero).
	AppliedMoney Money `json:"applied_money"`
	// The ID of the applied discount, if available. Discounts applied in older versions of
	// Square Register might not have an ID.
	DiscountID string `json:"discount_id"`
}

// Represents details of an item purchased in a payment.
type PaymentItemDetail struct {
	// The name of the item's merchant-defined category, if any.
	CategoryName string `json:"category_name"`
	// The item's merchant-defined SKU, if any.
	Sku string `json:"sku"`
	// The unique ID of the item purchased, if any.
	ItemID string `json:"item_id"`
	// The unique ID of the item variation purchased, if any.
	ItemVariationID string `json:"item_variation_id"`
}

// Represents an item, custom monetary amount,
// or other entity purchased as part of a payment.
type PaymentItemization struct {
	// The item's name.
	Name string `json:"name"`
	// The quantity of the item purchased. This can be a decimal value.
	Quantity float64 `json:"quantity"`
	// The type of purchase that the itemization represents, such as an ITEM or
	// CUSTOM_AMOUNT.
	ItemizationType string `json:"itemization_type"`
	// Details of the item, including its unique identifier and the identifier of the item
	// variation purchased.
	ItemDetail PaymentItemDetail `json:"item_detail"`
	// Notes entered by the merchant about the item at the time of payment, if any.
	Notes string `json:"notes"`
	// The name of the item variation purchased, if any.
	ItemVariationName string `json:"item_variation_name"`
	// The total cost of the item, including all taxes and discounts.
	TotalMoney Money `json:"total_money"`
	// The cost of a single unit of this item.
	SingleQuantityMoney Money `json:"single_quantity_money"`
	// The total cost of the itemization and its modifiers, not including taxes or
	// discounts.
	GrossSalesMoney Money `json:"gross_sales_money"`
	// The total of all discounts applied to the itemization. This value is always negative or
	// zero.
	DiscountMoney Money `json:"discount_money"`
	// The sum of gross_sales_money and discount_money.
	NetSalesMoney Money `json:"net_sales_money"`
	// All taxes applied to this itemization.
	Taxes []PaymentTax `json:"taxes"`
	// All discounts applied to this itemization.
	Discounts []PaymentDiscount `json:"discounts"`
	// All modifier options applied to this itemization.
	Modifiers []PaymentModifier `json:"modifiers"`
}

// Represents a modifier option applied to an itemization in a payment.
type PaymentModifier struct {
	// The modifier option's name.
	Name string `json:"name"`
	// The amount of money that this modifier option adds to the payment.
	AppliedMoney Money `json:"applied_money"`
	// The ID of the applied modifier option, if available. Modifier options applied in older
	// versions of Square Register might not have an ID.
	ModifierOptionID string `json:"modifier_option_id"`
}

// Represents a single tax applied to a payment.
type PaymentTax struct {
	// The merchant-defined name of the tax.
	Name string `json:"name"`
	// The amount of money that this tax adds to the payment.
	AppliedMoney Money `json:"applied_money"`
	// The rate of the tax, as a string representation of a decimal number. A value of
	// 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate"`
	// Whether the tax is an ADDITIVE tax or an INCLUSIVE tax.
	InclusionType string `json:"inclusion_type"`
	// The ID of the tax, if available. Taxes applied in older versions of Square Register
	// might not have an ID.
	FeeID string `json:"fee_id"`
}

// Represents a phone number.
type PhoneNumber struct {
	// The phone number's international calling code.
	// For US phone numbers, this value is +1.
	CallingCode string `json:"calling_code"`
	// The phone number.
	Number string `json:"number"`
}

// Represents a refund initiated by a Square merchant.
type Refund struct {
	// The type of refund (FULL or PARTIAL).
	Type string `json:"type"`
	// The merchant-specified reason for the refund.
	Reason string `json:"reason"`
	// The amount of money refunded. This amount is always negative.
	RefundedMoney Money `json:"refunded_money"`
	// The time when the merchant initiated the refund for Square to process, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The time when Square processed the refund on behalf of the merchant, in ISO 8601 format.
	ProcessedAt string `json:"processed_at"`
	// The Square-issued ID of the payment the refund is applied to.
	PaymentID string `json:"payment_id"`
}

// Represents a deposit or withdrawal made by Square to a merchant's bank account.
type Settlement struct {
	// The settlement's unique identifier.
	ID string `json:"id"`
	// The settlement's current status.
	Status string `json:"status"`
	// The time when the settlement was submitted for deposit or withdrawal, in ISO 8601 format.
	InitiatedAt string `json:"initiated_at"`
	// The Square-issued unique identifier for the bank account associated with the
	// settlement.
	BankAccountID string `json:"bank_account_id"`
	// The amount of money involved in the settlement. A positive amount indicates a deposit,
	// and a negative amount indicates a withdrawal. This amount is never zero.
	TotalMoney Money `json:"total_money"`
	// The entries included in this settlement.
	Entries []SettlementEntry `json:"entries"`
}

// Represents a single entry in a Settlement.
type SettlementEntry struct {
	// The type of activity this entry represents.
	Type string `json:"type"`
	// The payment associated with the settlement entry, if any.
	PaymentID string `json:"payment_id"`
	// The total amount of money this entry contributes to the total settlement amount.
	AmountMoney Money `json:"amount_money"`
	// The amount of all Square fees associated with this settlement entry. This value is
	// always negative or zero.This amount has already been applied to amount_money.
	FeeMoney Money `json:"fee_money"`
}

// Represents a merchant's subscription to an application.
type Subscription struct {
	// The subscription's unique ID.
	ID string `json:"id"`
	// The ID of the merchant with the subscription.
	MerchantID string `json:"merchant_id"`
	// The ID of the SubscriptionPlan the subscription belongs to.
	PlanID string `json:"plan_id"`
	// The subscription's status, such as active or canceled.
	Status string `json:"status"`
	// The method of payment used to pay the subscription's monthly fee.
	PaymentMethod string `json:"payment_method"`
	// The subscription's base monthly fee.
	FeeBaseMoney Money `json:"fee_base_money"`
	// The date when the subscription most recently became active, in YYYY-MM-DD format.
	ServiceStartDate string `json:"service_start_date"`
	// The history of subscription fees paid or pending for this subscription, in reverse
	// chronological order (newest first).
	Fees []SubscriptionFee `json:"fees"`
}

// Represents a single fee charged to a merchant for a Subscription.
type SubscriptionFee struct {
	// The date when the subscription fee was charged, in YYYY-MM-DD format.
	FeeDate string `json:"fee_date"`
	// The payment status of the subscription fee, such as PENDING or PAID.
	FeeStatus string `json:"fee_status"`
	// The subscription fee's base amount.
	FeeBaseMoney Money `json:"fee_base_money"`
	// The total of all taxes applied to the subscription fee.
	FeeTaxMoney Money `json:"fee_tax_money"`
	// The subscription fee's total amount.
	// This is always the sum of fee_base_money and fee_tax_money.
	FeeTotalMoney Money `json:"fee_total_money"`
}

// Represents an application subscription plan.
type SubscriptionPlan struct {
	// The plan's unique ID.
	ID string `json:"id"`
	// The plan's name.
	Name string `json:"name"`
	// The country the plan applies to, in ISO 3166-1-alpha-2 format.
	CountryCode string `json:"country_code"`
	// The plan's base monthly fee.
	FeeBaseMoney Money `json:"fee_base_money"`
}

// Represents a form and amount of tender provided for a payment. Multiple forms of tender can
// be provided for a single payment.
type Tender struct {
	// The tender's unique ID.
	ID string `json:"id"`
	// The type of tender.
	Type string `json:"type"`
	// A human-readable description of the tender.
	Name string `json:"name"`
	// The ID of the employee that processed the tender.
	// This field is included only if the associated merchant had employee
	// management features enabled at the time the tender was processed.
	EmployeeID string `json:"employee_id"`
	// The URL of the receipt for the tender.
	ReceiptUrl string `json:"receipt_url"`
	// The brand of credit card provided.Only present if the tender's type is CREDIT_CARD.
	CardBrand string `json:"card_brand"`
	// The last four digits of the provided credit card's account number.
	// Only present if the tender's type is CREDIT_CARD.
	PanSuffix string `json:"pan_suffix"`
	// The method with which the tender was entered.
	EntryMethod string `json:"entry_method"`
	// Notes entered by the merchant about the tender at the time of payment, if any.
	// Typically only present for tender with the typeOTHER.
	PaymentNote string `json:"payment_note"`
	// The total amount of money provided in this form of tender.
	TotalMoney Money `json:"total_money"`
	// The amount of total_money applied to the payment.
	TenderedMoney Money `json:"tendered_money"`
	// The amount of total_money returned to the buyer as change.
	ChangeBackMoney Money `json:"change_back_money"`
	// The total of all refunds applied to this tender. This amount is always negative or zero.
	RefundedMoney Money `json:"refunded_money"`
}

// Represents a timecard for an employee.
type Timecard struct {
	// The timecard's unique ID.
	ID string `json:"id"`
	// The ID of the employee the timecard is associated with.
	EmployeeID string `json:"employee_id"`
	// If true, the timecard was deleted by the merchant, and it is no longer
	// valid.
	Deleted bool `json:"deleted"`
	// The time the employee clocked in, in ISO 8601 format.
	ClockinTime string `json:"clockin_time"`
	// The time the employee clocked out, in ISO 8601 format.
	ClockoutTime string `json:"clockout_time"`
	// The ID of the location the employee clocked in from, if any.
	ClockinLocationID string `json:"clockin_location_id"`
	// The ID of the location the employee clocked out from, if any.
	ClockoutLocationID string `json:"clockout_location_id"`
	// The time when the timecard was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// The time when the timecard was most recently updated, in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
}

// Represents an event associated with a timecard, such as an employee clocking in.
type TimecardEvent struct {
	// The event's unique ID.
	ID string `json:"id"`
	// The type of action performed on the timecard, such as CLOCKIN or
	// API_CREATE.
	EventType string `json:"event_type"`
	// The time the employee clocked in, in ISO 8601 format.
	ClockinTime string `json:"clockin_time"`
	// The time the employee clocked out, in ISO 8601 format.
	ClockoutTime string `json:"clockout_time"`
	// The time when the event was created, in ISO 8601 format.
	CreatedAt string `json:"created_at"`
}
