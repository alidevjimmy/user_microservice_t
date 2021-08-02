package errors

const (
	DuplicateUsernameErrorMessage                                        = "این نام کاربری مطعلق به مشخص دیگیری است"
	DuplicatePhoneErrorMessage                                           = "این شماره مطعلق به شخص دیگیری است"
	InternalServerErrorMessage                                           = "خطایی رخ داده است لطفا با پشتیبانی تماس بگیرید"
	PhoneIsRequiredErrorMessage                                          = "شماره تلفن اجباری است"
	UsernameIsRequiredErrorMessage                                       = "نام کاربری اجباری است"
	NameIsRequiredErrorMessage                                           = "نام اجباری است"
	FamilyIsRequiredErrorMessage                                         = "نام خانوادگی اجباری است"
	AgeIsRequiredErrorMessage                                            = "سن اجباری است"
	PasswordIsRequiredErrorMessage                                       = "رمز عبور است"
	UsernameOnlyCanContainUnderlineAndEnglishWordsAndNumbersErrorMessage = "نام کاربری فقط می‌تواند شامل حروف انگلیسی، عدد و خط تیره باشه"
	PhoneOrUsernameIsRequiredErrorMessage                                = "نام کاربری یا شماره تماس اجباری است"
	UserNotFoundError                                                    = "کاربر یافت نشد"
	UserAlreadyActiveErrorMessage                                                    = "حساب کاربری شما فعال است"
	CodeOrPhoneDoesNotExistsErrorMessage                                             = "کد فعالسازی یا شماره تماس نادرست است"
	CodeIsExpiredErrorMessage                                                        = "کد فعالسازی شما منقضی شده است"
)
