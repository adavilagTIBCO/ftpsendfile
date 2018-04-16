package ftpsendfile

import (
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	server := "some.ftp.server"
	port := 21
	url := server + ":" + strconv.Itoa(port)

	tc.SetInput("server", server)
	tc.SetInput("port", port)
	tc.SetInput("username", "username")
	tc.SetInput("password", "password")
	tc.SetInput("pathsrc", "/tmp/")
	tc.SetInput("filesrc", "test.txt")
	tc.SetInput("pathdest", "/")
	tc.SetInput("filedest", "test.txt")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("output")
	assert.Equal(t, result, "Successfully connected to "+url)
}
