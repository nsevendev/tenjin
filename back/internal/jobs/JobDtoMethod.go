package jobs

// ToSummaryDTO Convertit un métier complet en version résumée
func (j *JobDTO) ToSummaryDTO() JobSummaryDTO {
	return JobSummaryDTO{
		ID:                      j.ID,
		Code:                    j.Code,
		Libelle:                 j.Libelle,
		Definition:              j.Definition,
		RiasecMajeur:            j.RiasecMajeur,
		RiasecMineur:            j.RiasecMineur,
		TransitionEcologique:    j.TransitionEcologique,
		TransitionNumerique:     j.TransitionNumerique,
		TransitionDemographique: j.TransitionDemographique,
		EmploiCadre:             j.EmploiCadre,
		EmploiReglemente:        j.EmploiReglemente,
		DomaineProfessionnel:    j.DomaineProfessionnel,
	}
}
