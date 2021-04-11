package cmd

import (
	"fmt"

	"github.com/FlorentinDUBOIS/go-template/pkg/libs/logutil"
	"github.com/FlorentinDUBOIS/go-template/pkg/mod"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() (*cobra.Command, error) {
	var err error

	cmd := &cobra.Command{
		Use:   "-",
		Short: "-",
		RunE:  rootCmd,
	}

	cmd.PersistentFlags().Int32P("log-level", "l", 4, "Change log verbosity level")
	cmd.PersistentFlags().StringP("config", "c", "", "Specify a configuration file")

	cobra.OnInitialize(initialize)
	if err = viper.BindPFlags(cmd.Flags()); nil != err {
		return nil, fmt.Errorf("could not flags, %w", err)
	}

	if err = viper.BindPFlags(cmd.PersistentFlags()); nil != err {
		return nil, fmt.Errorf("could not bind persistent flags, %w", err)
	}

	return cmd, nil
}

func initialize() {
	logrus.SetLevel(
		logutil.GetLogrusLevel(viper.GetInt("log.level")),
	)

	if "" == viper.GetString("config") {
		viper.AddConfigPath("/etc/-/config")
		viper.AddConfigPath("$HOME/.config/-/config")
		viper.AddConfigPath("config")
	} else {
		viper.AddConfigPath(viper.GetString("config"))
	}

	if err := viper.ReadInConfig(); nil != err {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.WithError(err).Fatal("could not read configuration")
		}
	}
}

func rootCmd(cmd *cobra.Command, args []string) error {
	configuration := mod.Configuration{}
	if err := viper.Unmarshal(&configuration); nil != err {
		return fmt.Errorf("failed to unmarshal the configuration, %w", err)
	}

	return nil
}
