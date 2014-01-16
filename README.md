sieci-projekt
=============

OKREŚLENIE PROBLEMU:
Tworzymy sieć komputerową o strukturze drzewowej, dla każdej maszyny istnieje z góry określona maksymalna pojemność - ilość maszyn, z którymi może być połączona (zakładamy, że jest to liczba większa od 2). Chcemy zachować najmniejszą możliwą średnicę sieci, tj. chcemy, aby odległość między najbardziej oddalonymi wierzchołkami była jak najmniejsza. Dochodzi nowy komputer, który musimy dodać on-line. Co robimy?

IDEA ROZWIĄZANIA:
Tworzymy sieć w formie ukorzenionego drzewa (korzeń stanowi pierwsza maszyna). Nowy komputer dołączamy do takiego wierzchołka, który nie osiągnął maksymalnej pojemności i ma najniższą głębokość. Jeśli takich wierzchołków jest więcej, wybieramy ten najbardziej na lewo.

CO PRZEMAWIA ZA UŻYCIEM WŁAŚNIE TEGO ALGORYTMU DODANIA KOMPUTERA:
1. Nowo dodany komputer połączony będzie z co najwyżej jednym komputerem dodanym wcześniej. Dlatego nigdy nie powstanie nam cykl i sieć zachowa strukturę drzewa.
2. Każde drzewo ma liście. Pojemność wierzchołków jest większa od 2, zatem zawsze istnieje wierzchołek, do którego możemy coś dodać.
3. Zauważmy, że średnica powiększa się, gdy tworzymy nowy poziom w drzewie. Jednak nasze postępowanie prowadzi do tego, by nowy poziom tworzony był tylko w wypadku całkowitego zapełnienia poprzedniego.

JEDNAK POJAWIA SIĘ PEWIEN PROBLEM: w związku z tym, że sieć budowana jest on-line, globalnie rozwiązanie nie jest optymalne -- gdy wierzchołki dochodzą rosnąco po pojemności (co w prawdziwym życiu również się zdarza - nowa maszyna ma spore szanse być lepszą), poziomów może być więcej niż potrzeba (przykład takiej sytuacji mamy, gdy przychodzące wierzchołki mają współczynnik pojemności równy odpowiednio 3, 5, 3, 3, 3. Zbudujemy wtedy drzewo o średnicy 3, podczas gdy możemy ukorzenić je w drugim wierzchołku i ma wówczas średnicę 2).

PROPONOWANE ROZWIĄZANIE: nasze drzewo będzie się automatycznie równoważyć po wykonaniu pewnej ilości ruchów. Ilość wykonanych ruchów powinna być asymptotycznie równa złożoności algorytmu równoważenia.

=============

SZCZEGÓŁY ROZWIĄZANIA

WYBRANY JĘZYK: go

DODAWANIE NOWEGO KOMPUTERA: