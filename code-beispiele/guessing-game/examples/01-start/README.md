# Zahlenratespiel - Erste Version

Die Datei `main.go` enthält eine sehr einfache Version des Spiels,
bei der die zu ratende Zahl im Code festgelegt ist und der Benutzer
einen einzigen Versuch bekommt, die Zahl zu raten.

Das Beispiel zeigt die Verwendung von Ein-/Ausgabe, einer Variablen
und einer `if`-Anweisung. Es ist ein erster Einstieg ins Programmieren.

## Erläuterungen

- Die Zeilen `fmt.Print(...)` und `fmt.Scan(...)` lesen
  die Benutzereingabe in eine Variable `guess` ein.
- Die `if`-Anweisung überprüft, ob die Eingabe mit der Zahl 42 übereinstimmt.
- Die `fmt.Println(...)`-Anweisungen geben eine Rückmeldung auf der Konsole aus.

## Beobachtungen:
- Das Programm funktioniert für einen Versuch, ist aber noch nicht flexibel.
  Es bietet keinerlei Möglichkeit, mehrere Versuche zu machen oder die Zahl zu ändern.
