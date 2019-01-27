package humphrey

import (
	"github.com/mastertinner/humphrey/internal/app/humphrey/img"
	"github.com/mastertinner/humphrey/pkg/game"
)

var escape = &game.Scene{
	Image: img.Box,
	Body:  "Mou, sicher. I bi aber nid ganz sicher wie. Da obe gsehsch mis Revier woni muess usbrächä. Da bruuchi dini Hilf! Bisch derbii?",
	Actions: map[string]game.Renderer{
		"Ja, eh!":      escape2,
		"Ke Luscht...": &game.End{},
	},
}

var escape2 = &game.Scene{
	Image: img.Sitting,
	Body:  "Wi machemers am beste?",
	Actions: map[string]game.Renderer{
		"Zersch mau öppis ässä!":     eat,
		"Ds ganze Revier umgrabe":    dig,
		"Mau chli ar Schibe chratze": scratchGlass,
	},
}

var eat = &game.Scene{
	Image: img.Eating,
	Body:  "Ok. Da im Egge ligt no chli Salat 🥬. Mmh, mmh, mmh! Iz müessemer nis aber ad Arbeit mache!",
	Actions: map[string]game.Renderer{
		"Ds ganze Revier umgrabe":    dig,
		"Mau chli ar Schibe chratze": scratchGlass,
	},
}

var dig = &game.Scene{
	Image: img.Shovel,
	Body:  "Dr Humphrey grabt stundelang sis ganze Revier um und bout sech e nöie Tunnel. När ischer ganz müed und duet zersch mau e Rundi schlafe... 🛌 Woner wider ufwacht erinneret er sech wider daser eingch mau het wöue usbrächä.",
	Actions: map[string]game.Renderer{
		"Doch mau chli ar Schibe chratze": scratchGlass,
		"No chli meh schlafe...":          &game.End{},
	},
}

var scratchGlass = &game.Scene{
	Image: img.Sitting,
	Body:  `Dr Humphy chratzet ar Schibe und scho glii drufabe chunnt d Säri agrennt. "Was isch de los, Humpheli? 😱 Lug, Nino, dr Humphy chratzet ar Schibe." Da steckt si ihri Hand ine um dr Humphrey z striichle. Was machemer itz?`,
	Actions: map[string]game.Renderer{
		"I liebe striichle! Lahmse mi la striichle!": pet,
		"I finger biisse": bite,
		"Das isch üsi Chance! Schnäu ar Hand usechlättere": climbOut,
	},
}

var pet = &game.Scene{
	Image: img.Hand,
	Body:  "D Säri striichlet dr Humphrey und är isch ganz glücklech 🐹 Leider bringt üs das üsem Ziel aber nid viu necher... Actions, we need actions!",
	Actions: map[string]game.Renderer{
		"Es undankbars Gschöpf sii und i Finger biisse!":     bite,
		"Vorsichtig ar Hand probiere usezchlättere":          climbOut,
		"Werum usbrächä weme sech no meh cha la striichle??": &game.End{},
	},
}

var bite = &game.Scene{
	Image: img.Hand,
	Body:  `Dr Humphy biisst d Säri mit voller Chraft i Finger. "Aaaau!" schreit si und fluechet vor sech härä 🤬. Ds Revier vom Humphrey isch aber no offe. Das isch doch die Chance!`,
	Actions: map[string]game.Renderer{
		"Schnäu usegumpe!": jumpOut,
		"Iz hani es schlächts Gwüsse... Schnäu vergrabe": &game.End{},
	},
}

var jumpOut = &game.Scene{
	Image: img.Box,
	Body:  "Dr Humphrey gumpet so schnäu winer cha us sim Revier und chlätteret a Bode. Vo dert us rennt er schäu am Nino verbii und ume Egge. Phuh, gschafft! Iz ganz still sii 🤫",
	Actions: map[string]game.Renderer{
		"Ke Mucks mache und im Egge blibe":          escaped,
		"Es isch gfürchig hie! Nach dr Säri rüefe!": &game.End{},
	},
}

var climbOut = &game.Scene{
	Image: img.Box,
	Body:  "Dr Humphrey chlätteret so schnäu winer cha am Arm vor Säri us sim Revier und gumpet a Bode. Vo dert us rennt er schäu am Nino verbii und ume Egge. Phuh, gschafft! Iz ganz still sii 🤫",
	Actions: map[string]game.Renderer{
		"Ke Mucks mache und im Egge blibe":          escaped,
		"Es isch gfürchig hie! Nach dr Säri rüefe!": &game.End{},
	},
}

var escaped = &game.Scene{
	Image: img.Sitting,
	Body:  "I gloub es gseht guet us. D Säri und dr Nino si mi im andere Ruum am sueche. Bündigi Sach! De chöimer iz hie mau chli afah erkunde. Wo weimer düre?",
	Actions: map[string]game.Renderer{
		"Unterem Bett düre":  groej,
		"Id Chuchi go luege": groej,
		"Ids Wohnzimmer gah": groej,
	},
}
