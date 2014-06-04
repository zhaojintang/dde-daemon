// This file is automatically generated, please don't edit manually.
package network

import (
	"fmt"
)

// Get key type
func getSettingPppoeKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_PPPOE_USERNAME:
		t = ktypeString
	case NM_SETTING_PPPOE_SERVICE:
		t = ktypeString
	case NM_SETTING_PPPOE_PASSWORD:
		t = ktypeString
	case NM_SETTING_PPPOE_PASSWORD_FLAGS:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting section
func isKeyInSettingPppoe(key string) bool {
	switch key {
	case NM_SETTING_PPPOE_USERNAME:
		return true
	case NM_SETTING_PPPOE_SERVICE:
		return true
	case NM_SETTING_PPPOE_PASSWORD:
		return true
	case NM_SETTING_PPPOE_PASSWORD_FLAGS:
		return true
	}
	return false
}

// Get key's default value
func getSettingPppoeDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_PPPOE_USERNAME:
		value = ""
	case NM_SETTING_PPPOE_SERVICE:
		value = ""
	case NM_SETTING_PPPOE_PASSWORD:
		value = ""
	case NM_SETTING_PPPOE_PASSWORD_FLAGS:
		value = uint32(0)
	}
	return
}

// Get JSON value generally
func generalGetSettingPppoeKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingPppoeKeyJSON: invalide key", key)
	case NM_SETTING_PPPOE_USERNAME:
		value = getSettingPppoeUsernameJSON(data)
	case NM_SETTING_PPPOE_SERVICE:
		value = getSettingPppoeServiceJSON(data)
	case NM_SETTING_PPPOE_PASSWORD:
		value = getSettingPppoePasswordJSON(data)
	case NM_SETTING_PPPOE_PASSWORD_FLAGS:
		value = getSettingPppoePasswordFlagsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingPppoeKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingPppoeKeyJSON: invalide key", key)
	case NM_SETTING_PPPOE_USERNAME:
		err = setSettingPppoeUsernameJSON(data, valueJSON)
	case NM_SETTING_PPPOE_SERVICE:
		err = setSettingPppoeServiceJSON(data, valueJSON)
	case NM_SETTING_PPPOE_PASSWORD:
		err = setSettingPppoePasswordJSON(data, valueJSON)
	case NM_SETTING_PPPOE_PASSWORD_FLAGS:
		err = setSettingPppoePasswordFlagsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingPppoeUsernameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME)
}
func isSettingPppoeServiceExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE)
}
func isSettingPppoePasswordExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD)
}
func isSettingPppoePasswordFlagsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS)
}

// Ensure section and key exists and not empty
func ensureSectionSettingPppoeExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_PPPOE_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_PPPOE_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_PPPOE_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_PPPOE_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_PPPOE_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_PPPOE_SETTING_NAME))
	}
}
func ensureSettingPppoeUsernameNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingPppoeUsernameExists(data) {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingPppoeUsername(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingPppoeServiceNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingPppoeServiceExists(data) {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingPppoeService(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingPppoePasswordNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingPppoePasswordExists(data) {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingPppoePassword(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingPppoePasswordFlagsNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingPppoePasswordFlagsExists(data) {
		rememberError(errs, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingPppoeUsername(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME)
	value = interfaceToString(ivalue)
	return
}
func getSettingPppoeService(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE)
	value = interfaceToString(ivalue)
	return
}
func getSettingPppoePassword(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD)
	value = interfaceToString(ivalue)
	return
}
func getSettingPppoePasswordFlags(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS)
	value = interfaceToUint32(ivalue)
	return
}

// Setter
func setSettingPppoeUsername(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME, value)
}
func setSettingPppoeService(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE, value)
}
func setSettingPppoePassword(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD, value)
}
func setSettingPppoePasswordFlags(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS, value)
}

// JSON Getter
func getSettingPppoeUsernameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME, getSettingPppoeKeyType(NM_SETTING_PPPOE_USERNAME))
	return
}
func getSettingPppoeServiceJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE, getSettingPppoeKeyType(NM_SETTING_PPPOE_SERVICE))
	return
}
func getSettingPppoePasswordJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD, getSettingPppoeKeyType(NM_SETTING_PPPOE_PASSWORD))
	return
}
func getSettingPppoePasswordFlagsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS, getSettingPppoeKeyType(NM_SETTING_PPPOE_PASSWORD_FLAGS))
	return
}

// JSON Setter
func setSettingPppoeUsernameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME, valueJSON, getSettingPppoeKeyType(NM_SETTING_PPPOE_USERNAME))
}
func setSettingPppoeServiceJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE, valueJSON, getSettingPppoeKeyType(NM_SETTING_PPPOE_SERVICE))
}
func setSettingPppoePasswordJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD, valueJSON, getSettingPppoeKeyType(NM_SETTING_PPPOE_PASSWORD))
}
func setSettingPppoePasswordFlagsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS, valueJSON, getSettingPppoeKeyType(NM_SETTING_PPPOE_PASSWORD_FLAGS))
}

// Logic JSON Setter

// Remover
func removeSettingPppoeUsername(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_USERNAME)
}
func removeSettingPppoeService(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_SERVICE)
}
func removeSettingPppoePassword(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD)
}
func removeSettingPppoePasswordFlags(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPPOE_SETTING_NAME, NM_SETTING_PPPOE_PASSWORD_FLAGS)
}
