package cmd

import (
	"client/proto"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete account by username.",
	Long:  `Delete account by username.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		username, _ := cmd.Flags().GetString("username")

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		c := proto.NewAccountsClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err = c.DeleteAccount(ctx, &proto.DeleteAccountRequest{Username: username})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Deleted: %s \n", username)
	},
}

func init() {
	deleteCmd.Flags().StringP("username", "U", "", "Account username.")
	deleteCmd.Flags().StringP("port", "P", "8081", "Host's port.")
	deleteCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	rootCmd.AddCommand(deleteCmd)
	err := deleteCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
}
