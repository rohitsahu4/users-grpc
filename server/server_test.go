package main

import (
	"context"
	"fmt"
	"testing"
	usr "users-grpc/uproto"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/golang/protobuf/ptypes"
)

var c usr.UserProfilesClient

func TestGetUserProfile(t *testing.T) {

	tests := []struct {
		in   *usr.GetUserProfileRequest
		want string
	}{
		{
			in:   &usr.GetUserProfileRequest{Id: "879e1d93-db7c-4a21-bb14-8b6b5ce08d2c"},
			want: `id:"879e1d93-db7c-4a21-bb14-8b6b5ce08d2c" email:"Rohit@appointy.com" first_name:"raman" last_name:"Sahu" birth_date:<seconds:1560763023 nanos:685608000 > `,
		},
		{
			in:   &usr.GetUserProfileRequest{Id: "dsss0c00e69-274b-4cbe-939c-df2e3244bb8e"},
			want: `sql: no rows in result set`,
		},
	}

	ss := server{}

	for _, tt := range tests {
		response, err := ss.GetUserProfile(context.Background(), tt.in)
		if err != nil {
			if fmt.Sprint(err) != tt.want {
				t.Errorf("GetUserProfile(%v)=%v, wanted %v", tt.in, fmt.Sprint(err), tt.want)
			}

		} else {
			if fmt.Sprint(response) != tt.want {
				t.Errorf("GetUserProfile(%v)=%v, wanted %v", tt.in, response, tt.want)
			}
		}
	}

}

func TestCreateUserProfile(t *testing.T) {
	btime := ptypes.TimestampNow()

	phones := []string{"+91 96440 36993", "+91 96440 36993"}
	tests := []struct {
		in   *usr.CreateUserProfileRequest
		want string
	}{
		{
			in:   &usr.CreateUserProfileRequest{UserProfile: &usr.UserProfile{FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}},
			want: ``,
		},
	}

	ss := server{}

	for _, tt := range tests {
		response, err := ss.CreateUserProfile(context.Background(), tt.in)
		if err != nil {
			if fmt.Sprint(err) != tt.want {
				t.Errorf("CreateUserProfile(%v)=%v, wanted %v", tt.in, fmt.Sprint(err), tt.want)
			}
		}
		_, errrr := ss.DeleteUserProfile(context.Background(), &usr.DeleteUserProfileRequest{Id: response.Id})
		if errrr != nil {
			t.Error(errrr)
		}
	}

}

func TestUpdateUserProfile(t *testing.T) {

	btime := ptypes.TimestampNow()
	phones := []string{"+91 96440 36993", "+91 96440 36993"}

	ss := server{}

	rr, err := ss.CreateUserProfile(context.Background(), &usr.CreateUserProfileRequest{UserProfile: &usr.UserProfile{FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}})
	if err != nil {
		t.Error(err)
	}
	id := rr.Id

	tests := []struct {
		in   *usr.UpdateUserProfileRequest
		want string
	}{
		{
			in:   &usr.UpdateUserProfileRequest{UserProfile: &usr.UserProfile{Id: id, FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}},
			want: fmt.Sprint(&usr.UserProfile{Id: id, FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}),
		},
		{
			in:   &usr.UpdateUserProfileRequest{UserProfile: &usr.UserProfile{Id: "asdfasdfjkasdfkljasdflkjh", FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}},
			want: "id not found!",
		},
	}

	for _, tt := range tests {
		response, err := ss.UpdateUserProfile(context.Background(), tt.in)
		if err != nil {
			if fmt.Sprint(err) != fmt.Sprint(tt.want) {
				t.Errorf("CreateUserProfile(%v)=%v, wanted %v", tt.in, fmt.Sprint(err), tt.want)
			}
		} else {
			if fmt.Sprint(response) != fmt.Sprint(tt.want) {
				t.Errorf("GetUserProfile(%v)=%v, wanted %v", tt.in, response, tt.want)
			}
		}

	}
	_, errrr := ss.DeleteUserProfile(context.Background(), &usr.DeleteUserProfileRequest{Id: id})
	if errrr != nil {
		t.Error(errrr)
	}

}
func TestDeleteUserProfile(t *testing.T) {

	btime := ptypes.TimestampNow()
	phones := []string{"+91 96440 36993", "+91 96440 36993"}

	ss := server{}

	rr, err := ss.CreateUserProfile(context.Background(), &usr.CreateUserProfileRequest{UserProfile: &usr.UserProfile{FirstName: "raman", LastName: "Sahu", Email: "Rohit@appointy.com", BirthDate: btime, Telephones: phones}})
	if err != nil {
		t.Error(err)
	}
	id := rr.Id

	tests := []struct {
		in   *usr.DeleteUserProfileRequest
		want string
	}{
		{
			in:   &usr.DeleteUserProfileRequest{Id: id},
			want: fmt.Sprint(&empty.Empty{}),
		},
		{
			in:   &usr.DeleteUserProfileRequest{Id: "asdfkhasjdflhjasdlfkjh"},
			want: "id not found!",
		},
	}

	for _, tt := range tests {
		response, err := ss.DeleteUserProfile(context.Background(), tt.in)
		if err != nil {
			if fmt.Sprint(err) != fmt.Sprint(tt.want) {
				t.Errorf("DeleteUserProfile(%v)=%v, wanted %v", tt.in, fmt.Sprint(err), tt.want)
			}
		} else {
			if fmt.Sprint(response) != fmt.Sprint(tt.want) {
				t.Errorf("DeleteUserProfile(%v)=%v, wanted %v", tt.in, response, tt.want)
			}
		}

	}

}

func TestListUsersProfiles(t *testing.T) {

	tests := []struct {
		in   *usr.ListUsersProfilesRequest
		want string
	}{
		{
			in:   &usr.ListUsersProfilesRequest{Query: "raman"},
			want: `profiles:<id:"879e1d93-db7c-4a21-bb14-8b6b5ce08d2c" email:"Rohit@appointy.com" first_name:"raman" last_name:"Sahu" birth_date:<seconds:1560763023 nanos:685608000 > > profiles:<id:"f735f0d2-df9a-43b9-b6ff-861b1a4996ba" first_name:"raman" last_name:"Sahu" birth_date:<> > `,
		},
		{
			in:   &usr.ListUsersProfilesRequest{Query: "asdfasdf"},
			want: ``,
		},
	}

	ss := server{}

	for _, tt := range tests {
		response, err := ss.ListUsersProfiles(context.Background(), tt.in)
		// fmt.Println(response)
		if err != nil {
			if fmt.Sprint(err) != tt.want {
				t.Errorf("GetUserProfile(%v)=%v, wanted %v", tt.in, fmt.Sprint(err), tt.want)
			}

		} else {
			if fmt.Sprint(response) != tt.want {
				t.Errorf("GetUserProfile(%v)=%v, wanted %v", tt.in, response, tt.want)
			}
		}
	}

}
