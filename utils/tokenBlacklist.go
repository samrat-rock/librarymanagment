
package utils

var BlacklistedTokens = make(map[string]bool)

func BlacklistToken(token string) {
    BlacklistedTokens[token] = true
}

func IsTokenBlacklisted(token string) bool {
    return BlacklistedTokens[token]
}
