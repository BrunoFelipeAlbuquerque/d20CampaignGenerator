package modifier

import (
	"fmt"
	"regexp"
	"strings"

	utils "d20campaing/internal/text"
)

var validModifierSource = regexp.MustCompile(`^[a-z0-9]+(?:_[a-z0-9]+)*(?:\.[a-z0-9]+(?:_[a-z0-9]+)*)+$`)

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
	r.register(SourceFlanking, "Flanking target")
	r.register(SourceHigherGround, "Higher ground advantage")
	r.register(SourceProneTarget, "Target is prone")
	r.register(SourceCharging, "Charging attack")
	r.register(SourceCover, "Target has cover")
	r.register(SourceSoftCover, "Target has soft cover")

	// PERCEPTION
	r.register(SourceConcealment, "Target has concealment")
	r.register(SourceTotalConcealment, "Target has total concealment")
	r.register(SourceInvisibleTarget, "Target is invisible")
	r.register(SourcePoorLighting, "Poor lighting conditions")

	// CONDITION
	r.register(SourceTargetHelpless, "Target is helpless")
	r.register(SourceTargetBlinded, "Target is blinded")
	r.register(SourceTargetStunned, "Target is stunned")
	r.register(SourceTargetFlatFooted, "Target is flat-footed")

	// ENVIRONMENT
	r.register(SourceSlipperySurface, "Slippery surface")
	r.register(SourceStrongWind, "Strong wind")
	r.register(SourceExtremeHeat, "Extreme heat")
	r.register(SourceExtremeCold, "Extreme cold")

	// SKILL
	r.register(SourceFavorableConditions, "Favorable conditions")
	r.register(SourceUnfavorableConditions, "Unfavorable conditions")
	r.register(SourceDistraction, "Distraction")

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
