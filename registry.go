package main

import (
	functions "go-learning/basics/1.10_method"
	variablescope "go-learning/basics/1.11_variable_scope"
	array "go-learning/basics/1.12_array"
	slice "go-learning/basics/1.13_slice"
	mapcollection "go-learning/basics/1.14_map"
	rangeiteration "go-learning/basics/1.15_range"
	typeconversion "go-learning/basics/1.16_type_conversion"
	interfaceexample "go-learning/basics/1.17_interface"
	concurrency "go-learning/basics/1.18_concurrency"
	blockchainio "go-learning/blockchain/1_blockchain_io"
	comparison "go-learning/blockchain/2_go_vs_nodejs"
	goexclusive "go-learning/blockchain/3_go_exclusive_scenarios"
	realscenarios "go-learning/blockchain/4_real_business_scenarios"
	pointers "go-learning/basics/1.4_pointer"
	structs "go-learning/basics/1.5_struct"
	constants "go-learning/basics/1.6_constants_enum"
	operators "go-learning/basics/1.7_operators"
	controlflow "go-learning/basics/1.8_control_flow"
	loopcontrol "go-learning/basics/1.9_loop_control"
	ginroutes "go-learning/gin/1_router_parameter"
	ginmiddleware "go-learning/gin/2_middleware"
	gormexamples "go-learning/gorm"
)

