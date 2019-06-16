package main

import (
	"log"
	usr "users/uproto"

	"golang.org/x/net/context"
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

	// btime := ptypes.TimestampNow()

	// phones := []string{"+91 96440 36993", "+91 96440 36993"}
	// response, err := c.UpdateUserProfile(context.Background(), &usr.UpdateUserProfileRequest{UserProfile: &usr.UserProfile{Id: "0affdaf9-1035-42a8-bf34-79fded940abb", FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %s", response)

	response, err := c.ListUsersProfiles(context.Background(), &usr.ListUsersProfilesRequest{Query: "hit@appointy.com"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response)
}
