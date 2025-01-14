package ubi

var characters = []string{}

func init() {
	var ubiCharacters = [][]string{
		raymanCharacters,
		r6Characters,
		acCharacters,
		princeOfPersiaCharacters,
		farCryCharacters,
		heroesOfMightAndMagicCharacters,
		bgeCharacters,
		grCharacters,
		spliterCellCharacters,
		wdCharacters,
	}
	for _, chars := range ubiCharacters {
		characters = append(characters, chars...)
	}
}

var raymanCharacters = []string{
	// Rayman variations
	"Rayman",
	"Raybox",
	"Raymesis",
	"Sir Rayelot",
	"Ray Plumber",
	"Ray Vaas",
	"Raymolk",
	"Rayomz",
	"Splinter Ray",
	"Assassin Ray",
	"Ray of Persia",
	"Funky Ray",
	"Champion Ray",
	"Santa Ray",
	"Ray-mas",
	"Ray-Fucius",
	"Shadow Ray",
	"Shayman",
	"Raywolf",
	"Raylatin",
	"Rayphael",
	"Raybolt",
	"Raybeard",
	"Raydoom",
	"Freeze Ray",
	"Raysiris",
	"Dr. Farayday",
	"Samuray",
	"Ray-One",
	"Raywiz",
	"Sorceray",
	"Ray-Gnarr",
	// Globox variations
	"Globox",
	"Glombrox",
	"Sir Globrax",
	"Glob Plumber",
	"Far Glob",
	"Globolk",
	"Poglox",
	"Glob Cell",
	"Grovebox",
	"Globeach",
	"Glombie",
	"Glombrero",
	"GlobXmas",
	"Sandbox",
	// Teensy variations
	"Green Teensy",
	"Goth Teensy",
	"Teensy Ray",
	"GlobTeen",
	"Teensy Queen",
	"Teensy Wizard",
	"Grand Minimus",
	"Ninja Teensy",
	"Flaming Teensy",
	"Teensy Hermit",
	"The Golden Teensy",
	"Teen Punk",
	"Rabbiteen",
	"Teenseer",
	"Teensy Elf",
	"Teensllo",
	// Barbara variations
	"Barbara",
	"Elysia",
	"Aurora",
	"Twila",
	"Estelia",
	"Selena",
	"Ursula",
	"Emma",
	"Olympia",
	"Sibylla",
	"Avelina",
	"Holly",
	"Amazonia",
	"Snowbrina",
	"Swifterella",
	"Hellen",
	"Sakura",
	"Candice",
	"Barr-Barr",
}

var r6Characters = []string{
	"Recruit",
	"Sledge",
	"Thatcher",
	"Ash",
	"Thermite",
	"Twitch",
	"Montagne",
	"Glaz",
	"Fuze",
	"Blitz",
	"IQ",
	"Buck",
	"Blackbeard",
	"Capitão",
	"Hibana",
	"Jackal",
	"Ying",
	"Zofia",
	"Dokkaebi",
	"Lion",
	"Finka",
	"Maverick",
	"Nomad",
	"Gridlock",
	"Nøkk",
	"Amaru",
	"Kali",
	"Iana",
	"Ace",
	"Zero",
	"Flores",
	"Osa",
	"Sens",
	"Smoke",
	"Mute",
	"Castle",
	"Pulse",
	"Doc",
	"Rook",
	"Kapkan",
	"Tachanka",
	"Jäger",
	"Bandit",
	"Frost",
	"Valkyrie",
	"Caveira",
	"Echo",
	"Mira",
	"Lesion",
	"Ela",
	"Vigil",
	"Alibi",
	"Maestro",
	"Clash",
	"Kaid",
	"Mozzie",
	"Warden",
	"Goyo",
	"Wamai",
	"Oryx",
	"Melusi",
	"Aruni",
	"Thunderbird",
	"Thorn",
	"Azami",
}

