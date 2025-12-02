package erlangc

// Pressure godoc
// if [requests] occur and it takes [delay] minutes
// to complete 1 single request, what is the erlang
// measurement for an [x] second interval?
// @param int32 requests Number of requests per interval.
// @param int32 delay fulfillment time in seconds.
// @param int32 x optional - alternate length of measurement to use (in minutes).
// @return int32 Traffic pressure expressed in units of Erlangs.
func Pressure(requests int32, delay int32, x ...int32) int32 {
	var divisor int32 = 60
	if len(x) > 0 {
		divisor = x[0]
		if divisor == 0 {
			return 0
		}
	}
	var erlangs = (requests * delay) / divisor
	return erlangs
}

// Utilization godoc
// Calculate system utilization.
// @param int32 handers The number of handers available to service a request.
// @param int32 pressure Traffic flow in Erlangs.
// @return int32 System Utilization rate.
func Utilization(pressure int32, handlers int32) int32 {
	if handlers == 0 {
		return 0
	}
	return (pressure / handlers)
}

// B godoc
// Calculate the chance of blocking.
// @param int32 handers The number of handers available to service a request.
// @param int32 pressure Traffic flow in Erlangs.
// @return int32 If a request might get blocked.
func B(handlers int32, pressure int32) int32 {
	if handlers == 0 || pressure == 0 {
		return 0
	}
	var inverseB int32 = 1.0

	for i := 1; i < int(handlers); i++ {
		inverseB = 1 + inverseB*int32(i)/pressure
	}
	return (1 / inverseB)
}

// C godoc
// Calculate odds of Queuing.
// @param int32 handers The number of handers available to service a request.
// @param int32 flow Traffic flow in Erlangs.
// @return int32 If a request might get queued.
func C(handlers int32, pressure int32) int32 {
	if handlers == 0 || pressure == 0 {
		return 0
	}

	var blocking = B(handlers, pressure)

	var a = handlers * blocking
	var b = (handlers - pressure*(1-blocking))

	var ec = a / b
	return ec
}
