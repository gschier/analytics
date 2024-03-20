package main

func TimezoneToCountryCode(tz string) string {
	return timezoneToCountryCode[tz]
}

// timezoneToCountryCode was parsed from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//
//	```js
//	 const mapping = {};
//	 document.querySelectorAll('table tbody')[0].querySelectorAll('tr').forEach(tr => {
//	   const tds = tr.querySelectorAll('td');
//	   mapping[tds[2].innerText] = tds[0].innerText;
//	 });
//	 ```
var timezoneToCountryCode = map[string]string{
	"Africa/Abidjan":                   "CI",
	"Africa/Accra":                     "GH",
	"Africa/Addis_Ababa":               "ET",
	"Africa/Algiers":                   "DZ",
	"Africa/Asmara":                    "ER",
	"Africa/Asmera":                    "ER",
	"Africa/Bamako":                    "ML",
	"Africa/Bangui":                    "CF",
	"Africa/Banjul":                    "GM",
	"Africa/Bissau":                    "GW",
	"Africa/Blantyre":                  "MW",
	"Africa/Brazzaville":               "CG",
	"Africa/Bujumbura":                 "BI",
	"Africa/Cairo":                     "EG",
	"Africa/Casablanca":                "MA",
	"Africa/Ceuta":                     "ES",
	"Africa/Conakry":                   "GN",
	"Africa/Dakar":                     "SN",
	"Africa/Dar_es_Salaam":             "TZ",
	"Africa/Djibouti":                  "DJ",
	"Africa/Douala":                    "CM",
	"Africa/El_Aaiun":                  "EH",
	"Africa/Freetown":                  "SL",
	"Africa/Gaborone":                  "BW",
	"Africa/Harare":                    "ZW",
	"Africa/Johannesburg":              "ZA",
	"Africa/Juba":                      "SS",
	"Africa/Kampala":                   "UG",
	"Africa/Khartoum":                  "SD",
	"Africa/Kigali":                    "RW",
	"Africa/Kinshasa":                  "CD",
	"Africa/Lagos":                     "NG",
	"Africa/Libreville":                "GA",
	"Africa/Lome":                      "TG",
	"Africa/Luanda":                    "AO",
	"Africa/Lubumbashi":                "CD",
	"Africa/Lusaka":                    "ZM",
	"Africa/Malabo":                    "GQ",
	"Africa/Maputo":                    "MZ",
	"Africa/Maseru":                    "LS",
	"Africa/Mbabane":                   "SZ",
	"Africa/Mogadishu":                 "SO",
	"Africa/Monrovia":                  "LR",
	"Africa/Nairobi":                   "KE",
	"Africa/Ndjamena":                  "TD",
	"Africa/Niamey":                    "NE",
	"Africa/Nouakchott":                "MR",
	"Africa/Ouagadougou":               "BF",
	"Africa/Porto-Novo":                "BJ",
	"Africa/Sao_Tome":                  "ST",
	"Africa/Timbuktu":                  "ML",
	"Africa/Tripoli":                   "LY",
	"Africa/Tunis":                     "TN",
	"Africa/Windhoek":                  "NA",
	"America/Adak":                     "US",
	"America/Anchorage":                "US",
	"America/Anguilla":                 "AI",
	"America/Antigua":                  "AG",
	"America/Araguaina":                "BR",
	"America/Argentina/Buenos_Aires":   "AR",
	"America/Argentina/Catamarca":      "AR",
	"America/Argentina/ComodRivadavia": "AR",
	"America/Argentina/Cordoba":        "AR",
	"America/Argentina/Jujuy":          "AR",
	"America/Argentina/La_Rioja":       "AR",
	"America/Argentina/Mendoza":        "AR",
	"America/Argentina/Rio_Gallegos":   "AR",
	"America/Argentina/Salta":          "AR",
	"America/Argentina/San_Juan":       "AR",
	"America/Argentina/San_Luis":       "AR",
	"America/Argentina/Tucuman":        "AR",
	"America/Argentina/Ushuaia":        "AR",
	"America/Aruba":                    "AW",
	"America/Asuncion":                 "PY",
	"America/Atikokan":                 "CA",
	"America/Atka":                     "US",
	"America/Bahia":                    "BR",
	"America/Bahia_Banderas":           "MX",
	"America/Barbados":                 "BB",
	"America/Belem":                    "BR",
	"America/Belize":                   "BZ",
	"America/Blanc-Sablon":             "CA",
	"America/Boa_Vista":                "BR",
	"America/Bogota":                   "CO",
	"America/Boise":                    "US",
	"America/Buenos_Aires":             "AR",
	"America/Cambridge_Bay":            "CA",
	"America/Campo_Grande":             "BR",
	"America/Cancun":                   "MX",
	"America/Caracas":                  "VE",
	"America/Catamarca":                "AR",
	"America/Cayenne":                  "GF",
	"America/Cayman":                   "KY",
	"America/Chicago":                  "US",
	"America/Chihuahua":                "MX",
	"America/Coral_Harbour":            "CA",
	"America/Cordoba":                  "AR",
	"America/Costa_Rica":               "CR",
	"America/Creston":                  "CA",
	"America/Cuiaba":                   "BR",
	"America/Curacao":                  "CW",
	"America/Danmarkshavn":             "GL",
	"America/Dawson":                   "CA",
	"America/Dawson_Creek":             "CA",
	"America/Denver":                   "US",
	"America/Detroit":                  "US",
	"America/Dominica":                 "DM",
	"America/Edmonton":                 "CA",
	"America/Eirunepe":                 "BR",
	"America/El_Salvador":              "SV",
	"America/Ensenada":                 "MX",
	"America/Fort_Nelson":              "CA",
	"America/Fort_Wayne":               "US",
	"America/Fortaleza":                "BR",
	"America/Glace_Bay":                "CA",
	"America/Godthab":                  "GL",
	"America/Goose_Bay":                "CA",
	"America/Grand_Turk":               "TC",
	"America/Grenada":                  "GD",
	"America/Guadeloupe":               "GP",
	"America/Guatemala":                "GT",
	"America/Guayaquil":                "EC",
	"America/Guyana":                   "GY",
	"America/Halifax":                  "CA",
	"America/Havana":                   "CU",
	"America/Hermosillo":               "MX",
	"America/Indiana/Indianapolis":     "US",
	"America/Indiana/Knox":             "US",
	"America/Indiana/Marengo":          "US",
	"America/Indiana/Petersburg":       "US",
	"America/Indiana/Tell_City":        "US",
	"America/Indiana/Vevay":            "US",
	"America/Indiana/Vincennes":        "US",
	"America/Indiana/Winamac":          "US",
	"America/Indianapolis":             "US",
	"America/Inuvik":                   "CA",
	"America/Iqaluit":                  "CA",
	"America/Jamaica":                  "JM",
	"America/Jujuy":                    "AR",
	"America/Juneau":                   "US",
	"America/Kentucky/Louisville":      "US",
	"America/Kentucky/Monticello":      "US",
	"America/Knox_IN":                  "US",
	"America/Kralendijk":               "BQ",
	"America/La_Paz":                   "BO",
	"America/Lima":                     "PE",
	"America/Los_Angeles":              "US",
	"America/Louisville":               "US",
	"America/Lower_Princes":            "SX",
	"America/Maceio":                   "BR",
	"America/Managua":                  "NI",
	"America/Manaus":                   "BR",
	"America/Marigot":                  "MF",
	"America/Martinique":               "MQ",
	"America/Matamoros":                "MX",
	"America/Mazatlan":                 "MX",
	"America/Mendoza":                  "AR",
	"America/Menominee":                "US",
	"America/Merida":                   "MX",
	"America/Metlakatla":               "US",
	"America/Mexico_City":              "MX",
	"America/Miquelon":                 "PM",
	"America/Moncton":                  "CA",
	"America/Monterrey":                "MX",
	"America/Montevideo":               "UY",
	"America/Montreal":                 "CA",
	"America/Montserrat":               "MS",
	"America/Nassau":                   "BS",
	"America/New_York":                 "US",
	"America/Nipigon":                  "CA",
	"America/Nome":                     "US",
	"America/Noronha":                  "BR",
	"America/North_Dakota/Beulah":      "US",
	"America/North_Dakota/Center":      "US",
	"America/North_Dakota/New_Salem":   "US",
	"America/Nuuk":                     "GL",
	"America/Ojinaga":                  "MX",
	"America/Panama":                   "PA",
	"America/Pangnirtung":              "CA",
	"America/Paramaribo":               "SR",
	"America/Phoenix":                  "US",
	"America/Port-au-Prince":           "HT",
	"America/Port_of_Spain":            "TT",
	"America/Porto_Acre":               "BR",
	"America/Porto_Velho":              "BR",
	"America/Puerto_Rico":              "PR",
	"America/Punta_Arenas":             "CL",
	"America/Rainy_River":              "CA",
	"America/Rankin_Inlet":             "CA",
	"America/Recife":                   "BR",
	"America/Regina":                   "CA",
	"America/Resolute":                 "CA",
	"America/Rio_Branco":               "BR",
	"America/Rosario":                  "AR",
	"America/Santa_Isabel":             "MX",
	"America/Santarem":                 "BR",
	"America/Santiago":                 "CL",
	"America/Santo_Domingo":            "DO",
	"America/Sao_Paulo":                "BR",
	"America/Scoresbysund":             "GL",
	"America/Shiprock":                 "US",
	"America/Sitka":                    "US",
	"America/St_Barthelemy":            "BL",
	"America/St_Johns":                 "CA",
	"America/St_Kitts":                 "KN",
	"America/St_Lucia":                 "LC",
	"America/St_Thomas":                "VI",
	"America/St_Vincent":               "VC",
	"America/Swift_Current":            "CA",
	"America/Tegucigalpa":              "HN",
	"America/Thule":                    "GL",
	"America/Thunder_Bay":              "CA",
	"America/Tijuana":                  "MX",
	"America/Toronto":                  "CA",
	"America/Tortola":                  "VG",
	"America/Vancouver":                "CA",
	"America/Virgin":                   "VI",
	"America/Whitehorse":               "CA",
	"America/Winnipeg":                 "CA",
	"America/Yakutat":                  "US",
	"America/Yellowknife":              "CA",
	"Antarctica/Casey":                 "AQ",
	"Antarctica/Davis":                 "AQ",
	"Antarctica/DumontDUrville":        "AQ",
	"Antarctica/Macquarie":             "AU",
	"Antarctica/Mawson":                "AQ",
	"Antarctica/McMurdo":               "AQ",
	"Antarctica/Palmer":                "AQ",
	"Antarctica/Rothera":               "AQ",
	"Antarctica/South_Pole":            "AQ",
	"Antarctica/Syowa":                 "AQ",
	"Antarctica/Troll":                 "AQ",
	"Antarctica/Vostok":                "AQ",
	"Arctic/Longyearbyen":              "SJ",
	"Asia/Aden":                        "YE",
	"Asia/Almaty":                      "KZ",
	"Asia/Amman":                       "JO",
	"Asia/Anadyr":                      "RU",
	"Asia/Aqtau":                       "KZ",
	"Asia/Aqtobe":                      "KZ",
	"Asia/Ashgabat":                    "TM",
	"Asia/Ashkhabad":                   "TM",
	"Asia/Atyrau":                      "KZ",
	"Asia/Baghdad":                     "IQ",
	"Asia/Bahrain":                     "BH",
	"Asia/Baku":                        "AZ",
	"Asia/Bangkok":                     "TH",
	"Asia/Barnaul":                     "RU",
	"Asia/Beirut":                      "LB",
	"Asia/Bishkek":                     "KG",
	"Asia/Brunei":                      "BN",
	"Asia/Calcutta":                    "IN",
	"Asia/Chita":                       "RU",
	"Asia/Choibalsan":                  "MN",
	"Asia/Chongqing":                   "CN",
	"Asia/Chungking":                   "CN",
	"Asia/Colombo":                     "LK",
	"Asia/Dacca":                       "BD",
	"Asia/Damascus":                    "SY",
	"Asia/Dhaka":                       "BD",
	"Asia/Dili":                        "TL",
	"Asia/Dubai":                       "AE",
	"Asia/Dushanbe":                    "TJ",
	"Asia/Famagusta":                   "CY",
	"Asia/Gaza":                        "PS",
	"Asia/Harbin":                      "CN",
	"Asia/Hebron":                      "PS",
	"Asia/Ho_Chi_Minh":                 "VN",
	"Asia/Hong_Kong":                   "HK",
	"Asia/Hovd":                        "MN",
	"Asia/Irkutsk":                     "RU",
	"Asia/Istanbul":                    "TR",
	"Asia/Jakarta":                     "ID",
	"Asia/Jayapura":                    "ID",
	"Asia/Jerusalem":                   "IL",
	"Asia/Kabul":                       "AF",
	"Asia/Kamchatka":                   "RU",
	"Asia/Karachi":                     "PK",
	"Asia/Kashgar":                     "CN",
	"Asia/Kathmandu":                   "NP",
	"Asia/Katmandu":                    "NP",
	"Asia/Khandyga":                    "RU",
	"Asia/Kolkata":                     "IN",
	"Asia/Krasnoyarsk":                 "RU",
	"Asia/Kuala_Lumpur":                "MY",
	"Asia/Kuching":                     "MY",
	"Asia/Kuwait":                      "KW",
	"Asia/Macao":                       "MO",
	"Asia/Macau":                       "MO",
	"Asia/Magadan":                     "RU",
	"Asia/Makassar":                    "ID",
	"Asia/Manila":                      "PH",
	"Asia/Muscat":                      "OM",
	"Asia/Nicosia":                     "CY",
	"Asia/Novokuznetsk":                "RU",
	"Asia/Novosibirsk":                 "RU",
	"Asia/Omsk":                        "RU",
	"Asia/Oral":                        "KZ",
	"Asia/Phnom_Penh":                  "KH",
	"Asia/Pontianak":                   "ID",
	"Asia/Pyongyang":                   "KP",
	"Asia/Qatar":                       "QA",
	"Asia/Qostanay":                    "KZ",
	"Asia/Qyzylorda":                   "KZ",
	"Asia/Rangoon":                     "MM",
	"Asia/Riyadh":                      "SA",
	"Asia/Saigon":                      "VN",
	"Asia/Sakhalin":                    "RU",
	"Asia/Samarkand":                   "UZ",
	"Asia/Seoul":                       "KR",
	"Asia/Shanghai":                    "CN",
	"Asia/Singapore":                   "SG",
	"Asia/Srednekolymsk":               "RU",
	"Asia/Taipei":                      "TW",
	"Asia/Tashkent":                    "UZ",
	"Asia/Tbilisi":                     "GE",
	"Asia/Tehran":                      "IR",
	"Asia/Tel_Aviv":                    "IL",
	"Asia/Thimbu":                      "BT",
	"Asia/Thimphu":                     "BT",
	"Asia/Tokyo":                       "JP",
	"Asia/Tomsk":                       "RU",
	"Asia/Ujung_Pandang":               "ID",
	"Asia/Ulaanbaatar":                 "MN",
	"Asia/Ulan_Bator":                  "MN",
	"Asia/Urumqi":                      "CN",
	"Asia/Ust-Nera":                    "RU",
	"Asia/Vientiane":                   "LA",
	"Asia/Vladivostok":                 "RU",
	"Asia/Yakutsk":                     "RU",
	"Asia/Yangon":                      "MM",
	"Asia/Yekaterinburg":               "RU",
	"Asia/Yerevan":                     "AM",
	"Atlantic/Azores":                  "PT",
	"Atlantic/Bermuda":                 "BM",
	"Atlantic/Canary":                  "ES",
	"Atlantic/Cape_Verde":              "CV",
	"Atlantic/Faeroe":                  "FO",
	"Atlantic/Faroe":                   "FO",
	"Atlantic/Jan_Mayen":               "SJ",
	"Atlantic/Madeira":                 "PT",
	"Atlantic/Reykjavik":               "IS",
	"Atlantic/South_Georgia":           "GS",
	"Atlantic/St_Helena":               "SH",
	"Atlantic/Stanley":                 "FK",
	"Australia/ACT":                    "AU",
	"Australia/Adelaide":               "AU",
	"Australia/Brisbane":               "AU",
	"Australia/Broken_Hill":            "AU",
	"Australia/Canberra":               "AU",
	"Australia/Currie":                 "AU",
	"Australia/Darwin":                 "AU",
	"Australia/Eucla":                  "AU",
	"Australia/Hobart":                 "AU",
	"Australia/LHI":                    "AU",
	"Australia/Lindeman":               "AU",
	"Australia/Lord_Howe":              "AU",
	"Australia/Melbourne":              "AU",
	"Australia/North":                  "AU",
	"Australia/NSW":                    "AU",
	"Australia/Perth":                  "AU",
	"Australia/Queensland":             "AU",
	"Australia/South":                  "AU",
	"Australia/Sydney":                 "AU",
	"Australia/Tasmania":               "AU",
	"Australia/Victoria":               "AU",
	"Australia/West":                   "AU",
	"Australia/Yancowinna":             "AU",
	"Brazil/Acre":                      "BR",
	"Brazil/DeNoronha":                 "BR",
	"Brazil/East":                      "BR",
	"Brazil/West":                      "BR",
	"Canada/Atlantic":                  "CA",
	"Canada/Central":                   "CA",
	"Canada/Eastern":                   "CA",
	"Canada/Mountain":                  "CA",
	"Canada/Newfoundland":              "CA",
	"Canada/Pacific":                   "CA",
	"Canada/Saskatchewan":              "CA",
	"Canada/Yukon":                     "CA",
	"CET":                              "",
	"Chile/Continental":                "CL",
	"Chile/EasterIsland":               "CL",
	"CST6CDT":                          "",
	"Cuba":                             "CU",
	"EET":                              "",
	"Egypt":                            "EG",
	"Eire":                             "IE",
	"EST":                              "",
	"EST5EDT":                          "",
	"Etc/GMT":                          "",
	"Etc/GMT+0":                        "",
	"Etc/GMT+1":                        "",
	"Etc/GMT+10":                       "",
	"Etc/GMT+11":                       "",
	"Etc/GMT+12":                       "",
	"Etc/GMT+2":                        "",
	"Etc/GMT+3":                        "",
	"Etc/GMT+4":                        "",
	"Etc/GMT+5":                        "",
	"Etc/GMT+6":                        "",
	"Etc/GMT+7":                        "",
	"Etc/GMT+8":                        "",
	"Etc/GMT+9":                        "",
	"Etc/GMT-0":                        "",
	"Etc/GMT-1":                        "",
	"Etc/GMT-10":                       "",
	"Etc/GMT-11":                       "",
	"Etc/GMT-12":                       "",
	"Etc/GMT-13":                       "",
	"Etc/GMT-14":                       "",
	"Etc/GMT-2":                        "",
	"Etc/GMT-3":                        "",
	"Etc/GMT-4":                        "",
	"Etc/GMT-5":                        "",
	"Etc/GMT-6":                        "",
	"Etc/GMT-7":                        "",
	"Etc/GMT-8":                        "",
	"Etc/GMT-9":                        "",
	"Etc/GMT0":                         "",
	"Etc/Greenwich":                    "",
	"Etc/UCT":                          "",
	"Etc/Universal":                    "",
	"Etc/UTC":                          "",
	"Etc/Zulu":                         "",
	"Europe/Amsterdam":                 "NL",
	"Europe/Andorra":                   "AD",
	"Europe/Astrakhan":                 "RU",
	"Europe/Athens":                    "GR",
	"Europe/Belfast":                   "GB",
	"Europe/Belgrade":                  "RS",
	"Europe/Berlin":                    "DE",
	"Europe/Bratislava":                "SK",
	"Europe/Brussels":                  "BE",
	"Europe/Bucharest":                 "RO",
	"Europe/Budapest":                  "HU",
	"Europe/Busingen":                  "DE",
	"Europe/Chisinau":                  "MD",
	"Europe/Copenhagen":                "DK",
	"Europe/Dublin":                    "IE",
	"Europe/Gibraltar":                 "GI",
	"Europe/Guernsey":                  "GG",
	"Europe/Helsinki":                  "FI",
	"Europe/Isle_of_Man":               "IM",
	"Europe/Istanbul":                  "TR",
	"Europe/Jersey":                    "JE",
	"Europe/Kaliningrad":               "RU",
	"Europe/Kiev":                      "UA",
	"Europe/Kirov":                     "RU",
	"Europe/Lisbon":                    "PT",
	"Europe/Ljubljana":                 "SI",
	"Europe/London":                    "GB",
	"Europe/Luxembourg":                "LU",
	"Europe/Madrid":                    "ES",
	"Europe/Malta":                     "MT",
	"Europe/Mariehamn":                 "AX",
	"Europe/Minsk":                     "BY",
	"Europe/Monaco":                    "MC",
	"Europe/Moscow":                    "RU",
	"Europe/Nicosia":                   "CY",
	"Europe/Oslo":                      "NO",
	"Europe/Paris":                     "FR",
	"Europe/Podgorica":                 "ME",
	"Europe/Prague":                    "CZ",
	"Europe/Riga":                      "LV",
	"Europe/Rome":                      "IT",
	"Europe/Samara":                    "RU",
	"Europe/San_Marino":                "SM",
	"Europe/Sarajevo":                  "BA",
	"Europe/Saratov":                   "RU",
	"Europe/Simferopol":                "UA",
	"Europe/Skopje":                    "MK",
	"Europe/Sofia":                     "BG",
	"Europe/Stockholm":                 "SE",
	"Europe/Tallinn":                   "EE",
	"Europe/Tirane":                    "AL",
	"Europe/Tiraspol":                  "MD",
	"Europe/Ulyanovsk":                 "RU",
	"Europe/Uzhgorod":                  "UA",
	"Europe/Vaduz":                     "LI",
	"Europe/Vatican":                   "VA",
	"Europe/Vienna":                    "AT",
	"Europe/Vilnius":                   "LT",
	"Europe/Volgograd":                 "RU",
	"Europe/Warsaw":                    "PL",
	"Europe/Zagreb":                    "HR",
	"Europe/Zaporozhye":                "UA",
	"Europe/Zurich":                    "CH",
	"Factory":                          "",
	"GB":                               "GB",
	"GB-Eire":                          "GB",
	"GMT":                              "",
	"GMT+0":                            "",
	"GMT-0":                            "",
	"GMT0":                             "",
	"Greenwich":                        "",
	"Hongkong":                         "HK",
	"HST":                              "",
	"Iceland":                          "IS",
	"Indian/Antananarivo":              "MG",
	"Indian/Chagos":                    "IO",
	"Indian/Christmas":                 "CX",
	"Indian/Cocos":                     "CC",
	"Indian/Comoro":                    "KM",
	"Indian/Kerguelen":                 "TF",
	"Indian/Mahe":                      "SC",
	"Indian/Maldives":                  "MV",
	"Indian/Mauritius":                 "MU",
	"Indian/Mayotte":                   "YT",
	"Indian/Reunion":                   "RE",
	"Iran":                             "IR",
	"Israel":                           "IL",
	"Jamaica":                          "JM",
	"Japan":                            "JP",
	"Kwajalein":                        "MH",
	"Libya":                            "LY",
	"MET":                              "",
	"Mexico/BajaNorte":                 "MX",
	"Mexico/BajaSur":                   "MX",
	"Mexico/General":                   "MX",
	"MST":                              "",
	"MST7MDT":                          "",
	"Navajo":                           "US",
	"NZ":                               "NZ",
	"NZ-CHAT":                          "NZ",
	"Pacific/Apia":                     "WS",
	"Pacific/Auckland":                 "NZ",
	"Pacific/Bougainville":             "PG",
	"Pacific/Chatham":                  "NZ",
	"Pacific/Chuuk":                    "FM",
	"Pacific/Easter":                   "CL",
	"Pacific/Efate":                    "VU",
	"Pacific/Enderbury":                "KI",
	"Pacific/Fakaofo":                  "TK",
	"Pacific/Fiji":                     "FJ",
	"Pacific/Funafuti":                 "TV",
	"Pacific/Galapagos":                "EC",
	"Pacific/Gambier":                  "PF",
	"Pacific/Guadalcanal":              "SB",
	"Pacific/Guam":                     "GU",
	"Pacific/Honolulu":                 "US",
	"Pacific/Johnston":                 "UM",
	"Pacific/Kiritimati":               "KI",
	"Pacific/Kosrae":                   "FM",
	"Pacific/Kwajalein":                "MH",
	"Pacific/Majuro":                   "MH",
	"Pacific/Marquesas":                "PF",
	"Pacific/Midway":                   "UM",
	"Pacific/Nauru":                    "NR",
	"Pacific/Niue":                     "NU",
	"Pacific/Norfolk":                  "NF",
	"Pacific/Noumea":                   "NC",
	"Pacific/Pago_Pago":                "AS",
	"Pacific/Palau":                    "PW",
	"Pacific/Pitcairn":                 "PN",
	"Pacific/Pohnpei":                  "FM",
	"Pacific/Ponape":                   "FM",
	"Pacific/Port_Moresby":             "PG",
	"Pacific/Rarotonga":                "CK",
	"Pacific/Saipan":                   "MP",
	"Pacific/Samoa":                    "WS",
	"Pacific/Tahiti":                   "PF",
	"Pacific/Tarawa":                   "KI",
	"Pacific/Tongatapu":                "TO",
	"Pacific/Truk":                     "FM",
	"Pacific/Wake":                     "UM",
	"Pacific/Wallis":                   "WF",
	"Pacific/Yap":                      "FM",
	"Poland":                           "PL",
	"Portugal":                         "PT",
	"PRC":                              "CN",
	"PST8PDT":                          "",
	"ROC":                              "TW",
	"ROK":                              "KR",
	"Singapore":                        "SG",
	"Turkey":                           "TR",
	"UCT":                              "",
	"Universal":                        "",
	"US/Alaska":                        "US",
	"US/Aleutian":                      "US",
	"US/Arizona":                       "US",
	"US/Central":                       "US",
	"US/East-Indiana":                  "US",
	"US/Eastern":                       "US",
	"US/Hawaii":                        "US",
	"US/Indiana-Starke":                "US",
	"US/Michigan":                      "US",
	"US/Mountain":                      "US",
	"US/Pacific":                       "US",
	"US/Samoa":                         "WS",
	"UTC":                              "",
	"W-SU":                             "RU",
	"WET":                              "",
	"Zulu":                             "",
}
