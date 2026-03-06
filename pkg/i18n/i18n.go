package i18n

// Resolve returns the appropriate bilingual string based on locale.
// If locale is "en" and the English value is non-empty, return the English value.
// Otherwise, fallback to the Indonesian value.
func Resolve(locale, valueID string, valueEN *string) string {
	if locale == "en" && valueEN != nil && *valueEN != "" {
		return *valueEN
	}
	return valueID
}

// ResolveOptional handles optional bilingual fields where both values may be nil/empty.
func ResolveOptional(locale string, valueID *string, valueEN *string) string {
	if locale == "en" && valueEN != nil && *valueEN != "" {
		return *valueEN
	}
	if valueID != nil {
		return *valueID
	}
	return ""
}
