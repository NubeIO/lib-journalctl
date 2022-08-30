package journalctl

import (
	"fmt"
	"testing"
)

func TestJournalCTL(*testing.T) {

	getLogs, err := NewJournalCTL().EntriesAfter("nubeio-flow-framework.service", "", "300")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range getLogs {
		fmt.Println(entry.Message)

	}

}
