package deltaTime

import "time"

var delta float64

// Time.deltaTime
var oldTime = time.Now().UnixNano()
var newTime = time.Now().UnixNano()

// deltaStart
func deltaTimeStart() {
	oldTime = time.Now().UnixNano()
}

//deltaTime
func deltaTime() float64 {

	newTime = time.Now().UnixNano()
	oldTime = time.Now().UnixNano()
	deltaTime := float64((newTime-oldTime)/1000000) * 0.001

	// fmt.Println(deltaTime * 10000)

	delta = deltaTime
	oldTime = newTime
	return deltaTime
}
