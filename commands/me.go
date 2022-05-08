package commands

import (
	"flag"
	"fmt"
	"main/commands/me_command"
	"main/commands/utils"
	"main/intrapi"
	"os"
)

type FtCommandMe struct {
	flags      *flag.FlagSet
	wallet     bool
	evalPoints bool
	grade      bool
	place      bool
	blackhole  bool
	prefix     bool
	tags       bool
	coalition  bool
}

func (me *FtCommandMe) GetCommand() string {
	return "me"
}

func (me *FtCommandMe) GetAlias() []string {
	return []string{"profile"}
}

func (me *FtCommandMe) GetFlags() *flag.FlagSet {
	return me.flags
}

func (me *FtCommandMe) GetDescription() string {
	return "Give your 42 profile"
}

func (me *FtCommandMe) Init() {
	me.flags = flag.NewFlagSet("me", flag.ExitOnError)
	me.flags.BoolVar(&me.wallet, "wallet", false, "Should display wallet points")
	me.flags.BoolVar(&me.evalPoints, "eval-points", true, "Should display evaluation points")
	me.flags.BoolVar(&me.grade, "grade", false, "Should display the current grade name")
	me.flags.BoolVar(&me.place, "place", false, "Should display your location in cluster/status")
	me.flags.BoolVar(&me.blackhole, "black-hole", false, "Should display your black hole duration/end date")
	me.flags.BoolVar(&me.prefix, "prefix", false, "Should display your pseudo prefix")
	me.flags.BoolVar(&me.tags, "tags", false, "Should display your tags")
	me.flags.BoolVar(&me.coalition, "coalition", false, "Should display your coalition name")
}

func (me *FtCommandMe) DefaultOutput() {
	user := intrapi.GetMe()

	fmt.Println()
	fmt.Printf("    ██╗  ██╗██████╗\n")
	fmt.Printf("    ██║  ██║╚════██╗	%s\n", user.UsualFullName)
	fmt.Printf("    ███████║ █████╔╝	%s\n", user.Login)
	fmt.Printf("    ╚════██║██╔═══╝ 	Level %.0f\n", user.CursusUsers[1].Level)
	fmt.Printf("         ██║███████╗	lvl 0 [%s] lvl 21\n", utils.LoadingBar(0, 21, int(user.CursusUsers[1].Level)))
	fmt.Printf("         ╚═╝╚══════╝\n")
	fmt.Println()
}

func (me *FtCommandMe) Handler(args []string) {
	if len(args) <= 0 {
		me.DefaultOutput()
		return
	}

	cmds := loadMeSubcommands()

	for _, command := range cmds {
		command.Init()
	}

	for _, command := range cmds {
		if command.GetCommand() == args[0] {
			command.Handler(args[1:])
			return
		}
		for _, alias := range command.GetAlias() {
			if alias == args[0] {
				command.Handler(args[1:])
				return
			}
		}
	}

	me.flags.Usage()
	os.Exit(1)
}

func loadMeSubcommands() []FtCommand {
	cmds := make([]FtCommand, 0)

	cmds = append(cmds, &me_command.FtCommandProject{})
	cmds = append(cmds, &me_command.FtCommandLogtime{})

	return cmds
}
