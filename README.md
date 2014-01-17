określenie problemu
=============

Tworzymy sieć komputerową o strukturze drzewowej, dla każdej maszyny istnieje z góry określona maksymalna pojemność - ilość maszyn, z którymi może być połączona (zakładamy, że jest to liczba większa od 2). Chcemy zachować najmniejszą możliwą średnicę sieci, tj. chcemy, aby odległość między najbardziej oddalonymi wierzchołkami była jak najmniejsza. Dochodzi nowy komputer, który musimy dodać on-line. Co robimy?

propozycja rozwiązania
=============

IDEA ROZWIĄZANIA:
Tworzymy sieć w formie ukorzenionego drzewa (korzeń stanowi pierwsza maszyna). Nowy komputer dołączamy do takiego wierzchołka, który nie osiągnął maksymalnej pojemności i ma najniższą głębokość. Jeśli takich wierzchołków jest więcej, wybieramy ten najbardziej na lewo.

CO PRZEMAWIA ZA UŻYCIEM WŁAŚNIE TEGO ALGORYTMU DODANIA KOMPUTERA:
1. Nowo dodany komputer połączony będzie z co najwyżej jednym komputerem dodanym wcześniej. Dlatego nigdy nie powstanie nam cykl i sieć zachowa strukturę drzewa.
2. Każde drzewo ma liście. Pojemność wierzchołków jest większa od 2, zatem zawsze istnieje wierzchołek, do którego możemy coś dodać.
3. Zauważmy, że średnica powiększa się, gdy tworzymy nowy poziom w drzewie. Jednak nasze postępowanie prowadzi do tego, by nowy poziom tworzony był tylko w wypadku całkowitego zapełnienia poprzedniego.

JEDNAK POJAWIA SIĘ PEWIEN PROBLEM: w związku z tym, że sieć budowana jest on-line, globalnie rozwiązanie nie jest optymalne -- gdy wierzchołki dochodzą rosnąco po pojemności (co w prawdziwym życiu również się zdarza - nowa maszyna ma spore szanse być lepszą), poziomów może być więcej niż potrzeba (przykład takiej sytuacji mamy, gdy przychodzące wierzchołki mają współczynnik pojemności równy odpowiednio 3, 5, 3, 3, 3. Zbudujemy wtedy drzewo o średnicy 3, podczas gdy możemy ukorzenić je w drugim wierzchołku i ma wówczas średnicę 2).

PROPONOWANE ROZWIĄZANIE: nasze drzewo będzie się automatycznie równoważyć po wykonaniu pewnej ilości ruchów. 

szczegóły rozwiązania
=============

DODAWANIE NOWEGO KOMPUTERA: na początku informację o tym, że przychodzi nowy komputer, otrzymuje korzeń. Następnie maszyny postępują według następującego schematu: 

jeśli Twoja maksymalna pojemność nie została osiągnięta, poślij do rodzica informację zwrotną postaci (twoje ID, twoja głębokość w drzewie). W przeciwnym wypadku, poślij zapytanie do każdego dziecka o wynik w jego podrzewie. Otrzymasz k różnych wyników, gdzie k jest liczbą twoich dzieci. Spośród nich wybierz najlepszy, tj. spośród wyników o minimalnej głębokości wybierz ten najbardziej po lewej. Przekaż wynik w informacji zwrotnej dla rodzica.

SAMORÓWNOWAŻENIE SIĘ: każdy komputer może przechowywać informację o tym, który wierzchołek z jego poddrzewa ma największą pojemność. Dzięki temu znalezienie najbardziej pojemnego wierzchołka jest logarytmiczne. Również zamieniając dziecko z rodzicem możemy w czasie logarytmicznym "wywindować" je na swoje miejsce. Procedurę przerywamy, gdy każdy wierzchołek wskazuje na siebie samego.

szczegóły implementacji
=============

WYBRANY JĘZYK: go

POŁĄCZENIE MIĘDZY KOMUPUTERAMI: socket

wybrane protokoły
=============
