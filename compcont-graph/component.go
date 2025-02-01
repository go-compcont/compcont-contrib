package compcontgraph

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-compcont/compcont-core"
)

type Config struct {
	PUMLFile string `ccf:"puml_file"`
	PNGFile  string `ccf:"png_file"`
}

type Component interface{}

const TypeID compcont.ComponentTypeID = "contrib.compcont-graph"

func absolutePath(path []compcont.ComponentName, sep string) string {
	ss := []string{}
	for _, n := range path {
		ss = append(ss, n.String())
	}
	return strings.Join(ss, sep)
}

func iter(
	w io.Writer,
	instance any, name compcont.ComponentName,
	parentPath []compcont.ComponentName, deps []compcont.ComponentName,
	deep int, indent string) error {
	prefix := strings.Repeat(indent, deep)
	currentPath := append(parentPath, name)
	if container, ok := instance.(compcont.IComponentContainer); !ok {
		fmt.Fprintf(w, prefix+"component [%s] as %s\n", name, absolutePath(currentPath, "_"))
	} else {
		fmt.Fprintf(w, prefix+`package %s as %s {`+"\n", strconv.Quote(name.String()), absolutePath(currentPath, "_"))

		for _, name := range container.LoadedComponentNames() {
			comp, err := container.GetComponent(name)
			if err != nil {
				return err
			}
			err = iter(w, comp.Instance, comp.BuildContext.Config.Name, currentPath, comp.BuildContext.Config.Deps, deep+1, indent)
			if err != nil {
				return err
			}
		}
		fmt.Fprintln(w, prefix+`}`)
	}
	for _, dep := range deps {
		fmt.Fprintf(w, prefix+"%s --> %s\n", absolutePath(currentPath, "_"), absolutePath(append(parentPath, dep), "_"))
	}
	return nil
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, Component]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance Component, err error) {
		buf := &bytes.Buffer{}
		fmt.Fprintln(buf, "@startuml")
		iter(buf, ctx.FindRoot().Mount.Instance, compcont.ComponentName("ROOT"), nil, nil, 0, "    ")
		fmt.Fprintln(buf, "@enduml")

		if config.PUMLFile != "" {
			err = os.WriteFile(config.PUMLFile, buf.Bytes(), 0644)
			if err != nil {
				return
			}
		}
		if config.PNGFile != "" {
			renderToPNG(buf.String(), config.PNGFile)
		}

		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
