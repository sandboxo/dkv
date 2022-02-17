package commands

import (
	// "fmt"

	// "github.com/flipkart-incubator/dkv/pkg/ctl"
	"fmt"
	"sort"
	"strings"

	"github.com/flipkart-incubator/dkv/pkg/ctl"
	"github.com/flipkart-incubator/dkv/pkg/serverpb"
	"github.com/spf13/cobra"
)

var (
	ClientCmd = &cobra.Command{
		Use:   "client",
		Short: "Command to operate the DKV client",
		Long: `The subcommands of ` + "`" + `dkv client` + "`" + ` operates the DKV client.
		For example, ` + "`" + `dkv client set <key> <value>` + "`" + ` sets a KV pair , 
		and ` + "`" + `dkv server get <key>` + "`" + ` gets the Value for a Key .`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// setUpClient()
		},
	}

	clientSetCmd = &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a key value pair",
		Long:  `Command to set <key> <value> pair`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientSet(args...)
		},
	}

	clientGetCmd = &cobra.Command{
		Use:   "get <key>",
		Short: "Get value for the given key",
		Long:  `Command to get value for the given key`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientGet(args...)
		},
	}

	clientDelCmd = &cobra.Command{
		Use:   "del <key>",
		Short: "Delete the given key",
		Long:  `Command to delete the given key`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientDelete(args...)
		},
	}

	clientIterCmd = &cobra.Command{
		Use:   "iter \"*\" | <prefix> [<startKey>]",
		Short: "Iterate keys matching the <prefix>, starting with <startKey> or \"*\" for all keys",
		Long:  `Command to iterate keys matching the <prefix>, starting with <startKey> or \"*\" for all keys`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientIterate(args...)
		},
	}

	clientKeysCmd = &cobra.Command{
		Use:   "keys \"*\" | <prefix> [<startKey>]",
		Short: "Get keys matching the <prefix>, starting with <startKey> or \"*\" for all keys",
		Long:  `Command to get keys matching the <prefix>, starting with <startKey> or \"*\" for all keys`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientKeys(args...)
		},
	}

	clientBackupCmd = &cobra.Command{
		Use:   "backup <path>",
		Short: "Backs up data to the given path",
		Long:  `Command to back up data to the given path`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientBackup(args...)
		},
	}

	clientRestoreCmd = &cobra.Command{
		Use:   "restore <path>",
		Short: "Restores data from the given path",
		Long:  `Command to restore data from the given path`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientRestore(args...)
		},
	}

	clientAddNodeCmd = &cobra.Command{
		Use:   "addNode <nexusUrl>",
		Short: "Add another master node to DKV cluster",
		Long:  `Command to add another master node to DKV cluster`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientAddNode(args...)
		},
	}

	clientRemoveNodeCmd = &cobra.Command{
		Use:   "removeNode <nexusUrl>",
		Short: "Remove a master node from DKV cluster",
		Long:  `Command to Remove a master node from DKV cluster`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientRemoveNode(args...)
		},
	}

	clientListNodesCmd = &cobra.Command{
		Use:   "listNodes",
		Short: "Lists the various DKV nodes that are part of the Nexus cluster",
		Long:  `Command to list the various DKV nodes that are part of the Nexus cluster`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientListNodes(args...)
		},
	}

	clientGetClusterInfoCmd = &cobra.Command{
		Use:   "getClusterInfo <dcId> <database> <vBucket>",
		Short: "Gets the latest cluster info",
		Long:  `Command to get the latest cluster info`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setUpClient()
			ClientGetClusteInfo(args...)
		},
	}
)

var (
	dkvAddr, dkvAuthority string
	client                *ctl.DKVClient
)

func init() {
	ClientCmd.AddCommand(clientSetCmd)
	ClientCmd.AddCommand(clientGetCmd)
	ClientCmd.AddCommand(clientDelCmd)
	ClientCmd.AddCommand(clientIterCmd)
	ClientCmd.AddCommand(clientKeysCmd)
	ClientCmd.AddCommand(clientBackupCmd)
	ClientCmd.AddCommand(clientRestoreCmd)
	ClientCmd.AddCommand(clientAddNodeCmd)
	ClientCmd.AddCommand(clientRemoveNodeCmd)
	ClientCmd.AddCommand(clientListNodesCmd)
	ClientCmd.AddCommand(clientGetClusterInfoCmd)
	RootCmd.AddCommand(ClientCmd)
}

func init() {
	addStringVarToFlagAndViper(ClientCmd, &dkvAddr, "dkvAddr", "127.0.0.1:8080", "<host>:<port> - DKV server address")
	addStringVarToFlagAndViper(ClientCmd, &dkvAuthority, "authority", "", "Override :authority pseudo header for routing purposes. Useful while accessing DKV via service mesh.")
}

func setUpClient() {
	fmt.Printf("Connecting to DKV service at %s", dkvAddr)
	if dkvAuthority = strings.TrimSpace(dkvAuthority); dkvAuthority != "" {
		fmt.Printf(" (:authority = %s)", dkvAuthority)
	}
	fmt.Printf("...")
	cl, err := ctl.NewInSecureDKVClient(dkvAddr, dkvAuthority)
	client = cl
	if err != nil {
		fmt.Printf("\nUnable to create DKV client. Error: %v\n", err)
		return
	}
	fmt.Println("DONE")
}

