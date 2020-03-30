package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var getDomainCmd = &cobra.Command{
	Use:   "get-domain",
	Short: "Get information about a crawled domain",
	Run:   doGetDomain,
	Args:  cobra.ExactArgs(1),
}

func doGetDomain(cmd *cobra.Command, args []string) {
	domain, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	resp, err := http.Get(fmt.Sprintf("%s/domains/%s", Server, domain))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}

func init() {
	getDomainCmd.Flags().StringVarP(&Server, "server", "s", "", "URL for crawler server")
	getDomainCmd.MarkFlagRequired("server")
	rootCmd.AddCommand(getDomainCmd)
}
