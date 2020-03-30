package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var Domain string
var Server string

type Body struct {
	URL string
}

func init() {
	crawlCmd.Flags().StringVarP(&Domain, "domain", "d", "", "Domain to crawl")
	crawlCmd.Flags().StringVarP(&Server, "server", "s", "", "URL for crawler server")
	crawlCmd.MarkFlagRequired("domain")
	crawlCmd.MarkFlagRequired("server")
	rootCmd.AddCommand(crawlCmd)
}

func doCrawl(cmd *cobra.Command, args []string) {

	body := &Body{
		URL: Domain,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/crawl", Server), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody))

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n--\n\n")
	fmt.Print("The spiders are now crawling your domain!\n")
	fmt.Print("Depending on the size of your site, this could take awhile.\n")
	fmt.Print("Please wait a bit and check back with\n")
	fmt.Print("\nosc get-domain <your domain>\n\n")
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Use the specified server to crawl a domain",
	Run:   doCrawl,
}