func ClientSet(args ...string) {
	if err := client.Put([]byte(args[0]), []byte(args[1])); err != nil {
		fmt.Printf("Unable to perform SET. Error: %v\n", err)
	} else {
		fmt.Println("OK")
	}
	defer client.Close()
}

func ClientGet(args ...string) {
	rc := serverpb.ReadConsistency_LINEARIZABLE
	if res, err := client.Get(rc, []byte(args[0])); err != nil {
		fmt.Printf("Unable to perform GET. Error: %v\n", err)
	} else {
		fmt.Println(string(res.Value))
	}
	defer client.Close()
}

func ClientDelete(args ...string) {
	if err := client.Delete([]byte(args[0])); err != nil {
		fmt.Printf("Unable to perform DEL. Error: %v\n", err)
	} else {
		fmt.Println("OK")
	}
	defer client.Close()
}

func ClientIterate(args ...string) {
	strtKy, kyPrfx := "", ""
	switch {
	case len(args) == 1:
		if strings.TrimSpace(args[0]) != "*" {
			kyPrfx = args[0]
		}
	case len(args) == 2:
		kyPrfx, strtKy = args[0], args[1]
	}
	if ch, err := client.Iterate([]byte(kyPrfx), []byte(strtKy)); err != nil {
		fmt.Printf("Unable to perform iteration. Error: %v\n", err)
	} else {
		for kvp := range ch {
			if kvp.ErrMsg != "" {
				fmt.Printf("Error: %s\n", kvp.ErrMsg)
			} else {
				fmt.Printf("%s => %s\n", kvp.Key, kvp.Val)
			}
		}
	}
}

func ClientKeys(args ...string) {
	strtKy, kyPrfx := "", ""
	switch {
	case len(args) == 1:
		if strings.TrimSpace(args[0]) != "*" {
			kyPrfx = args[0]
		}
	case len(args) == 2:
		kyPrfx, strtKy = args[0], args[1]
	}
	if ch, err := client.Iterate([]byte(kyPrfx), []byte(strtKy)); err != nil {
		fmt.Printf("Unable to perform iteration. Error: %v\n", err)
	} else {
		for kvp := range ch {
			if kvp.ErrMsg != "" {
				fmt.Printf("Error: %s\n", kvp.ErrMsg)
			} else {
				fmt.Printf("%s\n", kvp.Key)
			}
		}
	}
}

func ClientBackup(args ...string) {
	if err := client.Backup(args[0]); err != nil {
		fmt.Printf("Unable to perform backup. Error: %v\n", err)
	} else {
		fmt.Println("Successfully backed up")
	}
	defer client.Close()
}

func ClientRestore(args ...string) {
	if err := client.Restore(args[0]); err != nil {
		fmt.Printf("Unable to perform restore. Error: %v\n", err)
	} else {
		fmt.Println("Successfully restored")
	}
	defer client.Close()
}

func ClientAddNode(args ...string) {
	nodeURL := args[0]
	if err := client.AddNode(nodeURL); err != nil {
		fmt.Printf("Unable to add node with URL: %s. Error: %v\n", nodeURL, err)
	}
	defer client.Close()
}

func ClientRemoveNode(args ...string) {
	nodeURL := args[0]
	if err := client.RemoveNode(nodeURL); err != nil {
		fmt.Printf("Unable to add node with URL: %s. Error: %v\n", nodeURL, err)
	}
	defer client.Close()
}

func ClientListNodes(args ...string) {
	if leaderId, members, err := client.ListNodes(); err != nil {
		fmt.Printf("Unable to retrieve the nodes of DKV cluster. Error: %v\n", err)
	} else {
		var ids []uint64
		for id := range members {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
		if _, present := members[leaderId]; present {
			fmt.Println("Current DKV cluster members:")
		} else {
			fmt.Println("WARNING: DKV Cluster unhealthy, leader unknown")
			fmt.Println("Current DKV cluster members:")
		}
		for _, id := range ids {
			fmt.Printf("%x => %s (%s) \n", id, members[id].NodeUrl, members[id].Status)
		}
	}
}

func ClientGetClusteInfo(args ...string) {
	dcId := ""
	database := ""
	vBucket := ""
	if len(args) > 0 {
		dcId = args[0]
	}
	if len(args) > 1 {
		database = args[1]
	}
	if len(args) > 2 {
		vBucket = args[2]
	}
	vBuckets, err := client.GetClusterInfo(dcId, database, vBucket)
	if err != nil {
		fmt.Printf("Unable to get Status: Error: %v\n", err)
	} else {
		if len(vBuckets) == 0 {
			fmt.Println("Found no nodes with the provided filters")
		} else {
			fmt.Println("Current DKV cluster nodes:")
			for _, bucket := range vBuckets {
				fmt.Println(bucket.String())
			}
		}
	}
}
