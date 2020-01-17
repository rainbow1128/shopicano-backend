package errors

type ErrorCode string

const (
	UnableToParseBody                             ErrorCode = "400001"
	InvalidMultiPartBody                          ErrorCode = "400002"
	BearerTokenNotFound                           ErrorCode = "400003"
	InvalidRequest                                ErrorCode = "400005"
	CartMustHaveAllDigitalOrAllNonDigitalProducts ErrorCode = "400006"
	StoreCreationDataInvalid                      ErrorCode = "422001"
	UserLoginDataInvalid                          ErrorCode = "422002"
	UserSignUpDataInvalid                         ErrorCode = "422003"
	UserUpdateDataInvalid                         ErrorCode = "422004"
	CategoryCreationDataInvalid                   ErrorCode = "422005"
	CollectionCreationDataInvalid                 ErrorCode = "422006"
	ShippingMethodCreationDataInvalid             ErrorCode = "422007"
	PaymentMethodCreationDataInvalid              ErrorCode = "422008"
	ProductCreationDataInvalid                    ErrorCode = "422009"
	ProductAttributeCreationDataInvalid           ErrorCode = "422010"
	AddCollectionDataInvalid                      ErrorCode = "422011"
	AdditionalChargeDataInvalid                   ErrorCode = "422012"
	OrderDataInvalid                              ErrorCode = "422013"
	OrderPaymentDataInvalid                       ErrorCode = "422014"
	AddStoreStaffDataInvalid                      ErrorCode = "422015"
	StoreCreationQueryFailed                      ErrorCode = "500001"
	DatabaseQueryFailed                           ErrorCode = "500002"
	PasswordEncryptionFailed                      ErrorCode = "500003"
	UserSignUpFailed                              ErrorCode = "500004"
	UserUpdateFailed                              ErrorCode = "500005"
	UnableToReadMultiPartData                     ErrorCode = "500006"
	MinioServiceFailed                            ErrorCode = "500007"
	UserLoginFailed                               ErrorCode = "500008"
	PaymentGatewayFailed                          ErrorCode = "500009"
	PaymentProcessingFailed                       ErrorCode = "500009"
	FailedToEnqueueTask                           ErrorCode = "500010"
	StoreAlreadyExists                            ErrorCode = "409001"
	StoreMemberAlreadyExists                      ErrorCode = "409002"
	CategoryAlreadyExists                         ErrorCode = "409003"
	CollectionAlreadyExists                       ErrorCode = "409004"
	ShippingMethodAlreadyExists                   ErrorCode = "409005"
	PaymentMethodAlreadyExists                    ErrorCode = "409006"
	ProductVariantAlreadyExists                   ErrorCode = "409007"
	UserAlreadyExists                             ErrorCode = "409008"
	ProductAlreadyExists                          ErrorCode = "409009"
	AdditionalChargeAlreadyExists                 ErrorCode = "409010"
	PaymentAlreadyProcessed                       ErrorCode = "409011"
	ProductAttributeAlreadyExists                 ErrorCode = "409012"
	UserHasAStore                                 ErrorCode = "403001"
	UserSignUpDisabled                            ErrorCode = "403002"
	StoreCreationDisabled                         ErrorCode = "403003"
	AppIsNotActivated                             ErrorCode = "403004"
	AuthorizationTokenMissing                     ErrorCode = "403005"
	AuthorizationTokenInvalid                     ErrorCode = "403006"
	AuthorizationTokenExpired                     ErrorCode = "403007"
	UserScopeUnauthorized                         ErrorCode = "403008"
	StoreNotFound                                 ErrorCode = "404001"
	SettingsNotFound                              ErrorCode = "404002"
	UserNotFound                                  ErrorCode = "404003"
	UserGroupNotFound                             ErrorCode = "404004"
	CategoryNotFound                              ErrorCode = "404005"
	CollectionNotFound                            ErrorCode = "404006"
	ProductNotFound                               ErrorCode = "404007"
	ShippingMethodNotFound                        ErrorCode = "404008"
	PaymentMethodNotFound                         ErrorCode = "404009"
	ProductVariantNotFound                        ErrorCode = "404010"
	AddressNotFound                               ErrorCode = "404011"
	AdditionalChargeNotFound                      ErrorCode = "404012"
	ProductUnavailable                            ErrorCode = "404013"
	OrderNotFound                                 ErrorCode = "404014"
	LoginCredentialsInvalid                       ErrorCode = "401001"
	VerificationTokenIsInvalid                    ErrorCode = "401002"
	UserNotActive                                 ErrorCode = "401003"
)
