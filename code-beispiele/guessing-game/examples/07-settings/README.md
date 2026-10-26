# Zahlenratespiel - Einstellungen aus einer Datei

Diese Ausbaustufe liest die Anzahl der Versuche und den Zahlenbereich aus
einer Datei `settings.txt`.

## Einstellungen

Die Datei enthält drei ganze Zahlen: zuerst die maximale Anzahl der Versuche,
dann die Untergrenze und die Obergrenze. Beide Grenzen gehören zum Zahlenbereich.
Leerzeichen und Zeilenumbrüche trennen die Werte.

## Erläuterungen

- `os.Open` öffnet die Datei zum Lesen. `defer file.Close()` sorgt dafür,
  dass sie beim Verlassen von `readSettings` wieder geschlossen wird.
- `fmt.Fscan` liest wie `fmt.Scan` Werte ein, hier aber aus der Datei.
- `rand.Intn(max_number-min_number+1) + min_number` wählt eine Zahl im
  eingestellten Bereich. Das `+1` sorgt dafür, dass auch die Obergrenze
  gewählt werden kann.
- `run` bekommt zusätzlich die beiden Grenzen, damit die Eingabeaufforderung
  den eingestellten Bereich anzeigt.

## Fortgeschrittenes Konzept: Standardwerte bei Fehlern

`readSettings` kapselt das Öffnen, Lesen und Prüfen der Einstellungen und
liefert drei Rückgabewerte statt eines Fehlers. Wenn die Datei fehlt, nicht
lesbar ist oder ungültige Werte enthält, liefert die Funktion vollständig
die globalen Standardwerte: drei Versuche und den Zahlenbereich 1 bis 100.
Diese sind als Konstanten außerhalb der Funktionen definiert. Das Spiel
kann dadurch auch bei fehlerhaften Einstellungen starten. Mehrere Rückgabewerte,
`defer` und diese Fehlerbehandlung sind fortgeschrittene Konzepte.

Die Fehlerbehandlung beschränkt sich zunächst auf die Einstellungen.
Für die Tipps des Spielers wird weiterhin eine gültige Ganzzahleingabe vorausgesetzt.
