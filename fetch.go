package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const SqlduckCreate = `CREATE SEQUENCE amenity_season_id_seq START 1;
	CREATE TABLE amenity_season (
		id UINTEGER DEFAULT nextval('amenity_season_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(63) NOT NULL
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE SEQUENCE duration_id_seq START 1;
	CREATE TABLE duration (
		id UINTEGER DEFAULT nextval('duration_id_seq') PRIMARY KEY NOT NULL,
		name CHAR(1) NOT NULL
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE SEQUENCE fee_id_seq START 1;
	CREATE TABLE fee (
		id UINTEGER DEFAULT nextval('fee_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(63) NOT NULL
	);
	INSERT INTO fee (id, name) VALUES (0, 'None');
	CREATE SEQUENCE state_code_id_seq START 1;
	CREATE TABLE state_code (
		id UINTEGER DEFAULT nextval('state_code_id_seq') PRIMARY KEY NOT NULL,
		name CHAR(2) NOT NULL
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');`
const SqliteCreate = `CREATE TABLE amenity_season (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE TABLE duration (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE TABLE fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO fee (id, name) VALUES (0, 'None');
	CREATE TABLE state_code (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');`

var amenitySeasons = map[string]uint16{
	"No":               1,
	"Yes - seasonal":   2,
	"Yes - year round": 3,
}

var durations = map[string]uint16{
	"d": 1,
	"h": 2,
	"m": 3,
}

