// generated by stringer -type=State .; DO NOT EDIT

package ardop

import "fmt"

const _State_name = "UnknownOfflineDisconnectedISSIRSIdleFECSendFECReceive"

var _State_index = [...]uint8{0, 7, 14, 26, 29, 32, 36, 43, 53}

func (i State) String() string {
	if i+1 >= State(len(_State_index)) {
		return fmt.Sprintf("State(%d)", i)
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}
