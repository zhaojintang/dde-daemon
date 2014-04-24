// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnL2tpPppKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_NODEFLATE:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_NO_PCOMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE:
		t = ktypeUint32
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnL2tpPpp(key string) bool {
	switch key {
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40:
		return true
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128:
		return true
	case NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL:
		return true
	case NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_NODEFLATE:
		return true
	case NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_NO_PCOMP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP:
		return true
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE:
		return true
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnL2tpPppKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_NODEFLATE:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_NO_PCOMP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE:
		valueJSON = `0`
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL:
		valueJSON = `0`
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnL2tpPppKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingVpnL2tpPppKeyJSON: invalide key", key)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP:
		value = getSettingVpnL2tpKeyRefuseEapJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP:
		value = getSettingVpnL2tpKeyRefusePapJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP:
		value = getSettingVpnL2tpKeyRefuseChapJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP:
		value = getSettingVpnL2tpKeyRefuseMschapJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2:
		value = getSettingVpnL2tpKeyRefuseMschapv2JSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE:
		value = getSettingVpnL2tpKeyRequireMppeJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40:
		value = getSettingVpnL2tpKeyRequireMppe40JSON(data)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128:
		value = getSettingVpnL2tpKeyRequireMppe128JSON(data)
	case NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL:
		value = getSettingVpnL2tpKeyMppeStatefulJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP:
		value = getSettingVpnL2tpKeyNobsdcompJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_NODEFLATE:
		value = getSettingVpnL2tpKeyNodeflateJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP:
		value = getSettingVpnL2tpKeyNoVjCompJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_NO_PCOMP:
		value = getSettingVpnL2tpKeyNoPcompJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP:
		value = getSettingVpnL2tpKeyNoAccompJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE:
		value = getSettingVpnL2tpKeyLcpEchoFailureJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL:
		value = getSettingVpnL2tpKeyLcpEchoIntervalJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnL2tpPppKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingVpnL2tpPppKeyJSON: invalide key", key)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP:
		setSettingVpnL2tpKeyRefuseEapJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP:
		setSettingVpnL2tpKeyRefusePapJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP:
		setSettingVpnL2tpKeyRefuseChapJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP:
		setSettingVpnL2tpKeyRefuseMschapJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2:
		setSettingVpnL2tpKeyRefuseMschapv2JSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE:
		logicSetSettingVpnL2tpKeyRequireMppeJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40:
		setSettingVpnL2tpKeyRequireMppe40JSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128:
		setSettingVpnL2tpKeyRequireMppe128JSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL:
		setSettingVpnL2tpKeyMppeStatefulJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP:
		setSettingVpnL2tpKeyNobsdcompJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_NODEFLATE:
		setSettingVpnL2tpKeyNodeflateJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP:
		setSettingVpnL2tpKeyNoVjCompJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_NO_PCOMP:
		setSettingVpnL2tpKeyNoPcompJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP:
		setSettingVpnL2tpKeyNoAccompJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE:
		setSettingVpnL2tpKeyLcpEchoFailureJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL:
		setSettingVpnL2tpKeyLcpEchoIntervalJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnL2tpKeyRefuseEapExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP)
}
func isSettingVpnL2tpKeyRefusePapExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP)
}
func isSettingVpnL2tpKeyRefuseChapExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP)
}
func isSettingVpnL2tpKeyRefuseMschapExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP)
}
func isSettingVpnL2tpKeyRefuseMschapv2Exists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2)
}
func isSettingVpnL2tpKeyRequireMppeExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE)
}
func isSettingVpnL2tpKeyRequireMppe40Exists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40)
}
func isSettingVpnL2tpKeyRequireMppe128Exists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128)
}
func isSettingVpnL2tpKeyMppeStatefulExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL)
}
func isSettingVpnL2tpKeyNobsdcompExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP)
}
func isSettingVpnL2tpKeyNodeflateExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE)
}
func isSettingVpnL2tpKeyNoVjCompExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP)
}
func isSettingVpnL2tpKeyNoPcompExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP)
}
func isSettingVpnL2tpKeyNoAccompExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP)
}
func isSettingVpnL2tpKeyLcpEchoFailureExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE)
}
func isSettingVpnL2tpKeyLcpEchoIntervalExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnL2tpPppExists(data _ConnectionData, errs FieldKeyErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME))
	}
}
func ensureSettingVpnL2tpKeyRefuseEapNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRefuseEapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRefusePapNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRefusePapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRefuseChapNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRefuseChapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRefuseMschapNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRefuseMschapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRefuseMschapv2NoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRefuseMschapv2Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRequireMppeNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRequireMppeExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRequireMppe40NoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRequireMppe40Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyRequireMppe128NoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyRequireMppe128Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyMppeStatefulNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyMppeStatefulExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyNobsdcompNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyNobsdcompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyNodeflateNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyNodeflateExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyNoVjCompNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyNoVjCompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyNoPcompNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyNoPcompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyNoAccompNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyNoAccompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyLcpEchoFailureNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyLcpEchoFailureExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyLcpEchoIntervalNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnL2tpKeyLcpEchoIntervalExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingVpnL2tpKeyRefuseEap(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP).(bool)
	return
}
func getSettingVpnL2tpKeyRefusePap(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP).(bool)
	return
}
func getSettingVpnL2tpKeyRefuseChap(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP).(bool)
	return
}
func getSettingVpnL2tpKeyRefuseMschap(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP).(bool)
	return
}
func getSettingVpnL2tpKeyRefuseMschapv2(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2).(bool)
	return
}
func getSettingVpnL2tpKeyRequireMppe(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE).(bool)
	return
}
func getSettingVpnL2tpKeyRequireMppe40(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40).(bool)
	return
}
func getSettingVpnL2tpKeyRequireMppe128(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128).(bool)
	return
}
func getSettingVpnL2tpKeyMppeStateful(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL).(bool)
	return
}
func getSettingVpnL2tpKeyNobsdcomp(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP).(bool)
	return
}
func getSettingVpnL2tpKeyNodeflate(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE).(bool)
	return
}
func getSettingVpnL2tpKeyNoVjComp(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP).(bool)
	return
}
func getSettingVpnL2tpKeyNoPcomp(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP).(bool)
	return
}
func getSettingVpnL2tpKeyNoAccomp(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP).(bool)
	return
}
func getSettingVpnL2tpKeyLcpEchoFailure(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE).(uint32)
	return
}
func getSettingVpnL2tpKeyLcpEchoInterval(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL).(uint32)
	return
}

