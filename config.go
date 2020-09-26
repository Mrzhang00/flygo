package flygo

import (
	"html/template"
	"strings"
)

const (
	//global config
	configKeyGlobalWebRoot     = "global.webRoot"     //global webRoot
	configKeyGlobalContextPath = "global.contextPath" //global contextPath

	//static config
	configKeyStaticEnable  = "static.enable"  //static pattern enable
	configKeyStaticPattern = "static.pattern" //static pattern URL
	configKeyStaticPrefix  = "static.prefix"  //static prefix local
	configKeyStaticCache   = "static.cache"   //static cache enable

	//view config
	configKeyViewEnable = "view.enable" //view enable
	configKeyViewPrefix = "view.prefix" //view prefix local
	configKeyViewSuffix = "view.suffix" //view suffix local
	configKeyViewCache  = "view.cache"  //view cache enable

	//template config
	configKeyTemplateEnable     = "template.enable"       //template enable
	configKeyTemplateFuncs      = "template.funcs"        //template funcs
	configKeyTemplateDelimLeft  = "template.delims.left"  //template delims left
	configKeyTemplateDelimRight = "template.delims.right" //template delims right

	configKeySessionEnable = "session.enable" //session enable
	//validate err config
	configKeyValidateErrCode = "validate.err.code" //err code
)

//Set webRoot
func (a *App) SetWebRoot(webRoot string) *App {
	a.setConfig(configKeyGlobalWebRoot, webRoot)
	return a
}

//Get webRoot
func (a *App) GetWebRoot() string {
	return a.getConfig(configKeyGlobalWebRoot).(string)
}

//Set contextPath
func (a *App) SetContextPath(contextPath string) *App {
	path := contextPath
	if contextPath != "" {
		path = "/" + strings.Trim(contextPath, "/")
	}
	a.setConfig(configKeyGlobalContextPath, path)
	return a
}

//Get contextPath
func (a *App) GetContextPath() string {
	return a.getConfig(configKeyGlobalContextPath).(string)
}

//Set static enable
func (a *App) SetStaticEnable(enable bool) *App {
	a.setConfig(configKeyStaticEnable, enable)
	return a
}

//Get static enable
func (a *App) GetStaticEnable() bool {
	return a.getConfig(configKeyStaticEnable).(bool)
}

//Set static pattern
func (a *App) SetStaticPattern(pattern string) *App {
	a.setConfig(configKeyStaticPattern, pattern)
	return a
}

//Get static pattern
func (a *App) GetStaticPattern() string {
	return a.getConfig(configKeyStaticPattern).(string)
}

//Set static prefix
func (a *App) SetStaticPrefix(prefix string) *App {
	a.setConfig(configKeyStaticPrefix, prefix)
	return a
}

//Get static prefix
func (a *App) GetStaticPrefix() string {
	return a.getConfig(configKeyStaticPrefix).(string)
}

//Set static cache
func (a *App) SetStaticCache(cache bool) *App {
	a.setConfig(configKeyStaticCache, cache)
	return a
}

//Get static cache
func (a *App) GetStaticCache() bool {
	return a.getConfig(configKeyStaticCache).(bool)
}

//Set view enable
func (a *App) SetViewEnable(enable bool) *App {
	a.setConfig(configKeyViewEnable, enable)
	return a
}

//Get view enable
func (a *App) GetViewEnable() bool {
	return a.getConfig(configKeyViewEnable).(bool)
}

//Set view prefix
func (a *App) SetViewPrefix(prefix string) *App {
	a.setConfig(configKeyViewPrefix, prefix)
	return a
}

//Get view prefix
func (a *App) GetViewPrefix() string {
	return a.getConfig(configKeyViewPrefix).(string)
}

//Set view suffix
func (a *App) SetViewSuffix(suffix string) *App {
	a.setConfig(configKeyViewSuffix, suffix)
	return a
}

//Get view suffix
func (a *App) GetViewSuffix() string {
	return a.getConfig(configKeyViewSuffix).(string)
}

//Set Template enable
func (a *App) SetTemplateEnable(enable bool) *App {
	a.setConfig(configKeyTemplateEnable, enable)
	return a
}

//Get Template enable
func (a *App) GetTemplateEnable() bool {
	return a.getConfig(configKeyTemplateEnable).(bool)
}

//Set Template funcs
func (a *App) SetTemplateFuncs(funcMap template.FuncMap) *App {
	a.setConfig(configKeyTemplateFuncs, funcMap)
	return a
}

//Get Template funcs
func (a *App) GetTemplateFuncs() template.FuncMap {
	return a.getConfig(configKeyTemplateFuncs).(template.FuncMap)
}

//Set Template delims left
func (a *App) SetTemplateDelimLeft(left string) *App {
	a.setConfig(configKeyTemplateDelimLeft, left)
	return a
}

//Get Template delims left
func (a *App) GetTemplateDelimLeft() string {
	return a.getConfig(configKeyTemplateDelimLeft).(string)
}

//Set Template delims right
func (a *App) SetTemplateDelimRight(right string) *App {
	a.setConfig(configKeyTemplateDelimRight, right)
	return a
}

//Get Template delims right
func (a *App) GetTemplateDelimRight() string {
	return a.getConfig(configKeyTemplateDelimRight).(string)
}

//Set view cache
func (a *App) SetViewCache(cache bool) *App {
	a.setConfig(configKeyViewCache, cache)
	return a
}

//Get view cache
func (a *App) GetViewCache() bool {
	return a.getConfig(configKeyViewCache).(bool)
}

//Set session enable
func (a *App) SetSessionEnable(enable bool) *App {
	a.setConfig(configKeySessionEnable, enable)
	return a
}

//Get session enable
func (a *App) GetSessionEnable() bool {
	return a.getConfig(configKeySessionEnable).(bool)
}

//Set validate err code
func (a *App) SetValidateErrCode(code int) *App {
	a.setConfig(configKeyValidateErrCode, code)
	return a
}

//Get validate err code
func (a *App) GetValidateErrCode() int {
	return a.getConfig(configKeyValidateErrCode).(int)
}

//Set config
func (a *App) setConfig(key string, value interface{}) *App {
	a.configs[key] = value
	return a
}

//Get config
func (a *App) getConfig(key string) interface{} {
	return a.configs[key]
}

//Print all configs
func (a *App) printConfig() {
	for name, val := range a.configs {
		a.LogInfo("Config [%s] = %v", name, val)
	}
}
