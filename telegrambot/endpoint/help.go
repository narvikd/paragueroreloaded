package endpoint

func getHelp() string {
	const commands = "<b>Usuarios:</b>\n" +
		"<b>/help</b>\n" +
		"<b>/links</b>\n" +
		"<b>/java</b>\n" +
		"<b>abro paraguas (y la respuesta a su sticker)</b>\n" +
		"<b>pole (a las 00:00)</b>\n" +
		"<b>mensaje de bienvenida cuando entra un usuario</b>\n" +
		"\n" +
		"<b>Coñas:</b>\n" +
		"<b>/ban</b>\n" +
		"<pre>" +
		"/rust\n" +
		"/kotlin\n" +
		"/andalu frase\n" +
		"/choni frase\n" +
		"/sigarro frase\n" +
		"</pre>" +
		"\n" +
		"<b>Admins:</b>\n" +
		"<pre>" +
		"/daw frase\n" +
		"/group groupID frase\n" +
		"</pre>" +
		"\n" +
		"<b>Admin Debug:</b>\n" +
		"<b>/time</b>\n" +
		"<b>/chatid</b>\n"
	return commands
}
