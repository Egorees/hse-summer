package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account by username.",
	Long:  `Get account by username.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		username, _ := cmd.Flags().GetString("username")

		resp, err := http.Get(fmt.Sprintf("http://%s:%s/account?username=%s", host, port, username))
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		log.Printf(sb)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("username", "U", "", "Account username.")
	getCmd.Flags().StringP("port", "P", "8080", "Host's port.")
	getCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	err := getCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
}
