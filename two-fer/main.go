//Package twofer is being used to provide a function to share something with
package twofer

//ShareWith Returns a string using a provided name to share something with
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
