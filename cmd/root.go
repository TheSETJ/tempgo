/*
Copyright © 2025 Sayyed Ehsan Taheri Javdi ehsantaheri25@yahoo.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tempgo",
	Short: "A tool to work with temperature",
}

var convertCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Convert temperature between Celsius, Fahrenheit and Kelvin",
	Example: "convert <temperature> <from> <to>",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var temperature float64
		var from, to string

		temperature, _ = strconv.ParseFloat(args[0], 32)
		from = args[1]
		to = args[2]

		if from == to {
			fmt.Println(temperature)
			os.Exit(0)
		}

		if from == "C" && to == "F" {
			fmt.Println(temperature*9.0/5 + 32)
			os.Exit(0)
		}

		if from == "C" && to == "K" {
			fmt.Println(temperature + 273.15)
			os.Exit(0)
		}

		if from == "F" && to == "C" {
			fmt.Println((temperature - 32) * 5.0 / 9)
			os.Exit(0)
		}

		if from == "F" && to == "K" {
			fmt.Println(273.15 + (temperature-32)/1.8)
			os.Exit(0)
		}

		if from == "K" && to == "C" {
			fmt.Println(temperature - 273.15)
			os.Exit(0)
		}

		if from == "K" && to == "F" {
			fmt.Println((temperature-273.15)*1.8 + 32)
			os.Exit(0)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
