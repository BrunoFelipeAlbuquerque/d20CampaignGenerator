package modifier

import (
	"fmt"
	"regexp"
	"strings"

	utils "d20campaing/internal/text"
)

var validModifierSource = regexp.MustCompile(`^[a-z0-9]+(_[a-z0-9]+)*$`)

type ModifierSource string

type CircumstanceSourceInfo struct {
	ID          ModifierSource
	Description string
}

type CircumstanceSourceRegistry struct {
	sources map[ModifierSource]CircumstanceSourceInfo
}

func NormalizeModifierSource(s string) ModifierSource {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	return ModifierSource(s)
}

func validateModifierSource(id ModifierSource) error {
	s := string(id)

	if len(s) == 0 {
		return fmt.Errorf("modifier source cannot be empty")
	}

	if !validModifierSource.MatchString(s) {
		return fmt.Errorf("invalid modifier source format: %s", s)
	}

	return nil
}

func NewDefaultCircumstanceSourceRegistry() *CircumstanceSourceRegistry {
	r := &CircumstanceSourceRegistry{
		sources: make(map[ModifierSource]CircumstanceSourceInfo),
	}

	// COMBAT
	r.register("flanking", "Flanking target")
	r.register("higher_ground", "Higher ground advantage")
	r.register("target_prone", "Target is prone")
	r.register("charging", "Charging attack")
	r.register("cover", "Target has cover")
	r.register("soft_cover", "Target has soft cover")

	// PERCEPTION
	r.register("concealment", "Target has concealment")
	r.register("total_concealment", "Target has total concealment")
	r.register("invisible_target", "Target is invisible")
	r.register("poor_lighting", "Poor lighting conditions")

	// CONDITION
	r.register("target_helpless", "Target is helpless")
	r.register("target_blinded", "Target is blinded")
	r.register("target_stunned", "Target is stunned")
	r.register("target_flat_footed", "Target is flat-footed")

	// ENVIRONMENT
	r.register("slippery_surface", "Slippery surface")
	r.register("strong_wind", "Strong wind")
	r.register("extreme_heat", "Extreme heat")
	r.register("extreme_cold", "Extreme cold")

	// SKILL
	r.register("favorable_conditions", "Favorable conditions")
	r.register("unfavorable_conditions", "Unfavorable conditions")
	r.register("distraction", "Distraction")

	return r
}

func (r *CircumstanceSourceRegistry) Get(id ModifierSource) (CircumstanceSourceInfo, bool) {
	info, ok := r.sources[id]
	return info, ok
}

func (r *CircumstanceSourceRegistry) IsKnown(id ModifierSource) bool {
	_, ok := r.sources[id]
	return ok
}

func (r *CircumstanceSourceRegistry) Register(id ModifierSource, desc string) error {
	if err := validateModifierSource(id); err != nil {
		return err
	}

	if _, exists := r.sources[id]; exists {
		return fmt.Errorf("duplicate ModifierSource: %s", id)
	}

	if similar := r.findSimilarSource(id); len(similar) > 0 {
		return fmt.Errorf(
			"similar ModifierSource detected: %s (conflicts with %v)",
			id,
			similar,
		)
	}

	r.register(id, desc)
	return nil
}

func (r *CircumstanceSourceRegistry) register(id ModifierSource, desc string) {
	r.sources[id] = CircumstanceSourceInfo{
		ID:          id,
		Description: desc,
	}
}

func (r *CircumstanceSourceRegistry) findSimilarSource(id ModifierSource) []ModifierSource {
	var similar []ModifierSource

	for existing := range r.sources {
		dist := utils.Levenshtein(string(id), string(existing))

		if dist <= 1 {
			similar = append(similar, existing)
		}
	}

	return similar
}
