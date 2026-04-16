package creaturetype

type creatureSubtypeStructuralEffectID string
type CreatureSubtypeStructuralEffectID = creatureSubtypeStructuralEffectID

const (
	BreathesWaterTrait         CreatureTypeTraitID = "BreathesWater"
	SwimWithoutChecksTrait     CreatureTypeTraitID = "SwimWithoutChecks"
	SwimAlwaysClassSkillTrait  CreatureTypeTraitID = "SwimAlwaysClassSkill"
	ImmunityBleedTrait         CreatureTypeTraitID = "ImmunityBleed"
	NotSubjectToFlankingTrait  CreatureTypeTraitID = "NotSubjectToFlanking"
	PrecisionDamageImmuneTrait CreatureTypeTraitID = "PrecisionDamageImmune"
)

const (
	IncorporealBodyEffect CreatureSubtypeStructuralEffectID = "IncorporealBody"
	SwarmBodyEffect       CreatureSubtypeStructuralEffectID = "SwarmBody"
)

type creatureSubtypeEffect struct {
	baseProfile           CreatureTypeProfile
	addedTraitIDs         []CreatureTypeTraitID
	removedTraitIDs       []CreatureTypeTraitID
	structuralEffectIDs   []CreatureSubtypeStructuralEffectID
	preservedOriginalType *CreatureTypeID
}

type CreatureSubtypeEffect = creatureSubtypeEffect

type subtypeEffectSpec struct {
	requiredBaseType              *CreatureTypeID
	addedTraitIDs                 []CreatureTypeTraitID
	removedTraitIDs               []CreatureTypeTraitID
	structuralEffectIDs           []CreatureSubtypeStructuralEffectID
	preservesOriginalTypeMetadata bool
}

var subtypeEffectTable = map[CreatureSubtypeID]subtypeEffectSpec{
	AquaticSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			BreathesWaterTrait,
			SwimWithoutChecksTrait,
			SwimAlwaysClassSkillTrait,
		},
	},
	AugmentedSubtype: {
		preservesOriginalTypeMetadata: true,
	},
	ElementalSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			NoNeedToEatSleepBreatheTrait,
			ImmunityBleedTrait,
			ImmunityParalysisTrait,
			ImmunityPoisonTrait,
			ImmunitySleepTrait,
			ImmunityStunTrait,
			NotSubjectToCriticalHitsTrait,
			NotSubjectToFlankingTrait,
		},
		removedTraitIDs: []CreatureTypeTraitID{
			BreatheEatSleepTrait,
		},
	},
	IncorporealSubtype: {
		addedTraitIDs: []CreatureTypeTraitID{
			NotSubjectToCriticalHitsTrait,
			PrecisionDamageImmuneTrait,
		},
		structuralEffectIDs: []CreatureSubtypeStructuralEffectID{
			IncorporealBodyEffect,
		},
	},
	NativeSubtype: {
		requiredBaseType: creatureTypePointer(OutsiderType),
		addedTraitIDs: []CreatureTypeTraitID{
			BreatheEatSleepTrait,
		},
		removedTraitIDs: []CreatureTypeTraitID{
			NoNeedToEatSleepBreatheTrait,
		},
	},
	SwarmSubtype: {
		structuralEffectIDs: []CreatureSubtypeStructuralEffectID{
			SwarmBodyEffect,
		},
	},
}

func NewCreatureSubtypeEffect(
	classification CreatureClassification,
	baseProfile CreatureTypeProfile,
) (CreatureSubtypeEffect, bool) {
	effect := creatureSubtypeEffect{
		baseProfile: baseProfile,
	}

	var added []CreatureTypeTraitID
	var removed []CreatureTypeTraitID
	var structural []CreatureSubtypeStructuralEffectID

	for _, subtype := range classification.GetSubtypes() {
		spec, ok := subtypeEffectTable[subtype]
		if !ok {
			return creatureSubtypeEffect{}, false
		}

		if spec.requiredBaseType != nil && classification.GetBaseType() != *spec.requiredBaseType {
			return creatureSubtypeEffect{}, false
		}

		added = append(added, spec.addedTraitIDs...)
		removed = append(removed, spec.removedTraitIDs...)
		structural = append(structural, spec.structuralEffectIDs...)

		if spec.preservesOriginalTypeMetadata {
			originalType, ok := classification.GetAugmentedFrom()
			if !ok {
				return creatureSubtypeEffect{}, false
			}

			originalCopy := originalType
			effect.preservedOriginalType = &originalCopy
		}
	}

	dedupedAdded, ok := dedupeSubtypeEffectTraitIDs(added)
	if !ok {
		return creatureSubtypeEffect{}, false
	}

	dedupedRemoved, ok := dedupeSubtypeEffectTraitIDs(removed)
	if !ok {
		return creatureSubtypeEffect{}, false
	}

	dedupedStructural, ok := dedupeStructuralEffectIDs(structural)
	if !ok {
		return creatureSubtypeEffect{}, false
	}

	effect.addedTraitIDs = dedupedAdded
	effect.removedTraitIDs = dedupedRemoved
	effect.structuralEffectIDs = dedupedStructural
	return effect, true
}

