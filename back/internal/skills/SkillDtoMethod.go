package skills

// ToSummaryDTO Convertit une compétence complète en version résumée
func (s *SkillDTO) ToSummaryDTO() SkillSummaryDTO {
	return SkillSummaryDTO{
		ID:                   s.ID,
		Type:                 s.Type,
		Code:                 s.Code,
		Libelle:              s.Libelle,
		RiasecMajeur:         s.RiasecMajeur,
		RiasecMineur:         s.RiasecMineur,
		TransitionEcologique: s.TransitionEcologique,
		TransitionNumerique:  s.TransitionNumerique,
		Obsolete:             s.Obsolete,
		SousCategorie:        s.SousCategorie,
		Transferable:         s.Transferable,
	}
}
