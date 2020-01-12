package humphrey

import (
	"github.com/mastertinner/humphrey/internal/app/humphrey/img"
	"github.com/mastertinner/humphrey/pkg/game"
)

var groej = &game.Scene{
	Image: img.Dog,
	Body:  "Dr Humphrey louft witer und gseht uf ds Mau e Hund im Egge stah. Dr Hund gseht chli gfährlech us aber dr Humphy necheret sech ihm langsam und vorsichtig.",
	Actions: map[string]game.Renderer{
		"No necher zum Hund häregah 🐶":         groej2,
		"Vorsichtig hingerem Egge vüreluege 😰": groej2,
	},
}

var groej2 = &game.Scene{
	Image: img.Dog,
	Body:  "Dr Hund gseht dr chlii Hamster rächt schnäll und chunnt zuenim härä. Är luegt ne dummbatzig a, häbt d Zunge use und leit dr Chopf schreg.",
	Actions: map[string]game.Renderer{
		`"Hallo, I bi dr Humphrey"`: groej3,
		`"Hilfeeeee!"`:              groej3,
	},
}

var groej3 = &game.Scene{
	Image: img.Dog,
	Body:  `Dr Hund stellt sech vor: "Gröj, mi Name. Was machsch de du hie am Bode?"`,
	Actions: map[string]game.Renderer{
		`"I bi us mim Revier usbroche um mal d Wohnig chli z erkunde"`: groejExplore,
	},
}

var groejExplore = &game.Scene{
	Image: img.Dog,
	Body:  `"De chöimer ja zämä chli erkunde." seit dr Gröj.`,
	Actions: map[string]game.Renderer{
		`"Ja eh. Machemer doch."`:       groejTogether,
		`"Nei, i wett lieber allei..."`: groejAlone,
	},
}

var groejAlone = &game.Scene{
	Image: img.Together,
	Body:  `"Ok. De loufi dr eifach chli nache" antwortet dr Gröj und setzt sech hintere Humphrey.`,
	Actions: map[string]game.Renderer{
		`"Ääh... Ok... de göhmer"`: groejTogether,
	},
}

var groejTogether = &game.Scene{
	Image: img.Together,
	Body:  "Di zwe loufe los u chöme scho glii wider im Wohnzimmer verbii. Dert hocke d Säri u dr Nino und luege grad di nöisti Folg 'Sex Education' uf Netflix.",
	Actions: map[string]game.Renderer{
		"Zur Säri gah 👩": approachSarah,
		"Zum Nino gah 👱": approachNino,
	},
}
