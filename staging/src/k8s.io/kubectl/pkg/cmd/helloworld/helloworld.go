package helloworld

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type HelloWorldOptions struct {
	genericclioptions.IOStreams
}

// NewHelloWorldOptions creates the options for hello-world
func NewHelloWorldOptions(ioStreams genericclioptions.IOStreams) *HelloWorldOptions {
	return &HelloWorldOptions{
		IOStreams: ioStreams,
	}
}

func (o *HelloWorldOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	return nil
}

func (o HelloWorldOptions) Validate() error {
	return nil
}

func (o HelloWorldOptions) RunHelloWorld() error {
	fmt.Fprint(o.Out, "Hello World!")
	return nil
}

func NewCmdHelloWorld(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := NewHelloWorldOptions(ioStreams)
	cmd := &cobra.Command{
		Use: "hello-world",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(f, cmd, args))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.RunHelloWorld())
		},
	}
	return cmd
}
