package gobpmntypes

/*

@Interfaces

- IFBaseID
- IFBaseListenerID
- IFBaseName
- IFBaseElement
- IFBaseValue
- IFBaseEvent
- IFBaseClass
- IFBaseLabel
- IFBaseType
- IFBaseConfig
- IFBaseKey
- IFBaseDefaultValue
- IFBaseScriptFormat
- IFBaseLocalVariableName
- IFBaseVariableAssignmentValue

*/

// IFBaseID ...
type IFBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

// IFBaseListenerID ...
type IFBaseListenerID interface {
	SetListenerID(listenerID string)
	GetListenerID() STR_PTR
}

// IFBaseName ...
type IFBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

// IFBaseElement ...
type IFBaseElement interface {
	SetElement(typ string, suffix interface{})
	GetElement() STR_PTR
}

// IFBaseValue ...
type IFBaseValue interface {
	SetValue(value string)
	GetValue() STR_PTR
}

// IFBaseEvent ...
type IFBaseEvent interface {
	SetEvent(event string)
	GetEvent() STR_PTR
}

// IFBaseClass ...
type IFBaseClass interface {
	SetClass(class string)
	GetClass() STR_PTR
}

// IFBaseLabel ...
type IFBaseLabel interface {
	SetLabel(label string)
	GetLabel() STR_PTR
}

// IFBaseType ...
type IFBaseType interface {
	SetType(typ string)
	GetType() STR_PTR
}

// IFBaseConfig ...
type IFBaseConfig interface {
	SetConfig(config string)
	GetConfig() STR_PTR
}

// IFBaseKey ...
type IFBaseKey interface {
	SetKey(key string)
	GetKey() STR_PTR
}

// IFBaseDefaultValue ...
type IFBaseDefaultValue interface {
	SetDefaultValue(defaultValue string)
	GetDefaultValue() STR_PTR
}

// IFBaseScriptFormat ...
type IFBaseScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() STR_PTR
}

// IFBaseLocalVariableName ...
type IFBaseLocalVariableName interface {
	SetLocalVariableName(variable string)
	GetLocalVariableName() STR_PTR
}

// IFBaseVariableAssignmentValue ...
type IFBaseVariableAssignmentValue interface {
	SetVariableAssignmentValue(value string)
	GetVariableAssignmentValue() STR_PTR
}
