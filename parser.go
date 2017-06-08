package main

import (
	"errors"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

type Config struct {
	Option   *Option
	Instance *Instance
	Dns      *Dns
}

type Option struct {
	Rollback      bool
	SecurityGroup []string `hcl:"security_group"`
}

type Instance struct {
	Tags  map[string]string
	Class string
}

type Dns struct {
	Type   string
	Config map[string]string
}

func hclParser(hclText string) (*Config, error) {
	obj, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := hcl.DecodeObject(&config, obj); err != nil {
		return nil, err
	}

	list, ok := obj.Node.(*ast.ObjectList)
	if !ok {
		return nil, errors.New("failed to parse: file does not contain root object")
	}

	if o := list.Filter("dns"); len(o.Items) == 0 || len(o.Items) > 1 {
		return nil, errors.New("one dns block is required")
	} else {
		dnsParser(&config, o.Items[0])
	}

	return &config, nil
}

func dnsParser(config *Config, item *ast.ObjectItem) error {
	key := item.Keys[0].Token.Value().(string)

	var m map[string]string
	if err := hcl.DecodeObject(&m, item.Val); err != nil {
		return err
	}

	config.Dns = &Dns{
		Type:   key,
		Config: m,
	}

	return nil
}
