package functions

type AuthFunction struct {
}

func (a *AuthFunction) Exec() Function {
	return a
}
