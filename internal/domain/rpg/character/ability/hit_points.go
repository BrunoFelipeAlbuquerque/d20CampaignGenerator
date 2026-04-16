package ability

type hpSource struct {
	name  string
	value int
}
type HPSource = hpSource

type hitPointKind string
type HitPointKind = hitPointKind

type hitDieType string
type HitDieType = hitDieType

const (
	StandardHitPoints  HitPointKind = "Standard"
	UndeadHitPoints    HitPointKind = "Undead"
	ConstructHitPoints HitPointKind = "Construct"
)

const (
	D6HitDie  HitDieType = "d6"
	D8HitDie  HitDieType = "d8"
	D10HitDie HitDieType = "d10"
	D12HitDie HitDieType = "d12"
)

type hitDie struct {
	total int
	d6    int
	d8    int
	d10   int
	d12   int
}
type HitDie = hitDie

type hitPoints struct {
	total            int
	current          int
	temporary        int
	nonLethal        int
	deathThreshold   int
	hd               hitDie
	sources          []hpSource
	temporarySources []hpSource
	activeTemporary  string
	kind             hitPointKind
	size             size
	constitution     int
	charisma         int
}
type HitPoints = hitPoints

func NewHPSource(name string, value int) (HPSource, bool) {
	if len(name) == 0 || value < 0 {
		return hpSource{}, false
	}

	return hpSource{
		name:  name,
		value: value,
	}, true
}

func NewHitDie(d6 int, d8 int, d10 int, d12 int) (HitDie, bool) {
	if d6 < 0 || d8 < 0 || d10 < 0 || d12 < 0 {
		return hitDie{}, false
	}

	total := d6 + d8 + d10 + d12
	if total <= 0 {
		return hitDie{}, false
	}

	return hitDie{
		total: total,
		d6:    d6,
		d8:    d8,
		d10:   d10,
		d12:   d12,
	}, true
}

func NewStandardHitPoints(hd HitDie, constitutionScore int) (HitPoints, bool) {
	if !isValidLivingConstitutionScore(constitutionScore) || !isSemanticallyValidHitDie(hd) {
		return hitPoints{}, false
	}

	hp := hitPoints{
		hd:           hd,
		kind:         StandardHitPoints,
		constitution: constitutionScore,
	}
	hp.recalculate(true)
	return hp, true
}

func NewUndeadHitPoints(hd HitDie, charismaScore int) (HitPoints, bool) {
	if charismaScore < 0 || !isSemanticallyValidHitDie(hd) {
		return hitPoints{}, false
	}

	hp := hitPoints{
		hd:       hd,
		kind:     UndeadHitPoints,
		charisma: charismaScore,
	}
	hp.recalculate(true)
	return hp, true
}

func NewConstructHitPoints(hd HitDie, size Size) (HitPoints, bool) {
	if !isValidSize(size) || !isSemanticallyValidHitDie(hd) {
		return hitPoints{}, false
	}

	hp := hitPoints{
		hd:   hd,
		kind: ConstructHitPoints,
		size: size,
	}
	hp.recalculate(true)
	return hp, true
}

func (s hpSource) GetName() string {
	return s.name
}

func (s hpSource) GetValue() int {
	return s.value
}

func (h hitDie) GetTotal() int {
	return h.total
}

func (h hitDie) GetDieCount(kind HitDieType) (int, bool) {
	switch kind {
	case D6HitDie:
		return h.d6, true
	case D8HitDie:
		return h.d8, true
	case D10HitDie:
		return h.d10, true
	case D12HitDie:
		return h.d12, true
	default:
		return 0, false
	}
}

func (h hitDie) GetAverageBaseHP() int {
	return (h.d6 * 4) + (h.d8 * 5) + (h.d10 * 6) + (h.d12 * 7)
}

func (h hitPoints) GetTotal() int {
	return h.total
}

func (h hitPoints) GetCurrent() int {
	return h.current
}

func (h hitPoints) GetTemporary() int {
	return h.temporary
}

func (h hitPoints) GetNonLethal() int {
	return h.nonLethal
}

func (h hitPoints) GetDeathThreshold() int {
	return h.deathThreshold
}

func (h hitPoints) GetHitDie() HitDie {
	return h.hd
}

func (h hitPoints) GetSources() []HPSource {
	return append([]hpSource(nil), h.sources...)
}

func (h hitPoints) GetTemporarySources() []HPSource {
	return append([]hpSource(nil), h.temporarySources...)
}

func (h hitPoints) GetKind() HitPointKind {
	return h.kind
}

func (h hitPoints) IsNonLethalImmune() bool {
	return h.kind == UndeadHitPoints || h.kind == ConstructHitPoints
}

func (h *hitPoints) SetTemporaryHPSource(name string, value int) bool {
	if value < 0 || len(name) == 0 {
		return false
	}

	for i, source := range h.temporarySources {
		if source.name != name {
			continue
		}

		h.temporarySources[i].value = value
		h.recalculateTemporary(false)
		return true
	}

	h.temporarySources = append(h.temporarySources, hpSource{name: name, value: value})
	h.recalculateTemporary(false)
	return true
}

func (h *hitPoints) RemoveTemporaryHPSource(name string) bool {
	for i, source := range h.temporarySources {
		if source.name != name {
			continue
		}

		h.temporarySources = append(h.temporarySources[:i], h.temporarySources[i+1:]...)
		if h.activeTemporary == name {
			h.activeTemporary = ""
		}
		h.recalculateTemporary(false)
		return true
	}

	return false
}

func (h *hitPoints) TakeDamage(amount int, isNonLethal bool) bool {
	if amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	remaining := h.consumeTemporaryHP(amount)

	if remaining == 0 {
		return true
	}

	if !isNonLethal {
		h.current -= remaining
		return true
	}

	if h.IsNonLethalImmune() {
		return true
	}

	h.nonLethal += remaining
	return true
}

