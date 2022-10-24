package Test

import (
	"fmt"
	"strconv"
	"testing"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"

	uuid "github.com/satori/go.uuid"
)

var tests = []handler.User {
	{FirstName: "", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "", Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},
}

func TestInsertUser(t *testing.T) {
	randEmail := "insertUSer@"+uuid.NewV4().String()
	t.Run("Insert valid user to DB", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}
		sampleUser := &handler.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := DB.InsertUser(*sampleUser)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		rows, err := DB.DB.Query(`SELECT * FROM User WHERE Email = ?`, sampleUser.Email)
		var userId, sessionId, firstName, lastName, nickName, email, DOB, avatar, aboutMe, createdAt, isLoggedIn, isPublic, numFollowers, numFollowing, numPosts, password string
		var resultUser *handler.User

		for rows.Next() {
			rows.Scan(&userId, &sessionId, &firstName, &lastName, &nickName, &email, &DOB, &avatar, &aboutMe, &createdAt, &isLoggedIn, &isPublic, &numFollowers, &numFollowing, &numPosts, &password)

			resultUser = &handler.User{
				UserId:      "-",
				SessionId:   sessionId,
				FirstName:   firstName,
				LastName:    lastName,
				NickName:    nickName,
				Email:       email,
				DateOfBirth: DOB,
				Avatar:      avatar,
				AboutMe:     aboutMe,
				CreatedAt:   "-",
				Password:    password,
			}
		}

		sampleUser.Password = strconv.FormatBool(handler.CheckPasswordHash(sampleUser.Password, resultUser.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		resultUser.Password = "true"
		want := sampleUser
		got := resultUser

		if *got != *want {
			t.Errorf("want %v, \n got %v", want, got)
		}
	})
	t.Run("inserting a user wiht used email to the db", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}
		sampleUser := &handler.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := DB.InsertUser(*sampleUser)
		if err == nil {
			t.Errorf("Error Catching already used email %v", err)
		}
	})
	t.Run("Check the length of neccesary values cant be 0", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}
		sampleUser := &handler.User{
			FirstName: "", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
			DateOfBirth: "Length", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "Length", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := DB.InsertUser(*sampleUser)
		if err == nil {
			t.Errorf("Error Catching already used empty values %v", err)
		}
	})
	t.Run("Check the length of neccesary values cant be 0", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")
		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}

		for i, s := range tests {
			err := DB.InsertUser(s)
			fmt.Println("Index:", i, "CURRENT ", s)
			fmt.Println()
		if err == nil {
			t.Errorf("Error Catching empty values %v", err)
		}
		}
	})
}

func TestGetUser(t *testing.T) {
	testEmail := "GetUser@" + uuid.NewV4().String()
	t.Run("Check with correct credentials", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}

		sampleUser := &handler.User{
			FirstName: "GetUser", LastName: "GetUser", NickName: "GetUser", Email: testEmail, Password: "GetUser",
			DateOfBirth: "GetUser", AboutMe: "GetUser", Avatar: "GetUser", CreatedAt: "GetUser", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		err := DB.InsertUser(*sampleUser)

		if err != nil {
			t.Errorf("Error Inserting the test user to the db")
		}

		//Now check if the value is input to the db
		wantStr := "Valid Credentials"
		wantBool := true

		gotBool, gotStr := DB.GetUser(sampleUser.Email, sampleUser.Password)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}

	})
}
