package resource

import "github.com/ironarachne/world/pkg/profession"

func getPottery() []Pattern {
	producer := profession.ByName("potter")

	patterns := []Pattern{
		{
			Name:        "bowl",
			Description: "a clay bowl",
			Tags: []string{
				"food vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.Origin}} bowl",
				},
			},
			Value: 1,
		},
		{
			Name:        "mug",
			Description: "a clay mug",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.Origin}} mug",
				},
			},
			Value: 1,
		},
		{
			Name:        "pitcher",
			Description: "a clay pitcher",
			Tags: []string{
				"liquid vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.Origin}} pitcher",
				},
			},
			Value: 1,
		},
		{
			Name:        "vase",
			Description: "a clay vase",
			Tags: []string{
				"vase",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.Origin}} vase",
				},
			},
			Value: 1,
		},
	}

	return patterns
}