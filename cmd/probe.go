package cmd

import (
	"fmt"

	"github.com/LINBIT/linstor-remote-storage/iscsi"
	"github.com/LINBIT/linstor-remote-storage/linstorcontrol"
	term "github.com/LINBIT/linstor-remote-storage/termcontrol"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// probeCmd represents the probe command
var probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "Probes an iSCSI starget",
	Long: `Triggers Pacemaker to probe the resoruce primitives of this iSCSI target.
That means Pacemaker will run the status operation on the nodes where the
resource can run.
This makes sure that Pacemakers view of the world is updated to the state
of the world.

For example:
./linstor-iscsi probe --iqn=iqn.2019-08.com.libit:example --lun=0`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		linstorCfg := linstorcontrol.Linstor{
			Loglevel:     log.GetLevel().String(),
			ControllerIP: controller,
		}
		targetCfg := iscsi.Target{
			IQN: iqn,
			LUN: uint8(lun),
		}
		iscsiCfg := &iscsi.ISCSI{Linstor: linstorCfg, Target: targetCfg}
		rscStateMap, err := iscsiCfg.ProbeResource()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Current state of CRM resources\niSCSI resource %s, logical unit #%d:\n", iqn, uint8(lun))
		for rscName, runState := range *rscStateMap {
			label := term.COLOR_YELLOW + "Unknown" + term.COLOR_RESET
			if runState.HaveState {
				if runState.Running {
					label = term.COLOR_GREEN + "Running" + term.COLOR_RESET
				} else {
					label = term.COLOR_RED + "Stopped" + term.COLOR_RESET
				}
			}
			fmt.Printf("    %-40s %s\n", rscName, label)
		}
	},
}

func init() {
	rootCmd.AddCommand(probeCmd)
}
