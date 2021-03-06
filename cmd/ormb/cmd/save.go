/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/caicloud/ormb/pkg/ormb"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Store a copy of model in local registry cache",
	Long: `Store a copy of model in local registry cache.

Note: modifying the model after this operation will
not change the item as it exists in the cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO(gaocegege): Validate.
		o, err := ormb.NewDefaultOCIormb()
		if err != nil {
			panic(err)
		}

		if err := o.Save(args[0], args[1]); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
