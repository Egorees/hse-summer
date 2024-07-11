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

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change username by lastname and new-name.",
	Long:  `Change username by lastname and new-name.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		lastName, _ := cmd.Flags().GetString("last-name")
		newName, _ := cmd.Flags().GetString("new-name")

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
			return
		}

		defer func() {
			_ = conn.Close()
		}()

		c := proto.NewAccountsClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err = c.ChangeAccount(ctx, &proto.ChangeAccountRequest{LastName: lastName, NewName: newName})
		if err != nil {
			switch status.Code(err) {
			case codes.AlreadyExists:
				fmt.Printf("error: username %s already exist\n", newName)
			case codes.Unknown:
				fmt.Printf("error: username %s does not exist\n", lastName)
			default:
				fmt.Println(err)
			}
			return
		}
		fmt.Printf("Changed: %s->%s \n", lastName, newName)
	},
}

func init() {
	changeCmd.Flags().StringP("last-name", "L", "", "Account last name.")
	changeCmd.Flags().StringP("new-name", "N", "", "Account new name.")
	changeCmd.Flags().StringP("port", "P", "8081", "Host's port.")
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
