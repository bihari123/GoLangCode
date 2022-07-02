package erratum

func Use(opener ResourceOpener, input string) (e error) {
	file, err := opener()

	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(opener, input)
		default:
			return err
		}
	}

	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			switch recoveryMessage.(type) {
			case FrobError:
				e = recoveryMessage.(FrobError).inner
				file.Defrob(recoveryMessage.(FrobError).defrobTag)
			case error:
				e = recoveryMessage.(error)
			default:
				panic(recoveryMessage)
			}
		}
		file.Close()
	}()
	file.Frob(input)
	return nil
	panic("Please implement the Use function")
}