var parkCodes = map[string]uint16{
	"abli": 1,
	"acad": 2,
	"adam": 3,
	"afam": 4,
	"afbg": 5,
	"agfo": 6,
	"alag": 7,
	"alca": 8,
	"aleu": 9,
	"alfl": 10,
	"alka": 11,
	"alpo": 12,
	"amch": 13,
	"amis": 14,
	"amme": 15,
	"anac": 16,
	"anch": 17,
	"ande": 18,
	"ania": 19,
	"anjo": 20,
	"anti": 21,
	"apco": 22,
	"apis": 23,
	"appa": 24,
	"arch": 25,
	"arho": 26,
	"arpo": 27,
	"asis": 28,
	"azru": 29,
	"badl": 30,
	"band": 31,
	"bawa": 32,
	"bela": 33,
	"beol": 34,
	"bepa": 35,
	"bibe": 36,
	"bica": 37,
	"bicr": 38,
	"bicy": 39,
	"biho": 40,
	"bisc": 41,
	"biso": 42,
	"bith": 43,
	"blca": 44,
	"blri": 45,
	"blrv": 46,
	"blsc": 47,
	"blue": 48,
	"boaf": 49,
	"boha": 50,
	"bost": 51,
	"bowa": 52,
	"brca": 53,
	"brcr": 54,
	"brvb": 55,
	"buff": 56,
	"buis": 57,
	"buov": 58,
	"cabr": 59,
	"cach": 60,
	"cacl": 61,
	"caco": 62,
	"cagr": 63,
	"caha": 64,
	"cahi": 65,
	"cajo": 66,
	"cakr": 67,
	"cali": 68,
	"calo": 69,
	"came": 70,
	"camo": 71,
	"cana": 72,
	"cane": 73,
	"cany": 74,
	"care": 75,
	"cari": 76,
	"carl": 77,
	"casa": 78,
	"cato": 79,
	"cave": 80,
	"cavo": 81,
	"cawo": 82,
	"cbpo": 83,
	"cebe": 84,
	"cebr": 85,
	"cech": 86,
	"cham": 87,
	"chat": 88,
	"chch": 89,
	"chcu": 90,
	"chic": 91,
	"chir": 92,
	"chis": 93,
	"choh": 94,
	"chpi": 95,
	"chri": 96,
	"chsc": 97,
	"chyo": 98,
	"ciro": 99,
	"clba": 100,
	"coga": 101,
	"colm": 102,
	"colo": 103,
	"colt": 104,
	"cong": 105,
	"coro": 106,
	"cowp": 107,
	"crla": 108,
	"crmo": 109,
	"cuga": 110,
	"cuis": 111,
	"cure": 112,
	"cuva": 113,
	"cwdw": 114,
	"daav": 115,
	"ddem": 116,
	"dena": 117,
	"depo": 118,
	"deso": 119,
	"deto": 120,
	"deva": 121,
	"dewa": 122,
	"dino": 123,
	"drto": 124,
	"ebla": 125,
	"edal": 126,
	"edis": 127,
	"efmo": 128,
	"eise": 129,
	"elca": 130,
	"elis": 131,
	"elma": 132,
	"elmo": 133,
	"elro": 134,
	"elte": 135,
	"euon": 136,
	"ever": 137,
	"feha": 138,
	"fiis": 139,
	"fila": 140,
	"flfo": 141,
	"flni": 142,
	"fobo": 143,
	"fobu": 144,
	"foda": 145,
	"fodo": 146,
	"fodu": 147,
	"fofo": 148,
	"fofr": 149,
	"fola": 150,
	"fols": 151,
	"foma": 152,
	"fomc": 153,
	"fomr": 154,
	"fone": 155,
	"fopo": 156,
	"fopu": 157,
	"fora": 158,
	"fosc": 159,
	"fosm": 160,
	"fost": 161,
	"fosu": 162,
	"foth": 163,
	"foun": 164,
	"fous": 165,
	"fova": 166,
	"fowa": 167,
	"frde": 168,
	"frdo": 169,
	"frhi": 170,
	"frla": 171,
	"frri": 172,
	"frsp": 173,
	"frst": 174,
	"gaar": 175,
	"gari": 176,
	"gate": 177,
	"gegr": 178,
	"gero": 179,
	"gett": 180,
	"gewa": 181,
	"gicl": 182,
	"glac": 183,
	"glba": 184,
	"glca": 185,
	"glde": 186,
	"glec": 187,
	"goga": 188,
	"gois": 189,
	"gosp": 190,
	"grba": 191,
	"grca": 192,
	"gree": 193,
	"greg": 194,
	"grfa": 195,
	"grko": 196,
	"grpo": 197,
	"grsa": 198,
	"grsm": 199,
	"grsp": 200,
	"grte": 201,
	"guco": 202,
	"guis": 203,
	"gumo": 204,
	"gwca": 205,
	"gwmp": 206,
	"hafe": 207,
	"hafo": 208,
	"hagr": 209,
	"haha": 210,
	"hale": 211,
	"hamp": 212,
	"hart": 213,
	"hatu": 214,
	"havo": 215,
	"heho": 216,
	"hobe": 217,
	"hocu": 218,
	"hofr": 219,
	"hofu": 220,
	"home": 221,
	"hono": 222,
	"hosp": 223,
	"hove": 224,
	"hstr": 225,
	"hutr": 226,
	"iafl": 227,
	"iatr": 228,
	"inde": 229,
	"indu": 230,
	"inup": 231,
	"isro": 232,
	"jaga": 233,
	"jame": 234,
	"jazz": 235,
	"jeca": 236,
	"jeff": 237,
	"jela": 238,
	"jica": 239,
	"joda": 240,
	"jofi": 241,
	"jofl": 242,
	"jomu": 243,
	"jotr": 244,
	"juba": 245,
	"kaho": 246,
	"kala": 247,
	"katm": 248,
	"kaww": 249,
	"keaq": 250,
	"kefj": 251,
	"kemo": 252,
	"kewe": 253,
	"kimo": 254,
	"klgo": 255,
	"klse": 256,
	"knri": 257,
	"kova": 258,
	"kowa": 259,
	"labe": 260,
	"lacl": 261,
	"lake": 262,
	"lamr": 263,
	"laro": 264,
	"lavo": 265,
	"lecl": 266,
	"lewi": 267,
	"libi": 268,
	"libo": 269,
	"liho": 270,
	"linc": 271,
	"liri": 272,
	"lode": 273,
	"loea": 274,
	"long": 275,
	"lowe": 276,
	"lyba": 277,
	"lyjo": 278,
	"maac": 279,
	"mabi": 280,
	"maca": 281,
	"malu": 282,
	"mamc": 283,
	"mana": 284,
	"manz": 285,
	"mapr": 286,
	"mava": 287,
	"mawa": 288,
	"memy": 289,
	"meve": 290,
	"miin": 291,
	"mima": 292,
	"mimi": 293,
	"misp": 294,
	"miss": 295,
	"mlkm": 296,
	"mnrr": 297,
	"moca": 298,
	"mocr": 299,
	"moja": 300,
	"mono": 301,
	"mopi": 302,
	"mora": 303,
	"morr": 304,
	"moru": 305,
	"muwo": 306,
	"nabr": 307,
	"nace": 308,
	"nama": 309,
	"natc": 310,
	"natr": 311,
	"natt": 312,
	"nava": 313,
	"nebe": 314,
	"neen": 315,
	"nepe": 316,
	"neph": 317,
	"neri": 318,
	"nico": 319,
	"niob": 320,
	"nisi": 321,
	"noat": 322,
	"noca": 323,
	"noco": 324,
	"npnh": 325,
	"npsa": 326,
	"obed": 327,
	"ocmu": 328,
	"okci": 329,
	"olsp": 330,
	"olym": 331,
	"orca": 332,
	"oreg": 333,
	"orpi": 334,
	"ovvi": 335,
	"oxhi": 336,
	"ozar": 337,
	"paal": 338,
	"paav": 339,
	"pagr": 340,
	"pais": 341,
	"para": 342,
	"peco": 343,
	"pefo": 344,
	"peri": 345,
	"pete": 346,
	"petr": 347,
	"pevi": 348,
	"pine": 349,
	"pinn": 350,
	"pipe": 351,
	"piro": 352,
	"pisc": 353,
	"pisp": 354,
	"poch": 355,
	"poex": 356,
	"pohe": 357,
	"popo": 358,
	"pore": 359,
	"prsf": 360,
	"prwi": 361,
	"puhe": 362,
	"puho": 363,
	"pull": 364,
	"rabr": 365,
	"redw": 366,
	"reer": 367,
	"rich": 368,
	"rigr": 369,
	"rira": 370,
	"roca": 371,
	"rocr": 372,
	"romo": 373,
	"rori": 374,
	"rowi": 375,
	"ruca": 376,
	"saan": 377,
	"sacn": 378,
	"sacr": 379,
	"safe": 380,
	"safr": 381,
	"saga": 382,
	"sagu": 383,
	"sahi": 384,
	"sair": 385,
	"sajh": 386,
	"saju": 387,
	"sama": 388,
	"samo": 389,
	"sand": 390,
	"sapa": 391,
	"sapu": 392,
	"sara": 393,
	"sari": 394,
	"scbl": 395,
	"seki": 396,
	"semo": 397,
	"shen": 398,
	"shil": 399,
	"sitk": 400,
	"slbe": 401,
	"spar": 402,
	"stea": 403,
	"stge": 404,
	"stli": 405,
	"ston": 406,
	"stri": 407,
	"stsp": 408,
	"sucr": 409,
	"tapr": 410,
	"thco": 411,
	"this": 412,
	"thje": 413,
	"thko": 414,
	"thrb": 415,
	"thri": 416,
	"thro": 417,
	"thst": 418,
	"tica": 419,
	"till": 420,
	"timu": 421,
	"tont": 422,
	"tosy": 423,
	"trte": 424,
	"tuai": 425,
	"tuin": 426,
	"tule": 427,
	"tuma": 428,
	"tupe": 429,
	"tusk": 430,
	"tuzi": 431,
	"ulsg": 432,
	"upde": 433,
	"vafo": 434,
	"vall": 435,
	"valr": 436,
	"vama": 437,
	"vick": 438,
	"vicr": 439,
	"viis": 440,
	"vive": 441,
	"voya": 442,
	"waba": 443,
	"waca": 444,
	"waco": 445,
	"wamo": 446,
	"wapa": 447,
	"waro": 448,
	"wefa": 449,
	"whho": 450,
	"whis": 451,
	"whmi": 452,
	"whsa": 453,
	"wica": 454,
	"wicl": 455,
	"wicr": 456,
	"wiho": 457,
	"wing": 458,
	"wori": 459,
	"wotr": 460,
	"wrbr": 461,
	"wrst": 462,
	"wupa": 463,
	"wwii": 464,
	"wwim": 465,
	"yell": 466,
	"york": 467,
	"yose": 468,
	"yuch": 469,
	"yuho": 470,
	"zion": 471,
}