func (e creatureSubtypeEffect) GetBaseProfile() CreatureTypeProfile {
	return e.baseProfile
}

func (e creatureSubtypeEffect) GetAddedTraitIDs() []CreatureTypeTraitID {
	return append([]CreatureTypeTraitID(nil), e.addedTraitIDs...)
}

func (e creatureSubtypeEffect) GetRemovedTraitIDs() []CreatureTypeTraitID {
	return append([]CreatureTypeTraitID(nil), e.removedTraitIDs...)
}

func (e creatureSubtypeEffect) GetStructuralEffectIDs() []CreatureSubtypeStructuralEffectID {
	return append([]CreatureSubtypeStructuralEffectID(nil), e.structuralEffectIDs...)
}

func (e creatureSubtypeEffect) GetPreservedOriginalType() (CreatureTypeID, bool) {
	if e.preservedOriginalType == nil {
		return "", false
	}

	return *e.preservedOriginalType, true
}

func (e creatureSubtypeEffect) HasAddedTrait(traitID CreatureTypeTraitID) bool {
	for _, current := range e.addedTraitIDs {
		if current == traitID {
			return true
		}
	}

	return false
}

func (e creatureSubtypeEffect) HasRemovedTrait(traitID CreatureTypeTraitID) bool {
	for _, current := range e.removedTraitIDs {
		if current == traitID {
			return true
		}
	}

	return false
}

func (e creatureSubtypeEffect) HasStructuralEffect(effectID CreatureSubtypeStructuralEffectID) bool {
	for _, current := range e.structuralEffectIDs {
		if current == effectID {
			return true
		}
	}

	return false
}

func dedupeStructuralEffectIDs(
	effectIDs []CreatureSubtypeStructuralEffectID,
) ([]CreatureSubtypeStructuralEffectID, bool) {
	if len(effectIDs) == 0 {
		return nil, true
	}

	seen := make(map[CreatureSubtypeStructuralEffectID]struct{}, len(effectIDs))
	deduped := make([]CreatureSubtypeStructuralEffectID, 0, len(effectIDs))

	for _, effectID := range effectIDs {
		if !isValidCreatureSubtypeStructuralEffectID(effectID) {
			return nil, false
		}

		if _, ok := seen[effectID]; ok {
			continue
		}

		seen[effectID] = struct{}{}
		deduped = append(deduped, effectID)
	}

	return deduped, true
}

func dedupeSubtypeEffectTraitIDs(traitIDs []CreatureTypeTraitID) ([]CreatureTypeTraitID, bool) {
	if len(traitIDs) == 0 {
		return nil, true
	}

	seen := make(map[CreatureTypeTraitID]struct{}, len(traitIDs))
	deduped := make([]CreatureTypeTraitID, 0, len(traitIDs))

	for _, traitID := range traitIDs {
		if !isValidSubtypeEffectTraitID(traitID) {
			return nil, false
		}

		if _, ok := seen[traitID]; ok {
			continue
		}

		seen[traitID] = struct{}{}
		deduped = append(deduped, traitID)
	}

	return deduped, true
}

func isValidSubtypeEffectTraitID(value CreatureTypeTraitID) bool {
	if isValidCreatureTypeTraitID(value) {
		return true
	}

	switch value {
	case BreathesWaterTrait,
		SwimWithoutChecksTrait,
		SwimAlwaysClassSkillTrait,
		ImmunityBleedTrait,
		NotSubjectToFlankingTrait,
		PrecisionDamageImmuneTrait:
		return true
	default:
		return false
	}
}

func isValidCreatureSubtypeStructuralEffectID(value CreatureSubtypeStructuralEffectID) bool {
	switch value {
	case IncorporealBodyEffect, SwarmBodyEffect:
		return true
	default:
		return false
	}
}

func creatureTypePointer(value CreatureTypeID) *CreatureTypeID {
	return &value
}
