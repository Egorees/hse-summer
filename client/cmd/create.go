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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new account by username and amount.",
	Long:  `Create new account by username and amount.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		username, _ := cmd.Flags().GetString("username")
		amount, _ := cmd.Flags().GetInt("amount")

		postBody, _ := json.Marshal(dto.CreateAccountRequest{
			Username: username,
			Amount:   amount,
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(fmt.Sprintf("http://%s:%s/account/create", host, port), "application/json", responseBody)

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
	createCmd.Flags().StringP("username", "U", "", "Account username.")
	createCmd.Flags().IntP("amount", "A", 0, "Account amount.")
	createCmd.Flags().StringP("port", "P", "8080", "Host's port.")
	createCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	rootCmd.AddCommand(createCmd)
	err := createCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
	err = createCmd.MarkFlagRequired("amount")
	if err != nil {
		log.Fatalln(err)
	}
}