var fees = map[string]uint16{
	"Commercial Entrance - Mini-bus":            1,
	"Commercial Entrance - Motor Coach":         2,
	"Commercial Entrance - Per Person":          3,
	"Commercial Entrance - Sedan":               4,
	"Commercial Entrance - Van":                 5,
	"Entrance - Education/Academic Groups":      6,
	"Entrance - Motorcycle":                     7,
	"Entrance - Non-commercial Groups":          8,
	"Entrance - Per Person":                     9,
	"Entrance - Private Vehicle":                10,
	"Entrance - Snowmobile":                     11,
	"Park Entrance Fee":                         12,
	"Timed Entry Reservation - Location":        13,
	"Timed Entry Reservation - Park":            14,
	"Timed Entry Reservation - Park & Location": 15,
}

var stateCodes = map[string]uint16{
	"AL": 1,
	"AK": 2,
	"AZ": 3,
	"AR": 4,
	"CA": 5,
	"CO": 6,
	"CT": 7,
	"DE": 8,
	"DC": 9,
	"FL": 10,
	"GA": 11,
	"HI": 12,
	"ID": 13,
	"IL": 14,
	"IN": 15,
	"IA": 16,
	"KS": 17,
	"KY": 18,
	"LA": 19,
	"ME": 20,
	"MD": 21,
	"MA": 22,
	"MI": 23,
	"MN": 24,
	"MS": 25,
	"MO": 26,
	"MT": 27,
	"NE": 28,
	"NV": 29,
	"NH": 30,
	"NJ": 31,
	"NM": 32,
	"NY": 33,
	"NC": 34,
	"ND": 35,
	"OH": 36,
	"OK": 37,
	"OR": 38,
	"PA": 39,
	"PR": 40,
	"RI": 41,
	"SC": 42,
	"SD": 43,
	"TN": 44,
	"TX": 45,
	"UT": 46,
	"VI": 47,
	"VT": 48,
	"VA": 49,
	"WA": 50,
	"WV": 51,
	"WI": 52,
	"WY": 53,
}

