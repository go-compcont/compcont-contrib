package compcontgraph

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-compcont/compcont-core"
	compcontzap "github.com/go-compcont/compcont-std/compcont-zap"
	"go.uber.org/zap"
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

func iter(w io.Writer, ctx compcont.BuildContext, deep int, indent string) error {
	prefix := strings.Repeat(indent, deep)

	currentPath := ctx.GetAbsolutePath()
	if ctx.Mount == nil {
		// 根节点
		return nil
	}
	parentPath := ctx.Container.GetContext().GetAbsolutePath()
	name := ctx.Config.Name
	deps := ctx.Config.Deps
	if container, ok := ctx.Mount.Instance.(compcont.IComponentContainer); !ok {
		fmt.Fprintf(w, prefix+"component [%s] as %s\n", name, absolutePath(currentPath, "_"))
	} else {
		fmt.Fprintf(w, prefix+`package %s as %s {`+"\n", strconv.Quote(name.String()), absolutePath(currentPath, "_"))

		for _, name := range container.LoadedComponentNames() {
			comp, err := container.GetComponent(name)
			if err != nil {
				return err
			}
			err = iter(w, comp.BuildContext, deep+1, indent)
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
		iter(buf, ctx.FindRoot(), 0, "    ")
		fmt.Fprintln(buf, "@enduml")

		if config.PUMLFile != "" {
			err = os.WriteFile(config.PUMLFile, buf.Bytes(), 0644)
			if err != nil {
				compcontzap.GetDefault().Info("写入PUML文件失败", zap.Error(err))
				return
			}
			compcontzap.GetDefault().Info("写入PUML文件成功", zap.String("puml_file", config.PUMLFile))
		}
		if config.PNGFile != "" {
			err = renderToPNG(buf.String(), config.PNGFile)
			if err != nil {
				compcontzap.GetDefault().Info("写入PNG文件失败", zap.Error(err))
				return
			}
			compcontzap.GetDefault().Info("写入PNG文件成功", zap.String("png_file", config.PNGFile))
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
