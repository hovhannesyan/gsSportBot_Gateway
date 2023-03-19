package telegram

func (b *Bot) getDisciplineName(disciplineId int64) string {
	res := b.eventClient.GetDisciplines(disciplineId)
	return res[0].Name
}
