package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	filePath = "./dataset.xml"
)

var (
	cases = []SearchRequest{
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      0,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      0,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      -1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      26,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     -1,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "wait error plz",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
		SearchRequest{
			Limit:      1,
			Offset:     0,
			Query:      "1a6fa827-62f1-45f6-b579-aaead2b47169",
			OrderField: "id",
			OrderBy:    0,
		},
	}

	testClients = []SearchClient{
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1488",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "",
		},
		SearchClient{
			AccessToken: "1337",
			URL:         "nil",
		},
		SearchClient{
			AccessToken: "666",
			URL:         "",
		},
		SearchClient{
			AccessToken: "zhopa1",
			URL:         "",
		},
		SearchClient{
			AccessToken: "zhopa2",
			URL:         "",
		},
		SearchClient{
			AccessToken: "zhopa3",
			URL:         "",
		},
		SearchClient{
			AccessToken: "zhopa4",
			URL:         "",
		},
	}

	testResults = []SearchResponse{
		SearchResponse{
			Users: []User{
				User{
					Id:   0,
					Name: "Boyd Wolf",
					Age:  22,
					About: `Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.
`,
					Gender: "male",
				},
			},
			NextPage: false,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users: []User{
				User{
					Id:   0,
					Name: "Boyd Wolf",
					Age:  22,
					About: `Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.
`,
					Gender: "male",
				},
			},
			NextPage: false,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
		SearchResponse{
			Users:    []User{},
			NextPage: true,
		},
	}
)

type Row struct {
	ID            int    `xml:"id"`
	GUID          string `xml:"guid"`
	IsActive      string `xml:"isActive"`
	Balance       string `xml:"balance"`
	Picture       string `xml:"picture"`
	Age           int    `xml:"age"`
	EyeColor      string `xml:"eyeColor"`
	FirstName     string `xml:"first_name"`
	LastName      string `xml:"last_name"`
	Gender        string `xml:"gender"`
	Company       string `xml:"company"`
	Email         string `xml:"email"`
	Phone         string `xml:"phone"`
	Address       string `xml:"adresses"`
	About         string `xml:"about"`
	Registered    string `xml:"register"`
	FavoriteFruit string `xml:"favoriteFruit"`
}

type Rows struct {
	Root []Row `xml:"row"`
}

func match(r Row, query string) bool {
	v := reflect.ValueOf(r)

	for i := 0; i < v.NumField(); i++ {
		if strings.Contains(v.Field(i).String(), query) {
			return true

		}
	}
	return false
}

func compare(r1, r2 Row, orderField, orderBy string) bool {
	s1 := reflect.ValueOf(r1).FieldByName(orderField).String()
	s2 := reflect.ValueOf(r2).FieldByName(orderField).String()

	return (s1 < s2) != (orderBy == "1")
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	query := r.FormValue("query")
	orderField := r.FormValue("order_field")
	orderBy := r.FormValue("order_by")
	accessToken := r.Header.Get("AccessToken")

	if query == "wait error plz" {
		time.Sleep(time.Second)
	}

	//fmt.Println(offset, limit)
	switch accessToken {
	case "1488":
		w.WriteHeader(http.StatusUnauthorized)
		return
	case "666":
		w.WriteHeader(http.StatusInternalServerError)
		return
	case "zhopa1":
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "{}")
		return

	case "zhopa2":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "{[}]")
		return
	case "zhopa3":
		//fmt.Println("Zhopa3 is here!")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error": "ErrorBadOrderField"}`)
		return
	case "zhopa4":
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "{[}]")
		return
	}

	if accessToken != "1337" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if orderBy != "0" && orderBy != "1" && orderBy != "-1" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := Rows{}
	//var d interface{}
	file, _ := os.Open(filePath)
	xmlData, _ := ioutil.ReadAll(file)
	//fmt.Println(string(xmlData))
	err := xml.Unmarshal(xmlData, &data) //no error handling!
	//err := xml.Unmarshal(xmlData, &d)
	//fmt.Println("Unmarshaled data:", data)
	if err != nil {
		//fmt.Println("Error in unmarshaling!")
		panic(err)
	}
	//fmt.Println(data.root[0])
	result := Rows{}

	for _, singleRow := range data.Root {
		if match(singleRow, query) {
			result.Root = append(result.Root, singleRow)
		}
	}

	if orderBy != "0" {
		sort.SliceStable(result, func(i, j int) bool {
			return compare(result.Root[i], result.Root[j], orderField, orderBy)
		})
	}
	//make slice!!!
	//fmt.Println(len(result.Root), offset, limit)
	if offset >= len(result.Root) {
		offset = len(result.Root) - 1
	}
	if offset+limit >= len(result.Root) {
		result.Root = result.Root[offset:]
	} else {
		result.Root = result.Root[offset : offset+limit]
	}
	//fmt.Println(result.Root)
	j, _ := json.MarshalIndent(rootToUsers(result.Root), "\t", "")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(j))
	//fmt.Println(string(j))
}

func rootToUsers(rowsToConvert []Row) []User {
	result := make([]User, 0, len(rowsToConvert))

	for _, r := range rowsToConvert {
		result = append(result, User{
			Id:     r.ID,
			Name:   r.FirstName + " " + r.LastName,
			Age:    r.Age,
			About:  r.About,
			Gender: r.Gender,
		})
	}
	return result
}

func TestFirstTry(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	for i, item := range cases {
		fmt.Println("Test №", i)
		sc := testClients[i]
		if sc.URL == "" {
			sc.URL = ts.URL
		}
		usersRes, err := sc.FindUsers(item)
		if err != nil {
			//fmt.Println("FindUsers returns Error!")
			continue
		}
		//fmt.Println(usersRes.Users)
		//fmt.Println(testResults[i])
		//fmt.Println(testResults[i].Users[0])
		if len(usersRes.Users) != len(testResults[i].Users) {
			fmt.Println("Different number of found users")
			t.FailNow()
		}
		for j := 0; j < len(testResults[i].Users); j++ {
			if usersRes.Users[j] != testResults[i].Users[j] {
				fmt.Println("Different users i, j", i, j)
				//fmt.Println(usersRes.Users[0])
				//fmt.Println(testResults[j].Users[0])
				t.Fail()
			}
		}
		if usersRes.NextPage != testResults[i].NextPage {
			t.Fail()
		}
	}
	ts.Close()
}

func main() {
	return
}
