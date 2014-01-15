sieci-projekt
=============

OKREŚLENIE PROBLEMU:
Mamy sieć komputerową, dla każdej maszyny istnieje z góry określona maksymalna pojemność - ilość maszyn, z którymi może być połączona (zakładamy, że jest to liczba większa od 2). Dochodzi nowy komputer. Co robimy?

IDEA ROZWIĄZANIA:
Tworzymy sieć w formie ukorzenionego drzewa (korzeń stanowi pierwsza maszyna). Nowy komputer dołączamy do takiego wierzchołka, który nie osiągnął maksymalnej pojemności i ma najniższą głębokość. Jeśli takich wierzchołków jest więcej, wybieramy ten najbardziej na lewo.

DOWÓD POPRAWNOŚCI ALGORYTMU DODANIA KOMPUTERA:
1. Nowo dodany komputer połączony będzie z co najwyżej jednym komputerem dodanym wcześniej. Dlatego nigdy nie powstanie nam cykl i sieć zachowa strukturę drzewa.
2. Każde drzewo ma liście. Pojemność wierzchołków jest większa od 2, zatem zawsze istnieje wierzchołek, do którego możemy coś dodać.

=============

SZCZEGÓŁY ROZWIĄZANIA

WYBRANY JĘZYK: go