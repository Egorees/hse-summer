package cmd

import (
	"bytes"
	"client/dto"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete account by username.",
	Long:  `Delete account by username.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		username, _ := cmd.Flags().GetString("username")

		postBody, _ := json.Marshal(dto.DeleteAccountRequest{
			Username: username,
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(fmt.Sprintf("http://%s:%s/account/delete", host, port), "application/json", responseBody)

		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	},
}

func init() {
	deleteCmd.Flags().StringP("username", "U", "", "Account username.")
	deleteCmd.Flags().StringP("port", "P", "8080", "Host's port.")
	deleteCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	rootCmd.AddCommand(deleteCmd)
	err := deleteCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
}
