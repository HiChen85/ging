package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var RootCMD = &cobra.Command{
	Use:   "ging",
	Short: "ging is to create a gin dev template",
	Long: "This tool is built for those who are freshmen in gin, it will create a simple template which include " +
		"handlers, models, routers, templates and static directories. Besides, it will add some basic database connection configuration for users, some " +
		"simple handler functions and routers.",
	Run: func(cmd *cobra.Command, args []string) {},
}

// create subCommand for ging_local
var CreateCMD = &cobra.Command{
	Use:   "create",
	Short: "create a template project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("no arguments added... \n" +
				"Please add a name for you project.")
		}
		if len(args) > 1 {
			log.Fatal("too many arguments, only one argument supported")
		}
		_, err := os.ReadDir(args[0])
		// the project name not exist
		if err != nil {
			log.Printf("No %v Dir, creating it now!\n", args[0])
			e := os.Mkdir(args[0], os.ModePerm)
			if e != nil {
				log.Fatal(e.Error())
			}
			e = os.Chdir(args[0])
			if e != nil {
				log.Fatal(e.Error())
			}
			CreateDirs()
			files, err := os.ReadDir(".")
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				if file.IsDir() {
					createTemplate(file.Name())
				}
			}
			
			createMainFile(".")
			moduleName, err := cmd.Flags().GetString("module")
			if err != nil {
				log.Fatal(err)
			}
			InitProject(moduleName)
			InstallDependencies()
		} else {
			log.Println("Dir exists... do not need to create...")
		}
	},
}

func init() {
	RootCMD.AddCommand(CreateCMD)
	CreateCMD.Flags().StringP("module", "m", "demo", "give a go modules name for your project")
}

func Execute() {
	log.Println("Creating template....")
	if err := RootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Println("Congratulation, templates created successfully!")
}
