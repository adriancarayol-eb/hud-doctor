package spinner

type Spinner interface {
	Start()
	Stop()
	Update(task, status string)
}
