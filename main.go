package main

import "github.com/TinyMurky/search-restaurant/cmd"

func main() {
	rootCmd := cmd.NewRootCmd()
	rootCmd.Execute()
}
