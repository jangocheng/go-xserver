package main

import (
	"fmt"

	"github.com/fananchong/go-xserver/common"
)

// PluginObj : 代表一个插件对象
var PluginObj common.Plugin

// PluginType : 插件类型
var PluginType common.NodeType

// Ctx : 应用程序上下文
var Ctx *common.Context

func init() {
	fmt.Println("LOAD PLUGIN: LOBBY")
	PluginObj = &Plugin{}
	PluginType = common.Lobby
}

// Plugin : 插件类
type Plugin struct {
}

// Start : 插件类实现启动
func (plugin *Plugin) Start() bool {
	Ctx.Log.Infoln("Plugin Start")
	return true
}

// Close : 插件类实现关闭
func (plugin *Plugin) Close() {
	Ctx.Log.Infoln("Plugin Close")
}

// main : 作为插件包，该函数可以不存在。添加之，是避免 go-lint 烦人的错误提示
func main() {}
