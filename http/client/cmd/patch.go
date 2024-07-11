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

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch account amount by username and new amount.",
	Long:  `Patch account amount by username and new amount.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		username, _ := cmd.Flags().GetString("username")
		amount, _ := cmd.Flags().GetInt("amount")

		postBody, _ := json.Marshal(dto.PatchAccountRequest{
			Username: username,
			Amount:   amount,
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(fmt.Sprintf("http://%s:%s/account/patch", host, port), "application/json", responseBody)

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
	patchCmd.Flags().StringP("username", "U", "", "Account username.")
	patchCmd.Flags().IntP("amount", "A", 0, "Account amount.")
	patchCmd.Flags().StringP("port", "P", "8080", "Host's port.")
	patchCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	rootCmd.AddCommand(patchCmd)

	err := patchCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
	err = patchCmd.MarkFlagRequired("amount")
	if err != nil {
		log.Fatalln(err)
	}
}
