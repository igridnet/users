package models

import (
	"fmt"
	"regexp"
)

type NodeType int

const (
	SensorNode NodeType = iota
	ActuatorNode
	ControllerNode
)

type Admin struct {
	tableName struct{} `pg:"users,alias:users"`
	ID        string   `json:"id" pg:"id,pk,unique"`
	Name      string   `json:"name" pg:"name"`
	Email     string   `json:"email" pg:"email"`
	Password  string   `json:"password" pg:"password"`
	Created   int64    `json:"created" pg:"created"`
}

type Region struct {
	tableName struct{} `pg:"regions,alias:regions"`
	ID        string   `json:"id" pg:"id,pk,unique"`
	Name      string   `json:"name" pg:"name"`
	Desc      string   `json:"description" pg:"description"`
	Created   int64    `json:"created" pg:"created"`
}

type Node struct {
	// By default go-pg generates table name and alias from struct name.
	tableName struct{} `pg:"nodes,alias:nodes"`
	UUID      string   `json:"uuid,omitempty"  pg:"uuid,pk,unique"`
	Addr      string   `json:"addr" pg:"addr"`
	Key       string   `json:"key,omitempty" pg:"key"`
	Name      string   `json:"name" pg:"name"`
	Type      int      `json:"type" pg:"type"`
	Region    string   `json:"region" pg:"region"`
	Latd      string   `json:"latitude" pg:"latitude"`
	Long      string   `json:"longitude" pg:"longitude"`
	Created   int64    `json:"created,omitempty" pg:"created"`
	Master    string   `json:"master,omitempty" pg:"master"`
}

func (nt NodeType) String() string {
	values := []string{
		"sensor", "actuator", "controller",
	}

	return values[nt]
}

var macAddrRegex = regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4})$")

func (n *Node) Valid() (bool, error) {
	if n.Type > int(ControllerNode) {
		return false, fmt.Errorf("invalid node type")
	}

	if len(n.Addr) == 0 || n.Addr == "" {
		return false, fmt.Errorf("invalid mac address")
	}

	valid := macAddrRegex.MatchString(n.Addr)

	if !valid {
		return false, fmt.Errorf("invalid mac address")
	}

	return true, nil
}
