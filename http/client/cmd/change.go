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

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change username by lastname and new-name.",
	Long:  `Change username by lastname and new-name.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		lastName, _ := cmd.Flags().GetString("last-name")
		newName, _ := cmd.Flags().GetString("new-name")

		postBody, _ := json.Marshal(dto.ChangeAccountRequest{
			LastName: lastName,
			NewName:  newName,
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(fmt.Sprintf("http://%s:%s/account/change", host, port), "application/json", responseBody)

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
	changeCmd.Flags().StringP("last-name", "L", "", "Account last name.")
	changeCmd.Flags().StringP("new-name", "N", "", "Account new name.")
	changeCmd.Flags().StringP("port", "P", "8080", "Host's port.")
	changeCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	rootCmd.AddCommand(changeCmd)
	err := changeCmd.MarkFlagRequired("last-name")
	if err != nil {
		log.Fatalln(err)
	}
	err = changeCmd.MarkFlagRequired("new-name")
	if err != nil {
		log.Fatalln(err)
	}
}
