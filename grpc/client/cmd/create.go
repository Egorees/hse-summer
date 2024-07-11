package cmd

import (
	"client/proto"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
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

		_, err = c.CreateAccount(ctx, &proto.CreateAccountRequest{Username: username, Amount: int32(amount)})
		if err != nil {
			if status.Code(err) == codes.AlreadyExists {
				fmt.Printf("error: username %s already exist\n", username)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("Created: %s \n", username)
	},
}

func init() {
	createCmd.Flags().StringP("username", "U", "", "Account username.")
	createCmd.Flags().IntP("amount", "A", 0, "Account amount.")
	createCmd.Flags().StringP("port", "P", "8081", "Host's port.")
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
