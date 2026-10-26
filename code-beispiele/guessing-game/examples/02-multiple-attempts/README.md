# Zahlenratespiel - Mehrere Versuche

Die Datei `main.go` enthält eine leicht erweiterte Version des Spiels,
bei der der Benutzer mehrere Versuche bekommt, die Zahl zu raten.
Die Zahl ist weiterhin fest im Code hinterlegt.

Das Beispiel zeigt die Verwendung von Schleifen, um eine Aktion mehrfach auszuführen.

## Erläuterungen

- Eingabe und Überprüfung sind fast exakt wie im vorherigen Beispiel.
- Der Code steht aber in einer Schleife (`for range 3`),
  dadurch wird dieser Code drei Mal wiederholt.
- Bei richtigem Raten wird die Funktion vorzeitig abgebrochen (`return`)
- 

## Beobachtungen:
- Das ist insofern besser als es nun mehrere Versuche gibt.
- Eine hohe "Wiederspielbarkeit" ergibt sich aber nicht,
  weil die Zahl weiterhin fest codiert ist.
