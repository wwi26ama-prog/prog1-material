# Zahlenratespiel - Mehrere Versuche

Die Datei `main.go` enthält einen Zwischenschritt, bei dem einige zukünftige
Verbesserungen vorbereitet werden. An der Funktionalität hat sich aber gegenüber
dem vorherigen Beispiel nichts geändert.

Hier wird gezeigt, wie die Verwendung von Variablen und Funktionen
den Code lesbarer und flexibler machen kann.

## Erläuterungen

- Der eigentliche Code ist i.W. identisch mit dem vorherigen Beispiel.
- Die zu ratende Zahl und die Anzahl der Versuche sind aber nicht mehr in
  der Programmlogik direkt festgelegt, sondern die Programmlogik ist stattdessen
  in einer Funktion `run()` definiert, die diese Werte als _Parameter_ erwartet.
- In der `main`-Funktion werden die Werte in Variablen gespeichert
  und an die Funktion übergeben.

## Beobachtungen:

- Die Lesbarkeit wird insofern verbessert, als es nun nur eine zentrale Stelle
  gibt, an der Werte stehen, die man potenziell verändern können will.
- Es gibt noch keine Möglichkeit, diese Werte wirklich zu ändern,
  man kann aber schon erahnen, dass diese nun z.B. beim Spielstart abgefragt
  werden oder auf andere Weise bestimmt werden könnten.