type Address struct {
	City                  string `json:"city"`
	CountryCode           string `json:"countryCode"`
	Line1                 string `json:"line1"`
	Line2                 string `json:"line2"`
	Line3                 string `json:"line3"`
	PostalCode            string `json:"postalCode"`
	ProvinceTerritoryCode string `json:"provinceTerritoryCode"`
	StateCode             string `json:"stateCode"`
	Type                  string `json:"type"`
}

type Campground struct {
	Accessibility struct {
		AccessRoads      []string `json:"accessRoads"`
		AdaInfo          string   `json:"adainfo"`
		AdditionaInfo    string   `json:"additionaInfo"`
		CellPhoneInfo    string   `json:"cellPhoneInfo"`
		Classifications  []string `json:"classifications"`
		FireStovePolicy  string   `json:"fireStovePolicy"`
		InternetInfo     string   `json:"internetInfo"`
		RvAllowed        string   `json:"rvAllowed"`
		RvInfo           string   `json:"rvInfo"`
		RvMaxLength      string   `json:"rvMaxLength"`
		TrailerMaxLength string   `json:"trailerMaxLength"`
		WheelchairAccess string   `json:"wheelchairAccess"`
	} `json:"accessibility"`
	Addresses []Address `json:"addresses"`
	Amenities struct {
		Amphitheater               string   `json:"amphitheater"`
		CampStore                  string   `json:"campStore"`
		CellPhoneReception         string   `json:"cellPhoneReception"`
		DumpStation                string   `json:"dumpStation"`
		FirewoodForSale            string   `json:"firewoodForSale"`
		FoodStorageLockers         string   `json:"foodStorageLockers"`
		IceAvailableForSale        string   `json:"iceAvailableForSale"`
		InternetConnectivity       string   `json:"internetConnectivity"`
		Laundry                    string   `json:"laundry"`
		PotableWater               []string `json:"potableWater"`
		Showers                    []string `json:"showers"`
		StaffOrVolunteerHostOnSite string   `json:"staffOrVolunteerHostOnsite"`
		Toilets                    []string `json:"toilets"`
		TrashRecyclingCollection   string   `json:"trashRecyclingCollection"`
	} `json:"amenities"`
	Campsites struct {
		ElectricalHookups string `json:"electricalHookups"`
		Group             string `json:"group"`
		Horse             string `json:"horse"`
		Other             string `json:"other"`
		RvOnly            string `json:"rvOnly"`
		TentOnly          string `json:"tentOnly"`
		TotalSites        string `json:"totalSites"`
		WalkBoatTo        string `json:"walkBoatTo"`
	} `json:"campsites"`
	Contacts           Contacts `json:"contacts"`
	Description        string   `json:"description"`
	DirectionsOverview string   `json:"directionsOverview"`
	DirectionsUrl      string   `json:"directionsUrl"`
	Fees               []Fee    `json:"fees"`
	GeometryPoiId      string   `json:"geometryPoiId"`
	Id                 string   `json:"id"`
	Images             []Image  `json:"images"`
	LastIndexedDate    string   `json:"lastIndexedDate"`
	LatLong            string   `json:"latLong"`
	Latitude           string   `json:"latitude"`
	Longitude          string   `json:"longitude"`
	Multimedia         []struct {
		Title string `json:"title"`
		Id    string `json:"id"`
		Type  string `json:"type"`
		Url   string `json:"url"`
	} `json:"multimedia"`
	Name                             string          `json:"name"`
	NumberOfSitesReservable          string          `json:"numberOfSitesReservable"`
	NumberOfSitesFirstComeFirstServe string          `json:"numberOfSitesFirstComeFirstServe"`
	OperatingHours                   []OperatingHour `json:"operatingHours"`
	ParkCode                         string          `json:"parkCode"`
	RegulationsOverview              string          `json:"regulationsOverview"`
	RegulationsUrl                   string          `json:"regulationsUrl"`
	RelevanceScore                   float64         `json:"relevanceScore"`
	ReservationInfo                  string          `json:"reservationInfo"`
	ReservationUrl                   string          `json:"reservationUrl"`
	WeatherOverview                  string          `json:"weatherOverview"`
}

