package spinner

type Spinner interface {
	Start()
	Stop()
	Update(steps []string)
}
