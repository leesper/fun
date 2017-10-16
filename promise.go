package fun

const (
	pending = iota
	fulfilled
	rejected
)

type Thener interface {
	Then()
}

func Resolve() {}
func Reject()  {}

type Promise struct{}

func NewPromise()     {}
func ResolvePromise() {}
func RejectPromise()  {}
func AllPromises()    {}
func RacePromises()   {}

func (p *Promise) Then(f func())  {}
func (p *Promise) Catch(f func()) {}