func (h *hitPoints) Heal(amount int) bool {
	if amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	oldCurrent := h.current
	h.current += amount
	if h.current > h.total {
		h.current = h.total
	}

	healed := h.current - oldCurrent
	if healed > 0 {
		h.nonLethal -= healed
		if h.nonLethal < 0 {
			h.nonLethal = 0
		}
	}

	return true
}

func (h *hitPoints) UpdateConstitutionScore(score int) bool {
	if !isValidLivingConstitutionScore(score) || h.kind != StandardHitPoints {
		return false
	}

	h.constitution = score
	h.recalculate(false)
	return true
}

func isValidLivingConstitutionScore(score int) bool {
	return score >= 1
}

func (h *hitPoints) UpdateCharismaScore(score int) bool {
	if score < 0 || h.kind != UndeadHitPoints {
		return false
	}

	h.charisma = score
	h.recalculate(false)
	return true
}

func (h *hitPoints) UpdateSize(value Size) bool {
	if !isValidSize(value) || h.kind != ConstructHitPoints {
		return false
	}

	h.size = value
	h.recalculate(false)
	return true
}

func (h *hitPoints) recalculate(initial bool) {
	sources, total, deathThreshold := h.resolveCoreState()
	oldTotal := h.total

	h.sources = sources
	h.total = total
	h.deathThreshold = deathThreshold

	if initial {
		h.current = total
	} else {
		h.current += total - oldTotal
		if h.current > h.total {
			h.current = h.total
		}
	}
}

func (h hitPoints) resolveCoreState() ([]hpSource, int, int) {
	baseDice := h.hd.GetAverageBaseHP()
	sources := []hpSource{
		{name: "Base Dice", value: baseDice},
	}

	total := baseDice
	minimumHitPoints := h.hd.total

	switch h.kind {
	case UndeadHitPoints:
		charismaBonus := getAbilityModifier(h.charisma) * h.hd.total
		sources = append(sources, hpSource{name: "Charisma (Undead)", value: charismaBonus})
		total += charismaBonus
		total = applyMinimumHitPointFloor(total, minimumHitPoints, &sources)
		return sources, total, 0

	case ConstructHitPoints:
		sizeBonus, _ := h.size.GetConstructBonusHP()
		sources = append(sources, hpSource{name: "Construct Size Bonus", value: sizeBonus})
		total += sizeBonus
		return sources, total, 0

	default:
		constitutionBonus := getAbilityModifier(h.constitution) * h.hd.total
		sources = append(sources, hpSource{name: "Constitution", value: constitutionBonus})
		total += constitutionBonus
		total = applyMinimumHitPointFloor(total, minimumHitPoints, &sources)
		return sources, total, -h.constitution
	}
}

func (h *hitPoints) recalculateTemporary(preserveActive bool) {
	if len(h.temporarySources) == 0 {
		h.temporary = 0
		h.activeTemporary = ""
		return
	}

	if preserveActive {
		if activeIndex, ok := h.findTemporarySourceIndexByName(h.activeTemporary); ok {
			h.temporary = h.temporarySources[activeIndex].value
			return
		}
	}

	activeIndex := h.findHighestTemporarySourceIndex()
	h.activeTemporary = h.temporarySources[activeIndex].name
	h.temporary = h.temporarySources[activeIndex].value
}

func (h *hitPoints) consumeTemporaryHP(amount int) int {
	if h.temporary == 0 || len(h.temporarySources) == 0 {
		return amount
	}

	activeIndex, ok := h.findTemporarySourceIndexByName(h.activeTemporary)
	if !ok {
		h.recalculateTemporary(false)
		activeIndex, ok = h.findTemporarySourceIndexByName(h.activeTemporary)
		if !ok {
			return amount
		}
	}

	remaining := amount
	if h.temporarySources[activeIndex].value <= remaining {
		remaining -= h.temporarySources[activeIndex].value
		h.temporarySources = append(h.temporarySources[:activeIndex], h.temporarySources[activeIndex+1:]...)
		h.activeTemporary = ""
		h.recalculateTemporary(false)
		return remaining
	}

	h.temporarySources[activeIndex].value -= remaining
	remaining = 0
	h.recalculateTemporary(true)
	return remaining
}

func applyMinimumHitPointFloor(total int, minimum int, sources *[]hpSource) int {
	if total >= minimum {
		return total
	}

	adjustment := minimum - total
	*sources = append(*sources, hpSource{name: "Minimum 1 HP per Hit Die", value: adjustment})
	return minimum
}

func getAbilityModifier(score int) int {
	delta := score - 10
	if delta >= 0 || delta%2 == 0 {
		return delta / 2
	}

	return (delta / 2) - 1
}

func isSemanticallyValidHitDie(hd HitDie) bool {
	if hd.total <= 0 || hd.d6 < 0 || hd.d8 < 0 || hd.d10 < 0 || hd.d12 < 0 {
		return false
	}

	return hd.total == hd.d6+hd.d8+hd.d10+hd.d12
}

func (h hitPoints) findTemporarySourceIndexByName(name string) (int, bool) {
	if len(name) == 0 {
		return 0, false
	}

	for i, source := range h.temporarySources {
		if source.name == name {
			return i, true
		}
	}

	return 0, false
}

func (h hitPoints) findHighestTemporarySourceIndex() int {
	highestIndex := 0
	for i := 1; i < len(h.temporarySources); i++ {
		if h.temporarySources[i].value > h.temporarySources[highestIndex].value {
			highestIndex = i
		}
	}

	return highestIndex
}
