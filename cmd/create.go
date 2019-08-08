package cmd

import (
	"log"
	"net"
	"strconv"

	"github.com/LINBIT/linstor-remote-storage/application"
	"github.com/rck/unit"
	"github.com/spf13/cobra"
)

var ip net.IP
var nodes []string
var username, password, size, portals string
var sizeKiB uint64

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an iSCSI target",
	Long:
`Creates a highly available iSCSI target based on LINSTOR and Pacemaker.
At first it creates a new resouce within the linstor system, using the
specified resource group. The name of the linstor resources is derived
from the iqn and the lun number.
After that it creates resource primitives in the Pacemaker cluster including
all necessary order and location constraints. The Pacemaker primites are
prefixed with p_, contain the name and a resource type postfix.

For example:
linstor-iscsi create --iqn=iqn.2019-08.com.libit:example --ip=192.168.122.181 \
 -username=foo --lun=0 --password=bar --resource_group=ssd_thin_2way --size=2G

Creates linstor resources example_lu0 and
pacemaker primitives p_iscsi_example_ip, p_iscsi_example, p_iscsi_example_lu0`,

	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		// TODO directly use for size, unit fulfills the flag interface.
		units := unit.DefaultUnits
		units["KiB"] = units["K"]
		units["MiB"] = units["M"]
		units["GiB"] = units["G"]
		units["TiB"] = units["T"]
		units["PiB"] = units["P"]
		units["EiB"] = units["E"]
		u := unit.MustNewUnit(units)

		v, err := u.ValueFromString(size)
		if err != nil {
			log.Fatal(err)
		}
		if v.Value < 0 {
			log.Fatal("Negative sizes are not allowed")
		}
		sizeKiB = uint64(v.Value / unit.DefaultUnits["K"])

		if portals == "" {
			portals = ip.String() + ":" + strconv.Itoa(application.DFLT_ISCSI_PORTAL_PORT)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		_, err := application.CreateResource(
			iqn, uint8(lun), sizeKiB, nodes,
			// clientNodeList not supported yet
			make([]string, 0),
			ip, username, password, portals)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IPVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "Set the service IP of the target (required)")
	createCmd.Flags().StringVar(&portals, "portals", "", "Set up portals, if unset, the service ip and default port")
	createCmd.Flags().StringSliceVar(&nodes, "nodes", []string{}, "Set up a list of nodes (required)")
	createCmd.Flags().StringVarP(&username, "username", "u", "", "Set the username (required)")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "Set the password (required)")
	createCmd.Flags().StringVar(&size, "size", "1G", "Set the size (required)")

	createCmd.MarkFlagRequired("ip")
	createCmd.MarkFlagRequired("username")
	createCmd.MarkFlagRequired("password")
	createCmd.MarkFlagRequired("nodes")
	createCmd.MarkFlagRequired("size")
}