type Contacts struct {
	PhoneNumbers   []PhoneNumber  `json:"phoneNumbers"`
	EmailAddresses []EmailAddress `json:"emailAddresses"`
}

type Crop struct {
	AspectRatio float64 `json:"aspectRatio"`
	URL         string  `json:"url"`
}

type EmailAddress struct {
	Description  string `json:"description"`
	EmailAddress string `json:"emailAddress"`
}

type Fee struct {
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type IdName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	AltText string `json:"altText"`
	Caption string `json:"caption"`
	Credit  string `json:"credit"`
	Crops   []Crop `json:"crops"`
	Title   string `json:"title"`
	URL     string `json:"url"`
}

type OperatingHour struct {
	Description string `json:"description"`
	Exceptions  []struct {
		Name           string `json:"name"`
		StartDate      string `json:"startDate"`
		EndDate        string `json:"endDate"`
		ExceptionHours Week   `json:"exceptionHours"`
	} `json:"exceptions"`
	Name          string `json:"name"`
	StandardHours Week   `json:"standardHours"`
}

type Park struct {
	Activities     []IdName        `json:"activities"`
	Addresses      []Address       `json:"addresses"`
	Contacts       Contacts        `json:"contacts"`
	Description    string          `json:"description"`
	Designation    string          `json:"designation"`
	DirectionsInfo string          `json:"directionsInfo"`
	DirectionsUrl  string          `json:"directionsUrl"`
	EntranceFees   []Fee           `json:"entranceFees"`
	EntrancePasses []Fee           `json:"entrancePasses"`
	FullName       string          `json:"fullName"`
	ID             string          `json:"id"`
	Images         []Image         `json:"images"`
	LatLong        string          `json:"latLong"`
	Latitude       string          `json:"latitude"`
	Longitude      string          `json:"longitude"`
	Name           string          `json:"name"`
	OperatingHours []OperatingHour `json:"operatingHours"`
	ParkCode       string          `json:"parkCode"`
	RelevanceScore float64         `json:"relevanceScore"`
	States         string          `json:"states"`
	Topics         []IdName        `json:"topics"`
	URL            string          `json:"url"`
	WeatherInfo    string          `json:"weatherInfo"`
}

type PhoneNumber struct {
	Description string `json:"description"`
	Extension   string `json:"extension"`
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
}

type Tour struct {
	Activities   []IdName `json:"activities"`
	Description  string   `json:"description"`
	DurationMax  string   `json:"durationMax"`
	DurationMin  string   `json:"durationMin"`
	DurationUnit string   `json:"durationUnit"`
	Id           string   `json:"id"`
	Images       []Image  `json:"images"`
	Park         struct {
		States      string `json:"states"`
		Designation string `json:"designation"`
		ParkCode    string `json:"parkCode"`
		FullName    string `json:"fullName"`
		Url         string `json:"url"`
		Name        string `json:"name"`
	} `json:"park"`
	RelevanceScore float64 `json:"relevanceScore"`
	Stops          []struct {
		Significance        string `json:"significance"`
		AssetId             string `json:"assetId"`
		AssetName           string `json:"assetName"`
		AssetType           string `json:"assetType"`
		AudioFileUrl        string `json:"audioFileUrl"`
		Id                  string `json:"id"`
		Ordinal             string `json:"ordinal"`
		DirectionToNextStop string `json:"directionToNextStop"`
	} `json:"stops"`
	Title  string   `json:"title"`
	Topics []IdName `json:"topics"`
}

type Week struct {
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
}

// API resources

type Campgrounds struct {
	Total string       `json:"total"`
	Data  []Campground `json:"data"`
	Limit string       `json:"limit"`
	Start string       `json:"start"`
}

func (c Campgrounds) SizeCurrent() int {
	return len(c.Data)
}

