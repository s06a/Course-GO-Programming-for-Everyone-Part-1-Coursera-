/* to convert miles and yards to kilometers.
   In this case we assume that we do care
   about marathon!
*/

package main

import "fmt"

func main() {
	var miles, yards int32 = 26, 385
	var kilometers float32
	kilometers = 1.60934 * (float32(miles) + float32(yards)/1760)
	fmt.Printf("\nA marathon is %g kilometers\n", kilometers)
}
