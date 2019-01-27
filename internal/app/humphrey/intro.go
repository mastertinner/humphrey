package humphrey

import (
	"github.com/mastertinner/humphrey/internal/app/humphrey/img"
	"github.com/mastertinner/humphrey/pkg/game"
)

var manual = &game.Scene{
	Image: img.Console,
	Body:  "Willkomme zu däm interaktive Textabetüür vomne chliine Hamster. Um z spile, muesch eifach aube mit de Pfilitaste dini nächsti Antwort uswählä u när Enter drücke. Lug, so:",
	Actions: map[string]game.Renderer{
		"I has chegget 💡":   intro,
		"I chume druus 👌":   intro,
		"I weiss wi geit 😎": intro,
	},
}

var intro = &game.Scene{
	Image: img.Sitting,
	Body:  "Woni no e ganz chliine Hamster bi gsi bini zimlech oft allei gsi. Di meiste vo mine Gschwüsterti si elter gsi als ig und hei albe ohni mi gspilt. Aber i bi trotzdäm es sehr glücklechs Hämsterli gsi. Als intelligänts Wäsä hani mi guet chönne sälber beschäftige und ha vorallem mine Eltere ghulfe üse Bou iizrichte. Mi Name isch Humphrey aber chaschmer ou Humphy sägä.",
	Actions: map[string]game.Renderer{
		"Hoi Humphy 🐹": intro2,
	},
}

var intro2 = &game.Scene{
	Image: img.Eating,
	Body:  "Woni chli grösser bi gsi, ischs när drum gange dass mir Gschwüsterti ufteilt wärdä. Eis nachem andere sisi wäggange, di andere Hamster, nur ig bi am Schluss no allei dert gsi. Aber das isch scho OK. Mir Hamster si ja gärn allei. Eines Tages si när doch no zwöi cho luege wosech für mi interessiert hei 👱‍♀️🧑🏽. I ha zersch dänkt es sige chli komischi aber si si när eingch no lieb gsi. I gloub di einti hani rächt schnäu ume Finger gwicklet gha. Dasch gar nid so schwirig. Eifach chli lieb driluege und a öppisem umeknäbberle. De chunnt das scho guet.",
	Actions: map[string]game.Renderer{
		"Ächt? 😱 Ds funktioniert??": intro3,
	},
}

var intro3 = &game.Scene{
	Image: img.Eating,
	Body:  "Klar! Es isch di sprichwörtlech kitschig-romantischi Liebi ufe erscht Blick gsi 🥰. Di heimi grad sofort mitgnoh und i ihri Wohnig gstellt. I bi denn nonid so ganz sicher gsi wasi vo däm söu halte aber i bi mau mitgange. Vorauem wöusi mr ab und zue öppis feins hei gäh zum hinter d Bäckli stecke 🥕 Es hetmr när sehr guet afah gfalle bi dene! Si si mega liebi und düemer di ganz zit mis Revier putze.",
	Actions: map[string]game.Renderer{
		"Hesch nonie wölle usbrächä?": escape,
	},
}
