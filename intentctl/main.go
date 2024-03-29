package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gabiSmachado/intents/client"
	"github.com/gabiSmachado/intents/datamodel"
	//"golang.org/x/exp/slog"
)

var (
	_flagURI   = flag.String("uri", "smoapp-service.smo.svc.cluster.local:3000", "server uri")
	_flagDebug = flag.Bool("debug", false, "enable debugging log")
	_client    client.Client
)

func CreateIntent(args []string) error {
	//createIntentCmd := flag.NewFlagSet("create", flag.PanicOnError)
	//flagEvent := createIntentCmd.String("event", "undefined", "mnemonic for the intent")
	var name, ricId, serviceID, description string
	var policyId, policyTypeId int
	var intent datamodel.Intent

	fmt.Println("Name:")
	fmt.Scanln(&name)

	fmt.Println("Description:")
	fmt.Scanln(&description)

	fmt.Println("RIC Id:")
	fmt.Scanln(&ricId)

	fmt.Println("Policy Id:")
	fmt.Scanln(&policyId)

	fmt.Println("Service Id")
	fmt.Scanln(&serviceID)

	fmt.Println("Policy Type Id:")
	fmt.Scanln(&policyTypeId)

 	intent = datamodel.Intent{
		Name: name,
		Description: description,
		RicID: ricId,
		PolicyId: policyId,
		ServiceID: serviceID,
		PolicyTypeId: policyTypeId,
	} 

	intentID, err := _client.IntentCreate(intent)
	if err != nil {
		return err
	}
	fmt.Printf("Created intent '%s' with id %d\n", intent.Name, intentID)
	return nil
}

func ListIntents() error {
	intents, err := _client.IntentList()
	if err == nil {
		for _, intent := range intents {
			fmt.Printf("ID  NAME\n")
			fmt.Printf("%d   %s\n", intent.Idx, intent.Name)
		}
	}
	return err
}

func IntentShow(args []string) error {
	intentShowCmd := flag.NewFlagSet("intent show", flag.PanicOnError)
	flagIdx := intentShowCmd.Int("intent", -1, "id of the intent")
	intentShowCmd.Parse(args)

	intent, err := _client.IntentShow(*flagIdx)
	if err == nil {
		fmt.Println(intent)
	}
	return err
}

func IntentDelete(args []string) error {
	intentDeleteCmd := flag.NewFlagSet("intent delete", flag.PanicOnError)
	flagIdx := intentDeleteCmd.Int("intent", -1, "intent to delete")
	intentDeleteCmd.Parse(args)

	err := _client.IntentDelete(*flagIdx)

	return err
}

func main() {
	// Parse the command line options
	flag.Parse()
	_client = client.Client{
		Uri: *_flagURI,
	}

/* 	// Setup the logger
	var logLevel = slog.LevelError
	if *_flagDebug {
		logLevel = slog.LevelDebug
	}

	handlerOptions := slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stderr, &handlerOptions)
	slog.SetDefault(slog.New(handler))

	slog.Debug("Logger created") */

	// Execute the command
	command, args := flag.Args()[0], flag.Args()[1:]
	var err error
	switch command {
	case "create":
		err = CreateIntent(args)
	case "list":
		err = ListIntents()
	case "show":
		err = IntentShow(args)
	case "delete":
		err = IntentDelete(args)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "command failed %s\n", err)
	}
}
