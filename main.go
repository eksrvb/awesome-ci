package main

import (
	"awesome-ci-semver/gitcontroller"
	"awesome-ci-semver/semver"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func runcmd(cmd string, shell bool) string {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
			panic("some error found")
		}
		return string(out)
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	cienv := flag.String("cienv", "Github", "set your CI Environment for Special Featueres!\nAvalible: Jenkins, Github, Gitlab, Custom\nDefault: Github")

	flag.Parse()

	getVersion := flag.NewFlagSet("versioning", flag.ExitOnError)
	overrideVersion := getVersion.String("version", "", "override version to Update")
	getVersionIncrease := getVersion.String("level", "", "predefine version to Update")
	isDryRun := getVersion.Bool("dry-run", false, "Make dry-run before writing version to Git")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "versioning":
		getVersion.Parse(os.Args[2:])

		environment := "Github"

		var gitVersion string
		if *overrideVersion != "" {
			gitVersion = *overrideVersion
		} else {
			gitVersion = gitcontroller.GetLatestReleaseVersion(environment)
		}

		var patchLevel string
		if *getVersionIncrease != "" {
			patchLevel = *getVersionIncrease
		} else {
			if environment == "Github" {
				r := regexp.MustCompile(`[a-zA-z ]+#([0-9]+) from ([0-9a-zA-Z-\/]+)`)
				fmt.Printf("%#v\n", r.FindStringSubmatch(`Merge pull request #3 from ITC-TO-MT/bugfix/test-1`))

				// prAndBranch := r.FindStringSubmatch(`Merge pull request #3 from ITC-TO-MT/bugfix/test-1`)
				// fmt.Printf("%#v\n", r.SubexpNames())
				/* matched, err := regexp.MustCompile("Merge pull request (#[0-9]+) from ([0-9a-zA-Z-\/]+)", "Merge pull request #3 from ITC-TO-MT/bugfix/test-1")
				if err != nil {
					fmt.Println(err)
				} */
				// fmt.Println(matched)
			}
			patchLevel = "bugfix"
		}

		if *isDryRun {
			fmt.Printf("Old version: %s\n", gitVersion)
			fmt.Printf("New version: %s\n", semver.IncreaseSemVer(patchLevel, gitVersion))
		} else {
			fmt.Printf("Old version: %s\n", gitVersion)
			newVersion := semver.IncreaseSemVer(patchLevel, gitVersion)
			fmt.Printf("New version: %s\n", newVersion)
			gitcontroller.CreateNextGitHubRelease(*cienv, newVersion)
		}

	}
}
