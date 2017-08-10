package configs

var AllowedDomains = []string{"fyndiq",}
var FeedbackCommandName = "/thanks"


func DomainAllowed(domain string) bool {
    for _, allowedDomain := range AllowedDomains {
        if domain == allowedDomain {
            return true
        }
    }

    return false
}

