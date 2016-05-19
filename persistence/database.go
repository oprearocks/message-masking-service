package persistence

// Expressions represents the set of currenly available regular expressions
// Used for matching and replacing the private data within the messages
var Expressions = map[string]string{
	"CreditCard":           `\b(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6011[0-9]{12}|622((12[6-9]|1[3-9][0-9])|([2-8][0-9][0-9])|(9(([0-1][0-9])|(2[0-5]))))[0-9]{10}|64[4-9][0-9]{13}|65[0-9]{14}|3(?:0[0-5]|[68][0-9])[0-9]{11}|3[47][0-9]{13})+\b`,
	"NorthAmericanPhone":   `\b((([0-9]{1})*[- .(]*([0-9]{3})[- .)]*[0-9]{3}[- .]*[0-9]{4})+)+\b`,
	"SocialSecurityNumber": `\b([0-9]{3}[-]*[0-9]{2}[-]*[0-9]{4})+\b`,
}
