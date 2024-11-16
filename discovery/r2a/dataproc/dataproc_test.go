package dataproc

import (
	"testing"

	"github.com/xifanyan/opentext/discovery/r2a/data/arm"
)

func Test_DataProc_Initiallization(t *testing.T) {
	dataSet := arm.NewDataSetBuilder().WithBasePath("../testdata").WithID("db01").Build()
	proc := NewDataProc(dataSet)
	if err := proc.Initialize(); err != nil {
		t.Errorf("initialize error: %v", err)
	}

	if len(dataSet.Volumns) != 3 {
		t.Errorf("expected 3 volumns, but got %d", len(dataSet.Volumns))
	}
}

func Test_DataProc_TotalColumnValuesCount(t *testing.T) {

	dataSet := arm.NewDataSetBuilder().WithBasePath("../testdata").WithID("db01").Build()
	proc := NewDataProc(dataSet)
	if err := proc.Initialize(); err != nil { // TODO: Check
		t.Errorf("initialize error: %v", err)
	}

	if err := proc.countTotalColumnWithValue(); err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

}
