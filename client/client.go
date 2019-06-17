package main

import (
	"context"
	"log"
	usr "users-grpc/uproto"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := usr.NewUserProfilesClient(conn)

	//////////CreateUserProfile

	btime := ptypes.TimestampNow()

	phones := []string{"+91 96440 36993", "+91 96440 36993"}
	response, err := c.CreateUserProfile(context.Background(), &usr.CreateUserProfileRequest{UserProfile: &usr.UserProfile{FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}})
	if err != nil {
		log.Fatalf("Error : %s", err)
	}
	log.Printf("%s", response)

	////////GetUserProfile

	// response, err := c.GetUserProfile(context.Background(), &usr.GetUserProfileRequest{Id:""})
	// if err != nil {
	// 	log.Fatalf("Error : %s", err)
	// }
	// log.Printf("%s", response)

	////////////DeleteUserProfile

	// response, err := c.DeleteUserProfile(context.Background(), &usr.DeleteUserProfileRequest{Id:""})
	// if err != nil {
	// log.Fatalf("Error : %s", err)
	// }
	// log.Printf("%s", response)

	////////////UPDATE USR PROFILE
	// btime := ptypes.TimestampNow()

	// phones := []string{"+91 96440 36993", "+91 96440 36993"}
	// response, err := c.UpdateUserProfile(context.Background(), &usr.UpdateUserProfileRequest{UserProfile: &usr.UserProfile{Id: "0affdaf9-1035-42a8-bf34-79fded940abb", FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}})
	// if err != nil {
	// log.Fatalf("Error : %s", err)
	// }
	// log.Printf("%s", response)

	/////////////LIST USER PROFILE

	// response, err := c.ListUsersProfiles(context.Background(), &usr.ListUsersProfilesRequest{Query: "hit@appointy.com"})
	// if err != nil {
	//  	// log.Fatalf("Error : %s", err)
	// }
	// log.Printf("%s", response)
}
