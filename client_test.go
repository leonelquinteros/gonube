package gonube

func getTestClient() Client {
	c := New(NewClientConfig())
	//c.config.Debug = true
	return c
}
