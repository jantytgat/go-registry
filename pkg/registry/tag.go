package registry

import "reflect"

const (
	transcryptTag = "transcrypt"
)

type Tag struct {
	Enabled bool
}

func getTags(r any) (map[string]Tag, error) {
	var err error

	t := reflect.TypeOf(r)

	// Create a map with capacity of the number of fields in T
	var m = make(map[string]Tag, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name

		// Check if secure tag is set on field
		tag, ok := t.Field(i).Tag.Lookup(transcryptTag)
		if !ok {
			continue
		}
		m[fieldName] = parseTag(tag)
	}

	return m, err
}

func parseTag(t string) Tag {
	var tag Tag
	switch t {
	case "true":
		tag = Tag{Enabled: true}
	default:
		tag = Tag{Enabled: false}
	}
	return tag
}
