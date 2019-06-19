package methods

import (
	"strings"

	"github.com/mia0x75/pages/slug"
	"github.com/mia0x75/pages/structure"
)

// GenerateTagsFromCommaString TODO
func GenerateTagsFromCommaString(input string) []structure.Tag {
	output := make([]structure.Tag, 0)
	tags := strings.Split(input, ",")
	for index := range tags {
		tags[index] = strings.TrimSpace(tags[index])
	}
	for _, tag := range tags {
		if tag != "" {
			output = append(output, structure.Tag{Name: []byte(tag), Slug: slug.Generate(tag, "tags")})
		}
	}
	return output
}
