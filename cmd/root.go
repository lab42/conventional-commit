package cmd

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "conventional-commit",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {

		meta := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
		if len(meta) < 1 {
			os.Exit(1)
		}

		owner, repo := meta[0], meta[1]

		res := regexp.MustCompile(`/refs\/pull\/(\d+)\/merge/`).FindStringSubmatch(os.Getenv("GITHUB_REF"))
		if len(res) < 1 {
			os.Exit(1)
		}

		pr, err := strconv.Atoi(res[1])
		cobra.CheckErr(err)

		ctx := context.Background()

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)

		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		githubPr, _, err := client.PullRequests.Get(ctx, owner, repo, pr)
		cobra.CheckErr(err)

		if regexp.MustCompile(
			fmt.Sprintf(
				`^(%s){1}(\([\w\-\.]+\))?(!)?: %s`,
				viper.GetString("TYPES"),
				viper.GetString("DESCRIPTION"),
			)).
			Match([]byte(githubPr.GetTitle())) {
			os.Exit(0)
		}

		os.Exit(1)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.conventional-commit.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".conventional-commit")
	}

	viper.SetDefault("TYPES", `build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test`)
	viper.SetDefault("DESCRIPTION", `([\w ]+)`)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
