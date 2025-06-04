package shell

func detectFromEnv(shellPath, comspec string) Kind {
	if k := classify(fromBaseName(shellPath)); k != Unknown {
		return k
	}
	if k := classify(fromBaseName(comspec)); k != Unknown {
		return k
	}
	return Unknown
}
