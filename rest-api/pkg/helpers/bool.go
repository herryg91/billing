package helpers

func BoolPtr(in bool) *bool {
	var out *bool
	out = &in
	return out
}
