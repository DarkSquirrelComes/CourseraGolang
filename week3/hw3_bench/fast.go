package main

import (
	"bufio"
	json "encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

//easyjson:json
type userStruct struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3486653aDecodeCourseraWeek3Hw3Bench(in *jlexer.Lexer, out *userStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3486653aEncodeCourseraWeek3Hw3Bench(out *jwriter.Writer, in userStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v userStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3486653aEncodeCourseraWeek3Hw3Bench(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v userStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3486653aEncodeCourseraWeek3Hw3Bench(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *userStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeCourseraWeek3Hw3Bench(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *userStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3486653aDecodeCourseraWeek3Hw3Bench(l, v)
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""
	users := make([]userStruct, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//    fmt.Println(scanner.Text())
		//}
		//for _, line := range lines {
		//line := scanner.Text()
		user := &userStruct{}
		err := user.UnmarshalJSON([]byte(scanner.Text()))
		if err != nil {
			fmt.Println("Unmarshling error")
		}
		if err != nil {
			panic(err)
		}
		users = append(users, *user)
	}

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers
		for _, browserRaw := range browsers {
			browser := browserRaw
			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browserRaw := range browsers {
			browser := browserRaw
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

func main() {
	return
}
