/*
 * Mqtts API
 *
 * Mqtts service
 *
 * API version: 1.0.0
 */
package swagger

// SSL interface
type Security struct {
	// Indicate if SSl is Enabled or not
	SSL bool `json:"SSL"`
	// Indicates if Client verification is enabled or not
	ClientVerification bool `json:"clientVerification"`
}
