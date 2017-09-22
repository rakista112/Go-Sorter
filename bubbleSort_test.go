package sorters
import (
	"fmt"
	"testing"
	"reflect"
)
func TestBubbleSort(t *testing.T) {
	actual := []int{4, 1, 9, 6}
	expected := []int{1, 4, 6, 9}
	sorter := BubbleSort{}
	sorter.Sort(actual)

	if(!reflect.DeepEqual(expected, actual)){
		msg := fmt.Sprintf("Expected %v got %v\n", expected, actual)
		t.Error(msg)
	}
}