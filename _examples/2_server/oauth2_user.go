package main

// OAuth2User is the generic user interface
// for OAuth2 login check
type OAuth2User interface {
	// PasswordIs matches a string with the stored password.
	// If the stored password is hash, this function will apply to the
	// input before matching.
	PasswordIs(pass string) bool
}
