// Code generated by goderive DO NOT EDIT.

package customprefix

func eq(this, that *MyStruct) bool {
	return (this == nil && that == nil) || (this != nil) && (that != nil) &&
		this.Int64 == that.Int64 &&
		this.String == that.String
}
