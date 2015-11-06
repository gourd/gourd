package compile

// Context for template compilation
type Context map[string]interface{}

// Add stores a value to the key
func (c Context) Set(key string, v interface{}) {
	c[key] = v
}

// GetStr gets a value of string Type
func (c Context) GetStr(key string) string {
	if out, ok := c[key]; !ok {
		return ""
	} else if outStr, ok := out.(string); !ok {
		return ""
	} else {
		return outStr
	}
}

// GourdError is interface
type GourdError interface {
	Code() int
	Error() string
}
