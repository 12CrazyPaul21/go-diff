package diffmatchpatch

type ExposedPatch struct {
	Diffs   []Diff // exposed
	Start1  int
	Start2  int
	Length1 int
	Length2 int
}

func (p *ExposedPatch) Convert() Patch {
	return Patch{
		diffs:   p.Diffs,
		Start1:  p.Start1,
		Start2:  p.Start2,
		Length1: p.Length1,
		Length2: p.Length2,
	}
}

type ExposedPatches []ExposedPatch

func (ep ExposedPatches) Convert() []Patch {
	patches := []Patch{}
	for _, p := range ep {
		patches = append(patches, p.Convert())
	}
	return patches
}
