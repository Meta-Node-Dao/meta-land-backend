package utility

import "regexp"

var (
	EthAddressRegexp   = "^0x[a-fA-F0-9]{40}$"
	EthSignatureRegexp = "^0x[a-fA-F0-9]{40}$"
	UrlRegexp          = "(http|ftp|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?"
)

func ValidateUrl(input string) bool {
	reg := regexp.MustCompile(UrlRegexp)
	return reg.MatchString(input)
}
