package main

import (
	"database/sql"
	"time"

	// "time"
	"net"

	"github.com/golang/protobuf/ptypes"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	// "github.com/golang/protobuf/proto"
	"fmt"
	usr "users-grpc/uproto"

	"github.com/lib/pq"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

const (
	port     = ":7777"
	host     = "localhost"
	dbport   = 5432
	user     = "postgres"
	password = "#server444#"
	dbname   = "users"
)

var db *sql.DB

type server struct {
}

func (s *server) CreateUserProfile(c context.Context, req *usr.CreateUserProfileRequest) (*usr.UserProfile, error) {

	p := req.UserProfile
	timep, err := ptypes.Timestamp(p.BirthDate)
	if err != nil {
		return nil, err
	}
	u1, err := uuid.NewV4()
	if err != nil {

		return nil, err
	}

	sqlStatement := `
	INSERT INTO users (id,email, first_name, last_name,birth_date,telephones)
	VALUES ($1, $2, $3, $4 , $5, $6 )`
	_, err = db.Exec(sqlStatement, u1, p.Email, p.FirstName, p.LastName, timep, pq.Array(p.Telephones))
	if err != nil {
		return nil, err
	}
	req.UserProfile.Id = u1.String()
	return req.UserProfile, nil

}

func (s *server) GetUserProfile(c context.Context, req *usr.GetUserProfileRequest) (*usr.UserProfile, error) {

	id := req.Id
	res := usr.UserProfile{}
	times := time.Now()

	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&res.Id, &res.Email, &res.FirstName, &res.LastName, &times, pq.Array(res.Telephones))
	if err != nil {
		return nil, err
	}
	timeinbuf, err := ptypes.TimestampProto(times)
	if err != nil {
		return nil, err
	}
	res.BirthDate = timeinbuf

	return &res, nil

}

func (s *server) UpdateUserProfile(c context.Context, req *usr.UpdateUserProfileRequest) (*usr.UserProfile, error) {

	p := req.UserProfile
	timep, _ := ptypes.Timestamp(p.BirthDate)

	sqlStatement := `
	UPDATE  users 
	SET
		email = $1,
		first_name = $2,
		last_name = $3,
		birth_date = $4,
		telephones = $5
	WHERE 
		id =  $6`
	res, err := db.Exec(sqlStatement, p.Email, p.FirstName, p.LastName, timep, pq.Array(p.Telephones), p.Id)
	if err != nil {
		return nil, err
	}
	nrows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if nrows == 0 {
		return nil, fmt.Errorf("id not found!")
	}
	return req.UserProfile, nil

}

func (s *server) DeleteUserProfile(c context.Context, req *usr.DeleteUserProfileRequest) (*empty.Empty, error) {
	id := req.Id

	sqlStatement := `
	DELETE FROM users
	WHERE id=$1;`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	nrows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if nrows == 0 {
		return nil, fmt.Errorf("id not found!")
	}
	em := &empty.Empty{}
	return em, nil

}

func (s *server) ListUsersProfiles(c context.Context, req *usr.ListUsersProfilesRequest) (*usr.ListUsersProfilesResponse, error) {

	query := req.Query
	res := usr.ListUsersProfilesResponse{}
	sqlStatement := `
	SELECT * 
FROM users 
WHERE email LIKE $1 OR 
      first_name LIKE $1 OR  
      last_name LIKE $1 
	`

	// nos := []string{}
	rows, err := db.Query(sqlStatement, fmt.Sprintf("%%%s%%", query))
	if err != nil {

		return nil, err
	}
	times := time.Now()
	for rows.Next() {

		usrp := usr.UserProfile{}
		rows.Scan(&usrp.Id, &usrp.Email, &usrp.FirstName, &usrp.LastName, &times, pq.Array(usrp.Telephones))
		timeinbuf, err := ptypes.TimestampProto(times)
		if err != nil {
			return nil, err
		}
		usrp.BirthDate = timeinbuf

		res.Profiles = append(res.Profiles, &usrp)
	}

	return &res, nil

}

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {

		panic(err)
	}

}
func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)

	}
	s := grpc.NewServer()

	usr.RegisterUserProfilesServer(s, &server{})

	s.Serve(lis)

}
