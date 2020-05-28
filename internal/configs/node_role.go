package configs

import "github.com/iwind/TeaGo/lists"

type NodeRole = string

const (
	NodeRoleMaster = "MASTER"
	NodeRoleSlave  = "SLAVE"
)

func AllNodeRoles() []string {
	return []string{NodeRoleMaster, NodeRoleSlave}
}

func ExistNodeRole(role NodeRole) bool {
	return lists.ContainsString(AllNodeRoles(), role)
}
