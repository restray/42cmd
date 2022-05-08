package commands

import (
	"flag"
	"fmt"
	"main/commands/utils"
	"main/intrapi"
	"math"
	"strings"
)

type FtCommandUser struct {
	flags    *flag.FlagSet
	Basic    bool
	Detailed bool
	ImageURL bool
	Pool     bool
	Misc     bool
	All      bool
}

func (cmd *FtCommandUser) GetCommand() string {
	return "user"
}

func (cmd *FtCommandUser) GetAlias() []string {
	return []string{"u", "me"}
}

func (cmd *FtCommandUser) GetFlags() *flag.FlagSet {
	return cmd.flags
}

func (cmd *FtCommandUser) GetDescription() string {
	return "Give an user"
}

func (cmd *FtCommandUser) Init() {
	cmd.flags = flag.NewFlagSet(cmd.GetCommand(), flag.ExitOnError)
	cmd.flags.BoolVar(&cmd.All, "all", false, "Display all informations")
	cmd.flags.BoolVar(&cmd.Basic, "basic", false, "Hide Basic infos")
	cmd.flags.BoolVar(&cmd.Detailed, "detailed", false, "Show Detailed infos")
	cmd.flags.BoolVar(&cmd.ImageURL, "image-url", false, "Show the profile avatar URL")
	cmd.flags.BoolVar(&cmd.Pool, "pool", false, "Show Pool")
	cmd.flags.BoolVar(&cmd.Misc, "misc", false, "Show Misc")
}

func (cmd *FtCommandUser) DefaultOutput() {
	fmt.Println("Specify a user name, like:")
	fmt.Println("ftcli user tbelhomm")
}

func (cmd *FtCommandUser) Handler(args []string) {
	if len(args) < 1 {
		cmd.DefaultOutput()
		return
	}

	if args[0] != "me" {
		args = args[1:]
	}

	cmd.flags.Parse(args[1:])

	cmd.Basic = !cmd.Basic

	if cmd.All {
		cmd.Basic = true
		cmd.Detailed = true
		cmd.ImageURL = true
		cmd.Pool = true
		cmd.Misc = true
	}

	var user *intrapi.User42

	if args[0] == "me" {
		user = intrapi.GetMe()
	} else {
		user = intrapi.GetUserByName(args[0])
	}

	if user == nil {
		fmt.Println("This user doesn't exist...")
		return
	}

	if cmd.Basic {
		fmt.Println(user)
	}
	if cmd.Detailed {
		size := math.Max(float64(len(user.FirstName)), float64(len(user.LastName)))
		size = math.Max(size, float64(len(user.Email)))
		size = math.Max(size, float64(len(user.Phone)))

		prefix := strings.Repeat("-", int(size)+2)

		fmt.Printf("\t+-------------%s\n", prefix)
		fmt.Printf("\t| First Name: %s\n", user.FirstName)
		fmt.Printf("\t|---\n")
		fmt.Printf("\t| Last Name:  %s\n", user.LastName)
		fmt.Printf("\t|---\n")
		fmt.Printf("\t| Email:      %s\n", user.Email)
		fmt.Printf("\t|---\n")
		fmt.Printf("\t| Phone:      %s\n", user.Phone)
		fmt.Printf("\t+-------------%s\n", prefix)
	}
	if cmd.ImageURL {
		prefix := strings.Repeat("-", int(len(user.NewImageURL))+2)

		fmt.Printf("\t+--------+-%s\n", prefix)
		fmt.Printf("\t| Avatar | %s\n", user.NewImageURL)
		fmt.Printf("\t+--------+-%s\n", prefix)
	}
	if cmd.Pool {
		fmt.Printf("\t+------+--------------\n")
		fmt.Printf("\t| Pool | %s/%s\n", user.PoolMonth, user.PoolYear)
		fmt.Printf("\t+------+--------------\n")
	}
	if cmd.Misc {
		fmt.Printf("\t   ,==.-------.\n")
		fmt.Printf("\t  (    ) ====  \\\tWallet:\t\t%s$\n", fmt.Sprint(user.Wallet))
		fmt.Printf("\t  ||  | [][][] |\tAlumni:\t\t%s\n", utils.EmojiBool(user.Alumni))
		fmt.Printf("\t,8||  | [][][] |\tStaff:\t\t%s\n", utils.EmojiBool(user.Staff))
		fmt.Printf("\t8 ||  | [][][] |\tRegistred:\t%s\n", fmt.Sprint(user.CreatedAt.Format("02/01/2006")))
		fmt.Printf("\t8 (    ) O O O /\tLast Seen:\t%s\n", fmt.Sprint(user.UpdatedAt.Format("02/01/2006 15:04:05")))
		fmt.Printf("\t'88`=='-------'\n")
	}
}
