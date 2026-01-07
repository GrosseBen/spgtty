spgtty/
├── go.mod                # Go-Modul-Definition
├── go.sum                # Abhängigkeits-Checksummen
├── main.go               # Haupt-Einstiegspunkt (CLI/Logik)
├── cmd/                  # Befehle/Subkommandos (z.B. für CLI-Tools)
│   └── spgtty/           # Beispiel: `spgtty deploy`, `spgtty test`
│       └── main.go
├── pkg/                  # Wiederverwendbare Pakete (für die Community)
│   ├── core/             # Kernfunktionen (z.B. Shelly-API-Wrapper)
│   │   └── api.go
│   ├── utils/            # Helferfunktionen (z.B. Logging, Fehlerbehandlung)
│   │   └── logger.go
│   └── scripts/          # Vordefinierte Skript-Templates
│       └── template.go
├── internal/             # Interne Logik (nicht für die Community)
│   └── config/           # Konfigurationen (z.B. Geräte-Profile)
│       └── config.go
├── scripts/              # Beispielskripte für Nutzer
│   └── example.sh
├── tests/                # Tests (sekundär, aber sinnvoll für Kernfunktionen)
│   └── core_test.go
└── README.md             # Dokumentation für die Community
