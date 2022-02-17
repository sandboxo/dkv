package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	RootCmd = &cobra.Command{
		Use:   "dkv",
		Short: "DKV is a key value store",
		Long: `DKV is a distributed key value store server written in Go. 

It exposes all its functionality over gRPC & Protocol Buffers.
Complete documentation is available at https://github.com/flipkart-incubator/dkv`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			bindFlags(cmd)
			return nil
		},
	}
)

func Execute() error {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/dkvconfig.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "dkvconfig.yaml"
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("dkvconfig")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.SafeWriteConfigAs(home + "/dkvconfig.yaml")

}

// Bind each cobra flag to its associated viper configuration
func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func addStringVarToFlagAndViper(cmd *cobra.Command, p *string, name string, value string, usage string) {
	serverPFlagSet := cmd.PersistentFlags()
	serverPFlagSet.StringVar(p, name, value, usage)
	viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}

func addDurationVarToFlagAndViper(cmd *cobra.Command, p *time.Duration, name string, value time.Duration, usage string) {
	serverPFlagSet := cmd.PersistentFlags()
	serverPFlagSet.DurationVar(p, name, value, usage)
	viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}

func addUint64VarToFlagAndViper(cmd *cobra.Command, p *uint64, name string, value uint64, usage string) {
	serverPFlagSet := cmd.PersistentFlags()
	serverPFlagSet.Uint64Var(p, name, value, usage)
	viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}

func addBooleanVarToFlagAndViper(cmd *cobra.Command, p *bool, name string, value bool, usage string) {
	serverPFlagSet := cmd.PersistentFlags()
	serverPFlagSet.BoolVar(p, name, value, usage)
	viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}
