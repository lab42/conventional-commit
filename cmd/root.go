package cmd

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

var Log *slog.Logger

var rootCmd = &cobra.Command{
	Use:   "conventional-commit",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		meta := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
		owner, repo := meta[0], meta[1]

		res := regexp.MustCompile(`refs\/pull\/(\d+)\/merge`).FindStringSubmatch(os.Getenv("GITHUB_REF"))
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

		re := regexp.MustCompile(
			fmt.Sprintf(
				`^(%s){1}(\([\w\-\.]+\))?(!)?: (.+)`,
				strings.Join(
					strings.Split(
						strings.TrimRight(
							os.Getenv("INPUT_ALLOWED_TYPES"),
							"\n",
						),
						"\n",
					),
					"|",
				),
			),
		)

		if !re.Match([]byte(githubPr.GetTitle())) {
			Log.Info("type must be one of:\n%s", os.Getenv("INPUT_ALLOWED_TYPES"))
			os.Exit(100)
		}

		subMatches := re.FindStringSubmatch(githubPr.GetTitle())
		prScope := subMatches[1]
		prDescription := subMatches[3]

		if cast.ToBool(os.Getenv("INPUT_REQUIRE_SCOPE")) {
			if len(prScope) < 3 {
				Log.Info("scope is mandatory")
				os.Exit(101)
			}
		}

		if len(prScope) > 2 {
			if !regexp.MustCompile(os.Getenv("INPUT_SCOPE_REGEXP")).Match(
				[]byte(strings.TrimLeft(strings.TrimRight(prScope, ")"), "(")),
			) {
				Log.Info("scope validation failed")
				os.Exit(102)
			}
		}

		if !regexp.MustCompile(os.Getenv("INPUT_DESCRIPTION_REGEXP")).Match([]byte(prDescription)) {
			Log.Info("description validation failed")
			os.Exit(103)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	Log = slog.New(slog.NewTextHandler(os.Stdout))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
