package tools

import (
	"github.com/spf13/cobra"
)

type FileFlags struct {
	Files     []string
	Recursive bool
	Sort      bool
}

func (s *FileFlags) Set(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&s.Files, "file", "f", nil, "Set file (format: /tmp/foo, https://..., -) (can repeat)")
	cmd.Flags().BoolVarP(&s.Recursive, "recursive", "R", true, "Process directory used in -f recursively (deprecated; set to true by default)")
	cmd.Flags().BoolVar(&s.Sort, "sort", true, "Sort by namespace, name, etc.")
}

type FileFlags2 struct {
	Files     []string
	Recursive bool
}

func (s *FileFlags2) Set(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(&s.Files, "file2", nil, "Set second file (format: /tmp/foo, https://..., -) (can repeat)")
	cmd.Flags().BoolVar(&s.Recursive, "file2-recursive", true, "Process directory used in --file2 recursively (deprecated; set to true by default)")
}
