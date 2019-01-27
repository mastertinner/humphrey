package humphrey

import (
	"github.com/mastertinner/humphrey/internal/app/humphrey/img"
	"github.com/mastertinner/humphrey/pkg/game"
)

var approachNino = &game.Scene{
	Image: img.Eating,
	Body:  `"Säri, lug mau wän i gfunde ha!", seit dr Nino und chunnt di beide cho striichle. D Säri chunnt sofort ou. Si si froh dass di beide wider zrügg gfunde hei.`,
	Actions: map[string]game.Renderer{
		"Uf Säri's Hand chrabble": onSarahsHand,
		"Uf Nino's Hand chrabble": onNinosHand,
	},
}

var approachSarah = &game.Scene{
	Image: img.Eating,
	Body:  `"Lug, Nino. Dr Humphy isch wider da. Und är het sogar e Kolleg gfunde!", seit d Säri und chunnt di beide cho striichle. Dr Nino chunnt sofort ou. Si si froh dass di beide wider zrügg gfunde hei.`,
	Actions: map[string]game.Renderer{
		"Uf Säri's Hand chrabble": onSarahsHand,
		"Uf Nino's Hand chrabble": onNinosHand,
	},
}

var onSarahsHand = &game.Scene{
	Image: img.Sitting,
	Body:  "D Säri treit dr chlii Hamster wider zrügg i sis Revier. Är isch glücklech wider deheim z sii nach all dene Strapaze.",
	Actions: map[string]game.Renderer{
		"Öppis ässe": eatFin,
		"Go schlafe": sleep,
	},
}

var onNinosHand = &game.Scene{
	Image: img.Sitting,
	Body:  "Dr Nino treit dr chlii Hamster wider zrügg i sis Revier. Är isch glücklech wider deheim z sii nach all dene Strapaze.",
	Actions: map[string]game.Renderer{
		"Öppis ässe": eatFin,
		"Go schlafe": sleep,
	},
}

var eatFin = &game.Scene{
	Image: img.Eating,
	Body:  "Es git no es feins Rüebli 🥕 zur Belohnig für ds Tierli. När ischs aber müed und geit langsam go lige.",
	Actions: map[string]game.Renderer{
		"Go schlafe": sleep,
	},
}

var sleep = &game.Scene{
	Image: img.Eating,
	Body:  "Dr Humphy geit müed i si Bou und leit sech hära. Phuh, das isch e asträngende Tag gsi woner viu entdeckt het! Morn de wider!",
	Actions: map[string]game.Renderer{
		`"Schlaf guet, Humphy!"`: &game.End{},
	},
}
