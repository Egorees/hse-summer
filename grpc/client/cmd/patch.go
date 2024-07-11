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

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch account amount by username and new amount.",
	Long:  `Patch account amount by username and new amount.`,
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

		_, err = c.PatchAccount(ctx, &proto.PatchAccountRequest{Username: username, Amount: int32(amount)})
		if err != nil {
			if status.Code(err) == codes.Unknown {
				fmt.Printf("error: username %s does not exist\n", username)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("Patched: %s \n", username)
	},
}

func init() {
	patchCmd.Flags().StringP("username", "U", "", "Account username.")
	patchCmd.Flags().IntP("amount", "A", 0, "Account amount.")
	patchCmd.Flags().StringP("port", "P", "8081", "Host's port.")
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
