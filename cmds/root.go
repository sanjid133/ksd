package cmds

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/appscode/go/analytics"
	v "github.com/appscode/go/version"
	"github.com/jpillora/go-ogle-analytics"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)


func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{

	}


	return cmd
}