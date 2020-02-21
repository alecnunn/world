package music

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/template"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

// Instrument is a musical instrument
type Instrument struct {
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	Type                   string   `json:"type"`
	BaseMaterialOptions    []string `json:"base_material_options"`
	SupportMaterialOptions []string `json:"support_material_options"`
	BaseMaterial           string   `json:"base_material"`
	SupportMaterial        string   `json:"support_material"`
	DescriptionTemplate    string   `json:"description_template"`
}

const instrumentError = "Could not generate instruments: %w"

// GenerateInstruments generates a set of musical instruments for a climate
func GenerateInstruments(resources []resource.Resource) ([]Instrument, error) {
	var instrument Instrument
	var availableBaseMaterials []string
	var availableSupportMaterials []string
	var woodName string

	availableHides := []string{}
	availableMetals := []string{}
	availableWoods := []string{}
	availableMaterials := []string{}

	allInstruments := getAllInstruments()
	availableInstruments := []Instrument{}
	instruments := []Instrument{}

	metals := resource.ByTag("metal", resources)

	for _, i := range metals {
		availableMetals = append(availableMetals, i.Name)
	}

	wood := resource.ByTag("wood", resources)

	for _, i := range wood {
		woodName = i.Name
		if !strings.HasSuffix(i.Name, "wood") {
			woodName += "-wood"
		}
		availableWoods = append(availableWoods, woodName)
	}

	hides := resource.ByTag("hide", resources)
	for _, h := range hides {
		availableHides = append(availableHides, h.Name)
	}

	if len(availableHides) > 0 {
		availableMaterials = append(availableMaterials, "hide")
	}
	if len(availableMetals) > 0 {
		availableMaterials = append(availableMaterials, "metal")
	}
	if len(availableWoods) > 0 {
		availableMaterials = append(availableMaterials, "wood")
	}

	for _, i := range allInstruments {
		if slices.StringSlicePartlyWithin(i.BaseMaterialOptions, availableMaterials) {
			if slices.StringSlicePartlyWithin(i.SupportMaterialOptions, availableMaterials) {
				availableInstruments = append(availableInstruments, i)
			}
		}
	}

	numberOfInstruments := rand.Intn(3) + 1

	for i := 0; i < numberOfInstruments; i++ {
		instrument = availableInstruments[rand.Intn(len(availableInstruments))]
		availableBaseMaterials = []string{}
		availableSupportMaterials = []string{}

		for _, m := range instrument.BaseMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableBaseMaterials = append(availableBaseMaterials, m)
			}
		}

		for _, m := range instrument.SupportMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableSupportMaterials = append(availableSupportMaterials, m)
			}
		}

		materialType, err := random.String(availableBaseMaterials)
		if err != nil {
			err = fmt.Errorf("Could not generate instruments: %w", err)
			return []Instrument{}, err
		}
		if materialType == "hide" {
			hide, err := random.String(availableHides)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = hide
		} else if materialType == "metal" {
			metal, err := random.String(availableMetals)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = metal
		} else if materialType == "wood" {
			wood, err := random.String(availableWoods)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = wood
		}

		materialType, err = random.String(availableSupportMaterials)
		if err != nil {
			err = fmt.Errorf(instrumentError, err)
			return []Instrument{}, err
		}
		if materialType == "hide" {
			hide, err := random.String(availableHides)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = hide
		} else if materialType == "metal" {
			metal, err := random.String(availableMetals)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = metal
		} else if materialType == "wood" {
			wood, err := random.String(availableWoods)
			if err != nil {
				err = fmt.Errorf(instrumentError, err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = wood
		}

		instrument.Description = instrument.getDescription()

		instruments = append(instruments, instrument)
	}

	return instruments, nil
}

func (instrument Instrument) getDescription() string {
	t := template.New("instrument description")

	var err error
	t, err = t.Parse(instrument.DescriptionTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, instrument); err != nil {
		panic(err)
	}

	result := tpl.String()

	return result
}

func getAllInstruments() []Instrument {
	instruments := []Instrument{
		{
			Name:                   "short flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "long flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "twin flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "short harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "long harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "full harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "lyre",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "lijerica",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "long-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "pierced lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "short-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"sinew"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}}",
		},
		{
			Name:                   "single-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drone",
		},
		{
			Name:                   "multiple-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drones",
		},
		{
			Name:                   "hand drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}}",
		},
		{
			Name:                   "short drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}}",
		},
		{
			Name:                   "walking drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}}",
		},
		{
			Name:                   "heavy drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}}",
		},
	}

	return instruments
}