var acCharacters = []string{
	// Main characters
	"Altaïr Ibn-La'Ahad",
	"Ezio Auditore da Firenze",
	"Connor",
	"Haytham Kenway",
	"Aveline de Grandpré",
	"Edward James Kenway",
	"Adéwalé",
	"Shay Patrick Cormac",
	"Arno Victor Dorian",
	"Shao Jun",
	"Arbaaz Mir",
	"Nikolai Andreievich Orelov",
	"Jacob Frye",
	"Evie Frye",
	"Bayek",
	"Kassandra",
	"Alexios",
	"Eivor",
	"Layla Hassan",
	"Basim",
	"Desmond Miles",
	// Valhalla era
	"Sigurd",
	"Aelfred",
	"Ivarr",
	"Ubba",
	"Randvi",
	"Gunnar",
	"Dag",
	"Valka",
	"Hytham",
	"Styrbjorn",
	"Harald",
	"Soma",
	"Galinn",
	"Birna",
	"Lif",
	"Fulke",
	"Halfdan",
	"Faravid",
	"Ricsige",
	"Geadric",
	"Eadwyn",
	"Ceolbert",
	"Ceolwulf II",
	"Leofrith",
	"Burgred",
	"Rhodri",
	"Deorlaf",
	"Angharad",
	"Cynon",
	"Tewdwr",
	"Modron",
	"Gwenydd",
	"Stowe",
	"Erke",
	"Kjotve",
	"Ljufvina",
	"Bjarmasdotter",
	"Hjorr",
	"Vili",
	"Hemming",
	"Tyrgve",
	"Hunwald",
	"Acha",
	"Herefrith",
	"Aelfgar",
	"Rued",
	"Brothir",
	"Broder",
	"Oswald",
	"Valdis",
	"Finnr",
	"Cynebert",
	"Birstan",
	"Estrid",
	"Alfida",
	"Rollo",
	"Guthrum",
	"Goodwin",
	// Odyssey era
	// -- Sparta
	"Leonidas",
	"Archidamos",
	"Pausanias",
	"Myrrine",
	"Nikolaos",
	"Stentor",
	"Brasidas",
	"Testikles",
	"Lysander",
	// -- Athens
	"Perikles",
	"Aspasia",
	"Kleon",
	"Sokrates",
	"Alkibiades",
	"Aristophanes",
	"Demosthenes",
	"Kephallonia",
	"Markos",
	"Phoibe",
	"Cyclops",
	// -- Others
	"Aletheia",
	"Artaxerxes",
	"Pythagoras",
	"Skoura",
	"Darius",
}

var princeOfPersiaCharacters = []string{
	"Prince",
	"Sultan",
	"Jaffar",
	"Princess",
	"King Sharaman",
	"Tus",
	"Garsiv",
	"Princess Tamina",
	"Sheik Amar",
	"Nizam",
	"Malik",
	"Sultan of Azad",
	"Maharajah",
	"Farah",
	"Kalim",
	"Vizier",
}

var farCryCharacters = []string{
	// Villans
	"Vaas Montenegro",
	"Pagan Min",
	"The Jackal",
	"Faith Seed",
	"Jacob Seed",
	"Joseph Seed",
	"Hoyt Volker",
	"John Seed",
	"Antón Castillo",
	"Harland Doyle",
	"Buck Hughes",
	"Ike Sloan",
	"Paul Harmon",
	"Yuma Lau",
	"Ull",
	"Batari",
	// Playable characters
	"Security Captain",
	"Mercenary",
	"Junior Deputy",
	"Jack Carver",
	"Takkar",
	"Dani Rojas",
	"Guapo",
	"Chorizo",
	"Ajay Ghale",
	"Jason Brody",
	"Rex Colt",
}

var heroesOfMightAndMagicCharacters = []string{
	// HoMaM 3
	"Christian",
	"Edric",
	"Orrin",
	"Sorsha",
	"Sylvia",
	"Valeska",
	"Tyris",
	"Lord Haart",
	"Catherine",
	"Roland",
	"Sir Mullich",
	"Adela",
	"Adelaide",
	"Caitlin",
	"Cuthbert",
	"Ingham",
	"Loynis",
	"Rion",
	"Sanya",
	"Clancy",
	"Ivor",
	"Jenova",
	"Kyrre",
	"Mephala",
	"Ryland",
	"Thorgrim",
	"Ufretin",
	"Gelu",
	"Aeris",
	"Alagar",
	"Coronius",
	"Elleshar",
	"Gem",
	"Malcom",
	"Melodia",
	"Uland",
	"Fafner",
	"Iona",
	"Josephine",
	"Neela",
	"Piquedram",
	"Rissa",
	"Thane",
	"Torosar",
	"Aine",
	"Astral",
	"Cyra",
	"Daremyth",
	"Halon",
	"Serena",
	"Solmyr",
	"Theodorus",
	"Dracon",
	"Calh",
	"Fiona",
	"Ignatius",
	"Marius",
	"Nymus",
	"Octavia",
	"Pyre",
	"Rashka",
	"Xeron",
	"Ash",
	"Axsis",
	"Ayden",
	"Calid",
	"Olema",
	"Xarfax",
	"Xyron",
	"Zydar",
	"Charna",
	"Clavius",
	"Galthran",
	"Isra",
	"Moandor",
	"Straker",
	"Tamika",
	"Vokial",
	"Aislinn",
	"Nagash",
	"Nimbus",
	"Sandro",
	"Septienna",
	"Thant",
	"Vidomina",
	"Xsi",
	"Ajit",
	"Arlach",
	"Dace",
	"Damacon",
	"Gunnar",
	"Lorelei",
	"Shakti",
	"Synca",
	"Mutare",
	"Mutare Drake",
	"Alamar",
	"Darkstorn",
	"Deemer",
	"Geon",
	"Jaegar",
	"Jeddite",
	"Malekith",
	"Sephinroth",
	"Crag Hack",
	"Gretchin",
	"Gurnisson",
	"Jabarkas",
	"Krellion",
	"Shiva",
	"Tyraxor",
	"Yog",
	"Boragus",
	"Kilgor",
	"Dessa",
	"Gird",
	"Gundula",
	"Oris",
	"Saurug",
	"Terek",
	"Vey",
	"Zubin",
	"Alkin",
	"Broghild",
	"Bron",
	"Drakon",
	"Gerwulf",
	"Korbac",
	"Tazar",
	"Wystan",
	"Andra",
	"Merist",
	"Mirlanda",
	"Rosic",
	"Styg",
	"Tiva",
	"Verdish",
	"Voy",
	"Adrienne",
	"Erdamon",
	"Fiur",
	"Ignissa",
	"Kalt",
	"Lacus",
	"Monere",
	"Pasis",
	"Thunar",
	"Aenain",
	"Brissa",
	"Ciele",
	"Gelare",
	"Grindan",
	"Inteus",
	"Labetha",
	"Luna",
}

