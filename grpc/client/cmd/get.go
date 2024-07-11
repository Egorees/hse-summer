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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account by username.",
	Long:  `Get account by username.`,
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

		resp, err := c.GetAccount(ctx, &proto.GetAccountRequest{Username: username})
		if err != nil {
			if status.Code(err) == codes.Unknown {
				fmt.Printf("error: username %s does not exist\n", username)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("username", "U", "", "Account username.")
	getCmd.Flags().StringP("port", "P", "8081", "Host's port.")
	getCmd.Flags().StringP("host", "H", "0.0.0.0", "Host address")
	err := getCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalln(err)
	}
}
