package erratum

func RecoverError(resource Resource, originalErr *error) {
	if err := recover(); err != nil {
		switch val := err.(type) {
		case FrobError:
			resource.Defrob(val.defrobTag)
			*originalErr = val
		case error:
			*originalErr = val
		}
	}
}

func SafeUse(opener ResourceOpener, input string, outErr *error) {
	var resource Resource
	resource, *outErr = opener()
	if *outErr == nil {
		defer resource.Close()
		defer RecoverError(resource, outErr)
		resource.Frob(input)
	} else {
		switch (*outErr).(type) {
		case TransientError:
			SafeUse(opener, input, outErr)
		}
	}
}

func Use(opener ResourceOpener, input string) error {
	var err error
	SafeUse(opener, input, &err)
	return err
}