var bgeCharacters = []string{
	"Jade",
	// Companions
	"Pey'j",
	"Double H",
	// Allies
	"Hahn",
	"Meï",
	"Governor of Hillys",
	"Nino",
	"Secundo",
	"Peepers",
	// Antagonists
	"General Kehck",
	"High Priest",
	// Hillys Citizens
	"Ming-Tzu",
	"Nouri",
	"Xiao",
	"Mo",
	"Francis",
	"Rufus",
	"Seven",
	"Chico",
	"Rob",
	"Fehn Digler",
	"Mamma",
	"Issam",
	"Hal",
	"Babukar",
	// Residents of The Lighthouse Shelter
	"Yoa",
	"Zaza",
	"Oumi",
	"Fehn",
	"Pablo",
	"Kip",
	"Woof",
}

var grCharacters = []string{
	// Known characters
	"Nomad",
	"Cole D. Walker",
	"Midas",
	"Weaver",
	"Holt",
	"John Kozak",
	"El Sueño",
	"Karen Bowman",
}

var spliterCellCharacters = []string{
	"Andrei V. Alekseevich",
	"Bagrat",
	"John Baxter",
	"Robert Blaustein",
	"Bobrov",
	"David Bowers",
	"Frances Coen",
	"Varlam Cristavi",
	"David Tanahill",
	"Mitchell Dougherty",
	"Sam Fisher",
	"Sarah Fisher",
	"Bartholomew Fisk",
	"Goa",
	"Vyacheslav Grinko",
	"Anna Grímsdóttir",
	"Thomas Gurgenidze",
	"Hamlet",
	"Ivan",
	"Kong Feirong",
	"Irving Lambert",
	"Long Dan",
	"Alice Madison",
	"Philip Masse",
	"Murman Kobiashvili",
	"Kombayn Nikoladze",
	"Nosenko",
	"Morris Odell",
	"Piotr Lejava",
	"Ramaz Lortkipanidze",
	"Vernon Wilkes, Jr.",
}

var wdCharacters = []string{
	"Abigail Vega",
	"Adolpho C. Peccorino",
	"Aiden Pearce",
	"Alejandro Mendez",
	"Angela Balik",
	"Angelo Bryan",
	"Angelo Tucci",
	"Anton Chenkov",
	"Arvid Stegman",
	"Bedbug",
	"Candy Amos",
	"Carl Breenwood",
	"Charlotte Gardner",
	"Clara Lille",
	"Damien Brenks",
	"Danny SoSueMe",
	"Darius Peaston",
	"David R. Treaklesen",
	"Delford Wade",
	"Dermot Quinn",
	"Desipio",
	"Donna Dean",
	"Donovan Rushmore",
	"Edgar Noone",
	"Francis T. Barrthes",
	"Frank Janson",
	"Gary Diggs",
	"Gregory Foydalem",
	"Helena Tucci",
	"Jackson Pearce",
	"JB Markowicz",
	"Jedidiah Ferguson",
	"Jen Kramer",
	"Jimmy Baumbach",
	"Jordi Chin",
	"Joseph DeMarco",
	"Joshua Kramer",
	"Joy Ndidi Adebayo",
	"Lance Brenner",
	"Lena Pearce",
	"Malcolm Deodato",
	"Marcus Brenks",
	"Martin Graften",
	"Mary Blass",
	"Maurice Vega",
	"Michael Orleans",
	"Nicholas Crispin",
	"Nicole Pearce",
	"Paul Benedict Henfield",
	"Raul Lionzo",
	"Raymond Kenney",
	"Robert Racine",
	"Rose Washington",
	"Sandford Amos",
	"Tobias Frewer",
	"Tyrone Hayes",
	"Wreck",
	"Yolanda Mendez",
}
