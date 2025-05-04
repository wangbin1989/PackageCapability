package main

import (
	"github.com/wangbin1989/scm/src/common"
)

type Privilege struct {
	Defaults   Defaults     `json:"defaults,omitempty"`
	UserName   string       `json:"username,omitempty"`
	GroupName  string       `json:"groupname,omitempty"`
	CtrlScript []CtrlScript `json:"ctrl-script,omitempty"`
	Executable []Executable `json:"executable,omitempty"`
	Tool       []Tool       `json:"tool,omitempty"`
}

type Defaults struct {
	RunAs string `json:"run-as"`
}

type CtrlScript struct {
	Action string `json:"action"`
	RunAs  string `json:"run-as"`
}

type Executable struct {
	Relpath string `json:"relpath"`
	RunAs   string `json:"run-as"`
}

type Tool struct {
	Relpath      string `json:"relpath"`
	User         string `json:"user"`
	Group        string `json:"group"`
	Capabilities string `json:"capabilities"`
	Permission   string `json:"permission"`
}

type RunAs string

const (
	Root    RunAs = "root"
	Package RunAs = "package"
)

func getPrivilegePath(packageName string) string {
	path := "/var/packages/" + packageName + "/conf/privilege"
	return path
}

func updatePrivilegeFile(privilegePath string, relPath string, capabilities string) {
	privilege := common.ReadFromJson[Privilege](privilegePath)
	toolArray := privilege.Tool

	for i, tool := range toolArray {
		if tool.Relpath == relPath {
			privilege.Tool[i].Capabilities = capabilities
			return
		}
	}

	newTool := Tool{
		Relpath:      relPath,
		User:         "package",
		Group:        "package",
		Capabilities: capabilities,
		Permission:   "0700",
	}

	privilege.Tool = append(toolArray, newTool)

	common.WriteToJson(privilegePath, privilege)
}
