package cmd

type Command []string

func (c *Command) Has(key string) bool {
	for _, cmd := range *c {
		if cmd == key {
			return true
		}
	}
	return false
}