// Setter
func setSettingVpnL2tpKeyRefuseEap(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP, value)
}
func setSettingVpnL2tpKeyRefusePap(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP, value)
}
func setSettingVpnL2tpKeyRefuseChap(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP, value)
}
func setSettingVpnL2tpKeyRefuseMschap(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP, value)
}
func setSettingVpnL2tpKeyRefuseMschapv2(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2, value)
}
func setSettingVpnL2tpKeyRequireMppe(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE, value)
}
func setSettingVpnL2tpKeyRequireMppe40(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40, value)
}
func setSettingVpnL2tpKeyRequireMppe128(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128, value)
}
func setSettingVpnL2tpKeyMppeStateful(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL, value)
}
func setSettingVpnL2tpKeyNobsdcomp(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP, value)
}
func setSettingVpnL2tpKeyNodeflate(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE, value)
}
func setSettingVpnL2tpKeyNoVjComp(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP, value)
}
func setSettingVpnL2tpKeyNoPcomp(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP, value)
}
func setSettingVpnL2tpKeyNoAccomp(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP, value)
}
func setSettingVpnL2tpKeyLcpEchoFailure(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, value)
}
func setSettingVpnL2tpKeyLcpEchoInterval(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL, value)
}

// JSON Getter
func getSettingVpnL2tpKeyRefuseEapJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP))
	return
}
func getSettingVpnL2tpKeyRefusePapJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP))
	return
}
func getSettingVpnL2tpKeyRefuseChapJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP))
	return
}
func getSettingVpnL2tpKeyRefuseMschapJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP))
	return
}
func getSettingVpnL2tpKeyRefuseMschapv2JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2))
	return
}
func getSettingVpnL2tpKeyRequireMppeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE))
	return
}
func getSettingVpnL2tpKeyRequireMppe40JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40))
	return
}
func getSettingVpnL2tpKeyRequireMppe128JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128))
	return
}
func getSettingVpnL2tpKeyMppeStatefulJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL))
	return
}
func getSettingVpnL2tpKeyNobsdcompJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP))
	return
}
func getSettingVpnL2tpKeyNodeflateJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NODEFLATE))
	return
}
func getSettingVpnL2tpKeyNoVjCompJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP))
	return
}
func getSettingVpnL2tpKeyNoPcompJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_PCOMP))
	return
}
func getSettingVpnL2tpKeyNoAccompJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP))
	return
}
func getSettingVpnL2tpKeyLcpEchoFailureJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE))
	return
}
func getSettingVpnL2tpKeyLcpEchoIntervalJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL))
	return
}

