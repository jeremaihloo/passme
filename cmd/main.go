package main

import (
	"github.com/alecthomas/kingpin"
	"os"
)

var (
	app   = kingpin.New("passme", "A password manager.")
	debug = app.Flag("debug", "Enable debug mode").Bool()

	doInit         = app.Command("init", "Initialize passme")
	masterPassword = doInit.Arg("masterPassword", "Your master password").String()

	show         = app.Command("show", "Show password items.")
	showKey      = show.Arg("key", "Search key. Show by this key.").Default("*").String()
	showDomain   = show.Flag("domain", "Show by domain").String()
	showUsername = show.Flag("username", "Show by username").String()

	saveItem = app.Command("save", "Create or update one user and password")
	domain   = saveItem.Arg("domain", "Password item domain.").String()
	username = saveItem.Arg("username", "Password item username.").String()
	password = saveItem.Arg("password", "Password item password").String()

	deleteItem  = app.Command("delete", "Delete one or more user password item")
	dDomain     = deleteItem.Arg("domain", "Delete item domain.").String()
	dUserName   = deleteItem.Arg("username", "Delete item username").String()
	dIfDomain   = deleteItem.Flag("domain", "Delete items by domain").String()
	dIfUserName = deleteItem.Flag("username", "Delete items by username").String()

	storage       = app.Command("storage", "Password files storage info.")
	sname         = storage.Arg("name", "Storage name.").String()
	url           = storage.Arg("url", "Storage url.").String()
	provider      = storage.Arg("provider", "Storage provider name.").String()
	storageList   = storage.Flag("list", "List all storage infos.").Bool()
	storageDelete = storage.Flag("delete", "Delete one storage info.").Bool()

	push     = app.Command("push", "Push current password files to storage")
	pushName = push.Arg("name", "Target storage name.").String()

	fetch     = app.Command("fetch", "Fetch remote storage files into cache.")
	fetchName = fetch.Arg("name", "Target storage name.").String()

	merge            = app.Command("merge", "Merge two storage files.")
	mergeStorageName = merge.Arg("name", "The remote storage name.").String()

	version = app.Command("version", "Show command line interface version.")

	configs   = app.Command("configs", "Config manager.")
	autoPush  = configs.Flag("auto-push", "Config if auto push after changes.").String()
	autoFetch = configs.Flag("aoto-fetch", "Config if auto fetch every time.").String()
	autoName  = configs.Flag("auto-name", "The name of storage when fetch or posh.").String()
	fromUrl   = configs.Flag("config-url", "Init config from url.").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	default:
		kingpin.Usage()
	}
}
