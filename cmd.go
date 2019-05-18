package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	corebasedef "github.com/zhs007/jarviscore/basedef"
	jarvismarketbasedef "github.com/zhs007/jarvismarket/basedef"
)

func addStart(rootCmd *cobra.Command) {
	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start jarvismarket",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./jarvismarket", "start")
				err := command.Start()
				if err != nil {
					fmt.Printf("start jarvismarket error. %v \n", err)

					os.Exit(-1)

					return
				}

				// fmt.Printf("jarvissh start, [PID] %d running...\n", command.Process.Pid)
				// ioutil.WriteFile("jarvissh.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

				daemon = false

				os.Exit(0)

				return
			}

			fmt.Printf("jarvismarket start.\n")

			fmt.Printf("jarvismarket start, [PID] %d running...\n", os.Getpid())
			ioutil.WriteFile("jarvismarket.pid", []byte(fmt.Sprintf("%d", os.Getpid())), 0666)

			startServ()
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")

	rootCmd.AddCommand(startCmd)
}

func addStop(rootCmd *cobra.Command) {
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop jarvismarket",
		Run: func(cmd *cobra.Command, args []string) {
			strb, err := ioutil.ReadFile("jarvismarket.pid")
			if err != nil {
				fmt.Printf("read jarvismarket.pid error %v\n", err)

				os.Exit(-1)

				return
			}

			command := exec.Command("kill", string(strb))
			command.Start()

			time.Sleep(time.Duration(30) * time.Second)

			fmt.Printf("jarvismarket stop.\n")
		},
	}

	rootCmd.AddCommand(stopCmd)
}

func addRestart(rootCmd *cobra.Command) {
	var daemon bool

	restartCmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart jarvismarket",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				strb, err := ioutil.ReadFile("jarvismarket.pid")
				if err == nil {
					fmt.Printf("stop jarvismarket %v ... \n", string(strb))

					command := exec.Command("kill", string(strb))
					command.Start()

					time.Sleep(time.Duration(30) * time.Second)
				}

				command := exec.Command("./jarvismarket", "start")
				err = command.Start()
				if err != nil {
					fmt.Printf("start jarvismarket error. %v \n", err)

					os.Exit(-1)

					return
				}

				// fmt.Printf("jarvissh start, [PID] %d running...\n", command.Process.Pid)
				// ioutil.WriteFile("jarvissh.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

				daemon = false

				os.Exit(0)

				return
			}

			fmt.Printf("jarvismarket start.\n")

			fmt.Printf("jarvismarket start, [PID] %d running...\n", os.Getpid())
			ioutil.WriteFile("jarvismarket.pid", []byte(fmt.Sprintf("%d", os.Getpid())), 0666)

			startServ()
		},
	}

	restartCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")

	rootCmd.AddCommand(restartCmd)
}

func addVersion(rootCmd *cobra.Command) {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get jarvismarket version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("jarvismarket version is %v \n", jarvismarketbasedef.VERSION)
			fmt.Printf("jarvis core version is %v \n", corebasedef.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)
}

func startCmd() error {
	rootCmd := &cobra.Command{
		Use: "jarvismarket",
	}

	//--------------------------------------------------------------------------------------------------------------------
	// start

	addStart(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// stop

	addStop(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// restart

	addRestart(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// version

	addVersion(rootCmd)

	return rootCmd.Execute()
}
