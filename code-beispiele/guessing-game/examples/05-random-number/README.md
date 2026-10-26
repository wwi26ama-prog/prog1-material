# Zahlenratespiel - Mehrere Versuche

Die Datei `main.go` enthält eine verbesserte Version des Spiels,
bei der die zu erratende Zahl nun zufällig gewählt wird.

## Erläuterungen

- In der `main`-Funktion wird die zu erratende Zahl nun mit der
  Funktion `rand.Intn(100)` erzeugt.
  
## Beobachtungen:

- Das Spiel ist damit "feature-complete": Die Grundfunktionalität
ist gegeben und es gibt keine offensichtlichen Einschränkungen mehr.
- Als mögliche Erweiterung könnte man nun noch die Anzahl der Versuche
oder den Zahlenbereich variabel gestalten (siehe Beispiel `07-settings`).
