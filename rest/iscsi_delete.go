package rest

import (
	"net/http"

	"github.com/LINBIT/linstor-remote-storage/iscsi"
)

// ISCSIDelete deletes a highly-available iSCSI target via the REST-API
func ISCSIDelete(w http.ResponseWriter, r *http.Request) {
	var iscsiCfg iscsi.ISCSI
	if err := unmarshalBody(w, r, &iscsiCfg); err != nil {
		return
	}
	maybeSetLinstorController(&iscsiCfg)

	if err := iscsiCfg.DeleteResource(); err != nil {
		_, _ = Errorf(http.StatusInternalServerError, w, "Could not delete resource: %v", err)
		return
	}

	// json.NewEncoder(w).Encode(xxx)
}