// DemoRegistry 示例注册表 - key不带Demo后缀，直接使用简洁名称
var demoRegistry = map[string]interface{}{
	// 指针示例
	"Pointers": pointers.PointersDemo,
	// 函数示例
	"Functions": functions.FunctionsDemo,
	"Closure":   functions.ClosureDemo,
	"Method":    functions.MethodDemo,
	// 流程控制示例
	"IfStatement":     controlflow.IfStatementDemo,
	"SwitchStatement": controlflow.SwitchStatementDemo,
	"ForLoop":         loopcontrol.ForLoopDemo,
	"Break":           loopcontrol.BreakDemo,
	"Continue":        loopcontrol.ContinueDemo,
	"Goto":            loopcontrol.GotoDemo,
	// 变量作用域示例
	"LocalVariable":  variablescope.LocalVariableDemo,
	"GlobalVariable": variablescope.GlobalVariableDemo,
	// 数组示例
	"ArrayDeclaration":      array.ArrayDeclarationDemo,
	"ArrayAccess":           array.ArrayAccessDemo,
	"MultidimensionalArray": array.MultidimensionalArrayDemo,
	"ArrayAsParameter":      array.ArrayAsParameterDemo,
	// 切片示例
	"SliceDeclaration":         slice.SliceDeclarationDemo,
	"SliceUsage":               slice.SliceUsageDemo,
	"SliceUnderlyingPrinciple": slice.SliceUnderlyingPrincipleDemo,
	// map 示例
	"MapDeclaration": mapcollection.MapDeclarationDemo,
	"MapUsage":       mapcollection.MapUsageDemo,
	"MapAsParameter": mapcollection.MapAsParameterDemo,
	"MapConcurrent":  mapcollection.MapConcurrentDemo,
	// range 迭代示例
	"RangeString":     rangeiteration.RangeStringDemo,
	"RangeArraySlice": rangeiteration.RangeArraySliceDemo,
	"RangeChannel":    rangeiteration.RangeChannelDemo,
	"RangeMap":        rangeiteration.RangeMapDemo,
	// 类型转换示例
	"NumericConversion":   typeconversion.NumericConversionDemo,
	"StringConversion":    typeconversion.StringConversionDemo,
	"InterfaceConversion": typeconversion.InterfaceConversionDemo,
	"StructConversion":    typeconversion.StructConversionDemo,
	// 接口示例
	"InterfaceBasic":          interfaceexample.InterfaceBasicDemo,
	"InterfaceImplementation": interfaceexample.InterfaceImplementationDemo,
	"InterfaceReceiver":       interfaceexample.InterfaceReceiverDemo,
	"InterfaceNesting":        interfaceexample.InterfaceNestingDemo,
	"InterfaceEmpty":          interfaceexample.InterfaceEmptyDemo,
	// 并发示例
	"Goroutine":      concurrency.GoroutineDemo,
	"Channel":        concurrency.ChannelDemo,
	"LockAndChannel": concurrency.LockAndChannelDemo,
	// 区块链 I/O 示例
	"BlockchainIO": blockchainio.BlockchainIODemo,
	// Go vs Node.js 对比示例
	"GoVsNodejs": comparison.GoVsNodejsConcurrencyDemo,
	// Go 独占优势场景示例
	"GoExclusive": goexclusive.GoExclusiveWeb3ScenariosDemo,
	// 真实业务场景分析示例
	"BlockSyncNecessity": realscenarios.BlockSyncNecessityDemo,
	// 结构体示例
	"AnonymousStruct":  structs.AnonymousStructDemo,
	"NestedStruct":     structs.NestedStructDemo,
	"StructMethods":    structs.StructMethodsDemo,
	"CrossFileUsage":   structs.CrossFileUsageDemo,
	"LowercaseStruct":  structs.LowercaseStructDemo,
	"RealWorldExample": structs.RealWorldExampleDemo,
	// 常量示例
	"Constants": constants.ConstantsDemo,
	"Enums":     constants.EnumsDemo,

	// 运算符示例
	"ArithmeticOperators": operators.ArithmeticOperatorsDemo,
	"Operators":           operators.OperatorsDemo,
	// Gin路由示例
	"BasicRoutes":             ginroutes.BasicRoutesDemo,
	"RESTfulRoutes":           ginroutes.RESTfulRoutesDemo,
	"PathParameter":           ginroutes.PathParameterDemo,
	"QueryParameter":          ginroutes.QueryParameterDemo,
	"JSONBinding":             ginroutes.JSONBindingDemo,
	"FormBinding":             ginroutes.FormBindingDemo,
	"RouteConflict":           ginroutes.RouteConflictDemo,
	"RouteGroup":              ginroutes.RouteGroupDemo,
	"RegexRoute":              ginroutes.RegexRouteDemo,
	"MiddlewareRoute":         ginroutes.MiddlewareRouteDemo,
	"StaticFiles":             ginroutes.StaticFilesDemo,
	"CustomValidation":        ginroutes.CustomValidationDemo,
	"ValidationErrorHandling": ginroutes.ValidationErrorHandlingDemo,
	"BuiltinValidationTags":   ginroutes.BuiltinValidationTagsDemo,
	"UnifiedResponse":         ginroutes.UnifiedResponseDemo,
	"SensitiveDataFilter":     ginroutes.SensitiveDataFilterDemo,
	"RateLimit":               ginroutes.RateLimitDemo,
	"VersionControl":          ginroutes.VersionControlDemo,
	"SwaggerIntegration":      ginroutes.SwaggerIntegrationDemo,
	"SwaggerAnnotations":      ginroutes.SwaggerAnnotationsDemo,
	"SwaggerSecurity":         ginroutes.SwaggerSecurityDemo,
	// Gin中间件示例
	"MiddlewareFlow":          ginmiddleware.MiddlewareFlowDemo,
	"JWTAuth":                 ginmiddleware.JWTAuthDemo,
	"CORSMiddleware":          ginmiddleware.CORSMiddlewareDemo,
	"MiddlewareDebug":         ginmiddleware.MiddlewareDebugDemo,
	"MiddlewareBestPractices": ginmiddleware.MiddlewareBestPracticesDemo,
	"MiddlewareTest":          ginmiddleware.MiddlewareTestDemo,
	"GinRouter":               ginmiddleware.GinRouterDemo,
	// GORM 示例（基于 fuyelead 项目）
	"GormBasics":              gormexamples.GormBasicsDemo,
	"GormRelationships":       gormexamples.GormRelationshipsDemo,
	"GormQueryOptimization":   gormexamples.GormQueryOptimizationDemo,
	"GormDatabaseConfig":      gormexamples.GormDatabaseConfigDemo,
	"GormPreloadExplanation":  gormexamples.GormPreloadExplanationDemo,
}
