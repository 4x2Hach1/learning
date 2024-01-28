// Code generated by goa v3.14.6, DO NOT EDIT.
//
// server HTTP client CLI support package
//
// Command:
// $ goa gen github.com/4x2Hach1/learning/next-go/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	serverc "github.com/4x2Hach1/learning/next-go/api/gen/http/server/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `server (hello|login|auth-user|users|user-by-id|new-user|update-user|memories|memory-by-id|new-memory|delete-memory|update-memory|new-heavy|check-heavy|delete-heavy)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` server hello --name "Doloribus velit rerum eveniet quaerat ut aut."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		serverFlags = flag.NewFlagSet("server", flag.ContinueOnError)

		serverHelloFlags    = flag.NewFlagSet("hello", flag.ExitOnError)
		serverHelloNameFlag = serverHelloFlags.String("name", "REQUIRED", "Name")

		serverLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		serverLoginBodyFlag = serverLoginFlags.String("body", "REQUIRED", "")

		serverAuthUserFlags     = flag.NewFlagSet("auth-user", flag.ExitOnError)
		serverAuthUserTokenFlag = serverAuthUserFlags.String("token", "REQUIRED", "")

		serverUsersFlags     = flag.NewFlagSet("users", flag.ExitOnError)
		serverUsersTokenFlag = serverUsersFlags.String("token", "REQUIRED", "")

		serverUserByIDFlags     = flag.NewFlagSet("user-by-id", flag.ExitOnError)
		serverUserByIDIDFlag    = serverUserByIDFlags.String("id", "REQUIRED", "ID")
		serverUserByIDTokenFlag = serverUserByIDFlags.String("token", "REQUIRED", "")

		serverNewUserFlags    = flag.NewFlagSet("new-user", flag.ExitOnError)
		serverNewUserBodyFlag = serverNewUserFlags.String("body", "REQUIRED", "")

		serverUpdateUserFlags     = flag.NewFlagSet("update-user", flag.ExitOnError)
		serverUpdateUserBodyFlag  = serverUpdateUserFlags.String("body", "REQUIRED", "")
		serverUpdateUserTokenFlag = serverUpdateUserFlags.String("token", "REQUIRED", "")

		serverMemoriesFlags     = flag.NewFlagSet("memories", flag.ExitOnError)
		serverMemoriesTokenFlag = serverMemoriesFlags.String("token", "REQUIRED", "")

		serverMemoryByIDFlags     = flag.NewFlagSet("memory-by-id", flag.ExitOnError)
		serverMemoryByIDIDFlag    = serverMemoryByIDFlags.String("id", "REQUIRED", "ID")
		serverMemoryByIDTokenFlag = serverMemoryByIDFlags.String("token", "REQUIRED", "")

		serverNewMemoryFlags     = flag.NewFlagSet("new-memory", flag.ExitOnError)
		serverNewMemoryBodyFlag  = serverNewMemoryFlags.String("body", "REQUIRED", "")
		serverNewMemoryTokenFlag = serverNewMemoryFlags.String("token", "REQUIRED", "")

		serverDeleteMemoryFlags     = flag.NewFlagSet("delete-memory", flag.ExitOnError)
		serverDeleteMemoryIDFlag    = serverDeleteMemoryFlags.String("id", "REQUIRED", "ID")
		serverDeleteMemoryTokenFlag = serverDeleteMemoryFlags.String("token", "REQUIRED", "")

		serverUpdateMemoryFlags     = flag.NewFlagSet("update-memory", flag.ExitOnError)
		serverUpdateMemoryBodyFlag  = serverUpdateMemoryFlags.String("body", "REQUIRED", "")
		serverUpdateMemoryIDFlag    = serverUpdateMemoryFlags.String("id", "REQUIRED", "ID")
		serverUpdateMemoryTokenFlag = serverUpdateMemoryFlags.String("token", "REQUIRED", "")

		serverNewHeavyFlags     = flag.NewFlagSet("new-heavy", flag.ExitOnError)
		serverNewHeavyTokenFlag = serverNewHeavyFlags.String("token", "REQUIRED", "")

		serverCheckHeavyFlags     = flag.NewFlagSet("check-heavy", flag.ExitOnError)
		serverCheckHeavyKeyFlag   = serverCheckHeavyFlags.String("key", "REQUIRED", "Key")
		serverCheckHeavyTokenFlag = serverCheckHeavyFlags.String("token", "REQUIRED", "")

		serverDeleteHeavyFlags     = flag.NewFlagSet("delete-heavy", flag.ExitOnError)
		serverDeleteHeavyKeyFlag   = serverDeleteHeavyFlags.String("key", "REQUIRED", "Key")
		serverDeleteHeavyTokenFlag = serverDeleteHeavyFlags.String("token", "REQUIRED", "")
	)
	serverFlags.Usage = serverUsage
	serverHelloFlags.Usage = serverHelloUsage
	serverLoginFlags.Usage = serverLoginUsage
	serverAuthUserFlags.Usage = serverAuthUserUsage
	serverUsersFlags.Usage = serverUsersUsage
	serverUserByIDFlags.Usage = serverUserByIDUsage
	serverNewUserFlags.Usage = serverNewUserUsage
	serverUpdateUserFlags.Usage = serverUpdateUserUsage
	serverMemoriesFlags.Usage = serverMemoriesUsage
	serverMemoryByIDFlags.Usage = serverMemoryByIDUsage
	serverNewMemoryFlags.Usage = serverNewMemoryUsage
	serverDeleteMemoryFlags.Usage = serverDeleteMemoryUsage
	serverUpdateMemoryFlags.Usage = serverUpdateMemoryUsage
	serverNewHeavyFlags.Usage = serverNewHeavyUsage
	serverCheckHeavyFlags.Usage = serverCheckHeavyUsage
	serverDeleteHeavyFlags.Usage = serverDeleteHeavyUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "server":
			svcf = serverFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "server":
			switch epn {
			case "hello":
				epf = serverHelloFlags

			case "login":
				epf = serverLoginFlags

			case "auth-user":
				epf = serverAuthUserFlags

			case "users":
				epf = serverUsersFlags

			case "user-by-id":
				epf = serverUserByIDFlags

			case "new-user":
				epf = serverNewUserFlags

			case "update-user":
				epf = serverUpdateUserFlags

			case "memories":
				epf = serverMemoriesFlags

			case "memory-by-id":
				epf = serverMemoryByIDFlags

			case "new-memory":
				epf = serverNewMemoryFlags

			case "delete-memory":
				epf = serverDeleteMemoryFlags

			case "update-memory":
				epf = serverUpdateMemoryFlags

			case "new-heavy":
				epf = serverNewHeavyFlags

			case "check-heavy":
				epf = serverCheckHeavyFlags

			case "delete-heavy":
				epf = serverDeleteHeavyFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "server":
			c := serverc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "hello":
				endpoint = c.Hello()
				data, err = serverc.BuildHelloPayload(*serverHelloNameFlag)
			case "login":
				endpoint = c.Login()
				data, err = serverc.BuildLoginPayload(*serverLoginBodyFlag)
			case "auth-user":
				endpoint = c.AuthUser()
				data, err = serverc.BuildAuthUserPayload(*serverAuthUserTokenFlag)
			case "users":
				endpoint = c.Users()
				data, err = serverc.BuildUsersPayload(*serverUsersTokenFlag)
			case "user-by-id":
				endpoint = c.UserByID()
				data, err = serverc.BuildUserByIDPayload(*serverUserByIDIDFlag, *serverUserByIDTokenFlag)
			case "new-user":
				endpoint = c.NewUser()
				data, err = serverc.BuildNewUserPayload(*serverNewUserBodyFlag)
			case "update-user":
				endpoint = c.UpdateUser()
				data, err = serverc.BuildUpdateUserPayload(*serverUpdateUserBodyFlag, *serverUpdateUserTokenFlag)
			case "memories":
				endpoint = c.Memories()
				data, err = serverc.BuildMemoriesPayload(*serverMemoriesTokenFlag)
			case "memory-by-id":
				endpoint = c.MemoryByID()
				data, err = serverc.BuildMemoryByIDPayload(*serverMemoryByIDIDFlag, *serverMemoryByIDTokenFlag)
			case "new-memory":
				endpoint = c.NewMemory()
				data, err = serverc.BuildNewMemoryPayload(*serverNewMemoryBodyFlag, *serverNewMemoryTokenFlag)
			case "delete-memory":
				endpoint = c.DeleteMemory()
				data, err = serverc.BuildDeleteMemoryPayload(*serverDeleteMemoryIDFlag, *serverDeleteMemoryTokenFlag)
			case "update-memory":
				endpoint = c.UpdateMemory()
				data, err = serverc.BuildUpdateMemoryPayload(*serverUpdateMemoryBodyFlag, *serverUpdateMemoryIDFlag, *serverUpdateMemoryTokenFlag)
			case "new-heavy":
				endpoint = c.NewHeavy()
				data, err = serverc.BuildNewHeavyPayload(*serverNewHeavyTokenFlag)
			case "check-heavy":
				endpoint = c.CheckHeavy()
				data, err = serverc.BuildCheckHeavyPayload(*serverCheckHeavyKeyFlag, *serverCheckHeavyTokenFlag)
			case "delete-heavy":
				endpoint = c.DeleteHeavy()
				data, err = serverc.BuildDeleteHeavyPayload(*serverDeleteHeavyKeyFlag, *serverDeleteHeavyTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// serverUsage displays the usage of the server command and its subcommands.
func serverUsage() {
	fmt.Fprintf(os.Stderr, `Server Service for front.
Usage:
    %[1]s [globalflags] server COMMAND [flags]

COMMAND:
    hello: Hello implements hello.
    login: Login implements login.
    auth-user: AuthUser implements authUser.
    users: Users implements users.
    user-by-id: UserByID implements userById.
    new-user: NewUser implements newUser.
    update-user: UpdateUser implements updateUser.
    memories: Memories implements memories.
    memory-by-id: MemoryByID implements memoryById.
    new-memory: NewMemory implements newMemory.
    delete-memory: DeleteMemory implements deleteMemory.
    update-memory: UpdateMemory implements updateMemory.
    new-heavy: NewHeavy implements newHeavy.
    check-heavy: CheckHeavy implements checkHeavy.
    delete-heavy: DeleteHeavy implements deleteHeavy.

Additional help:
    %[1]s server COMMAND --help
`, os.Args[0])
}
func serverHelloUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server hello -name STRING

Hello implements hello.
    -name STRING: Name

Example:
    %[1]s server hello --name "Doloribus velit rerum eveniet quaerat ut aut."
`, os.Args[0])
}

func serverLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server login -body JSON

Login implements login.
    -body JSON: 

Example:
    %[1]s server login --body '{
      "email": "Ea qui.",
      "password": "Et saepe non."
   }'
`, os.Args[0])
}

func serverAuthUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server auth-user -token STRING

AuthUser implements authUser.
    -token STRING: 

Example:
    %[1]s server auth-user --token "Est ut."
`, os.Args[0])
}

func serverUsersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server users -token STRING

Users implements users.
    -token STRING: 

Example:
    %[1]s server users --token "Non cupiditate nulla error consequatur."
`, os.Args[0])
}

func serverUserByIDUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server user-by-id -id INT -token STRING

UserByID implements userById.
    -id INT: ID
    -token STRING: 

Example:
    %[1]s server user-by-id --id 8257531658622185016 --token "Velit quam aperiam consectetur omnis necessitatibus quasi."
`, os.Args[0])
}

func serverNewUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server new-user -body JSON

NewUser implements newUser.
    -body JSON: 

Example:
    %[1]s server new-user --body '{
      "email": "A enim praesentium modi dicta cumque.",
      "name": "Autem tempore fugit.",
      "password": "At molestiae aliquam pariatur."
   }'
`, os.Args[0])
}

func serverUpdateUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server update-user -body JSON -token STRING

UpdateUser implements updateUser.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s server update-user --body '{
      "email": "Iure aliquid deserunt enim.",
      "name": "Molestiae nostrum."
   }' --token "Est minima."
`, os.Args[0])
}

func serverMemoriesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server memories -token STRING

Memories implements memories.
    -token STRING: 

Example:
    %[1]s server memories --token "Sit nihil tenetur placeat in."
`, os.Args[0])
}

func serverMemoryByIDUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server memory-by-id -id INT -token STRING

MemoryByID implements memoryById.
    -id INT: ID
    -token STRING: 

Example:
    %[1]s server memory-by-id --id 1226034219362355291 --token "Maxime molestias illum."
`, os.Args[0])
}

func serverNewMemoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server new-memory -body JSON -token STRING

NewMemory implements newMemory.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s server new-memory --body '{
      "date": "2004-11-09",
      "detail": "Reiciendis suscipit possimus eligendi.",
      "location": "Aspernatur repellendus veritatis.",
      "title": "Et et."
   }' --token "Est error quasi sint eius."
`, os.Args[0])
}

func serverDeleteMemoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server delete-memory -id INT -token STRING

DeleteMemory implements deleteMemory.
    -id INT: ID
    -token STRING: 

Example:
    %[1]s server delete-memory --id 3090224288635727171 --token "Quia molestiae."
`, os.Args[0])
}

func serverUpdateMemoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server update-memory -body JSON -id INT -token STRING

UpdateMemory implements updateMemory.
    -body JSON: 
    -id INT: ID
    -token STRING: 

Example:
    %[1]s server update-memory --body '{
      "date": "1974-01-04",
      "detail": "Nihil officia officiis dignissimos magni ea sed.",
      "location": "Harum sapiente dolor ut provident pariatur non.",
      "title": "Veniam id neque nam sunt eos vero."
   }' --id 3150585145125593148 --token "Sed earum cum eum doloribus."
`, os.Args[0])
}

func serverNewHeavyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server new-heavy -token STRING

NewHeavy implements newHeavy.
    -token STRING: 

Example:
    %[1]s server new-heavy --token "Unde enim et adipisci praesentium ratione nihil."
`, os.Args[0])
}

func serverCheckHeavyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server check-heavy -key STRING -token STRING

CheckHeavy implements checkHeavy.
    -key STRING: Key
    -token STRING: 

Example:
    %[1]s server check-heavy --key "Est aut aut quia aut." --token "Rerum rerum dicta numquam recusandae ea."
`, os.Args[0])
}

func serverDeleteHeavyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] server delete-heavy -key STRING -token STRING

DeleteHeavy implements deleteHeavy.
    -key STRING: Key
    -token STRING: 

Example:
    %[1]s server delete-heavy --key "Explicabo similique ratione beatae consequatur animi consequatur." --token "Aliquam voluptatem unde voluptate nihil qui."
`, os.Args[0])
}
