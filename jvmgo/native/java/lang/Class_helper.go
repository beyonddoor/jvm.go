package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func getParameterTypeArr(method *rtc.Method) *rtc.Object {
	paramTypes := method.ParameterTypes()
	paramCount := len(paramTypes)

	classClass := rtc.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.Refs()
		for i, paramType := range paramTypes {
			classObjs[i] = paramType.JClass()
		}
	}

	return classArr
}

func getReturnType(method *rtc.Method) *rtc.Object {
	goReturnType := method.ReturnType()
	return goReturnType.JClass()
}

func getExceptionTypeArr(method *rtc.Method) *rtc.Object {
	exTypes := method.ExceptionTypes()
	exCount := len(exTypes)

	classClass := rtc.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.Refs()
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass()
		}
	}

	return classArr
}

func getAnnotationByteArr(goBytes []byte) *rtc.Object {
	if goBytes != nil {
		jBytes := jutil.CastUint8sToInt8s(goBytes)
		return rtc.NewByteArray(jBytes)
	}
	return nil
}

func getSignatureStr(signature string) *rtc.Object {
	if signature != "" {
		return rtda.JString(signature)
	}
	return nil
}
