package types

// AttributesWithOverrides the document attributes with some overrides provided by the CLI (for example)
type AttributesWithOverrides struct {
	Content   map[string]interface{}
	Overrides map[string]string
}

// All returns all attributes
func (a AttributesWithOverrides) All() Attributes {
	result := Attributes{}
	for k, v := range a.Content {
		result[k] = v
	}
	for k, v := range a.Overrides {
		result[k] = v
	}
	return result
}

// Set sets the given attribute
func (a AttributesWithOverrides) Set(key string, value interface{}) {
	a.Content[key] = value
}

// Add adds the given attributes
func (a AttributesWithOverrides) Add(attrs map[string]interface{}) {
	for k, v := range attrs {
		a.Content[k] = v
	}
}

// Delete deletes the given attribute
func (a AttributesWithOverrides) Delete(key string) {
	delete(a.Content, key)
}

// GetAsString gets the string value for the given key (+ `true`),
// or empty string (+ `false`) if none was found
func (a AttributesWithOverrides) GetAsString(key string) (string, bool) {
	// if value is overridden
	if value, found := a.Overrides[key]; found {
		return value, true
	}
	// check in predefined attributes
	if value, found := Predefined[key]; found {
		return value, true
	}
	// if value is reset
	if _, found := a.Overrides["!"+key]; found {
		return "", false
	}
	if value, found := a.Content[key].(string); found {
		return value, true
	}
	// TODO: raise a warning if there was no entry found
	return "", false
}

// GetAsStringWithDefault gets the string value for the given key,
// or returns the given default value
func (a AttributesWithOverrides) GetAsStringWithDefault(key, defaultValue string) string {
	if value, found := a.Overrides[key]; found {
		return value
	}
	// check in predefined attributes
	if value, found := Predefined[key]; found {
		return value
	}
	if value, found := a.Content[key].(string); found {
		return value
	}
	// TODO: raise a warning if there was no entry found
	return defaultValue
}
