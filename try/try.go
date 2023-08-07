package try

type Exception struct {
	Id  int
	Msg string
}

type ExceptionHandle func(Exception)

type TryStruct struct {
	catchs map[int]ExceptionHandle
	try    func()
}

func Try(tryHandle func()) *TryStruct {
	trystruct := TryStruct{
		catchs: make(map[int]ExceptionHandle),
		try:    tryHandle,
	}
	return &trystruct
}

func (this *TryStruct) Catch(exceptionId int, catch func(Exception)) *TryStruct {
	this.catchs[exceptionId] = catch
	return this
}

func (this *TryStruct) Finally(finally func()) {
	defer func() {
		if e := recover(); e != nil {
			exception := e.(Exception)
			if catch, ok := this.catchs[exception.Id]; ok {
				catch(exception)
			}
		}
	}()

	this.try()
}

func Throw(id int, msg string) Exception {
	panic(Exception{id, msg})
}