// JSON Setter
func setSettingVpnL2tpKeyRefuseEapJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP))
}
func setSettingVpnL2tpKeyRefusePapJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP))
}
func setSettingVpnL2tpKeyRefuseChapJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP))
}
func setSettingVpnL2tpKeyRefuseMschapJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP))
}
func setSettingVpnL2tpKeyRefuseMschapv2JSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2))
}
func setSettingVpnL2tpKeyRequireMppeJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE))
}
func setSettingVpnL2tpKeyRequireMppe40JSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40))
}
func setSettingVpnL2tpKeyRequireMppe128JSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128))
}
func setSettingVpnL2tpKeyMppeStatefulJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL))
}
func setSettingVpnL2tpKeyNobsdcompJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP))
}
func setSettingVpnL2tpKeyNodeflateJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NODEFLATE))
}
func setSettingVpnL2tpKeyNoVjCompJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP))
}
func setSettingVpnL2tpKeyNoPcompJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_PCOMP))
}
func setSettingVpnL2tpKeyNoAccompJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP))
}
func setSettingVpnL2tpKeyLcpEchoFailureJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE))
}
func setSettingVpnL2tpKeyLcpEchoIntervalJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL, valueJSON, getSettingVpnL2tpPppKeyType(NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL))
}

// Remover
func removeSettingVpnL2tpKeyRefuseEap(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_EAP)
}
func removeSettingVpnL2tpKeyRefusePap(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_PAP)
}
func removeSettingVpnL2tpKeyRefuseChap(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_CHAP)
}
func removeSettingVpnL2tpKeyRefuseMschap(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAP)
}
func removeSettingVpnL2tpKeyRefuseMschapv2(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REFUSE_MSCHAPV2)
}
func removeSettingVpnL2tpKeyRequireMppe(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE)
}
func removeSettingVpnL2tpKeyRequireMppe40(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_40)
}
func removeSettingVpnL2tpKeyRequireMppe128(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_REQUIRE_MPPE_128)
}
func removeSettingVpnL2tpKeyMppeStateful(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_MPPE_STATEFUL)
}
func removeSettingVpnL2tpKeyNobsdcomp(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NOBSDCOMP)
}
func removeSettingVpnL2tpKeyNodeflate(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NODEFLATE)
}
func removeSettingVpnL2tpKeyNoVjComp(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_VJ_COMP)
}
func removeSettingVpnL2tpKeyNoPcomp(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_PCOMP)
}
func removeSettingVpnL2tpKeyNoAccomp(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_NO_ACCOMP)
}
func removeSettingVpnL2tpKeyLcpEchoFailure(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE)
}
func removeSettingVpnL2tpKeyLcpEchoInterval(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_INTERVAL)
}