func (c Campgrounds) SizeTotal() int {
	total, err := strconv.Atoi(c.Total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

func (c Campgrounds) SqlduckCreate() string {
	return `CREATE SEQUENCE campground_id_seq START 1;
		CREATE TABLE campground (
		id UINTEGER DEFAULT nextval('campground_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		campsites_electrical_hookups UINTEGER NOT NULL,
		campsites_first_come_first_serve UINTEGER NOT NULL,
		campsites_reservable UINTEGER NOT NULL,
		campsites_total UINTEGER NOT NULL,
		has_camp_store_id UINTEGER NOT NULL,
		has_cell_phone_reception_id UINTEGER NOT NULL,
		has_laundry_id UINTEGER NOT NULL,
		is_rv_allowed UINTEGER NOT NULL,
		park_id UINTEGER NOT NULL,
		reservation_url VARCHAR(255) NOT NULL,
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);`
}

func (c Campgrounds) SqliteCreate() string {
	return `CREATE TABLE campground (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		campsites_electrical_hookups INTEGER NOT NULL,
		campsites_first_come_first_serve INTEGER NOT NULL,
		campsites_reservable INTEGER NOT NULL,
		campsites_total INTEGER NOT NULL,
		has_camp_store_id INTEGER NOT NULL,
		has_cell_phone_reception_id INTEGER NOT NULL,
		has_laundry_id INTEGER NOT NULL,
		is_rv_allowed INTEGER NOT NULL,
		park_id INTEGER NOT NULL,
		reservation_url TEXT NOT NULL,
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);`
}

func (c Campgrounds) SqlInsert(idStart int) []string {
	numQueries := len(c.Data)
	for _, campground := range c.Data {
		numQueries += len(campground.Fees)
	}
	queries := make([]string, 0, numQueries)
	campgroundId := idStart + 1
	for _, campground := range c.Data {
		campsitesElectricalHookups, err := strconv.Atoi(campground.Campsites.ElectricalHookups)
		if err != nil {
			log.Fatal(err)
		}
		campsitesFirstComeFirstServe, err := strconv.Atoi(campground.NumberOfSitesFirstComeFirstServe)
		if err != nil {
			log.Fatal(err)
		}
		campsitesReservable, err := strconv.Atoi(campground.NumberOfSitesReservable)
		if err != nil {
			log.Fatal(err)
		}
		campsitesTotal, err := strconv.Atoi(campground.Campsites.TotalSites)
		if err != nil {
			log.Fatal(err)
		}
		hasCampStore, ok := amenitySeasons[campground.Amenities.CampStore]
		if !ok {
			hasCampStore = 0
		}
		hasCellPhoneReception, ok := amenitySeasons[campground.Amenities.CellPhoneReception]
		if !ok {
			hasCellPhoneReception = 0
		}
		hasLaundry, ok := amenitySeasons[campground.Amenities.Laundry]
		if !ok {
			hasLaundry = 0
		}
		rvAllowed, err := strconv.Atoi(campground.Accessibility.RvAllowed)
		if err != nil {
			log.Fatal(err)
		}
		var isRvAllowed int
		if rvAllowed == 1 {
			isRvAllowed = 1
		} else {
			isRvAllowed = 0
		}
		parkId, ok := parkCodes[campground.ParkCode]
		if !ok {
			parkId = 0
		}
		queries = append(queries, fmt.Sprintf(
			"INSERT INTO campground (id, name, campsites_electrical_hookups, campsites_first_come_first_serve, campsites_reservable, campsites_total, has_camp_store_id, has_cell_phone_reception_id, has_laundry_id, is_rv_allowed, park_id, reservation_url) VALUES (%d, '%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, '%s');",
			campgroundId,
			strings.ReplaceAll(campground.Name, "'", "''"),
			campsitesElectricalHookups,
			campsitesFirstComeFirstServe,
			campsitesReservable,
			campsitesTotal,
			hasCampStore,
			hasCellPhoneReception,
			hasLaundry,
			isRvAllowed,
			parkId,
			campground.ReservationUrl,
		))
		campgroundId++
	}
	return queries
}

type Parks struct {
	Total string `json:"total"`
	Data  []Park `json:"data"`
	Limit string `json:"limit"`
	Start string `json:"start"`
}

func (p Parks) SizeCurrent() int {
	return len(p.Data)
}

func (p Parks) SizeTotal() int {
	total, err := strconv.Atoi(p.Total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

func (p Parks) SqlduckCreate() string {
	return `CREATE SEQUENCE park_id_seq START 1;
		CREATE TABLE park (
		id UINTEGER DEFAULT nextval('park_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		city VARCHAR(255) NOT NULL,
		state_code_id UINTEGER NOT NULL,
		url VARCHAR(255) NOT NULL,
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE SEQUENCE park_fee_id_seq START 1;
	CREATE TABLE park_fee (
		id UINTEGER DEFAULT nextval('park_fee_id_seq') PRIMARY KEY NOT NULL,
		park_id UINTEGER NOT NULL,
		fee_id UINTEGER NOT NULL,
		cost_cents UINTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (fee_id) REFERENCES fee(id)
	);`
}

func (p Parks) SqliteCreate() string {
	return `CREATE TABLE park (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		city TEXT NOT NULL,
		state_code_id INTEGER NOT NULL,
		url TEXT NOT NULL,
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE TABLE park_fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		park_id INTEGER NOT NULL,
		fee_id INTEGER NOT NULL,
		cost_cents INTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (fee_id) REFERENCES fee(id)
	);`
}

func (p Parks) SqlInsert(idStart int) []string {
	numQueries := len(p.Data)
	for _, park := range p.Data {
		numQueries += len(park.EntranceFees)
	}
	queries := make([]string, 0, numQueries)
	for _, park := range p.Data {
		city, stateCode := "", ""
		for _, address := range park.Addresses {
			if address.Type == "Physical" {
				city = address.City
				stateCode = address.StateCode
				break
			}
		}
		parkCode, ok := parkCodes[park.ParkCode]
		if !ok {
			parkCode = 0
		}
		state, ok := stateCodes[stateCode]
		if !ok {
			state = 0
		}
		queries = append(queries, fmt.Sprintf(
			"INSERT INTO park (id, name, city, state_code_id, url) VALUES (%d, '%s', '%s', %d, '%s');",
			parkCode,
			strings.ReplaceAll(park.Name, "'", "''"),
			strings.ReplaceAll(city, "'", "''"),
			state,
			park.URL,
		))
		for _, fee := range park.EntranceFees {
			costDollars, err := strconv.ParseFloat(fee.Cost, 32)
			if err != nil {
				log.Fatal(err)
			}
			costCents := int(costDollars * 100)
			fee, ok := fees[fee.Title]
			if !ok {
				fee = 0
			}
			queries = append(queries, fmt.Sprintf(
				"INSERT INTO park_fee (park_id, fee_id, cost_cents) VALUES (%d, %d, %d);",
				parkCode,
				fee,
				costCents,
			))
		}
	}
	return queries
}

type Tours struct {
	Total string `json:"total"`
	Data  []Tour `json:"data"`
	Limit string `json:"limit"`
	Start string `json:"start"`
}

func (t Tours) SizeCurrent() int {
	return len(t.Data)
}

func (t Tours) SizeTotal() int {
	total, err := strconv.Atoi(t.Total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

func (t Tours) SqlduckCreate() string {
	return `CREATE SEQUENCE tour_id_seq START 1;
		CREATE TABLE tour (
		id UINTEGER DEFAULT nextval('tour_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		duration_max UINTEGER NOT NULL,
		duration_min UINTEGER NOT NULL,
		duration_id UINTEGER NOT NULL,
		park_id UINTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);`
}

func (t Tours) SqliteCreate() string {
	return `CREATE TABLE tour (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		duration_max INTEGER NOT NULL,
		duration_min INTEGER NOT NULL,
		duration_id INTEGER NOT NULL,
		park_id INTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);`
}

func (t Tours) SqlInsert(idStart int) []string {
	numQueries := len(t.Data)
	for _, tour := range t.Data {
		numQueries += len(tour.Stops)
	}
	queries := make([]string, 0, numQueries)
	for _, tour := range t.Data {
		durationMax, err := strconv.Atoi(tour.DurationMax)
		if err != nil {
			log.Fatal(err)
		}
		durationMin, err := strconv.Atoi(tour.DurationMin)
		if err != nil {
			log.Fatal(err)
		}
		queries = append(queries, fmt.Sprintf(
			"INSERT INTO tour (name, duration_max, duration_min, duration_id, park_id) VALUES ('%s', %d, %d, %d, %d);",
			strings.ReplaceAll(tour.Title, "'", "''"),
			durationMax,
			durationMin,
			durations[tour.DurationUnit],
			parkCodes[tour.Park.ParkCode],
		))
	}
	return queries
}

type Resource interface {
	Campgrounds | Parks | Tours
	SizeCurrent() int
	SizeTotal() int
	SqlduckCreate() string
	SqliteCreate() string
	SqlInsert(idStart int) []string
}

type NpsClient struct {
	Client *http.Client
	Key    string
	base   string
}

func makeNpsClient(key string) *NpsClient {
	return &NpsClient{
		Client: &http.Client{},
		Key:    key,
		base:   "https://developer.nps.gov/api/v1/",
	}
}

func (c *NpsClient) buildUrl(resource string, start int) string {
	u, err := url.Parse(c.base + resource)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("api_key", c.Key)
	q.Set("limit", "50")
	q.Set("start", strconv.Itoa(start))
	u.RawQuery = q.Encode()
	return u.String()
}

func getOne[R Resource](c *NpsClient, resource string, start int) R {
	req, err := http.NewRequest("GET", c.buildUrl(resource, start), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	time.Sleep(1 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var r R
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func writeJson[R Resource](client *NpsClient, resource string) string {
	filename := fmt.Sprintf("data/%s.jsonl", resource)
	if _, err := os.Stat(filename); err == nil {
		log.Printf("File %s already exists", filename)
		return filename
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	current := 0
	one := getOne[R](client, resource, current)
	err = encoder.Encode(one)
	if err != nil {
		log.Fatal(err)
	}
	total := one.SizeTotal()
	current = one.SizeCurrent()
	fmt.Printf("%s: %d/%d\n", resource, current, total)
	for current < total {
		one := getOne[R](client, resource, current)
		err = encoder.Encode(one)
		if err != nil {
			log.Fatal(err)
		}
		current += one.SizeCurrent()
		fmt.Printf("%s: %d/%d\n", resource, current, total)
	}
	return filename
}

func writeSql[R Resource](resource string) {
	jsonFile, err := os.Open(fmt.Sprintf("data/%s.jsonl", resource))
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	sqlduckCreateFile, err := os.OpenFile("data/duckdb/create.sql", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlduckCreateFile.Close()
	sqliteCreateFile, err := os.OpenFile("data/sqlite/create.sql", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqliteCreateFile.Close()
	sqlInsertFile, err := os.OpenFile(fmt.Sprintf("data/insert/%s.sql", resource), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlInsertFile.Close()
	current := 0
	scanner := bufio.NewScanner(jsonFile)
	const maxCapacity = 512 * 1024 // 512KB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		line := scanner.Text()
		var r R
		err = json.Unmarshal([]byte(line), &r)
		if err != nil {
			log.Fatal(err)
		}
		if current == 0 {
			_, err = sqlduckCreateFile.WriteString(r.SqlduckCreate() + "\n")
			if err != nil {
				log.Fatal(err)
			}
			_, err = sqliteCreateFile.WriteString(r.SqliteCreate() + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, query := range r.SqlInsert(current) {
			_, err = sqlInsertFile.WriteString(query + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		current += r.SizeCurrent()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeSqlCreate() {
	sqlduckFile, err := os.OpenFile("data/duckdb/create.sql", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlduckFile.Close()
	_, err = sqlduckFile.WriteString(SqlduckCreate)
	if err != nil {
		log.Fatal(err)
	}
	sqliteFile, err := os.OpenFile("data/sqlite/create.sql", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqliteFile.Close()
	_, err = sqliteFile.WriteString(SqliteCreate)
	if err != nil {
		log.Fatal(err)
	}
	sqlInsertFile, err := os.OpenFile("data/insert/create.sql", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlInsertFile.Close()
	for key, value := range amenitySeasons {
		_, err = sqlInsertFile.WriteString(fmt.Sprintf(
			"INSERT INTO amenity_season (id, name) VALUES (%d,'%s');\n",
			value,
			key,
		))
		if err != nil {
			log.Fatal(err)
		}
	}
	for key, value := range durations {
		_, err = sqlInsertFile.WriteString(fmt.Sprintf(
			"INSERT INTO duration (id, name) VALUES (%d,'%s');\n",
			value,
			key,
		))
		if err != nil {
			log.Fatal(err)
		}
	}
	for key, value := range fees {
		_, err = sqlInsertFile.WriteString(fmt.Sprintf(
			"INSERT INTO fee (id, name) VALUES (%d,'%s');\n",
			value,
			key,
		))
		if err != nil {
			log.Fatal(err)
		}
	}
	for key, value := range stateCodes {
		_, err = sqlInsertFile.WriteString(fmt.Sprintf(
			"INSERT INTO state_code (id, name) VALUES (%d,'%s');\n",
			value,
			key,
		))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	npsClient := makeNpsClient(os.Getenv("NPS_API_KEY"))
	err := os.MkdirAll("data/duckdb", 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("data/insert", 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("data/sqlite", 0755)
	if err != nil {
		log.Fatal(err)
	}
	writeJson[Parks](npsClient, "parks")
	writeJson[Campgrounds](npsClient, "campgrounds")
	writeJson[Tours](npsClient, "tours")
	writeSqlCreate()
	writeSql[Parks]("parks")
	writeSql[Campgrounds]("campgrounds")
	writeSql[Tours]("tours")
}
